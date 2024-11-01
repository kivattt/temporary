package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func main() {
	w, _ := fsnotify.NewWatcher()
	defer w.Close()

	fmt.Println("Adding .")
	w.Add(".")
	fmt.Println("Removing .")
	w.Remove(".")
	fmt.Println("Done!")
}
