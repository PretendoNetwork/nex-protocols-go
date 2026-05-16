// Package types implements all the types used by the NintendoNotifications protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NintendoNotificationEventProfile is a type within the NintendoNotifications protocol
type NintendoNotificationEventProfile struct {
	types.Structure
	types.Data
	Region   types.UInt8
	Country  types.UInt8
	Area     types.UInt8
	Language types.UInt8
	Platform types.UInt8
}

// WriteTo writes the NintendoNotificationEventProfile to the given writable
func (nnep NintendoNotificationEventProfile) WriteTo(writable types.Writable) {
	nnep.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	nnep.Region.WriteTo(contentWritable)
	nnep.Country.WriteTo(contentWritable)
	nnep.Area.WriteTo(contentWritable)
	nnep.Language.WriteTo(contentWritable)
	nnep.Platform.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nnep.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NintendoNotificationEventProfile from the given readable
func (nnep *NintendoNotificationEventProfile) ExtractFrom(readable types.Readable) error {
	if err := nnep.Data.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventProfile.Data. %s", err.Error())
	}

	if err := nnep.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventProfile header. %s", err.Error())
	}

	if err := nnep.Region.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventProfile.Region. %s", err.Error())
	}

	if err := nnep.Country.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventProfile.Country. %s", err.Error())
	}

	if err := nnep.Area.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventProfile.Area. %s", err.Error())
	}

	if err := nnep.Language.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventProfile.Language. %s", err.Error())
	}

	if err := nnep.Platform.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract NintendoNotificationEventProfile.Platform. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoNotificationEventProfile
func (nnep NintendoNotificationEventProfile) Copy() types.RVType {
	copied := NewNintendoNotificationEventProfile()

	copied.StructureVersion = nnep.StructureVersion
	copied.Data = nnep.Data.Copy().(types.Data)
	copied.Region = nnep.Region
	copied.Country = nnep.Country
	copied.Area = nnep.Area
	copied.Language = nnep.Language
	copied.Platform = nnep.Platform

	return copied
}

// Equals checks if the given NintendoNotificationEventProfile contains the same data as the current NintendoNotificationEventProfile
func (nnep NintendoNotificationEventProfile) Equals(o types.RVType) bool {
	if _, ok := o.(NintendoNotificationEventProfile); !ok {
		return false
	}

	other := o.(NintendoNotificationEventProfile)

	if nnep.StructureVersion != other.StructureVersion {
		return false
	}

	if !nnep.Data.Equals(other.Data) {
		return false
	}

	if nnep.Region != other.Region {
		return false
	}

	if nnep.Country != other.Country {
		return false
	}

	if nnep.Area != other.Area {
		return false
	}

	if nnep.Language != other.Language {
		return false
	}

	return nnep.Platform == other.Platform
}

// CopyRef copies the current value of the NintendoNotificationEventProfile
// and returns a pointer to the new copy
func (nnep NintendoNotificationEventProfile) CopyRef() types.RVTypePtr {
	copied := nnep.Copy().(NintendoNotificationEventProfile)
	return &copied
}

// Deref takes a pointer to the NintendoNotificationEventProfile
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (nnep *NintendoNotificationEventProfile) Deref() types.RVType {
	return *nnep
}

// String returns the string representation of the NintendoNotificationEventProfile
func (nnep NintendoNotificationEventProfile) String() string {
	return nnep.FormatToString(0)
}

// FormatToString pretty-prints the NintendoNotificationEventProfile using the provided indentation level
func (nnep NintendoNotificationEventProfile) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NintendoNotificationEventProfile{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, nnep.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sRegion: %s,\n", indentationValues, nnep.Region))
	b.WriteString(fmt.Sprintf("%sCountry: %s,\n", indentationValues, nnep.Country))
	b.WriteString(fmt.Sprintf("%sArea: %s,\n", indentationValues, nnep.Area))
	b.WriteString(fmt.Sprintf("%sLanguage: %s,\n", indentationValues, nnep.Language))
	b.WriteString(fmt.Sprintf("%sPlatform: %s\n", indentationValues, nnep.Platform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoNotificationEventProfile returns a new NintendoNotificationEventProfile
func NewNintendoNotificationEventProfile() NintendoNotificationEventProfile {
	return NintendoNotificationEventProfile{
		Region:   types.NewUInt8(0),
		Country:  types.NewUInt8(0),
		Area:     types.NewUInt8(0),
		Language: types.NewUInt8(0),
		Platform: types.NewUInt8(0),
	}

}
