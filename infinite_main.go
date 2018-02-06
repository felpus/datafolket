package main

import (
	"./infinite"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c1 := make(chan os.Signal, 1)
	signal.Notify(c1)

	exit_chan := make(chan int)
	go func() {
		for {
			s := <- c1
			switch s {
			// taskkill -SIGHUP XXXX
			case syscall.SIGHUP:
				fmt.Println("Hungup")
				os.Exit(0)

				// taskkill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				fmt.Println("Interrupt")
				os.Exit(0)

				// taskkill -SIGTERM XXXX
			case syscall.SIGTERM:
				fmt.Println("Force stop")
				exit_chan <- 0
				os.Exit(0)

				// taskkill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				fmt.Println("Stop and core dump")
				exit_chan <- 0
				os.Exit(0)

			default:
				fmt.Println("Unknown signal.")
				exit_chan <- 1
				os.Exit(0)
			}
		}

	}()
	Loop()
}

func Loop() {
	infinite.Infinite01("Look, I can count forever:")
}
