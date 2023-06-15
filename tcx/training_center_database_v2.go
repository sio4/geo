// Code generated by xsdgen. DO NOT EDIT.

package tcx

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AbstractSourcet struct {
	Name string `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Name"`
}

type AbstractStept struct {
	StepId int `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 StepId"`
}

type ActivityLapt struct {
	TotalTimeSeconds    float64                    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 TotalTimeSeconds"`
	DistanceMeters      float64                    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 DistanceMeters"`
	MaximumSpeed        float64                    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 MaximumSpeed,omitempty"`
	Calories            uint                       `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Calories"`
	AverageHeartRateBpm HeartRateInBeatsPerMinutet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 AverageHeartRateBpm,omitempty"`
	MaximumHeartRateBpm HeartRateInBeatsPerMinutet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 MaximumHeartRateBpm,omitempty"`
	Intensity           Intensityt                 `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Intensity"`
	Cadence             byte                       `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Cadence,omitempty"`
	TriggerMethod       TriggerMethodt             `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 TriggerMethod"`
	Track               []Trackt                   `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Track,omitempty"`
	Notes               string                     `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Notes,omitempty"`
	Extensions          Extensionst                `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
	StartTime           time.Time                  `xml:"StartTime,attr"`
}

func (t *ActivityLapt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ActivityLapt
	var layout struct {
		*T
		StartTime *xsdDateTime `xml:"StartTime,attr"`
	}
	layout.T = (*T)(t)
	layout.StartTime = (*xsdDateTime)(&layout.T.StartTime)
	return e.EncodeElement(layout, start)
}
func (t *ActivityLapt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ActivityLapt
	var overlay struct {
		*T
		StartTime *xsdDateTime `xml:"StartTime,attr"`
	}
	overlay.T = (*T)(t)
	overlay.StartTime = (*xsdDateTime)(&overlay.T.StartTime)
	return d.DecodeElement(&overlay, &start)
}

type ActivityListt struct {
	Activity          []Activityt          `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Activity,omitempty"`
	MultiSportSession []MultiSportSessiont `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 MultiSportSession,omitempty"`
}

type ActivityReferencet struct {
	Id time.Time `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Id"`
}

func (t *ActivityReferencet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ActivityReferencet
	var layout struct {
		*T
		Id *xsdDateTime `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Id"`
	}
	layout.T = (*T)(t)
	layout.Id = (*xsdDateTime)(&layout.T.Id)
	return e.EncodeElement(layout, start)
}
func (t *ActivityReferencet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ActivityReferencet
	var overlay struct {
		*T
		Id *xsdDateTime `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Id"`
	}
	overlay.T = (*T)(t)
	overlay.Id = (*xsdDateTime)(&overlay.T.Id)
	return d.DecodeElement(&overlay, &start)
}

type Activityt struct {
	Id         time.Time       `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Id"`
	Lap        []ActivityLapt  `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Lap"`
	Notes      string          `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Notes,omitempty"`
	Training   Trainingt       `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Training,omitempty"`
	Creator    AbstractSourcet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Creator,omitempty"`
	Extensions Extensionst     `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
	Sport      Sportt          `xml:"Sport,attr"`
}

func (t *Activityt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Activityt
	var layout struct {
		*T
		Id *xsdDateTime `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Id"`
	}
	layout.T = (*T)(t)
	layout.Id = (*xsdDateTime)(&layout.T.Id)
	return e.EncodeElement(layout, start)
}
func (t *Activityt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Activityt
	var overlay struct {
		*T
		Id *xsdDateTime `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Id"`
	}
	overlay.T = (*T)(t)
	overlay.Id = (*xsdDateTime)(&overlay.T.Id)
	return d.DecodeElement(&overlay, &start)
}

// Identifies a PC software application.
type Applicationt struct {
	Name       string      `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Name"`
	Build      Buildt      `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Build"`
	LangID     string      `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 LangID"`
	PartNumber PartNumbert `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 PartNumber"`
}

// May be one of Internal, Alpha, Beta, Release
type BuildTypet string

type Buildt struct {
	Version Versiont   `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Version"`
	Type    BuildTypet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Type,omitempty"`
	Time    string     `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Time,omitempty"`
	Builder string     `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Builder,omitempty"`
}

type Cadencet struct {
	Low  float64 `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Low"`
	High float64 `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 High"`
}

type CaloriesBurnedt struct {
	Calories uint `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Calories"`
}

