// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package imaging

import (
	"context"
	"encoding/xml"
	imaging "github.com/AminN77/customGoOnvif/Imaging"
	"github.com/juju/errors"
	"github.com/AminN77/customGoOnvif"
	"github.com/AminN77/customGoOnvif/sdk"
)

// Call_GetImagingSettings forwards the call to dev.CallMethod() then parses the payload of the reply as a GetImagingSettingsResponse.
func Call_GetImagingSettings(ctx context.Context, dev *onvif.Device, request imaging.GetImagingSettingsRequest) (CustomImagingSettings, error) {
	var reply Envelope2
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetImagingSettingsResponse.ImagingSettings, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetImagingSettings")
		return reply.Body.GetImagingSettingsResponse.ImagingSettings, errors.Annotate(err, "reply")
	}
}

type Envelope2 struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  string   `xml:"Header"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	XMLName                    xml.Name `xml:"Body"`
	Text                       string   `xml:",chardata"`
	GetImagingSettingsResponse struct {
		Text            string                `xml:",chardata"`
		Xmlns           string                `xml:"xmlns,attr"`
		ImagingSettings CustomImagingSettings `xml:"ImagingSettings"`
	} `xml:"GetImagingSettingsResponse"`
}

type CustomImagingSettings struct {
	XMLName               xml.Name `xml:"ImagingSettings"`
	Text                  string   `xml:",chardata"`
	BacklightCompensation struct {
		Mode  string `xml:"Mode"`
		Level string `xml:"Level"`
	} `xml:"BacklightCompensation"`
	Brightness      string `xml:"Brightness"`
	ColorSaturation string `xml:"ColorSaturation"`
	Contrast        string `xml:"Contrast"`
	Exposure        struct {
		Text            string `xml:",chardata"`
		Mode            string `xml:"Mode"`
		Priority        string `xml:"Priority"`
		Window          string `xml:"Window"`
		MinExposureTime string `xml:"MinExposureTime"`
		MaxExposureTime string `xml:"MaxExposureTime"`
		MinGain         string `xml:"MinGain"`
		MaxGain         string `xml:"MaxGain"`
		MinIris         string `xml:"MinIris"`
		MaxIris         string `xml:"MaxIris"`
		ExposureTime    string `xml:"ExposureTime"`
		Gain            string `xml:"Gain"`
		Iris            float64 `xml:"Iris"`
	} `xml:"Exposure"`
	Focus struct {
		Text          string `xml:",chardata"`
		AFMode        string `xml:"AFMode"`
		AutoFocusMode string `xml:"AutoFocusMode"`
		DefaultSpeed  string `xml:"DefaultSpeed"`
		NearLimit     string `xml:"NearLimit"`
		FarLimit      string `xml:"FarLimit"`
		Extension     string `xml:"Extension"`
	} `xml:"Focus"`
	IrCutFilter      string `xml:"IrCutFilter"`
	Sharpness        string `xml:"Sharpness"`
	WideDynamicRange struct {
		Text  string `xml:",chardata"`
		Mode  string `xml:"Mode"`
		Level string `xml:"Level"`
	} `xml:"WideDynamicRange"`
	WhiteBalance struct {
		Text      string  `xml:",chardata"`
		Mode      string  `xml:"Mode"`
		CrGain    float64 `xml:"CrGain"`
		CbGain    float64 `xml:"CbGain"`
		Extension string  `xml:"Extension"`
	} `xml:"WhiteBalance"`
	Extension struct {
		Text               string `xml:",chardata"`
		ImageStabilization struct {
			Text      string `xml:",chardata"`
			Mode      string `xml:"Mode"`
			Level     string `xml:"Level"`
			Extension string `xml:"Extension"`
		} `xml:"ImageStabilization"`
		Extension struct {
			Text      string `xml:",chardata"`
			Extension struct {
				Text             string `xml:",chardata"`
				ToneCompensation struct {
					Text  string `xml:",chardata"`
					Mode  string `xml:"Mode"`
					Level string `xml:"Level"`
				} `xml:"ToneCompensation"`
				Defogging struct {
					Text  string `xml:",chardata"`
					Mode  string `xml:"Mode"`
					Level string `xml:"Level"`
				} `xml:"Defogging"`
			} `xml:"Extension"`
		} `xml:"Extension"`
	} `xml:"Extension"`
}