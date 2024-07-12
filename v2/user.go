package v2

import (
	"context"
	"fmt"
	"log"

	"github.com/OneSignal/onesignal-go-api/v2"
)

func CreateUser(appID, appKey, userID string) {
	// Initialize the OneSignal client
	apiClient := onesignal.NewAPIClient(onesignal.NewConfiguration())

	user := onesignal.NewUser()

	// Define the user identity and properties
	userIdentity := map[string]interface{}{
		"external_id": userID,
	}
	// Define the user subscription
	subscriptions := []onesignal.SubscriptionObject{}
	subscription := onesignal.NewSubscriptionObject()
	subscription.SetType("iOSPush")
	subscriptions = append(subscriptions, *subscription)

	user.SetIdentity(userIdentity)
	user.SetSubscriptions(subscriptions)

	appAuth := context.WithValue(context.Background(), onesignal.AppAuth, appKey)
	response, _, err := apiClient.DefaultApi.CreateUser(appAuth, appID).User(*user).Execute()
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	// Print the response
	fmt.Printf("User created successfully: %+v\n", response)
}
