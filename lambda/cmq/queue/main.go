package main

import (
	"context"
	"fmt"
	"go-resources/lambda/example"
	"go-resources/lambda/helpers/queueh"
	"go-resources/lambda/httph"
	"os"
	"runtime/debug"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func main() {
	lg := zerolog.New(os.Stdout)
	defer func() { lg.Info().Msg("Shutting down lambda instance") }()

	lg.Info().Msg("Starting lambda instance")

	//
	// your setup here: aws session, connections etc
	//

	// ⬆️ Everything above here is called once per instance, not once per execution (invocation).
	// This means we don't waste time initializing things every time.
	// ⬇️ Everything below here is called for each execution (invocation).
	lambda.Start(NewLambdaHandlerSQS( /* your depts here */ lg))
}

type LambdaHandlerSQS func(context.Context, events.SQSEvent) error

// nolint: funlen
func NewLambdaHandlerSQS( /* your depts here */ lg zerolog.Logger) LambdaHandlerSQS {
	return func(ctx context.Context, evts events.SQSEvent) error {
		// init logging, recovering, timing
		lg = httph.CtxLogger(ctx, lg)
		t0 := time.Now()

		defer func() {
			lg.Info(). // TBD: handle error case
					Float64("duration_ms", float64(time.Since(t0).Milliseconds())).
					Msg(fmt.Sprintf("done in %v", time.Since(t0)))
		}()

		defer func() { // recover panic and log them properly, just in case
			if r := recover(); r != nil {
				lg.Fatal().Fields(map[string]interface{}{
					"trace": debug.Stack(),
				}).Msg(fmt.Sprintf("recovered paninc: %v", r))
			}
		}()

		lg.Info().Int("rcv_evts_count", len(evts.Records)).
			Msg("new message")

		// Shoud we get only 1 message at a time?
		// 		PRO: easier error handling since an error would reject the whole batch,
		// 			event the already processed messages
		// 		CONS: single message execution means more executions, less performance
		// 		DECISION: for the moment, multiple messages since we are idempotent:
		// 					worst case we save the message again with the same primary keys (overwrite)
		// 					just changing the store timestamp
		// Should we parallelize the insert?
		// 		PRO: faster and cheaper (shorter execution duration)
		// 		CONS: less readable code
		//		DECISION: for the moment, serial code, multiple lambdas, in the future we'll re-evaluate the strategy
		// Could be nice to create a foreach here?
		// 		PRO: nicer, clearer
		// 		CONS: more difficult to understand and change
		// 		DECISION: seems we would be trading simplicity for aestetic, so for now stays this way.
		for i, msg := range evts.Records {
			// I may want, in the future, check the event type.
			var in example.In

			err := queueh.Unwrap(msg.Body, &in)
			if err != nil {
				return AndLog(lg, errors.Wrap(err, "can't extract original evento from SQS message"))
			}

			_, err = example.DoStuff(ctx, in)
			if err != nil {
				err = errors.Wrap(err, "can't collect event")
				lg.Error().Int("prev_msgs_ok", i). // how many messages were correctly processed before
									Fields(map[string]interface{}{"msg_failed": in}).Msg(err.Error())

				return err
			}
		}

		return nil
	}
}

func AndLog(lg zerolog.Logger, err error) error {
	lg.Error().Msg(err.Error())

	return err
}
