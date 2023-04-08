// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"encoding/xml"
	"github.com/juju/errors"
	"github.com/AminN77/customGoOnvif"
	"github.com/AminN77/customGoOnvif/device"
	"io"
)

// Call_GetSystemDateAndTimeNew forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemDateAndTimeResponse.
func Call_GetSystemDateAndTimeNew(ctx context.Context, dev *onvif.Device, request device.GetSystemDateAndTime) (device.GetSystemDateAndTimeResponseNew, error) {
	type Envelope struct {
		XMLName xml.Name `xml:"Envelope"`
		//Text    string   `xml:",chardata"`
		//Sc      string   `xml:"sc,attr"`
		//S       string   `xml:"s,attr"`
		//Tt      string   `xml:"tt,attr"`
		//Tds     string   `xml:"tds,attr"`
		Header  string   `xml:"Header"`
		Body device.GetSystemDateAndTimeResponseNew `xml:"Body"`
	}

	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body, errors.Annotate(err, "call")
	} else {
		bo := &Envelope{}
		b, err := io.ReadAll(httpReply.Body)
		err = xml.Unmarshal(b, bo)
		return bo.Body, errors.Annotate(err, "reply")
	}
}


