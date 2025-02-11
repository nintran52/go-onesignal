package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	androidChannelID := os.Getenv("ANDROID_CHANNEL_ID")

	// v1.SenderV1(appId, subscriptionID)
	// v2.CreateSubscription(appId, appKey, userID)
	// v2.CreateUser(appId, appKey, userID, os.Getenv("SUBSCRIPTION_ID"))
	v2.SenderV2(appId, appKey, userID, androidChannelID)
}
