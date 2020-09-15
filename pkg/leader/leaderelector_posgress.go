package leader

import (
	"cirello.io/pglock"
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)


type DbElector struct {
	log           logrus.FieldLogger
	db			  *gorm.DB
	config        Config
	isLeader      bool
	leaderName    string
}



func NewDbElector(db *gorm.DB, config Config, leaderName string, logger logrus.FieldLogger) *DbElector {
	return &DbElector{db: db, log: logger, config: config, isLeader: false, leaderName: leaderName}
}

func (l *DbElector) IsLeader() bool {
	return l.isLeader
}

func (l *DbElector) StartLeaderElection(ctx context.Context) error {

	c, err := pglock.New(l.db.DB(),
		pglock.WithLeaseDuration(l.config.LeaseDuration),
		pglock.WithHeartbeatFrequency(l.config.RetryInterval),
		pglock.WithCustomTable(l.leaderName))
	if err != nil {
		l.log.WithError(err).Error("Failed to create db lock")
	}
	err = c.CreateTable()
	if err != nil {
		if p, ok := errors.Unwrap(err).(*pq.Error); !ok || p.Code.Name() != "duplicate_table" {
			l.log.WithError(err).Infof("CCCCCCCCCCCCCCCCCCCCCCCCCCC")
			return err
		}
	}

	go func() {
		var lock *pglock.Lock
		defer func() {
			if lock != nil {
				lock.Close()
			}
		}()
		for {
			if ctx.Err() != nil {
				return
			}
			l.log.Infof("BBBBBBBBBBBBBBBBBBBBBBBBB")
			err = c.Do(ctx, l.leaderName, l.locked)
			l.log.WithError(err).Infof("AAAAAAAAAAAAAAAAAAAAAAAA")
		}
	}()

	return nil
}
func (l *DbElector) locked(ctx context.Context, lock *pglock.Lock) error{
	l.log.Infof("GGGGGGGGGGGGGGGGGGGGGGGGGGGG")
	l.isLeader = true
	<-ctx.Done()
	l.isLeader = false
	return nil
}