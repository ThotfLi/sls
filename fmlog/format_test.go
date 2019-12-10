package fmlog

import "testing"

func TestNew(t *testing.T) {
	n:=New("这是一条消息",TIME|FILE|STACK)
	println(n)
}

