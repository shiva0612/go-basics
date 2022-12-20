package main

import (
	"log"
	"time"
)

func main() {

	<-time.After(2 * time.Second)

	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	timer.Stop() //prevents populating in the channel (goroutines reading the channel will wait eternal)
	timer.Reset(4 * time.Second)

	timer = time.AfterFunc(2*time.Second, func() {
		log.Println("im executed after the time")
	})
	timer.Stop() //not if I stopped :-(

	ticker := time.NewTicker(2 * time.Second)
	<-ticker.C
	ticker.Stop()
	ticker.Reset(4 * time.Second)

	t1 := time.Now()
	t2 := time.Now()
	t1.Before(t1)
	t1.After(t2)
	t1.Equal(t2)

	duration := t2.Sub(t1)
	duration.Seconds()
	duration.Minutes()
	duration.Hours()

	future_time := time.Now().Add(2 * time.Second)
	_ = future_time

	time.Now().Format(time.RFC3339)
	string_time := time.Now().Format("2006-01-02 15:04:05")

	timee, _ := time.Parse("2006-01-02 15:04:05", string_time)
	_ = timee

}
