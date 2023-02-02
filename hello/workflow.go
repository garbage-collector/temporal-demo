package hello

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// Workflow is a Hello World workflow definition.
func Workflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumInterval: 10 * time.Second,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("HelloWorld workflow started", "name", name)

	var result string
	err := workflow.ExecuteActivity(ctx, ActivityHello, name).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, ActivityPing, name).Get(ctx, nil)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("HelloWorld workflow completed.", "result", result)

	return result, nil
}

func ActivityHello(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "Hello " + name + "!", nil
}

func ActivityPing(ctx context.Context, name string) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := client.Get("http://localhost:8080/ping")
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected API response with code: %d", response.StatusCode)
	}

	return nil
}
