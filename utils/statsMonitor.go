package utils

import (
	"log"
	"math"
)

type StatsMonitor struct {
	TotalTransmittedPacket, SuccessfulTransmittedPacket int

}

func NewStatusMonitor() *StatsMonitor{
	return &StatsMonitor{0,0}
}

func (s * StatsMonitor)Success(){
	s.SuccessfulTransmittedPacket++
	s.TotalTransmittedPacket++
}
func (s *StatsMonitor)Transmitted(cnt int){
	s.TotalTransmittedPacket+=cnt
}

func (s *StatsMonitor)Stats(timetaken float64){
	//timetaken in microseconds
	log.Println("Efficiency: ",  float64(s.SuccessfulTransmittedPacket)/
									float64(s.TotalTransmittedPacket))
	//bits transferred per second
	throughput := float64(s.SuccessfulTransmittedPacket * FRAMELENGTH) / timetaken //in bits per micro second
	throughput = throughput * math.Pow(10, -6) * math.Pow(10, 6)
	//print(float64(s.SuccessfulTransmittedPacket)/
	//								float64(s.TotalTransmittedPacket), ",",throughput)
	log.Println("Throughput: ", throughput," Mbps")
}
