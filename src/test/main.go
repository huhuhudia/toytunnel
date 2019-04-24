package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main()  {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan)
	for sig := range sigChan{
		fmt.Println()
		fmt.Println(sig)
	}
}
