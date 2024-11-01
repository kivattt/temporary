package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	w, _ := fsnotify.NewWatcher()

	go func() {
		defer fmt.Println("Channel handler quit!")
		for {
			select {
			case _, ok := <-w.Events:
				if !ok {
					return
				}
			case _, ok := <-w.Errors:
				if !ok {
					fmt.Println("error channel closed!")
					return
				}
			}
		}
	}()

	fmt.Println("Adding .")
	w.Add(".")
	fmt.Println("Removing .")
	w.Remove(".")
	fmt.Println("Done!")

	w.Add(".")
	w.Remove(".")

	w.Close()

	time.Sleep(800 * time.Millisecond)
}
