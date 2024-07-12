package v2

import (
	"context"
	"fmt"
	"log"

	"github.com/OneSignal/onesignal-go-api/v2"
)

func CreateSubscription(appID, appKey, userID string) {
	aliasLabel := "external_id"
	createSubscriptionRequestBody := *onesignal.NewCreateSubscriptionRequestBody()
	subscription := onesignal.NewSubscriptionObject()
	subscription.SetType("iOSPush")
	createSubscriptionRequestBody.SetSubscription(*subscription)

	apiClient := onesignal.NewAPIClient(onesignal.NewConfiguration())

	appAuth := context.WithValue(context.Background(), onesignal.AppAuth, appKey)

	response, _, err := apiClient.DefaultApi.CreateSubscription(appAuth, appID, aliasLabel, userID).CreateSubscriptionRequestBody(createSubscriptionRequestBody).Execute()

	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	// Print the response
	fmt.Printf("User created successfully: %+v\n", response)
}
