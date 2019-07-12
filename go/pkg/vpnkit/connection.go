package vpnkit

import (
	"context"
	"io"
)

// Client exposes and unexposes ports on vpnkit.
type Client interface {
	Expose(context.Context, *Port) error
	Unexpose(context.Context, *Port) error
	ListExposed(context.Context) ([]Port, error)
	DumpState(context.Context, io.Writer) error
}
