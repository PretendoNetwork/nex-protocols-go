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
	ipStart        types.UInt32
	countryCode    types.String
	countryName    types.String
	regionCode     types.String
	regionName     types.String
	cityName       types.String
	latitude       types.Float
	longitude      types.Float
	countryGeoCode types.UInt32
	regionGeoCode  types.UInt32
	cityGeoCode    types.UInt32
}

// WriteTo writes the IPLocation to the given writable
func (ipl IPLocation) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ipl.ipStart.WriteTo(contentWritable)
	ipl.countryCode.WriteTo(contentWritable)
	ipl.countryName.WriteTo(contentWritable)
	ipl.regionCode.WriteTo(contentWritable)
	ipl.regionName.WriteTo(contentWritable)
	ipl.cityName.WriteTo(contentWritable)
	ipl.latitude.WriteTo(contentWritable)
	ipl.longitude.WriteTo(contentWritable)
	ipl.countryGeoCode.WriteTo(contentWritable)
	ipl.regionGeoCode.WriteTo(contentWritable)
	ipl.cityGeoCode.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ipl.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

func (ipl *IPLocation) ExtractFrom(readable types.Readable) error {
	if err := ipl.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation header. %s", err.Error())
	}

	if err := ipl.ipStart.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_ipStart. %s", err.Error())
	}

	if err := ipl.countryCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_countryCode. %s", err.Error())
	}

	if err := ipl.countryName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_countryName. %s", err.Error())
	}

	if err := ipl.regionCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_regionCode. %s", err.Error())
	}

	if err := ipl.regionName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_regionName. %s", err.Error())
	}

	if err := ipl.latitude.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_latitude. %s", err.Error())
	}

	if err := ipl.longitude.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_longitude. %s", err.Error())
	}

	if err := ipl.countryGeoCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_countryGeoCode. %s", err.Error())
	}

	if err := ipl.regionGeoCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_regionGeoCode. %s", err.Error())
	}

	if err := ipl.cityGeoCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract IPLocation.m_cityGeoCode. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instanced of IPLocation
func (ipl IPLocation) Copy() types.RVType {
	copied := NewIPLocation()

	copied.StructureVersion = ipl.StructureVersion
	copied.ipStart = ipl.ipStart
	copied.countryCode = ipl.countryCode
	copied.countryName = ipl.countryName
	copied.regionCode = ipl.regionCode
	copied.regionName = ipl.regionName
	copied.cityName = ipl.cityName
	copied.latitude = ipl.latitude
	copied.longitude = ipl.longitude
	copied.countryGeoCode = ipl.countryGeoCode
	copied.regionGeoCode = ipl.regionGeoCode
	copied.cityGeoCode = ipl.cityGeoCode

	return copied
}

// Equal checks if the given IPLocation contains the same data as the current IPLocation
func (ipl IPLocation) Equals(o types.RVType) bool {
	if _, ok := o.(IPLocation); !ok {
		return false
	}

	other := o.(IPLocation)

	if ipl.StructureVersion != other.StructureVersion {
		return false
	}

	if ipl.ipStart != other.ipStart {
		return false
	}

	if ipl.countryCode != other.countryCode {
		return false
	}

	if ipl.countryName != other.countryName {
		return false
	}

	if ipl.regionCode != other.regionName {
		return false
	}

	if ipl.cityName != other.cityName {
		return false
	}

	if ipl.latitude != other.latitude {
		return false
	}

	if ipl.longitude != other.longitude {
		return false
	}

	if ipl.countryGeoCode != other.countryGeoCode {
		return false
	}

	if ipl.regionGeoCode != other.regionGeoCode {
		return false
	}

	return ipl.cityGeoCode == other.cityGeoCode
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

func (ipl IPLocation) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("IPlocation{\n}")
	fmt.Fprintf(&b, "%sm_ipStart: %s,\n", indentationValues, ipl.ipStart)
	fmt.Fprintf(&b, "%sm_countryCode: %s,\n", indentationValues, ipl.countryCode)
	fmt.Fprintf(&b, "%sm_countryName: %s,\n", indentationValues, ipl.countryName)
	fmt.Fprintf(&b, "%sm_regionCode: %s,\n", indentationValues, ipl.regionCode)
	fmt.Fprintf(&b, "%sm_regionName: %s,\n", indentationValues, ipl.regionName)
	fmt.Fprintf(&b, "%sm_cityName: %s,\n", indentationValues, ipl.cityName)
	fmt.Fprintf(&b, "%sm_latitude: %s,\n", indentationValues, ipl.latitude)
	fmt.Fprintf(&b, "%sm_longitude: %s,\n", indentationValues, ipl.longitude)
	fmt.Fprintf(&b, "%sm_countryGeoCode: %s,\n", indentationValues, ipl.countryGeoCode)
	fmt.Fprintf(&b, "%sm_regionGeoCode: %s,\n", indentationValues, ipl.regionGeoCode)
	fmt.Fprintf(&b, "%sm_cityGeoCode: %s,\n", indentationValues, ipl.cityGeoCode)

	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewIPLocation returns a new IPLocation
func NewIPLocation() IPLocation {
	return IPLocation{
		ipStart:        types.NewUInt32(0),
		countryCode:    types.NewString(""),
		countryName:    types.NewString(""),
		regionCode:     types.NewString(""),
		regionName:     types.NewString(""),
		cityName:       types.NewString(""),
		latitude:       types.NewFloat(0),
		longitude:      types.NewFloat(0),
		countryGeoCode: types.NewUInt32(0),
		regionGeoCode:  types.NewUInt32(0),
		cityGeoCode:    types.NewUInt32(0),
	}

}
