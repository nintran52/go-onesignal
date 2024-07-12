package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	v1 "github.com/nintran52/onesignal-v2/v1"
	v2 "github.com/nintran52/onesignal-v2/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	appId := os.Getenv("APP_ID")
	appKey := os.Getenv("APP_KEY")
	userID := os.Getenv("USER_ID")
	subscriptionID := os.Getenv("SUBSCRIPTION_ID")

	v1.SenderV1(appId, subscriptionID)
	v2.CreateUser(appId, appKey, userID)
	v2.CreateSubscription(appId, appKey, userID)
	v2.SenderV2(appId, appKey, userID)
}
