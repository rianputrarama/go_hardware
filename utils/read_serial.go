package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tarm/serial"
)

func main() {
	// pass the port value as flag
	port := flag.String("port", "/dev/ttyUSB0", "The serial port the device is connected to.")
	flag.Parse()

	config := &serial.Config{
		Name:        *port,
		Baud:        9600,
		ReadTimeout: time.Second * 5,
		Size:        8,
		Parity:      serial.ParityNone,
		StopBits:    serial.Stop1,
	}

	stream, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}

	// close properly the stream when ctrl+c is hit
	closePort(stream)

	// scan the stream from the serial port
	scanner := bufio.NewScanner(stream)
	// read the stream and print it line by line
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func closePort(port *serial.Port) {
	c := make(chan os.Signal)
	// listen to any os interruption, SIGTERM signal and notify in the channel
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		// read signal from the channel
		<-c
		fmt.Println("\nCiao!")
		// close the port properly
		port.Close()
		// exit the program
		os.Exit(1)
	}()
}
