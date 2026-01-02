// Package types implements all the types used by the TicketGranting protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ValidateAndRequestTicketParam is a type within the TicketGranting protocol
type ValidateAndRequestTicketParam struct {
	types.Structure
	enableCrossplay            bool
	PlatformType               types.UInt32
	UserName                   types.String
	ExtraData                  types.DataHolder
	IgnoreAPIVersionCheck      types.Bool
	APIVersionGeneral          types.UInt32
	APIVersionCustom           types.UInt32
	PlatformTypeForPlatformPID types.UInt8
}

// WriteTo writes the ValidateAndRequestTicketParam to the given writable
func (rs ValidateAndRequestTicketParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rs.PlatformType.WriteTo(contentWritable)
	rs.UserName.WriteTo(contentWritable)
	rs.ExtraData.WriteTo(contentWritable)
	rs.IgnoreAPIVersionCheck.WriteTo(contentWritable)
	rs.APIVersionGeneral.WriteTo(contentWritable)
	rs.APIVersionCustom.WriteTo(contentWritable)

	if rs.enableCrossplay {
		rs.PlatformTypeForPlatformPID.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rs.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ValidateAndRequestTicketParam from the given readable
func (rs *ValidateAndRequestTicketParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = rs.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam header. %s", err.Error())
	}

	err = rs.PlatformType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.PlatformType. %s", err.Error())
	}

	err = rs.UserName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.UserName. %s", err.Error())
	}

	err = rs.ExtraData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.ExtraData. %s", err.Error())
	}

	err = rs.IgnoreAPIVersionCheck.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.IgnoreAPIVersionCheck. %s", err.Error())
	}

	err = rs.APIVersionGeneral.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.APIVersionGeneral. %s", err.Error())
	}

	err = rs.APIVersionCustom.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.APIVersionCustom. %s", err.Error())
	}

	if rs.enableCrossplay {
		err = rs.PlatformTypeForPlatformPID.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ValidateAndRequestTicketParam.PlatformTypeForPlatformPID. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of ValidateAndRequestTicketParam
func (rs ValidateAndRequestTicketParam) Copy() types.RVType {
	copied := NewValidateAndRequestTicketParam(rs.enableCrossplay)

	copied.StructureVersion = rs.StructureVersion
	copied.PlatformType = rs.PlatformType.Copy().(types.UInt32)
	copied.UserName = rs.UserName.Copy().(types.String)
	copied.ExtraData = rs.ExtraData.Copy().(types.DataHolder)
	copied.IgnoreAPIVersionCheck = rs.IgnoreAPIVersionCheck.Copy().(types.Bool)
	copied.APIVersionGeneral = rs.APIVersionGeneral.Copy().(types.UInt32)
	copied.APIVersionCustom = rs.APIVersionCustom.Copy().(types.UInt32)
	copied.PlatformTypeForPlatformPID = rs.PlatformTypeForPlatformPID.Copy().(types.UInt8)

	return copied
}

// Equals checks if the given ValidateAndRequestTicketParam contains the same data as the current ValidateAndRequestTicketParam
func (rs ValidateAndRequestTicketParam) Equals(o types.RVType) bool {
	if _, ok := o.(ValidateAndRequestTicketParam); !ok {
		return false
	}

	other := o.(ValidateAndRequestTicketParam)

	if rs.StructureVersion != other.StructureVersion {
		return false
	}

	if !rs.PlatformType.Equals(other.PlatformType) {
		return false
	}

	if !rs.UserName.Equals(other.UserName) {
		return false
	}

	if !rs.ExtraData.Equals(other.ExtraData) {
		return false
	}

	if !rs.IgnoreAPIVersionCheck.Equals(other.IgnoreAPIVersionCheck) {
		return false
	}

	if !rs.APIVersionGeneral.Equals(other.APIVersionGeneral) {
		return false
	}

	if !rs.APIVersionCustom.Equals(other.APIVersionCustom) {
		return false
	}

	return rs.PlatformTypeForPlatformPID.Equals(other.PlatformTypeForPlatformPID)
}

// CopyRef copies the current value of the ValidateAndRequestTicketParam
// and returns a pointer to the new copy
func (rs ValidateAndRequestTicketParam) CopyRef() types.RVTypePtr {
	copied := rs.Copy().(ValidateAndRequestTicketParam)
	return &copied
}

// Deref takes a pointer to the ValidateAndRequestTicketParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rs *ValidateAndRequestTicketParam) Deref() types.RVType {
	return *rs
}

// String returns the string representation of the ValidateAndRequestTicketParam
func (rs ValidateAndRequestTicketParam) String() string {
	return rs.FormatToString(0)
}

// FormatToString pretty-prints the ValidateAndRequestTicketParam using the provided indentation level
func (rs ValidateAndRequestTicketParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ValidateAndRequestTicketParam{\n")
	b.WriteString(fmt.Sprintf("%sPlatformType: %s,\n", indentationValues, rs.PlatformType))
	b.WriteString(fmt.Sprintf("%sUserName: %s,\n", indentationValues, rs.UserName))
	b.WriteString(fmt.Sprintf("%sExtraData: %s,\n", indentationValues, rs.ExtraData.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sIgnoreAPIVersionCheck: %s,\n", indentationValues, rs.IgnoreAPIVersionCheck))
	b.WriteString(fmt.Sprintf("%sAPIVersionGeneral: %s,\n", indentationValues, rs.APIVersionGeneral))
	b.WriteString(fmt.Sprintf("%sAPIVersionCustom: %s,\n", indentationValues, rs.APIVersionCustom))
	b.WriteString(fmt.Sprintf("%sPlatformTypeForPlatformPID: %s,\n", indentationValues, rs.PlatformTypeForPlatformPID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewValidateAndRequestTicketParam returns a new ValidateAndRequestTicketParam
func NewValidateAndRequestTicketParam(enableCrossplay bool) ValidateAndRequestTicketParam {
	return ValidateAndRequestTicketParam{
		enableCrossplay:            enableCrossplay,
		PlatformType:               types.NewUInt32(0),
		UserName:                   types.NewString(""),
		ExtraData:                  types.NewDataHolder(),
		IgnoreAPIVersionCheck:      types.NewBool(false),
		APIVersionGeneral:          types.NewUInt32(0),
		APIVersionCustom:           types.NewUInt32(0),
		PlatformTypeForPlatformPID: types.NewUInt8(0),
	}

}
