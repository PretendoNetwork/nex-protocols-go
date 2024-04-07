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
	Language           *types.String
	Offset             *types.PrimitiveU32
	Size               *types.PrimitiveU32
	IsBalanceAvailable *types.PrimitiveBool
	UniqueID           *types.PrimitiveU32
	Platform           *types.PrimitiveU8 // * Revision 1
}

// WriteTo writes the ServiceItemListServiceItemParam to the given writable
func (silsip *ServiceItemListServiceItemParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	silsip.Language.WriteTo(writable)
	silsip.Offset.WriteTo(writable)
	silsip.Size.WriteTo(writable)
	silsip.IsBalanceAvailable.WriteTo(writable)
	silsip.UniqueID.WriteTo(writable)

	if silsip.StructureVersion >= 1 {
		silsip.Platform.WriteTo(writable)
	}

	content := contentWritable.Bytes()

	silsip.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemListServiceItemParam from the given readable
func (silsip *ServiceItemListServiceItemParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = silsip.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam header. %s", err.Error())
	}

	err = silsip.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Language. %s", err.Error())
	}

	err = silsip.Offset.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Offset. %s", err.Error())
	}

	err = silsip.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Size. %s", err.Error())
	}

	err = silsip.IsBalanceAvailable.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.IsBalanceAvailable. %s", err.Error())
	}

	err = silsip.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.UniqueID. %s", err.Error())
	}

	if silsip.StructureVersion >= 1 {
		err = silsip.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Platform. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemListServiceItemParam
func (silsip *ServiceItemListServiceItemParam) Copy() types.RVType {
	copied := NewServiceItemListServiceItemParam()

	copied.StructureVersion = silsip.StructureVersion
	copied.Language = silsip.Language.Copy().(*types.String)
	copied.Offset = silsip.Offset.Copy().(*types.PrimitiveU32)
	copied.Size = silsip.Size.Copy().(*types.PrimitiveU32)
	copied.IsBalanceAvailable = silsip.IsBalanceAvailable.Copy().(*types.PrimitiveBool)
	copied.UniqueID = silsip.UniqueID.Copy().(*types.PrimitiveU32)
	copied.Platform = silsip.Platform.Copy().(*types.PrimitiveU8)

	return copied
}

// Equals checks if the given ServiceItemListServiceItemParam contains the same data as the current ServiceItemListServiceItemParam
func (silsip *ServiceItemListServiceItemParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemListServiceItemParam); !ok {
		return false
	}

	other := o.(*ServiceItemListServiceItemParam)

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

// String returns the string representation of the ServiceItemListServiceItemParam
func (silsip *ServiceItemListServiceItemParam) String() string {
	return silsip.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemListServiceItemParam using the provided indentation level
func (silsip *ServiceItemListServiceItemParam) FormatToString(indentationLevel int) string {
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
func NewServiceItemListServiceItemParam() *ServiceItemListServiceItemParam {
	silsip := &ServiceItemListServiceItemParam{
		Language:           types.NewString(""),
		Offset:             types.NewPrimitiveU32(0),
		Size:               types.NewPrimitiveU32(0),
		IsBalanceAvailable: types.NewPrimitiveBool(false),
		UniqueID:           types.NewPrimitiveU32(0),
		Platform:           types.NewPrimitiveU8(0),
	}

	return silsip
}
