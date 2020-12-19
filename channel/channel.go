package channel

import (
	. "CSMA/utils"
	"log"
	"sync"
	"time"
)


type Channel struct {
	C2S chan string
	C2R chan string
	Buffer * SyncBuffer
	monitor * StatsMonitor
}

func NewChannel(chc2s, chc2r chan Bytestream, Monitor *StatsMonitor) * Channel{
	return &Channel{
		C2S: chc2s,
		C2R: chc2r,
		Buffer: NewSyncBuffer(),
		monitor: Monitor,
	}
}

func (c * Channel)Init(){
	log.Println("Channel Initialized")
	var wgi sync.WaitGroup
	wgi.Add(1)
	//go Teller()
	go c.send2receiver(&wgi)
	c.receivefromSender()

}

func (c* Channel)receivefromSender(){
	for {
		//make the collision flag unset
		SetNotCollided()

		if SenseMedium() == BUSY{
			//if more than one frames in the pipe then collision

			if len(c.C2S) > 1 {
				//collision
				SetCollided()

				//log into monitor for throughput and efficiency
				c.monitor.Transmitted(len(c.C2S))

				log.Println("Collision")
				//free the pipe, data is unusable
				for len(c.C2S) > 0 {
					<-c.C2S
				}

			}else{
				//there is a single frame inside the channel
				//putting it into buffer
				bitstream := <-c.C2S
				c.Buffer.Push(bitstream)

				//log into monitor for throughput and efficiency
				c.monitor.Success()
			}

			//the data needs to be captured from medium which will incur some time penalty
			time.Sleep(800* time.Microsecond)

			//making the medium idle again
			SetMediumIdle()
		}

	}
}

func (c *Channel)send2receiver(wg  *sync.WaitGroup){
	defer wg.Done()
	for {

		if c.Buffer.Len()==0{
			time.Sleep(500*time.Microsecond)
			continue
		}

		bytestream := c.Buffer.Pop()
		c.C2R <- bytestream
	}
	log.Println("exited from s2r")

}