package imaging

import (
	"context"
	"encoding/xml"
	onvif "github.com/AminN77/customGoOnvif"
	imaging "github.com/AminN77/customGoOnvif/Imaging"
	"github.com/AminN77/customGoOnvif/sdk"
	"github.com/juju/errors"
)

// Call_GetStatus forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStatusResponse.
func Call_GetStatus(ctx context.Context, dev *onvif.Device, request imaging.GetStatus) (GetStatusResponse, error) {
	var reply Envelope3
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetStatusResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetStatus")
		return reply.Body.GetStatusResponse, errors.Annotate(err, "reply")
	}
}

type Envelope3 struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  string   `xml:"Header"`
	Body    Body2    `xml:"Body"`
}

type Body2 struct {
	XMLName           xml.Name          `xml:"Body"`
	Text              string            `xml:",chardata"`
	GetStatusResponse GetStatusResponse `xml:"GetStatusResponse"`
}

type GetStatusResponse struct {
	Text   string `xml:",chardata"`
	Status struct {
		Text          string `xml:",chardata"`
		FocusStatus20 struct {
			Text       string  `xml:",chardata"`
			Position   float64 `xml:"Position"`
			MoveStatus string  `xml:"MoveStatus"`
		} `xml:"FocusStatus20"`
	} `xml:"Status"`
}
