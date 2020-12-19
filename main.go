package main

import (
	ch "CSMA/channel"
	rcv "CSMA/receiver"
	snd "CSMA/sender"
	"CSMA/utils"
	"log"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func init(){
	go utils.TellerMedium() //initiate the authority of Medium status control i.e IDLE or BUSY
	go utils.TellerCollision() //initiate the authority of Collision status control i.e IDLE or BUSY

	//initiate the random seed
	rand.Seed(time.Now().UnixNano())
}

func cleanup() {
	log.Println("Cleaning Up...")
	time.Sleep(1*time.Second)
}

const (
	inf = 1<<16
	MAXSIMULATETIME= 10 //simulate for this number of second Seconds

	NUM_NODES = 1 //Number of Sender Node

	PERSISTENCE_STRATAGY = "PP" //Choose Betweeen "1P", "NP", "PP"
	PACKET_ARRIVAL_RATE=7//Number of packets arriving per Second

)
var wg sync.WaitGroup

func main(){
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	//setting up interthread communication
	S2C := make(chan string, inf) //communication between sender and channel has infinite buffer capacity
	C2R := make(chan string)


	//Evaluation Metrics counter (Throughput and Efficiency)
	statsMonitor := utils.NewStatusMonitor()


	//Initializatoin of Channel and Single receiver
	channel := ch.NewChannel(S2C, C2R, statsMonitor)
	receiverMac := rand.Intn(1 << 8)
	receiver := rcv.NewReceiver(C2R)


	wg.Add(NUM_NODES)

	//Running Channel, Receiver in different goroutines
	go channel.Init()
	go receiver.Init()

	//initiating the Time counter
	start := time.Now()

	for nodeNo := 0; nodeNo < NUM_NODES; nodeNo++ {

		go func(nodeNo, MAXSIMULATETIME, receiverMac int, group *sync.WaitGroup, PERSISTENCE_STRATAGY string) {
			//each node In a concurrent thread

			sender := snd.NewSenderNode(nodeNo+1, receiverMac, S2C)
			sender.Init(MAXSIMULATETIME, PACKET_ARRIVAL_RATE, group, PERSISTENCE_STRATAGY)

		}(nodeNo, MAXSIMULATETIME, receiverMac, &wg, PERSISTENCE_STRATAGY)
	}

	wg.Wait()
	//Upon Transmission close the channel
	close(S2C)

	timeTakeninDuration := time.Since(start)
	timeTaken := float64(timeTakeninDuration / time.Microsecond)

	//Printing Metrics

	//print(NUM_NODES,",",timeTaken,",")

	log.Println("time Taken: ", timeTaken*math.Pow(10, -6), " Sec")

	statsMonitor.Stats(timeTaken)
	print("\n")
}