package sender

import (
	. "CSMA/utils"
	"log"
	"math/rand"
	"sync"
)

//maximumm collision before a frame drops
const MAX_COLLISIONS = 12
const TIMESLOTS = 500 //at least 500 milliseconds

type Sender struct {
	buffer [] Bytestream
	S2C chan <- Bytestream
	collisions int
	NodeNo int
	srcMac int
	destMac int
}

func NewSenderNode(nodeno int, destMac int, chc2s chan <- Bytestream) *Sender{
	return &Sender{
		buffer: []Bytestream{},
		collisions: 0,
		S2C : chc2s,
		NodeNo: nodeno,
		destMac: destMac,
		srcMac: rand.Intn(1<<7),

	}
}

func (s * Sender)Init(MAXSIMULATETIME, Arrival_rate int, wg *sync.WaitGroup, method string){
	//notify about the sender thread completion
	defer wg.Done()
	//generate frames
	s.GenerateFrames(MAXSIMULATETIME, Arrival_rate)

	if method=="1P" {
		s.OnePersistent()
	}else if method =="NP" {
		s.NonPersistent()
	}else {
		s.PPersistent()
	}
	log.Println("Exiting sender", s.NodeNo)

}

func (s* Sender)GenerateFrames(MAXSIMULATETIME, ARRIVAL_RATE int){
	timeNow :=0.0
	for timeNow < float64(MAXSIMULATETIME) {
		//Time penalty for frame generation
		timeNow += rand.Float64() / float64(ARRIVAL_RATE)
		//appending the frames
		data := rand.Intn(1<<15)
		genFrame := MakeFrame(s.srcMac, s.destMac, data, s.NodeNo)
		s.buffer = append(s.buffer, genFrame)
	}
}