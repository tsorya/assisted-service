package thread

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type PeriodicTask struct {
	Name     string
	Interval time.Duration
	Exec     func()
}

type PeriodicTasks struct {
	log   logrus.FieldLogger
	done  chan struct{}
	tasks []PeriodicTask
}

func NewPeriodicTasksRunner(log logrus.FieldLogger, monitors []PeriodicTask) *PeriodicTasks {
	return &PeriodicTasks{
		log:   log,
		tasks: monitors,
	}
}

// Start monitors thread
func (m *PeriodicTasks) Start() {
	m.done = make(chan struct{})
	m.log.Infof("Started monitors thread")
	go m.runTasks()
}

// Stop monitors
func (m *PeriodicTasks) Stop() {
	m.log.Infof("Stopping all monitors")
	close(m.done)
	m.log.Infof("Stopped all monitors")
}

func (m *PeriodicTasks) runTasks() {
	var wg sync.WaitGroup
	wg.Add(len(m.tasks))
	for _, task := range m.tasks {
		go m.taskRunnerLoop(m.done, &wg, task)
	}
	m.log.Infof("Waiting for all tasks to finish")
	wg.Wait()
	m.log.Infof("Finished running all tasks")
}

func (m *PeriodicTasks) taskRunnerLoop(done <-chan struct{}, wg *sync.WaitGroup, task PeriodicTask) {
	m.log.Infof("Start running %s task", task.Name)
	defer wg.Done()
	ticker := time.NewTicker(task.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			m.log.Infof("Stop running %s task", task.Name)
			return
		case <-ticker.C:
			task.Exec()
		}
	}
}
