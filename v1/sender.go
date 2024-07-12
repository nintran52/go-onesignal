package v1

import (
	"context"
	"fmt"
	"log"

	"github.com/OneSignal/onesignal-go-api/v2"
)

func SenderV1(appID, subscriptionID string) {
	// Initialize the OneSignal client
	apiClient := onesignal.NewAPIClient(onesignal.NewConfiguration())

	// Create a new notification
	notification := onesignal.NewNotification(appID)

	title := "Test Notification"
	message := "Hello, this is a test notification!"

	notification.SetHeadings(onesignal.StringMap{En: &title})
	notification.SetContents(onesignal.StringMap{En: &message})

	ids := []string{subscriptionID}
	notification.SetIncludePlayerIds(ids)

	// Send the notification
	// appAuth := context.WithValue(context.Background(), onesignal.AppAuth, appKey)
	response, _, err := apiClient.DefaultApi.CreateNotification(context.Background()).Notification(*notification).Execute()
	if err != nil {
		log.Fatalf("Error sending notification: %v", err)
	}

	// Print the response
	fmt.Printf("Notification sent successfully: %+v\n", response)
}
