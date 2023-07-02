package imaging

import (
	"context"
	onvif "github.com/AminN77/customGoOnvif"
	imaging "github.com/AminN77/customGoOnvif/Imaging"
	"github.com/juju/errors"
)

// Call_SetSharpnessImagingSettings forwards the call to dev.CallMethod().
func Call_SetSharpnessImagingSettings(ctx context.Context, dev *onvif.Device, request imaging.SetSharpnessImagingSettings) error {
	if _, err := dev.CallMethod(request); err != nil {
		return errors.Annotate(err, "call")
	}

	return nil
}
