package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"notification-parser/app"
	"sync"
	"time"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	a := app.App{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		_, err := s.Every(5).Seconds().Do(a.Start)
		if err != nil {
			fmt.Errorf("%s error", err)
		}
		s.StartBlocking()
		defer wg.Done()
	}()
	wg.Wait()
}
