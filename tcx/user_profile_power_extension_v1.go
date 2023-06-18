// user_profile_power_extension_v1.go: based on the code generated by xsdgen.
// http://www.garmin.com/xmlschemas/UserProfilePowerExtensionv1.xsd

package tcx

type ProfilePowerZonest struct {
	FTP       uint                `xml:"http://www.garmin.com/xmlschemas/ProfileExtension/v1 FTP,omitempty"`
	PowerZone []ProfilePowerZonet `xml:"http://www.garmin.com/xmlschemas/ProfileExtension/v1 PowerZone"`
}

type ProfilePowerZonet struct {
	Number int  `xml:"http://www.garmin.com/xmlschemas/ProfileExtension/v1 Number"`
	Low    uint `xml:"http://www.garmin.com/xmlschemas/ProfileExtension/v1 Low"`
	High   uint `xml:"http://www.garmin.com/xmlschemas/ProfileExtension/v1 High"`
}
