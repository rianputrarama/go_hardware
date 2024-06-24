package main

import (
	"machine"
	"time"
	"tinygo.org/x/drivers/dht"
)

func main() {
	dataPin := machine.D12
	dataPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// Inisialisasi sensor DHT11
	dhtSensor := dht.New(dataPin, dht.DHT11)
	for {
		temp, humidity, _ := dhtSensor.Measurements()
		println("Temperature:", temp, "°C")
		println("Humidity:", humidity, "%")
		time.Sleep(2 * time.Second)
	}
}
