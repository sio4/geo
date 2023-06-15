// Code generated by xsdgen. DO NOT EDIT.

package tcx

import "encoding/xml"

type ActivityLapExtensiont struct {
	AvgSpeed       float64     `xml:"http://www.garmin.com/xmlschemas/ActivityExtension/v2 AvgSpeed,omitempty"`
	MaxBikeCadence byte        `xml:"http://www.garmin.com/xmlschemas/ActivityExtension/v2 MaxBikeCadence,omitempty"`
	AvgRunCadence  byte        `xml:"http://www.garmin.com/xmlschemas/ActivityExtension/v2 AvgRunCadence,omitempty"`
	MaxRunCadence  byte        `xml:"http://www.garmin.com/xmlschemas/ActivityExtension/v2 MaxRunCadence,omitempty"`
	Steps          uint        `xml:"http://www.garmin.com/xmlschemas/ActivityExtension/v2 Steps,omitempty"`
	AvgWatts       uint        `xml:"http://www.garmin.com/xmlschemas/ActivityExtension/v2 AvgWatts,omitempty"`
	MaxWatts       uint        `xml:"http://www.garmin.com/xmlschemas/ActivityExtension/v2 MaxWatts,omitempty"`
	Extensions     Extensionst `xml:"http://www.garmin.com/xmlschemas/ActivityExtension/v2 Extensions,omitempty"`
}

type ActivityTrackpointExtensiont struct {
	Speed         float64            `xml:"http://www.garmin.com/xmlschemas/ActivityExtension/v2 Speed,omitempty"`
	RunCadence    byte               `xml:"http://www.garmin.com/xmlschemas/ActivityExtension/v2 RunCadence,omitempty"`
	Watts         uint               `xml:"http://www.garmin.com/xmlschemas/ActivityExtension/v2 Watts,omitempty"`
	Extensions    Extensionst        `xml:"http://www.garmin.com/xmlschemas/ActivityExtension/v2 Extensions,omitempty"`
	CadenceSensor CadenceSensorTypet `xml:"CadenceSensor,attr,omitempty"`
}

// May be one of Footpod, Bike
type CadenceSensorTypet string

type Extensionst []string

func (a Extensionst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ArrayType string   `xml:"http://schemas.xmlsoap.org/wsdl/ arrayType,attr"`
		Items     []string `xml:" item"`
	}
	output.Items = []string(a)
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{"", "xmlns:ns1"}, Value: "http://www.w3.org/2001/XMLSchema"})
	output.ArrayType = "ns1:anyType[]"
	return e.EncodeElement(&output, start)
}
func (a *Extensionst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var tok xml.Token
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if tok, ok := tok.(xml.StartElement); ok {
			var item string
			if err = d.DecodeElement(&item, &tok); err == nil {
				*a = append(*a, item)
			}
		}
		if _, ok := tok.(xml.EndElement); ok {
			break
		}
	}
	return err
}

// Token must be defined as a type because of a bug in the MSXML parser which does not correctly process xsd:token using the whiteSpace value
// of "collapse"
type Tokent string
