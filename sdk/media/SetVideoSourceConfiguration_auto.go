// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/AminN77/customGoOnvif"
	"github.com/AminN77/customGoOnvif/sdk"
	"github.com/AminN77/customGoOnvif/media"
)

// Call_SetVideoSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetVideoSourceConfigurationResponse.
func Call_SetVideoSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.SetVideoSourceConfiguration) (media.SetVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoSourceConfigurationResponse media.SetVideoSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetVideoSourceConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetVideoSourceConfiguration")
		return reply.Body.SetVideoSourceConfigurationResponse, errors.Annotate(err, "reply")
	}
}