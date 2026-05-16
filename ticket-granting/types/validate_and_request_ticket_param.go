// Package types implements all the types used by the TicketGranting protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/ticket-granting/constants"
)

// ValidateAndRequestTicketParam is a type within the TicketGranting protocol
type ValidateAndRequestTicketParam struct {
	types.Structure
	PlatformType               types.UInt32
	Username                   types.String
	ExtraData                  types.DataHolder // * NullData or AuthenticationInfo
	IgnoreAPIVersionCheck      types.Bool
	APIVersionGeneral          types.UInt32
	APIVersionCustom           types.UInt32
	PlatformTypeForPlatformPID constants.PlatformType // * Only present on games with crossplay between Switch and 3DS/Wii U
}

// WriteTo writes the ValidateAndRequestTicketParam to the given writable
func (vartp ValidateAndRequestTicketParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	vartp.PlatformType.WriteTo(contentWritable)
	vartp.Username.WriteTo(contentWritable)
	vartp.ExtraData.WriteTo(contentWritable)
	vartp.IgnoreAPIVersionCheck.WriteTo(contentWritable)
	vartp.APIVersionGeneral.WriteTo(contentWritable)
	vartp.APIVersionCustom.WriteTo(contentWritable)

	// * This enum starts at 1, so if it's 0 we can safely
	// * assume it isn't being used. This field is only
	// * present on games with crossplay between Switch and
	// * 3DS/Wii U
	if vartp.PlatformTypeForPlatformPID != 0 {
		vartp.PlatformTypeForPlatformPID.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	vartp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ValidateAndRequestTicketParam from the given readable
func (vartp *ValidateAndRequestTicketParam) ExtractFrom(readable types.Readable) error {
	if err := vartp.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam header. %s", err.Error())
	}

	if err := vartp.PlatformType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.PlatformType. %s", err.Error())
	}

	if err := vartp.Username.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.Username. %s", err.Error())
	}

	if err := vartp.ExtraData.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.ExtraData. %s", err.Error())
	}

	if err := vartp.IgnoreAPIVersionCheck.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.IgnoreAPIVersionCheck. %s", err.Error())
	}

	if err := vartp.APIVersionGeneral.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.APIVersionGeneral. %s", err.Error())
	}

	if err := vartp.APIVersionCustom.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.APIVersionCustom. %s", err.Error())
	}

	// * This is a hack. This field is not based on a NEX
	// * version difference or a structure version update.
	// * It is only present on games with crossplay between
	// * Switch and 3DS/Wii U. Since we don't know whether
	// * or not a game has crossplay at this stage, this is
	// * the best we can do
	if readable.Remaining() != 0 {
		if err := vartp.PlatformTypeForPlatformPID.ExtractFrom(readable); err != nil {
			return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.PlatformTypeForPlatformPID. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of ValidateAndRequestTicketParam
func (vartp ValidateAndRequestTicketParam) Copy() types.RVType {
	copied := NewValidateAndRequestTicketParam()

	copied.StructureVersion = vartp.StructureVersion
	copied.PlatformType = vartp.PlatformType
	copied.Username = vartp.Username
	copied.ExtraData = vartp.ExtraData.Copy().(types.DataHolder)
	copied.IgnoreAPIVersionCheck = vartp.IgnoreAPIVersionCheck
	copied.APIVersionGeneral = vartp.APIVersionGeneral
	copied.APIVersionCustom = vartp.APIVersionCustom
	copied.PlatformTypeForPlatformPID = vartp.PlatformTypeForPlatformPID

	return copied
}

// Equals checks if the given ValidateAndRequestTicketParam contains the same data as the current ValidateAndRequestTicketParam
func (vartp ValidateAndRequestTicketParam) Equals(o types.RVType) bool {
	if _, ok := o.(ValidateAndRequestTicketParam); !ok {
		return false
	}

	other := o.(ValidateAndRequestTicketParam)

	if vartp.StructureVersion != other.StructureVersion {
		return false
	}

	if vartp.PlatformType != other.PlatformType {
		return false
	}

	if vartp.Username != other.Username {
		return false
	}

	if !vartp.ExtraData.Equals(other.ExtraData) {
		return false
	}

	if vartp.IgnoreAPIVersionCheck != other.IgnoreAPIVersionCheck {
		return false
	}

	if vartp.APIVersionGeneral != other.APIVersionGeneral {
		return false
	}

	if vartp.APIVersionCustom != other.APIVersionCustom {
		return false
	}

	return vartp.PlatformTypeForPlatformPID == other.PlatformTypeForPlatformPID
}

// CopyRef copies the current value of the ValidateAndRequestTicketParam
// and returns a pointer to the new copy
func (vartp ValidateAndRequestTicketParam) CopyRef() types.RVTypePtr {
	copied := vartp.Copy().(ValidateAndRequestTicketParam)
	return &copied
}

// Deref takes a pointer to the ValidateAndRequestTicketParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (vartp *ValidateAndRequestTicketParam) Deref() types.RVType {
	return *vartp
}

// String returns the string representation of the ValidateAndRequestTicketParam
func (vartp ValidateAndRequestTicketParam) String() string {
	return vartp.FormatToString(0)
}

// FormatToString pretty-prints the ValidateAndRequestTicketParam using the provided indentation level
func (vartp ValidateAndRequestTicketParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ValidateAndRequestTicketParam{\n")
	b.WriteString(fmt.Sprintf("%sPlatformType: %s,\n", indentationValues, vartp.PlatformType))
	b.WriteString(fmt.Sprintf("%sUsername: %s,\n", indentationValues, vartp.Username))
	b.WriteString(fmt.Sprintf("%sExtraData: %s,\n", indentationValues, vartp.ExtraData.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sIgnoreAPIVersionCheck: %s,\n", indentationValues, vartp.IgnoreAPIVersionCheck))
	b.WriteString(fmt.Sprintf("%sAPIVersionGeneral: %s,\n", indentationValues, vartp.APIVersionGeneral))

	// * This is a hack. This field is not based on a NEX
	// * version difference or a structure version update.
	// * It is only present on games with crossplay between
	// * Switch and 3DS/Wii U. Since we don't know whether
	// * or not a game has crossplay at this stage, this is
	// * the best we can do
	if vartp.PlatformTypeForPlatformPID != 0 {
		b.WriteString(fmt.Sprintf("%sAPIVersionCustom: %s,\n", indentationValues, vartp.APIVersionCustom))
		b.WriteString(fmt.Sprintf("%sPlatformTypeForPlatformPID: %s\n", indentationValues, vartp.PlatformTypeForPlatformPID))
	} else {
		b.WriteString(fmt.Sprintf("%sAPIVersionCustom: %s\n", indentationValues, vartp.APIVersionCustom))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewValidateAndRequestTicketParam returns a new ValidateAndRequestTicketParam
func NewValidateAndRequestTicketParam() ValidateAndRequestTicketParam {
	return ValidateAndRequestTicketParam{
		PlatformType:               types.NewUInt32(0),
		Username:                   types.NewString(""),
		ExtraData:                  types.NewDataHolder(),
		IgnoreAPIVersionCheck:      types.NewBool(false),
		APIVersionGeneral:          types.NewUInt32(0),
		APIVersionCustom:           types.NewUInt32(0),
		PlatformTypeForPlatformPID: constants.PlatformType(0), // TODO - What is the real default?
	}

}
