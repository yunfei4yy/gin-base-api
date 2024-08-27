package crontab

import "github.com/robfig/cron/v3"

// Crontab 类型
type Crontab struct {
	cron *cron.Cron
}

// TaskInterface 任务接口
type TaskInterface interface {
	Spec() string
	Fn() func()
}

// Init 函数用于初始化一个Crontab实例
func Init() *Crontab {
	c := &Crontab{
		cron: cron.New(),
	}
	return c
}

// Start 启动crontab服务
func (c *Crontab) Start() {
	c.cron.Start()
}

// Stop 停止crontab服务
func (c *Crontab) Stop() {
	c.cron.Stop()
}

// AddTask 添加任务
func (c *Crontab) AddTask(tasks ...TaskInterface) {
	if len(tasks) == 0 {
		return
	}

	for _, task := range tasks {
		_, err := c.cron.AddFunc(task.Spec(), task.Fn())
		if err != nil {
			return
		}
	}
}
