package test

import (
	"testing"
	"yunfei/app/event/event"
	"yunfei/internal/global"
)

func TestEvent(t *testing.T) {
	global.EventDispatcher.Dispatch(&event.FooEvent{})
}
