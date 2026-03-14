package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// create a new Cron Masnsger
	c := cron.New()

	// Add a job that runs every minute
	c.AddFunc("* * * * *", func() {
		fmt.Println("Job 1: Job is running every minute", time.Now())
	})

	// Add a job that runs every day at 9:00 AM
	c.AddFunc("0 9 * * *", func() {
		fmt.Println("Job 2: Job is running everyday by 9am", time.Now())
	})

	c.Start()
	select {}
}
