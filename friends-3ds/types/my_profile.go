// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MyProfile is a type within the Friends3DS protocol
type MyProfile struct {
	types.Structure
	types.Data
	Region   types.UInt8
	Country  types.UInt8
	Area     types.UInt8
	Language types.UInt8
	Platform types.UInt8
	Unknown1 types.UInt64
	Unknown2 types.String
	Unknown3 types.String
}

// WriteTo writes the MyProfile to the given writable
func (mp MyProfile) WriteTo(writable types.Writable) {
	mp.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	mp.Region.WriteTo(contentWritable)
	mp.Country.WriteTo(contentWritable)
	mp.Area.WriteTo(contentWritable)
	mp.Language.WriteTo(contentWritable)
	mp.Platform.WriteTo(contentWritable)
	mp.Unknown1.WriteTo(contentWritable)
	mp.Unknown2.WriteTo(contentWritable)
	mp.Unknown3.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MyProfile from the given readable
func (mp *MyProfile) ExtractFrom(readable types.Readable) error {
	var err error

	err = mp.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Data. %s", err.Error())
	}

	err = mp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile header. %s", err.Error())
	}

	err = mp.Region.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Region. %s", err.Error())
	}

	err = mp.Country.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Country. %s", err.Error())
	}

	err = mp.Area.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Area. %s", err.Error())
	}

	err = mp.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Language. %s", err.Error())
	}

	err = mp.Platform.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Platform. %s", err.Error())
	}

	err = mp.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Unknown1. %s", err.Error())
	}

	err = mp.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Unknown2. %s", err.Error())
	}

	err = mp.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Unknown3. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MyProfile
func (mp MyProfile) Copy() types.RVType {
	copied := NewMyProfile()

	copied.StructureVersion = mp.StructureVersion
	copied.Data = mp.Data.Copy().(types.Data)
	copied.Region = mp.Region.Copy().(types.UInt8)
	copied.Country = mp.Country.Copy().(types.UInt8)
	copied.Area = mp.Area.Copy().(types.UInt8)
	copied.Language = mp.Language.Copy().(types.UInt8)
	copied.Platform = mp.Platform.Copy().(types.UInt8)
	copied.Unknown1 = mp.Unknown1.Copy().(types.UInt64)
	copied.Unknown2 = mp.Unknown2.Copy().(types.String)
	copied.Unknown3 = mp.Unknown3.Copy().(types.String)

	return copied
}

// Equals checks if the given MyProfile contains the same data as the current MyProfile
func (mp MyProfile) Equals(o types.RVType) bool {
	if _, ok := o.(*MyProfile); !ok {
		return false
	}

	other := o.(*MyProfile)

	if mp.StructureVersion != other.StructureVersion {
		return false
	}

	if !mp.Data.Equals(other.Data) {
		return false
	}

	if !mp.Region.Equals(other.Region) {
		return false
	}

	if !mp.Country.Equals(other.Country) {
		return false
	}

	if !mp.Area.Equals(other.Area) {
		return false
	}

	if !mp.Language.Equals(other.Language) {
		return false
	}

	if !mp.Platform.Equals(other.Platform) {
		return false
	}

	if !mp.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !mp.Unknown2.Equals(other.Unknown2) {
		return false
	}

	return mp.Unknown3.Equals(other.Unknown3)
}

// CopyRef copies the current value of the MyProfile
// and returns a pointer to the new copy
func (mp MyProfile) CopyRef() types.RVTypePtr {
	copied := mp.Copy().(MyProfile)
	return &copied
}

// Deref takes a pointer to the MyProfile
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (mp *MyProfile) Deref() types.RVType {
	return *mp
}

// String returns the string representation of the MyProfile
func (mp MyProfile) String() string {
	return mp.FormatToString(0)
}

// FormatToString pretty-prints the MyProfile using the provided indentation level
func (mp MyProfile) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MyProfile{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, mp.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sRegion: %s,\n", indentationValues, mp.Region))
	b.WriteString(fmt.Sprintf("%sCountry: %s,\n", indentationValues, mp.Country))
	b.WriteString(fmt.Sprintf("%sArea: %s,\n", indentationValues, mp.Area))
	b.WriteString(fmt.Sprintf("%sLanguage: %s,\n", indentationValues, mp.Language))
	b.WriteString(fmt.Sprintf("%sPlatform: %s,\n", indentationValues, mp.Platform))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, mp.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, mp.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %s,\n", indentationValues, mp.Unknown3))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMyProfile returns a new MyProfile
func NewMyProfile() MyProfile {
	return MyProfile{
		Data:     types.NewData(),
		Region:   types.NewUInt8(0),
		Country:  types.NewUInt8(0),
		Area:     types.NewUInt8(0),
		Language: types.NewUInt8(0),
		Platform: types.NewUInt8(0),
		Unknown1: types.NewUInt64(0),
		Unknown2: types.NewString(""),
		Unknown3: types.NewString(""),
	}

}