type CourseFoldert struct {
	Folder        []CourseFoldert     `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Folder,omitempty"`
	CourseNameRef []NameKeyReferencet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 CourseNameRef,omitempty"`
	Notes         string              `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Notes,omitempty"`
	Extensions    Extensionst         `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
	Name          string              `xml:"Name,attr"`
}

type CourseLapt struct {
	TotalTimeSeconds    float64                    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 TotalTimeSeconds"`
	DistanceMeters      float64                    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 DistanceMeters"`
	BeginPosition       Positiont                  `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 BeginPosition,omitempty"`
	BeginAltitudeMeters float64                    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 BeginAltitudeMeters,omitempty"`
	EndPosition         Positiont                  `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 EndPosition,omitempty"`
	EndAltitudeMeters   float64                    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 EndAltitudeMeters,omitempty"`
	AverageHeartRateBpm HeartRateInBeatsPerMinutet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 AverageHeartRateBpm,omitempty"`
	MaximumHeartRateBpm HeartRateInBeatsPerMinutet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 MaximumHeartRateBpm,omitempty"`
	Intensity           Intensityt                 `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Intensity"`
	Cadence             byte                       `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Cadence,omitempty"`
	Extensions          Extensionst                `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
}

type CourseListt struct {
	Course []Courset `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Course,omitempty"`
}

// May be no more than 10 items long
type CoursePointNamet string

// May be one of Generic, Summit, Valley, Water, Food, Danger, Left, Right, Straight, First Aid, 4th Category, 3rd Category, 2nd Category, 1st Category, Hors Category, Sprint
type CoursePointTypet string

type CoursePointt struct {
	Name           CoursePointNamet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Name"`
	Time           time.Time        `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Time"`
	Position       Positiont        `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Position"`
	AltitudeMeters float64          `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 AltitudeMeters,omitempty"`
	PointType      CoursePointTypet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 PointType"`
	Notes          string           `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Notes,omitempty"`
	Extensions     Extensionst      `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
}

func (t *CoursePointt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T CoursePointt
	var layout struct {
		*T
		Time *xsdDateTime `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Time"`
	}
	layout.T = (*T)(t)
	layout.Time = (*xsdDateTime)(&layout.T.Time)
	return e.EncodeElement(layout, start)
}
func (t *CoursePointt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CoursePointt
	var overlay struct {
		*T
		Time *xsdDateTime `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Time"`
	}
	overlay.T = (*T)(t)
	overlay.Time = (*xsdDateTime)(&overlay.T.Time)
	return d.DecodeElement(&overlay, &start)
}

type Coursest struct {
	CourseFolder CourseFoldert `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 CourseFolder"`
	Extensions   Extensionst   `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
}

type Courset struct {
	Name        RestrictedTokent `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Name"`
	Lap         []CourseLapt     `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Lap,omitempty"`
	Track       []Trackt         `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Track,omitempty"`
	Notes       string           `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Notes,omitempty"`
	CoursePoint []CoursePointt   `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 CoursePoint,omitempty"`
	Creator     AbstractSourcet  `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Creator,omitempty"`
	Extensions  Extensionst      `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
}

type CustomHeartRateZonet struct {
	Low  HeartRateValuet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Low"`
	High HeartRateValuet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 High"`
}

type CustomSpeedZonet struct {
	ViewAs                SpeedTypet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 ViewAs"`
	LowInMetersPerSecond  float64    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 LowInMetersPerSecond"`
	HighInMetersPerSecond float64    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 HighInMetersPerSecond"`
}

// Identifies the originating GPS device that tracked a run or
// used to identify the type of device capable of handling
// the data for loading.
type Devicet struct {
	Name      string   `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Name"`
	UnitId    uint     `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 UnitId"`
	ProductID uint     `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 ProductID"`
	Version   Versiont `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Version"`
}

type Distancet struct {
	Meters uint `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Meters"`
}

type Durationt struct {
}

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

type FirstSportt struct {
	Activity Activityt `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Activity"`
}

type Folderst struct {
	History  Historyt  `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 History,omitempty"`
	Workouts Workoutst `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Workouts,omitempty"`
	Courses  Coursest  `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Courses,omitempty"`
}

// May be one of Male, Female
type Gendert string

type HeartRateAbovet struct {
	HeartRate HeartRateValuet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 HeartRate"`
}

type HeartRateAsPercentOfMaxt struct {
	Value byte `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Value"`
}

type HeartRateBelowt struct {
	HeartRate HeartRateValuet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 HeartRate"`
}

type HeartRateInBeatsPerMinutet struct {
	Value byte `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Value"`
}

