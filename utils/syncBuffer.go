package utils

import "sync"

type Bytestream=string

type SyncBuffer struct {
	Buf  []Bytestream
	lock sync.Mutex
	cond *sync.Cond
}

func NewSyncBuffer() *SyncBuffer{
	sb := &SyncBuffer{
		Buf: []Bytestream{},
		lock: sync.Mutex{},
	}
	sb.cond = sync.NewCond(&sb.lock)
	return sb
}

func (sb *SyncBuffer)Push(bytestream Bytestream){
	sb.cond.L.Lock()
	sb.Buf = append(sb.Buf, bytestream)
	sb.cond.Broadcast()
	sb.cond.L.Unlock()
}

func (sb *SyncBuffer)Len() int{
	sb.lock.Lock()
	le := len(sb.Buf)
	sb.lock.Unlock()
	return le
}

func (sb *SyncBuffer)Pop() Bytestream{
	sb.cond.L.Lock()
	for len(sb.Buf)==0{
		sb.cond.Wait()
	}
	bytestream := sb.Buf[0]
	sb.Buf = sb.Buf[1:]
	sb.cond.L.Unlock()
	return bytestream
}

func (sb *SyncBuffer)Clear(){
	sb.lock.Lock()
	sb.Buf = []Bytestream{}
	sb.lock.Unlock()
}