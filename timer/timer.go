package timer

import (
	"sync"

	"github.com/robfig/cron/v3"
)

// spec参考：https://pkg.go.dev/github.com/robfig/cron

type Timer interface {
	// AddTaskByFunc 通过函数的方式添加任务
	AddTaskByFunc(taskName string, spec string, task func()) (cron.EntryID, error)

	// AddTaskByJob 通过接口的方法添加任务
	AddTaskByJob(taskName string, spec string, job interface{ Run() }) (cron.EntryID, error)

	// FindTaskCron 获取对应的taskName的cron，可能为空
	FindTaskCron(taskName string) (*cron.Cron, bool)

	// StartTask 开始任务
	StartTask(taskName string)

	// StopTask 停止任务
	StopTask(taskName string)

	// RemoveTask 删除指定任务
	RemoveTask(taskName string, id int)

	// ClearTask 清除任务
	ClearTask(taskName string)

	// Close 释放资源
	Close()
}

type timer struct {
	taskList map[string]*cron.Cron
	sync.Mutex
}

// AddTaskByFunc 通过函数的方式添加任务
func (t *timer) AddTaskByFunc(taskName string, spec string, task func()) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[taskName]; !ok {
		t.taskList[taskName] = cron.New(cron.WithSeconds())
	}
	id, err := t.taskList[taskName].AddFunc(spec, task)
	t.taskList[taskName].Start()
	return id, err
}

// AddTaskByJob 通过接口的方法添加任务
func (t *timer) AddTaskByJob(taskName string, spec string, job interface{ Run() }) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[taskName]; !ok {
		t.taskList[taskName] = cron.New(cron.WithSeconds())
	}
	id, err := t.taskList[taskName].AddJob(spec, job)
	t.taskList[taskName].Start()
	return id, err
}

// FindTaskCron 获取对应的taskName的cron，可能为空
func (t *timer) FindTaskCron(taskName string) (*cron.Cron, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.taskList[taskName]
	return v, ok
}

// StartTask 开始任务
func (t *timer) StartTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.Start()
	}
}

// StopTask 停止任务
func (t *timer) StopTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.Stop()
	}
}

// RemoveTask 删除指定任务
func (t *timer) RemoveTask(taskName string, id int) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.Remove(cron.EntryID(id))
	}
}

// ClearTask 清除任务
func (t *timer) ClearTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.taskList[taskName]; ok {
		v.Stop()
		delete(t.taskList, taskName)
	}
}

// Close 释放资源
func (t *timer) Close() {
	t.Lock()
	defer t.Unlock()
	for _, v := range t.taskList {
		v.Stop()
	}
}

func NewTimerTask() Timer {
	return &timer{taskList: make(map[string]*cron.Cron)}
}
