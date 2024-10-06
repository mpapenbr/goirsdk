package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mpapenbr/goirsdk/irsdk"
)

func main() {
	// setup signal handler to catch Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	defer func() {
		signal.Stop(sigChan)
	}()

	go mainLoop()
	fmt.Println("Use Ctrl+C to exit")
	select {
	case <-sigChan:
	}
	fmt.Println("Exiting...")
}

func mainLoop() {
	client := http.Client{Timeout: time.Second * 10}
	for {

		// loop until iRacing is running
		for {
			fmt.Println("Waiting for connection to iRacing")
			if simRunning, err := irsdk.IsSimRunning(context.Background(), &client); err != nil {
				// exit with panic here if we can't connect to iRacing
				panic(err)
			} else if simRunning {
				fmt.Println("iRacing Simulation is running")
				break
			}
			time.Sleep(time.Second * 2)
		}

		api := irsdk.NewIrsdk()
		for {
			// wait for iRacing to signal that new data is available
			if api.WaitForValidData() {
				// retrieve the telemetry data
				api.GetData()
				// basic example of how to retrieve a value, here: the
				if sessionTime, err := api.GetDoubleValue("SessionTime"); err == nil {
					fmt.Printf("Session time: %f\n", sessionTime)
				} else {
					fmt.Println("Error getting session time", err)
				}
			}
			// sleep until we check again for new data
			time.Sleep(time.Second)
		}
	}
}
