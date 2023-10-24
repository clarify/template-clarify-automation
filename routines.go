package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/clarify/clarify-go/automation"
	"github.com/clarify/clarify-go/fields"
	"github.com/clarify/clarify-go/views"
)

// routines contain a tree-structure of named routines that can be matched
// by the CLI.
var routines = automation.Routines{
	"publish": automation.Routines{
		"example": publishExample,
	},
	"evaluate": automation.Routines{
		"high-temperature": detectHighTemperature,
	},
	"debug": automation.Routines{
		"test-slack-webhook": sendToSlack("Sorry for the disturbance; just testing that this woks."),
	},
}

// publishExample includes a routine that publish items from signals. To make it
// work for your own organization, you must at least replace the integration
// IDs.
var publishExample = automation.PublishSignals{
	// List of integration IDs to publish signals from using this routine.
	// You can find your integration ID(s) in the admin panel.
	Integrations: []string{"INSERT_INTEGRATION_ID"},
	// TransformVersion should be updated when altering transforms to correctly
	// detect if already published items are up-to-date.
	TransformVersion: "v0",
	// List of transforms to apply.
	Transforms: []func(item *views.ItemSave){
		// Transform functions can be inlined. This could be useful if you
		// want to create a transform that's not reusable by other
		// configurations.
		func(item *views.ItemSave) {
			// Detect names of format "device-id/measurement-name/eng-unit".
			// Note that production code may want to use regular expressions
			// to specify more precise rules.
			comps := strings.Split(item.Name, "/")
			if len(comps) == 3 {
				item.Name = comps[1]
				item.EngUnit = comps[2]
				item.Labels.Add("device-id", comps[0])
			}
		},
		// You can also refer to transforms you have written elsewhere. The
		// following transform functions are defined in the
		// publish_transforms.go file.
		prettifyEngUnit,
		addISQLabels,
	},
}

// detectHighTemperature includes a routine that sends a slack message when the
// inside temperature goes above 26 degrees. To make it work for you, you must
// at least replace the item ID.
var detectHighTemperature = automation.EvaluateActions{
	Evaluation: automation.Evaluation{
		Items: []fields.ItemAggregation{
			{Alias: "t0", ID: "INSERT_ITEM_ID", Aggregation: fields.AggregateAvg, Lag: 1},
			{Alias: "t1", ID: "INSERT_ITEM_ID", Aggregation: fields.AggregateAvg},
		},
		Calculations: []fields.Calculation{
			{Alias: "trigger", Formula: "gapfill(t0) < 26 && gapfill(t1) >=26"}, // Returns 1.0 when true.
		},
	},
	RollupBucket: fields.FixedCalendarDuration(15 * time.Minute),
	Actions: []automation.ActionFunc{
		automation.ActionSeriesContains("trigger", 1),
		automation.ActionRoutine(sendToSlack("Average inside temperature above 26°C")),
	},
}

// sendToSlack contains a simple example of how to send a message to slack
// using incoming webhooks. See https://api.slack.com/messaging/webhooks for how
// to set up an application. Once configured, enable incoming webhooks set the
// CLARIFY_SLACK_WEBHOOK_URL environment variable for the action to work.
func sendToSlack(msg string) automation.RoutineFunc {
	url := os.Getenv("CLARIFY_SLACK_WEBHOOK_URL")
	if url == "" {
		// When URL is not set, return a place-holder routine that will always
		// fail.
		return func(ctx context.Context, cfg *automation.Config) error {
			return fmt.Errorf("CLARIFY_SLACK_WEBHOOK_URL not set")
		}
	}

	// Return slack web-hook routine.
	return func(ctx context.Context, cfg *automation.Config) error {
		var body bytes.Buffer
		enc := json.NewEncoder(&body)

		type slackMessage struct {
			Text string `json:"text"`
		}
		if err := enc.Encode(slackMessage{Text: msg}); err != nil {
			return err
		}

		resp, err := http.Post(url, "application/json", &body)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status code: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		}
		return nil
	}
}
