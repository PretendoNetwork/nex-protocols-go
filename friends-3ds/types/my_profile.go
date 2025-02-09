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
	Region              types.UInt8
	Country             types.UInt8
	Area                types.UInt8
	Language            types.UInt8
	Platform            types.UInt8
	LocalFriendCodeSeed types.UInt64
	MACAddress          types.String
	SerialNumber        types.String
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
	mp.LocalFriendCodeSeed.WriteTo(contentWritable)
	mp.MACAddress.WriteTo(contentWritable)
	mp.SerialNumber.WriteTo(contentWritable)

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

	err = mp.LocalFriendCodeSeed.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.LocalFriendCodeSeed. %s", err.Error())
	}

	err = mp.MACAddress.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.MACAddress. %s", err.Error())
	}

	err = mp.SerialNumber.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.SerialNumber. %s", err.Error())
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
	copied.LocalFriendCodeSeed = mp.LocalFriendCodeSeed.Copy().(types.UInt64)
	copied.MACAddress = mp.MACAddress.Copy().(types.String)
	copied.SerialNumber = mp.SerialNumber.Copy().(types.String)

	return copied
}

// Equals checks if the given MyProfile contains the same data as the current MyProfile
func (mp MyProfile) Equals(o types.RVType) bool {
	if _, ok := o.(MyProfile); !ok {
		return false
	}

	other := o.(MyProfile)

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

	if !mp.LocalFriendCodeSeed.Equals(other.LocalFriendCodeSeed) {
		return false
	}

	if !mp.MACAddress.Equals(other.MACAddress) {
		return false
	}

	return mp.SerialNumber.Equals(other.SerialNumber)
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
	b.WriteString(fmt.Sprintf("%sLocalFriendCodeSeed: %s,\n", indentationValues, mp.LocalFriendCodeSeed))
	b.WriteString(fmt.Sprintf("%sMACAddress: %s,\n", indentationValues, mp.MACAddress))
	b.WriteString(fmt.Sprintf("%sSerialNumber: %s,\n", indentationValues, mp.SerialNumber))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMyProfile returns a new MyProfile
func NewMyProfile() MyProfile {
	return MyProfile{
		Data:                types.NewData(),
		Region:              types.NewUInt8(0),
		Country:             types.NewUInt8(0),
		Area:                types.NewUInt8(0),
		Language:            types.NewUInt8(0),
		Platform:            types.NewUInt8(0),
		LocalFriendCodeSeed: types.NewUInt64(0),
		MACAddress:          types.NewString(""),
		SerialNumber:        types.NewString(""),
	}

}
