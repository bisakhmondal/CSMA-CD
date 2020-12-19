package utils

const (
	BUSY = 1
	IDLE = 0
)
var addcount = make(chan int) // set current medium status
var counter = make(chan int) //  get current medium status

func SetMediumBusy() { addcount <- BUSY }
func SetMediumIdle() { addcount <- IDLE }

func SenseMedium() int { return <-counter }


func TellerMedium() {
	var curStatus int = IDLE // Status of medium is confined to only Teller goroutine
	for {
		select {
		case curStatus = <-addcount:
		case counter <- curStatus:
		}
	}
}
