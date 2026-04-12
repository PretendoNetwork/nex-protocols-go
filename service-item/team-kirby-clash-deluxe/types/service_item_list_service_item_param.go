// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemListServiceItemParam is a type within the ServiceItem protocol
type ServiceItemListServiceItemParam struct {
	types.Structure
	Language           types.String
	Offset             types.UInt32
	Size               types.UInt32
	IsBalanceAvailable types.Bool
	UniqueID           types.UInt32
	Platform           types.UInt8 // * Revision 1
}

// WriteTo writes the ServiceItemListServiceItemParam to the given writable
func (silsip ServiceItemListServiceItemParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	silsip.Language.WriteTo(contentWritable)
	silsip.Offset.WriteTo(contentWritable)
	silsip.Size.WriteTo(contentWritable)
	silsip.IsBalanceAvailable.WriteTo(contentWritable)
	silsip.UniqueID.WriteTo(contentWritable)

	if silsip.StructureVersion >= 1 {
		silsip.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	silsip.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemListServiceItemParam from the given readable
func (silsip *ServiceItemListServiceItemParam) ExtractFrom(readable types.Readable) error {
	if err := silsip.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam header. %s", err.Error())
	}

	if err := silsip.Language.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Language. %s", err.Error())
	}

	if err := silsip.Offset.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Offset. %s", err.Error())
	}

	if err := silsip.Size.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Size. %s", err.Error())
	}

	if err := silsip.IsBalanceAvailable.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.IsBalanceAvailable. %s", err.Error())
	}

	if err := silsip.UniqueID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.UniqueID. %s", err.Error())
	}

	if silsip.StructureVersion >= 1 {
		if err := silsip.Platform.ExtractFrom(readable); err != nil {
			return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Platform. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemListServiceItemParam
func (silsip ServiceItemListServiceItemParam) Copy() types.RVType {
	copied := NewServiceItemListServiceItemParam()

	copied.StructureVersion = silsip.StructureVersion
	copied.Language = silsip.Language.Copy().(types.String)
	copied.Offset = silsip.Offset.Copy().(types.UInt32)
	copied.Size = silsip.Size.Copy().(types.UInt32)
	copied.IsBalanceAvailable = silsip.IsBalanceAvailable.Copy().(types.Bool)
	copied.UniqueID = silsip.UniqueID.Copy().(types.UInt32)
	copied.Platform = silsip.Platform.Copy().(types.UInt8)

	return copied
}

// Equals checks if the given ServiceItemListServiceItemParam contains the same data as the current ServiceItemListServiceItemParam
func (silsip ServiceItemListServiceItemParam) Equals(o types.RVType) bool {
	if _, ok := o.(ServiceItemListServiceItemParam); !ok {
		return false
	}

	other := o.(ServiceItemListServiceItemParam)

	if silsip.StructureVersion != other.StructureVersion {
		return false
	}

	if !silsip.Language.Equals(other.Language) {
		return false
	}

	if !silsip.Offset.Equals(other.Offset) {
		return false
	}

	if !silsip.Size.Equals(other.Size) {
		return false
	}

	if !silsip.IsBalanceAvailable.Equals(other.IsBalanceAvailable) {
		return false
	}

	if !silsip.UniqueID.Equals(other.UniqueID) {
		return false
	}

	return silsip.Platform.Equals(other.Platform)
}

// CopyRef copies the current value of the ServiceItemListServiceItemParam
// and returns a pointer to the new copy
func (silsip ServiceItemListServiceItemParam) CopyRef() types.RVTypePtr {
	copied := silsip.Copy().(ServiceItemListServiceItemParam)
	return &copied
}

// Deref takes a pointer to the ServiceItemListServiceItemParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (silsip *ServiceItemListServiceItemParam) Deref() types.RVType {
	return *silsip
}

// String returns the string representation of the ServiceItemListServiceItemParam
func (silsip ServiceItemListServiceItemParam) String() string {
	return silsip.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemListServiceItemParam using the provided indentation level
func (silsip ServiceItemListServiceItemParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemListServiceItemParam{\n")
	b.WriteString(fmt.Sprintf("%sLanguage: %s,\n", indentationValues, silsip.Language))
	b.WriteString(fmt.Sprintf("%sOffset: %s,\n", indentationValues, silsip.Offset))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, silsip.Size))
	b.WriteString(fmt.Sprintf("%sIsBalanceAvailable: %s,\n", indentationValues, silsip.IsBalanceAvailable))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, silsip.UniqueID))
	b.WriteString(fmt.Sprintf("%sPlatform: %s,\n", indentationValues, silsip.Platform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemListServiceItemParam returns a new ServiceItemListServiceItemParam
func NewServiceItemListServiceItemParam() ServiceItemListServiceItemParam {
	return ServiceItemListServiceItemParam{
		Language:           types.NewString(""),
		Offset:             types.NewUInt32(0),
		Size:               types.NewUInt32(0),
		IsBalanceAvailable: types.NewBool(false),
		UniqueID:           types.NewUInt32(0),
		Platform:           types.NewUInt8(0),
	}

}
