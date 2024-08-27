package task

import "yunfei/internal/crontab"

func Task() []crontab.TaskInterface {
	return []crontab.TaskInterface{
		&FooTask{},
	}
}
