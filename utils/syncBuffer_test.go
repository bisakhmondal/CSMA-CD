package utils

import (
	"testing"
)

func TestNewSyncBuffer(t *testing.T) {
	buf := NewSyncBuffer()
	go func(){
		data := []Bytestream{"Bisakh", "Mondal", "Abc", "def", "ghi", "jkl"}
		for _, j := range data{
			buf.Push(j)
			//time.Sleep(4*time.Second)
		}
	}()

	for i:=0; i<6;i++{
		t.Log(buf.Pop()+"\n")
		t.Log("waiting\n")
	}
}
func TestNewSyncBuffer2(t *testing.T){
	t.Log("Nice")
}

