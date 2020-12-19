package sender

import (
	. "CSMA/utils"
	"math"
	"math/rand"
	"time"
)

func (s * Sender)NonPersistent(){
	curFrame := s.buffer[0]

	for len(s.buffer)!=0 {

		if SenseMedium() ==IDLE {
			//marking the medium busy to gain access
			SetMediumBusy()
			//sending the frame
			s.S2C <- curFrame
			//wait for collision signal during propagation
			time.Sleep(500*time.Microsecond)

			if CollisionStatus() == COLLIDED{
				//if collided
				s.collisions++

				if s.collisions >= MAX_COLLISIONS {
					//drop the frame
					s.buffer = s.buffer[1:]
					//reset collision count
					s.collisions=0
					if len(s.buffer) !=0{
						curFrame = s.buffer[0]
					}
					continue //no need to wait further start afresh
				}

				//wait for timeslots depending the back-off algorithm
				k := rand.Intn(int(math.Pow(2, float64(s.collisions))))

				time.Sleep(time.Duration(k*TIMESLOTS)* time.Microsecond)

			}else{
				//transmission successful
				s.buffer = s.buffer[1:]
				//reset collision count
				s.collisions=0
				if len(s.buffer) !=0{
					curFrame = s.buffer[0]
				}
			}

		} else{
			//Wait Non Aggressive algorithm wait for random amount of time
			randn := rand.Intn(5) //within 2.5 second
			time.Sleep(time.Duration(randn* TIMESLOTS)* time.Microsecond)
		}
	}
}
