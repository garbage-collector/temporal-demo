package emancipation

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// Workflow is a workflow definition for the emancipation of Miley Cyrus.
func Workflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumInterval: 10 * time.Second,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Emancipation workflow started", "name", name)

	var result string
	err := workflow.ExecuteActivity(ctx, BuyMyselfFlowers, 5).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, WriteMyNameInTheSand, name).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, TalkToMyselfForHours).Get(ctx, nil)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, SayThingsYouDontUnderstand, "gloubi-boulga").Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, TakeMyselfDancing).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	err = workflow.ExecuteActivity(ctx, HoldMyOwnHand).Get(ctx, nil)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("Emancipation workflow completed.", "result", result)

	return result, nil
}
