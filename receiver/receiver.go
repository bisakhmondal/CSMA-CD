package receiver

import (
	. "CSMA/utils"
	"log"
)

type Receiver struct{
	C2R chan Bytestream
}

func NewReceiver(c chan Bytestream) *Receiver{
	return &Receiver{
		C2R: c,
	}
}

func (r * Receiver)Init(){
	for {
		bytestream := <- r.C2R
		//if bytestream ==""{
		//	break
		//}
		log.Println("Frame Received: Node", GetNodeInfo(bytestream), " | Data: ", GetData(bytestream))
	}
}
