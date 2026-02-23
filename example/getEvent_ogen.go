package main

import (
	"context"
	"os"

	api "github.com/fingerprintjs/go-sdk/v8/sdk"
	"github.com/joho/godotenv"
)

type BearerTokenAuth struct {
	token string
}

func NewBearerTokenAuth(token string) *BearerTokenAuth {
	return &BearerTokenAuth{token: token}
}

func (b *BearerTokenAuth) BearerAuth(ctx context.Context, operationName api.OperationName) (api.BearerAuth, error) {
	return api.BearerAuth{
		Token: b.token,
	}, nil
}

func main() {
	// Load environment variables
	godotenv.Load()

	client, err := api.NewClient("https://api.fpjs.io/v4", NewBearerTokenAuth(os.Getenv("FINGERPRINT_API_KEY")))
	event, err := client.GetEvent(context.Background(), api.GetEventParams{
		EventID:   os.Getenv("EVENT_ID"),
		RulesetID: api.NewOptString(os.Getenv("RULESET_ID")),
	})

	if err != nil {
		panic(err)
	}

	switch p := event.(type) {
	case *api.Event:
		json, err := p.MarshalJSON()
		if err == nil {
			print(string(json))
		}
		ruleAction, ok := p.RuleAction.Get()
		if ok {
			if ruleAction.OneOf.IsEventRuleActionBlock() {
				// Handle block
			} else if ruleAction.OneOf.IsEventRuleActionAllow() {
				// Handle allow
			}
		}

	case *api.GetEventNotFound:
		println(p.Error.Message)
	}
}
