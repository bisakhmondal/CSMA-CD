package utils

import "testing"

func TestFrame(t *testing.T) {
	frame := MakeFrame(1, 1, 4566, 56)
	t.Log(frame)
	t.Log(GetData(frame))
	t.Log(GetNodeInfo(frame))

}