type HeartRateValuet struct {
}

type HeartRatet struct {
	HeartRateZone Zonet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 HeartRateZone"`
}

type HistoryFoldert struct {
	Folder      []HistoryFoldert     `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Folder,omitempty"`
	ActivityRef []ActivityReferencet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 ActivityRef,omitempty"`
	Week        []Weekt              `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Week,omitempty"`
	Notes       string               `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Notes,omitempty"`
	Extensions  Extensionst          `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
	Name        string               `xml:"Name,attr"`
}

type Historyt struct {
	Running    HistoryFoldert    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Running"`
	Biking     HistoryFoldert    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Biking"`
	Other      HistoryFoldert    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Other"`
	MultiSport MultiSportFoldert `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 MultiSport"`
	Extensions Extensionst       `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
}

// May be one of Active, Resting
type Intensityt string

type MultiSportFoldert struct {
	Folder                []MultiSportFoldert  `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Folder,omitempty"`
	MultisportActivityRef []ActivityReferencet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 MultisportActivityRef,omitempty"`
	Week                  []Weekt              `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Week,omitempty"`
	Notes                 string               `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Notes,omitempty"`
	Extensions            Extensionst          `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
	Name                  string               `xml:"Name,attr"`
}

type MultiSportSessiont struct {
	Id         time.Time    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Id"`
	FirstSport FirstSportt  `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 FirstSport"`
	NextSport  []NextSportt `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 NextSport,omitempty"`
	Notes      string       `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Notes,omitempty"`
}

func (t *MultiSportSessiont) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T MultiSportSessiont
	var layout struct {
		*T
		Id *xsdDateTime `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Id"`
	}
	layout.T = (*T)(t)
	layout.Id = (*xsdDateTime)(&layout.T.Id)
	return e.EncodeElement(layout, start)
}
func (t *MultiSportSessiont) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MultiSportSessiont
	var overlay struct {
		*T
		Id *xsdDateTime `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Id"`
	}
	overlay.T = (*T)(t)
	overlay.Id = (*xsdDateTime)(&overlay.T.Id)
	return d.DecodeElement(&overlay, &start)
}

type NameKeyReferencet struct {
	Id RestrictedTokent `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Id"`
}

type NextSportt struct {
	Transition ActivityLapt `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Transition,omitempty"`
	Activity   Activityt    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Activity"`
}

type Nonet struct {
}

// Must match the pattern [\p{Lu}\d]{3}-[\p{Lu}\d]{5}-[\p{Lu}\d]{2}
type PartNumbert string

type Plant struct {
	Name            RestrictedTokent `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Name,omitempty"`
	Extensions      Extensionst      `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
	Type            TrainingTypet    `xml:"Type,attr"`
	IntervalWorkout bool             `xml:"IntervalWorkout,attr"`
}

type Positiont struct {
	LatitudeDegrees  float64 `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 LatitudeDegrees"`
	LongitudeDegrees float64 `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 LongitudeDegrees"`
}

type PredefinedHeartRateZonet struct {
	Number int `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Number"`
}

type PredefinedSpeedZonet struct {
	Number int `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Number"`
}

type QuickWorkoutt struct {
	TotalTimeSeconds float64 `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 TotalTimeSeconds"`
	DistanceMeters   float64 `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 DistanceMeters"`
}

type Repeatt struct {
	StepId      int             `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 StepId"`
	Repetitions int             `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Repetitions"`
	Child       []AbstractStept `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Child"`
}

// May be no more than 15 items long
type RestrictedTokent string

// May be one of Present, Absent
type SensorStatet string

// May be one of Pace, Speed
type SpeedTypet string

type Speedt struct {
	SpeedZone Zonet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 SpeedZone"`
}

// May be one of Running, Biking, Other
type Sportt string

type Stept struct {
	StepId    int              `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 StepId"`
	Name      RestrictedTokent `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Name,omitempty"`
	Duration  Durationt        `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Duration"`
	Intensity Intensityt       `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Intensity"`
	Target    Targett          `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Target"`
}

type Targett struct {
}

type Timet struct {
	Seconds uint `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Seconds"`
}

// Token must be defined as a type because of a bug in the MSXML parser which
// does not correctly process xsd:token using the whiteSpace value of "collapse"
type Tokent string

