package event

import (
	"sync"
	"yunfei/app/event/listener"
	"yunfei/internal/event"
	"yunfei/internal/global"
)

var listenerList = []event.ListenerInterface{
	listener.FooListener{},
	listener.BarListener{},
}
var once sync.Once

func init() {
	once.Do(func() {
		global.EventDispatcher = event.New()
	})
	for _, l := range listenerList {
		global.EventDispatcher.Register(l)
	}
}
