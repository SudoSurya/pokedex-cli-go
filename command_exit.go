package main

import (
	"fmt"
	"os"
)

func callbackExit(cfg *config) error{
	fmt.Println("Exited From Terminal")
	os.Exit(0)
	return nil
}
