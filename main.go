package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	params := &api.CreateMessageParams{}
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetTo(os.Getenv("RECIPIENT_PHONE_NUMBER"))
	params.SetBody("Hi there")
	params.SetMessagingServiceSid(os.Getenv("TWILIO_MESSAGING_SERVICE_SID"))

	mediaItems := []string{
		"https://cdn.freebiesupply.com/logos/large/2x/twilio-logo-png-transparent.png",
	}
	params.SetMediaUrl(mediaItems[:])

	start := time.Now()
	duration, err := time.ParseDuration("17m")
	if err != nil {
		log.Fatalf("Could not determine the time to schedule the message. Reason: %v", err)
	}
	params.SetSendAt(start.Add(duration))
	params.SetScheduleType("fixed")

	client := twilio.NewRestClient()
	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Fatalf("Could not schedule the message. Reason: %v", err)
	}

	if resp.Sid != nil {
		fmt.Printf("%s,%s\n", *resp.Sid, *resp.Status)
	} else {
		fmt.Printf("%s,%s\n", resp.Sid, resp.Status)
	}
}