type Trackpointt struct {
	Time           time.Time                  `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Time"`
	Position       Positiont                  `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Position,omitempty"`
	AltitudeMeters float64                    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 AltitudeMeters,omitempty"`
	DistanceMeters float64                    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 DistanceMeters,omitempty"`
	HeartRateBpm   HeartRateInBeatsPerMinutet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 HeartRateBpm,omitempty"`
	Cadence        byte                       `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Cadence,omitempty"`
	SensorState    SensorStatet               `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 SensorState,omitempty"`
	Extensions     Extensionst                `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
}

func (t *Trackpointt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Trackpointt
	var layout struct {
		*T
		Time *xsdDateTime `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Time"`
	}
	layout.T = (*T)(t)
	layout.Time = (*xsdDateTime)(&layout.T.Time)
	return e.EncodeElement(layout, start)
}
func (t *Trackpointt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Trackpointt
	var overlay struct {
		*T
		Time *xsdDateTime `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Time"`
	}
	overlay.T = (*T)(t)
	overlay.Time = (*xsdDateTime)(&overlay.T.Time)
	return d.DecodeElement(&overlay, &start)
}

type Trackt struct {
	Trackpoint []Trackpointt `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Trackpoint"`
}

type TrainingCenterDatabaset struct {
	Folders    Folderst        `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Folders,omitempty"`
	Activities ActivityListt   `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Activities,omitempty"`
	Workouts   WorkoutListt    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Workouts,omitempty"`
	Courses    CourseListt     `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Courses,omitempty"`
	Author     AbstractSourcet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Author,omitempty"`
	Extensions Extensionst     `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
}

// May be one of Workout, Course
type TrainingTypet string

type Trainingt struct {
	QuickWorkoutResults QuickWorkoutt `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 QuickWorkoutResults,omitempty"`
	Plan                Plant         `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Plan,omitempty"`
	VirtualPartner      bool          `xml:"VirtualPartner,attr"`
}

// May be one of Manual, Distance, Location, Time, HeartRate
type TriggerMethodt string

type UserInitiatedt struct {
}

type Versiont struct {
	VersionMajor uint `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 VersionMajor"`
	VersionMinor uint `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 VersionMinor"`
	BuildMajor   uint `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 BuildMajor,omitempty"`
	BuildMinor   uint `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 BuildMinor,omitempty"`
}

type Weekt struct {
	Notes    string    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Notes,omitempty"`
	StartDay time.Time `xml:"StartDay,attr"`
}

func (t *Weekt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Weekt
	var layout struct {
		*T
		StartDay *xsdDate `xml:"StartDay,attr"`
	}
	layout.T = (*T)(t)
	layout.StartDay = (*xsdDate)(&layout.T.StartDay)
	return e.EncodeElement(layout, start)
}
func (t *Weekt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Weekt
	var overlay struct {
		*T
		StartDay *xsdDate `xml:"StartDay,attr"`
	}
	overlay.T = (*T)(t)
	overlay.StartDay = (*xsdDate)(&overlay.T.StartDay)
	return d.DecodeElement(&overlay, &start)
}

type WorkoutFoldert struct {
	Folder         []WorkoutFoldert    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Folder,omitempty"`
	WorkoutNameRef []NameKeyReferencet `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 WorkoutNameRef,omitempty"`
	Extensions     Extensionst         `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
	Name           string              `xml:"Name,attr"`
}

type WorkoutListt struct {
	Workout []Workoutt `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Workout,omitempty"`
}

type Workoutst struct {
	Running    WorkoutFoldert `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Running"`
	Biking     WorkoutFoldert `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Biking"`
	Other      WorkoutFoldert `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Other"`
	Extensions Extensionst    `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
}

type Workoutt struct {
	Name        RestrictedTokent `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Name"`
	Step        []AbstractStept  `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Step"`
	ScheduledOn []time.Time      `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 ScheduledOn,omitempty"`
	Notes       string           `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Notes,omitempty"`
	Creator     AbstractSourcet  `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Creator,omitempty"`
	Extensions  Extensionst      `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 Extensions,omitempty"`
	Sport       Sportt           `xml:"Sport,attr"`
}

func (t *Workoutt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Workoutt
	var layout struct {
		*T
		ScheduledOn *xsdDate `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 ScheduledOn,omitempty"`
	}
	layout.T = (*T)(t)
	layout.ScheduledOn = (*xsdDate)(&layout.T.ScheduledOn)
	return e.EncodeElement(layout, start)
}
func (t *Workoutt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Workoutt
	var overlay struct {
		*T
		ScheduledOn *xsdDate `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 ScheduledOn,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.ScheduledOn = (*xsdDate)(&overlay.T.ScheduledOn)
	return d.DecodeElement(&overlay, &start)
}

type Zonet struct {
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}