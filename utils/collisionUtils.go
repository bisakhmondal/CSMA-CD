package utils


const (
	COLLIDED = 0
	NCOLLIDED = 1
)

var addcountC = make(chan int) // set current collision status
var counterC = make(chan int) //  get current collsion track

func SetNotCollided() { addcountC <- NCOLLIDED }
func SetCollided() { addcountC <- COLLIDED }

func CollisionStatus() int { return <-counterC }


func TellerCollision() {
	var curStatus int = NCOLLIDED // Status of transmision
	for {
		select {
		case curStatus = <-addcountC:
		case counterC <- curStatus:
		}
	}
}

