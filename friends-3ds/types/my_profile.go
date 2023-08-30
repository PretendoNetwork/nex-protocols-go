// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// MyProfile is a data structure used by the Friends 3DS protocol to hold user profile information
type MyProfile struct {
	nex.Structure
	*nex.Data
	Region   uint8
	Country  uint8
	Area     uint8
	Language uint8
	Platform uint8
	Unknown1 uint64
	Unknown2 string
	Unknown3 string
}

// ExtractFromStream extracts a MyProfile from a stream
func (myProfile *MyProfile) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	myProfile.Region, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Region. %s", err.Error())
	}

	myProfile.Country, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Country. %s", err.Error())
	}

	myProfile.Area, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Area. %s", err.Error())
	}

	myProfile.Language, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Language. %s", err.Error())
	}

	myProfile.Platform, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Platform. %s", err.Error())
	}

	myProfile.Unknown1, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Unknown1. %s", err.Error())
	}

	myProfile.Unknown2, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Unknown2. %s", err.Error())
	}

	myProfile.Unknown3, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract MyProfile.Unknown3. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MyProfile
func (myProfile *MyProfile) Copy() nex.StructureInterface {
	copied := NewMyProfile()

	copied.SetStructureVersion(myProfile.StructureVersion())

	copied.Data = myProfile.ParentType().Copy().(*nex.Data)
	copied.SetParentType(copied.Data)

	copied.Region = myProfile.Region
	copied.Country = myProfile.Country
	copied.Area = myProfile.Area
	copied.Language = myProfile.Language
	copied.Platform = myProfile.Platform
	copied.Unknown1 = myProfile.Unknown1
	copied.Unknown2 = myProfile.Unknown2
	copied.Unknown3 = myProfile.Unknown3

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (myProfile *MyProfile) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MyProfile)

	if myProfile.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !myProfile.ParentType().Equals(other.ParentType()) {
		return false
	}

	if myProfile.Region != other.Region {
		return false
	}

	if myProfile.Country != other.Country {
		return false
	}

	if myProfile.Area != other.Area {
		return false
	}

	if myProfile.Language != other.Language {
		return false
	}

	if myProfile.Platform != other.Platform {
		return false
	}

	if myProfile.Unknown1 != other.Unknown1 {
		return false
	}

	if myProfile.Unknown2 != other.Unknown2 {
		return false
	}

	if myProfile.Unknown3 != other.Unknown3 {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (myProfile *MyProfile) String() string {
	return myProfile.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (myProfile *MyProfile) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MyProfile{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, myProfile.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sRegion: %d,\n", indentationValues, myProfile.Region))
	b.WriteString(fmt.Sprintf("%sCountry: %d,\n", indentationValues, myProfile.Country))
	b.WriteString(fmt.Sprintf("%sArea: %d,\n", indentationValues, myProfile.Area))
	b.WriteString(fmt.Sprintf("%sLanguage: %d,\n", indentationValues, myProfile.Language))
	b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, myProfile.Platform))
	b.WriteString(fmt.Sprintf("%sUnknown1: %d,\n", indentationValues, myProfile.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %q,\n", indentationValues, myProfile.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %q\n", indentationValues, myProfile.Unknown3))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMyProfile returns a new MyProfile
func NewMyProfile() *MyProfile {
	return &MyProfile{}
}
