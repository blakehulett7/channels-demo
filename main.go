package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	lines := []string{"Jesus", "is", "my", "Lord", "and", "Savior!"}

	simple_speech(lines)

	var waitgroup sync.WaitGroup
	for _, line := range lines {
		waitgroup.Add(1)
		go func() {
			fmt.Println(line)
			time.Sleep(time.Second)
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
}

func simple_speech(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
		time.Sleep(time.Second)
	}
}

func channel_speech(lines []string, channel chan string) {
	for _, line := range lines {
		channel <- line
		time.Sleep(time.Second * 1)
	}
	close(channel)
}
