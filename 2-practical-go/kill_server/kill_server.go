package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
)

func main() {
	// when using GoLand: kill_server/server.pid
	// fmt.Println(KillServer("server.pid"))
	err := KillServer("server.pid")
	if err != nil {
		fmt.Println("ERROR: ", err)
		if errors.Is(err, fs.ErrExist) {
			fmt.Println("not found")
		}
		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Printf("> %s\n", e)
		}
	}
}
func KillServer(pipFile string) error {
	
	file, err := os.Open(pipFile)
	if err != nil {

		return fmt.Errorf("%q - error opening file %s", pipFile, err)
	}
	/*
		- defer happens when function exits, no matter what ( even when a panic 😱 happens)
		- defer works at the function  level ( don't us in a for loop ➰)
		- defer works in reverse other (kind of like a Stack 📚 or think LIFO 🥞)
		- Idiomatic use: try to to first acquire a resource, check for error, defer and release
	*/
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("%q - bad pid format: %w", pipFile, err)
	}

	slog.Info("Killing", "pid", pid)
// optional deleting the file
	if err := os.Remove(pipFile); err != nil {
		slog.Warn("delete", "file", pipFile, "error", err)
	}

	return nil
}
