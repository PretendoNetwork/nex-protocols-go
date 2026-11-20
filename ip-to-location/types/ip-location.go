// Package types implements all the types used by the IPToLocation protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// IPLocation is a type with the IPToLocation protocol
type IPLocation struct {
	types.Structure
	IPStart        types.UInt32
	CountryCode    types.String
	CountryName    types.String
	RegionCode     types.String
	RegionName     types.String
	CityName       types.String
	Latitude       types.Float
	Longitude      types.Float
	CountryGeoCode types.UInt32
	RegionGeoCode  types.UInt32
	CityGeoCode    types.UInt32
}

// WriteTo writes the IPLocation to the given writable
func (ipl IPLocation) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ipl.IPStart.WriteTo(contentWritable)
	ipl.CountryCode.WriteTo(contentWritable)
	ipl.CountryName.WriteTo(contentWritable)
	ipl.RegionCode.WriteTo(contentWritable)
	ipl.RegionName.WriteTo(contentWritable)
	ipl.CityName.WriteTo(contentWritable)
	ipl.Latitude.WriteTo(contentWritable)
	ipl.Longitude.WriteTo(contentWritable)
	ipl.CountryGeoCode.WriteTo(contentWritable)
	ipl.RegionGeoCode.WriteTo(contentWritable)
	ipl.CityGeoCode.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ipl.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the IPLocation from the given readable
func (ipl *IPLocation) ExtractFrom(readable types.Readable) error {
	if err := ipl.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation header. %s", err.Error())
	}

	if err := ipl.IPStart.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_ipStart. %s", err.Error())
	}

	if err := ipl.CountryCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_countryCode. %s", err.Error())
	}

	if err := ipl.CountryName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_countryName. %s", err.Error())
	}

	if err := ipl.RegionCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_regionCode. %s", err.Error())
	}

	if err := ipl.RegionName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_regionName. %s", err.Error())
	}

	if err := ipl.CityName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_cityName. %s", err.Error())
	}

	if err := ipl.Latitude.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_latitude. %s", err.Error())
	}

	if err := ipl.Longitude.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_longitude. %s", err.Error())
	}

	if err := ipl.CountryGeoCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_countryGeoCode. %s", err.Error())
	}

	if err := ipl.RegionGeoCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_regionGeoCode. %s", err.Error())
	}

	if err := ipl.CityGeoCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_cityGeoCode. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instanced of IPLocation
func (ipl IPLocation) Copy() types.RVType {
	copied := NewIPLocation()

	copied.StructureVersion = ipl.StructureVersion
	copied.IPStart = ipl.IPStart
	copied.CountryCode = ipl.CountryCode
	copied.CountryName = ipl.CountryName
	copied.RegionCode = ipl.RegionCode
	copied.RegionName = ipl.RegionName
	copied.CityName = ipl.CityName
	copied.Latitude = ipl.Latitude
	copied.Longitude = ipl.Longitude
	copied.CountryGeoCode = ipl.CountryGeoCode
	copied.RegionGeoCode = ipl.RegionGeoCode
	copied.CityGeoCode = ipl.CityGeoCode

	return copied
}

// Equals checks if the given IPLocation contains the same data as the current IPLocation
func (ipl IPLocation) Equals(o types.RVType) bool {
	if _, ok := o.(IPLocation); !ok {
		return false
	}

	other := o.(IPLocation)

	if ipl.StructureVersion != other.StructureVersion {
		return false
	}

	if ipl.IPStart != other.IPStart {
		return false
	}

	if ipl.CountryCode != other.CountryCode {
		return false
	}

	if ipl.CountryName != other.CountryName {
		return false
	}

	if ipl.RegionCode != other.RegionCode {
		return false
	}

	if ipl.RegionName != other.RegionName {
		return false
	}

	if ipl.CityName != other.CityName {
		return false
	}

	if ipl.Latitude != other.Latitude {
		return false
	}

	if ipl.Longitude != other.Longitude {
		return false
	}

	if ipl.CountryGeoCode != other.CountryGeoCode {
		return false
	}

	if ipl.RegionGeoCode != other.RegionGeoCode {
		return false
	}

	return ipl.CityGeoCode == other.CityGeoCode
}

// CopyRef copies the current value of the IPLocation
// and returns a pointer to the new copy
func (ipl IPLocation) CopyRef() types.RVTypePtr {
	copied := ipl.Copy().(IPLocation)
	return &copied
}

// Deref takes a pointer of IPLocation
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (ipl *IPLocation) Deref() types.RVType {
	return *ipl
}

// String returns the string representation of the IPLocation
func (ipl IPLocation) String() string {
	return ipl.FormatToString(0)
}

// FormatToString pretty-prints the IPLocation using the provided indentation level
func (ipl IPLocation) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("IPlocation{\n}")
	fmt.Fprintf(&b, "%sm_ipStart: %s,\n", indentationValues, ipl.IPStart)
	fmt.Fprintf(&b, "%sm_countryCode: %s,\n", indentationValues, ipl.CountryCode)
	fmt.Fprintf(&b, "%sm_countryName: %s,\n", indentationValues, ipl.CountryName)
	fmt.Fprintf(&b, "%sm_regionCode: %s,\n", indentationValues, ipl.RegionCode)
	fmt.Fprintf(&b, "%sm_regionName: %s,\n", indentationValues, ipl.RegionName)
	fmt.Fprintf(&b, "%sm_cityName: %s,\n", indentationValues, ipl.CityName)
	fmt.Fprintf(&b, "%sm_latitude: %s,\n", indentationValues, ipl.Latitude)
	fmt.Fprintf(&b, "%sm_longitude: %s,\n", indentationValues, ipl.Longitude)
	fmt.Fprintf(&b, "%sm_countryGeoCode: %s,\n", indentationValues, ipl.CountryGeoCode)
	fmt.Fprintf(&b, "%sm_regionGeoCode: %s,\n", indentationValues, ipl.RegionGeoCode)
	fmt.Fprintf(&b, "%sm_cityGeoCode: %s,\n", indentationValues, ipl.CityGeoCode)

	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewIPLocation returns a new IPLocation
func NewIPLocation() IPLocation {
	return IPLocation{
		IPStart:        types.NewUInt32(0),
		CountryCode:    types.NewString(""),
		CountryName:    types.NewString(""),
		RegionCode:     types.NewString(""),
		RegionName:     types.NewString(""),
		CityName:       types.NewString(""),
		Latitude:       types.NewFloat(0),
		Longitude:      types.NewFloat(0),
		CountryGeoCode: types.NewUInt32(0),
		RegionGeoCode:  types.NewUInt32(0),
		CityGeoCode:    types.NewUInt32(0),
	}
}
