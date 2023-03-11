package cron

import (
	"github.com/robfig/cron/v3"
	"log"
)

func StartScheduleLogging() {
	c := cron.New()

	_, err := c.AddFunc("@every 10m", func() {
		log.Printf("--- running scheduler ---")
	})

	if err != nil {
		log.Printf("--- scheduler error ---: %s", err.Error())
		return
	}

	c.Start()
}
