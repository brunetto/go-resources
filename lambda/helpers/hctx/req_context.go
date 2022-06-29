package hctx

import (
	"context"
	"time"

	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/rs/zerolog"

	"github.com/gofrs/uuid"
	"github.com/rs/xid"
)

func CtxLogger(ctx context.Context, lg zerolog.Logger) zerolog.Logger {
	lc, ok := core.GetRuntimeContextFromContext(ctx)
	if ok {
		return lg.With().Str("req_id", lc.AwsRequestID).Logger()
	}

	return lg.With().Str("req_id", xid.New().String()).Logger()
}

// If you want uuidv4 instead of xid
// but see https://encore.dev/blog/go-1.18-generic-identifiers
func id() string {
	id, err := uuid.NewV4()
	if err != nil {
		return "missing-req-id-" + time.Now().Format(time.RFC3339Nano)
	}

	return id.String()
}
