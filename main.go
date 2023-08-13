package main

import (
	"time"

	"github.com/francisko-rezende/sensor-simulator/simulator"
)

func main() {

	ticker := time.NewTicker(time.Second)

	go func() {
		for range ticker.C {
			simulator.SensorsSimulator()
		}
	}()

	select {}
}
