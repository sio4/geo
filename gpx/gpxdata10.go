// gpxdata10.go: based on the code generated by xsdgen.
// http://www.cluetrust.com/Schemas/gpxdata10.xsd

package gpx

import (
	"encoding/xml"
	"time"

	"github.com/sio4/geo/xsd"
)

// The GPXData schema is a mix-in schema for use with the GPX 1.1 schema from
// Topografix.  It was designed with the intent of expanding GPX (using the
// existing mechanisms) to contain additional content from GPS receivers
// commonly used for sporting and geocaching activities.

type (
	HR      uint    // hr, beats per minute
	Cadence uint    // cadence, revolutions per minute
	Temp    float64 // temp, degrees celcius
)

// dataPt is used to store information about non-geographic data point.
// This is useful when HR or other data is present, but no latitude or
// longitude data is.
type DataPt struct {
	XMLName xml.Name `xml:"http://www.cluetrust.com/XML/GPXDATA/1/0 dataPt"`

	Time     time.Time `xml:"time"`
	Ele      float64   `xml:"ele,omitempty"`
	HR       HR        `xml:"hr,omitempty"`
	Cadence  Cadence   `xml:"cadence,omitempty"`
	Temp     Temp      `xml:"temp,omitempty"`
	Distance Distance  `xml:"distance,omitempty"`
	Sensor   Sensor    `xml:"sensor,omitempty"`
}

func (t *DataPt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T DataPt
	var layout struct {
		*T
		Time *xsd.DateTime `xml:"time"`
	}
	layout.T = (*T)(t)
	layout.Time = (*xsd.DateTime)(&layout.T.Time)
	return e.EncodeElement(layout, start)
}

func (t *DataPt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T DataPt
	var overlay struct {
		*T
		Time *xsd.DateTime `xml:"time"`
	}
	overlay.T = (*T)(t)
	overlay.Time = (*xsd.DateTime)(&overlay.T.Time)
	return d.DecodeElement(&overlay, &start)
}

type Run struct {
	Sport       SportType `xml:"sport,omitempty"`       // The type of sport being participated in
	ProgramType string    `xml:"programType,omitempty"` // The type of program being executed
	Laps        Laps      `xml:"laps,omitempty"`        // Group of laps that make up the run
}

type CourseType struct {
	Index int  `xml:"index,omitempty"` // Index of the course in the GPS
	Laps  Laps `xml:"laps,omitempty"`  // Group of laps that make up the course
}

type Laps struct {
	Lap []Lap `xml:"lap,omitempty"`
}

type Lap struct {
	Index          int            `xml:"index,omitempty"`
	StartPoint     Location       `xml:"startPoint,omitempty"`
	EndPoint       Location       `xml:"endPoint,omitempty"`
	StartTime      time.Time      `xml:"startTime,omitempty"`
	ElapsedTime    float32        `xml:"elapsedTime,omitempty"`
	Calories       uint           `xml:"calories,omitempty"`
	Distance       float32        `xml:"distance,omitempty"`
	TrackReference TrackReference `xml:"trackReference,omitempty"`
	Summary        []Summary      `xml:"summary,omitempty"`
	Trigger        Trigger        `xml:"trigger,omitempty"`
	Intensity      IntensityKind  `xml:"intensity,omitempty"`
}

func (t *Lap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Lap
	var layout struct {
		*T
		StartTime *xsd.DateTime `xml:"startTime,omitempty"`
	}
	layout.T = (*T)(t)
	layout.StartTime = (*xsd.DateTime)(&layout.T.StartTime)
	return e.EncodeElement(layout, start)
}

func (t *Lap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Lap
	var overlay struct {
		*T
		StartTime *xsd.DateTime `xml:"startTime,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.StartTime = (*xsd.DateTime)(&overlay.T.StartTime)
	return d.DecodeElement(&overlay, &start)
}

type SportType string

const (
	SportTypeBike  SportType = "bike"
	SportTypeRun   SportType = "run"
	SportTypeOther SportType = "other"
)

type Distance struct {
	Sensor SensorKind `xml:"sensor,attr,omitempty"`

	Value float64 `xml:",chardata"`
}

type Sensor struct {
	Kind SensorKind `xml:"kind,attr"`
}

type SensorKind string

const (
	SensorKindWheel     SensorKind = "wheel"
	SensorKindPedometer SensorKind = "pedometer"
)

type Summary struct {
	Name string      `xml:"name,attr"`
	Kind SummaryKind `xml:"kind,attr"`

	Value float64 `xml:",chardata"`
}

type SummaryKind string

const (
	SummaryKindMax SummaryKind = "max"
	SummaryKindMin SummaryKind = "min"
	SummaryKindAvg SummaryKind = "avg"
)

type Trigger struct {
	Kind TriggerKind `xml:"kind,attr"`

	Value string `xml:",chardata"`
}

type TriggerKind string

const (
	TriggerKindManual   TriggerKind = "manual"
	TriggerKindTime     TriggerKind = "time"
	TriggerKindDistance TriggerKind = "distance"
	TriggerKindLocation TriggerKind = "location"
	TriggerKindHR       TriggerKind = "hr"
)

type IntensityKind string

const (
	IntensityKindRest   IntensityKind = "rest"
	IntensityKindActive IntensityKind = "active"
)

type Location struct {
	Lat Latitude  `xml:"lat,attr"` // type from GPX
	Lon Longitude `xml:"lon,attr"` // type from GPX
}

type TrackReference struct {
	TrackNumber uint   `xml:"trackNumber,attr,omitempty"`
	TrackName   string `xml:"trackName,attr,omitempty"`

	StartPoint uint `xml:"startPoint,omitempty"`
	EndPoint   uint `xml:"endPoint,omitempty"`
}
