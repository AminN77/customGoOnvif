package imaging

import (
	"github.com/AminN77/customGoOnvif/xsd"
	"github.com/AminN77/customGoOnvif/xsd/onvif"
)

type Capabilities struct {
	EFlip                       xsd.Boolean `xml:"EFlip,attr"`
	Reverse                     xsd.Boolean `xml:"Reverse,attr"`
	GetCompatibleConfigurations xsd.Boolean `xml:"GetCompatibleConfigurations,attr"`
	MoveStatus                  xsd.Boolean `xml:"MoveStatus,attr"`
	StatusPosition              xsd.Boolean `xml:"StatusPosition,attr"`
}

type GetServiceCapabilities struct {
	XMLName string `xml:"timg:GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

type GetImagingSettings struct {
	XMLName          string               `xml:"timg:GetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type SetExposureImagingSettings struct {
	XMLName          string                          `xml:"timg:SetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken            `xml:"timg:VideoSourceToken"`
	ImagingSettings  onvif.ExposureImagingSettings20 `xml:"timg:ImagingSettings"`
	ForcePersistence xsd.Boolean                     `xml:"timg:ForcePersistence"`
}

type SetSaturationImagingSettings struct {
	XMLName          string                            `xml:"timg:SetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken              `xml:"timg:VideoSourceToken"`
	ImagingSettings  onvif.SaturationImagingSettings20 `xml:"timg:ImagingSettings"`
	ForcePersistence xsd.Boolean                       `xml:"timg:ForcePersistence"`
}

type SetBrightnessImagingSettings struct {
	XMLName          string                            `xml:"timg:SetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken              `xml:"timg:VideoSourceToken"`
	ImagingSettings  onvif.BrightnessImagingSettings20 `xml:"timg:ImagingSettings"`
	ForcePersistence xsd.Boolean                       `xml:"timg:ForcePersistence"`
}

type SetContrastImagingSettings struct {
	XMLName          string                          `xml:"timg:SetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken            `xml:"timg:VideoSourceToken"`
	ImagingSettings  onvif.ContrastImagingSettings20 `xml:"timg:ImagingSettings"`
	ForcePersistence xsd.Boolean                     `xml:"timg:ForcePersistence"`
}

type SetSharpnessImagingSettings struct {
	XMLName          string                           `xml:"timg:SetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken             `xml:"timg:VideoSourceToken"`
	ImagingSettings  onvif.SharpnessImagingSettings20 `xml:"timg:ImagingSettings"`
	ForcePersistence xsd.Boolean                      `xml:"timg:ForcePersistence"`
}

type GetOptions struct {
	XMLName          string               `xml:"timg:GetOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type MoveFocusAbs struct {
	XMLName          string               `xml:"timg:Move"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	Focus            onvif.FocusMoveAbs   `xml:"timg:Focus"`
}

type MoveFocusRel struct {
	XMLName          string               `xml:"timg:Move"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	Focus            onvif.FocusMoveRel   `xml:"timg:Focus"`
}

type MoveFocusCon struct {
	XMLName          string               `xml:"timg:Move"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	Focus            onvif.FocusMoveCon   `xml:"timg:Focus"`
}

type GetMoveOptions struct {
	XMLName          string               `xml:"timg:GetMoveOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type Stop struct {
	XMLName          string               `xml:"timg:Stop"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetStatus struct {
	XMLName          string               `xml:"timg:GetStatus"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetPresets struct {
	XMLName          string               `xml:"timg:GetPresets"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetCurrentPreset struct {
	XMLName          string               `xml:"timg:GetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type SetCurrentPreset struct {
	XMLName          string               `xml:"timg:SetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	PresetToken      onvif.ReferenceToken `xml:"timg:PresetToken"`
}
