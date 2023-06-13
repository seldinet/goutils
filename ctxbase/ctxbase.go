package ctxbase

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

const ContextBaseName = "ContextBase"

type ContextBase struct {
	RequestID string
	ActionID  string
}

func FromCtx(ctx context.Context) *ContextBase {
	if ctx == nil {
		return nil
	}
	c, ok := ctx.Value(ContextBaseName).(*ContextBase)
	if !ok {
		return nil
	}
	return c
}

func NewID() string {
	u := uuid.NewV1()
	return u.String()
}
