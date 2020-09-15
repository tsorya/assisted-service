package db

import (
	"time"

	"cirello.io/pglock"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func CreateDbLockClient(db *gorm.DB, log logrus.FieldLogger) (*pglock.Client, error) {
	c, err := pglock.New(db.DB(),
		pglock.WithLeaseDuration(15*time.Second),
		pglock.WithHeartbeatFrequency(1*time.Second),
		pglock.WithCustomTable("aaaa"))
	if err != nil {
		log.WithError(err).Error("Failed to create db lock")
		return nil, err
	}
	err = c.CreateTable()
	if err != nil {
		if p, ok := errors.Unwrap(err).(*pq.Error); !ok || p.Code.Name() != "duplicate_table" {
			log.WithError(err).Infof("Failed to create lock table")
			return nil, err
		}
	}
	return c, nil
}
