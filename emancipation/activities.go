package emancipation

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.temporal.io/sdk/activity"
)

func BuyMyselfFlowers(ctx context.Context, quantity int) (string, error) {
	activity.GetLogger(ctx).Info("Welcome to the florist")
	return fmt.Sprintf("Thanks for buying %d orchids", quantity), nil
}

func WriteMyNameInTheSand(ctx context.Context, name string) (string, error) {
	activity.GetLogger(ctx).Info("Welcome to Bray-Dunes")
	return fmt.Sprintf("%s is written in the sand", name), nil
}

func TalkToMyselfForHours(ctx context.Context) error {
	activity.GetLogger(ctx).Info(fmt.Sprintf("It's %s, let's have a talk", time.Now().Format("03:04PM")))

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := client.Get("http://localhost:8080/talk")
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected API response with code: %d", response.StatusCode)
	}

	return nil
}

func SayThingsYouDontUnderstand(ctx context.Context, things string) (string, error) {
	activity.GetLogger(ctx).Info("What are you saying?")
	return fmt.Sprintf("I'm sorry, I don't understand this: %s", things), nil
}

func TakeMyselfDancing(ctx context.Context) (string, error) {
	activity.GetLogger(ctx).Info("Welcome to the ball")
	return fmt.Sprintf("What a dance!"), nil
}

func HoldMyOwnHand(ctx context.Context) error {
	activity.GetLogger(ctx).Info("Let's shake your hand!")
	return nil
}
