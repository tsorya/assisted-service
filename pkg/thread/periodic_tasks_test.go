package thread

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJob(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Monitor tests")
}

type Test struct {
	index int
}

func (t *Test) IncreaseIndex() {
	t.index += 1
}

var _ = Describe("Monitor tests", func() {
	It("testing monitor runner", func() {
		test1 := Test{}
		test2 := Test{}
		monitors := []PeriodicTask{
			{Name: "Test", Interval: 10 * time.Millisecond, Exec: test1.IncreaseIndex},
			{Name: "Test2", Interval: 30 * time.Millisecond, Exec: test2.IncreaseIndex},
		}
		monitorsRunner := NewPeriodicTasksRunner(log.WithField("pkg", "cluster-monitor"), monitors)
		monitorsRunner.Start()
		time.Sleep(100 * time.Millisecond)
		monitorsRunner.Stop()
		index := test1.index
		Expect(test1.index > 7 && test1.index < 12).Should(Equal(true))
		index2 := test2.index
		Expect(test2.index > 2 && test2.index < 4).Should(Equal(true))

		By("Verifying all tasks were stopped and are not increasing anymore")
		time.Sleep(100 * time.Millisecond)
		Expect(test1.index).Should(Equal(index))
		Expect(test2.index).Should(Equal(index2))
	})
})
