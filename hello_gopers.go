//package main
//
//import (
//	"machine"
//	"time"
//)
//
//func main() {
//	machine.
//	// setup the LED as an output
//	led := machine.D13
//	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
//	ledSwitch := true
//	start := 0
//	for {
//		// turn the light on
//		led.Set(ledSwitch)
//		ledSwitch = !ledSwitch
//		println(start, " Hello Tech Makers!")
//		time.Sleep(2 * time.Second)
//		start += 1
//	}
//}
