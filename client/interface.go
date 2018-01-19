package client

import (
	"golang.org/x/net/context"
)
type TestInfoClient interface {
	TestInfo(ctx context.Context) (error)
}
