package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func main() {
	w, _ := fsnotify.NewWatcher()
	defer w.Close()

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
}
