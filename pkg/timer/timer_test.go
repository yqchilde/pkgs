package timer

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var job = mockJob{}

type mockJob struct{}

func (job mockJob) Run() {
	mockFunc()
}

func mockFunc() {
	time.Sleep(time.Second)
	fmt.Println("1s...")
}

func TestNewTimerTask(t *testing.T) {
	timerTask := NewTimerTask()

	t.Run("AddTaskByFunc", func(t *testing.T) {
		_, err := timerTask.AddTaskByFunc("func", "@every 1s", mockFunc)
		assert.Nil(t, err)

		_, ok := timerTask.(*timer).taskList["func"]
		if !ok {
			t.Error("no find func")
		}
	})

	t.Run("AddTaskByJob", func(t *testing.T) {
		_, err := timerTask.AddTaskByJob("job", "@every 1s", job)
		assert.Nil(t, err)

		_, ok := timerTask.(*timer).taskList["job"]
		if !ok {
			t.Error("no find job")
		}
	})

	t.Run("FindTaskCron", func(t *testing.T) {
		_, ok := timerTask.FindTaskCron("func")
		if !ok {
			t.Error("no find func")
		}
		_, ok = timerTask.FindTaskCron("job")
		if !ok {
			t.Error("no find job")
		}
		_, ok = timerTask.FindTaskCron("none")
		if ok {
			t.Error("find none")
		}
	})

	t.Run("ClearTask", func(t *testing.T) {
		timerTask.ClearTask("func")
		_, ok := timerTask.FindTaskCron("func")
		if ok {
			t.Error("find func")
		}
	})
}
