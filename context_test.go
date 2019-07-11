package gos

import (
	"context"
	"testing"
)

var c = &Context{
	Ctx:      context.Background(),
	NextFlag: true,
}

func TestNext(t *testing.T) {
	if !c.Next() {
		t.Error("get next error result")
	}
}

func TestFail(t *testing.T) {
	c.Fail()
	if c.NextFlag {
		t.Error("Fail() run error")
	}
}
