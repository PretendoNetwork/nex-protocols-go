package friends_3ds_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

type MyProfile struct {
	nex.Structure
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

// NewMyProfile returns a new MyProfile
func NewMyProfile() *MyProfile {
	return &MyProfile{}
}
