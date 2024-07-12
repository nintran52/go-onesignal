package v2

import (
	"context"
	"fmt"
	"log"

	"github.com/OneSignal/onesignal-go-api/v2"
)

func SenderV2(appID, appKey, userID, androidChannelID string) {
	// Initialize the OneSignal client
	apiClient := onesignal.NewAPIClient(onesignal.NewConfiguration())

	// Create a new notification
	notification := onesignal.NewNotification(appID)

	title := "Test Notification"
	message := "Hello, this is a test notification!"

	notification.SetHeadings(onesignal.StringMap{En: &title})
	notification.SetContents(onesignal.StringMap{En: &message})

	notification.SetTargetChannel("push")
	notification.SetAndroidChannelId(androidChannelID)
	notification.AdditionalProperties = map[string]interface{}{"include_aliases": map[string]interface{}{"external_id": []string{userID}}}

	// Send the notification
	appAuth := context.WithValue(context.Background(), onesignal.AppAuth, appKey)
	response, _, err := apiClient.DefaultApi.CreateNotification(appAuth).Notification(*notification).Execute()
	if err != nil {
		log.Fatalf("Error sending notification: %v", err)
	}

	// Print the response
	fmt.Printf("Notification sent successfully: %+v\n", response)
}
