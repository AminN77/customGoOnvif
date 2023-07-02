// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package imaging

import (
"context"
imaging "github.com/AminN77/customGoOnvif/Imaging"
"github.com/juju/errors"
"github.com/AminN77/customGoOnvif"
)

// Call_MoveFocusAbs forwards the call to dev.CallMethod().
func Call_MoveFocusAbs(ctx context.Context, dev *onvif.Device, request imaging.MoveFocusAbs) error {
	if _, err := dev.CallMethod(request); err != nil {
		return errors.Annotate(err, "call")
	}

	return nil
}
