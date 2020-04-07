package utils

import (
	"context"

	"github.com/google/uuid"
	"github.com/micro/go-micro/v2/metadata"
	"golang.org/x/net/trace"
)

// CreateTracing 设置链路跟踪信息
func CreateTracing(context context.Context, family, title string) (context.Context, trace.Trace) {
	tr := trace.New(family, title)

	ctx := trace.NewContext(context, tr)

	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = metadata.Metadata{}
	}

	traceID := uuid.New()

	tmd := metadata.Metadata{}
	for k, v := range md {
		tmd[k] = v
	}

	tmd["traceID"] = traceID.String()
	tmd["fromName"] = family
	ctx = metadata.NewContext(ctx, tmd)

	return ctx, tr
}
