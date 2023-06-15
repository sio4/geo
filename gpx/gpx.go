package gpx

import (
	"encoding/xml"
	"os"
)

func ParseFile(filename string) (*GPX, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return &GPX{}, err
	}

	return Parse(bytes)
}

func Parse(bytes []byte) (*GPX, error) {
	gpx := &GPX{}

	err := xml.Unmarshal(bytes, gpx)
	if err != nil {
		return gpx, err
	}

	return gpx, nil
}

func (x *GPX) WriteGPX(filename string) error {
	data, err := x.ToXML()
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (x *GPX) ToXML() ([]byte, error) {
	return xml.MarshalIndent(x, "", "  ")
}
