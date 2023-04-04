package ptz

import (
	"context"
	"github.com/AminN77/customGoOnvif"
	"github.com/AminN77/customGoOnvif/ptz"
	"github.com/AminN77/customGoOnvif/sdk"
	"github.com/juju/errors"
)

// Call_AbsoluteMoveZoom forwards the call to dev.CallMethod() then parses the payload of the reply as a AbsoluteMoveResponse.
func Call_AbsoluteMoveZoom(ctx context.Context, dev *onvif.Device, request ptz.AbsoluteMoveZoom) (ptz.AbsoluteMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AbsoluteMoveResponse ptz.AbsoluteMoveResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AbsoluteMoveResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "AbsoluteMove")
		return reply.Body.AbsoluteMoveResponse, errors.Annotate(err, "reply")
	}
}
