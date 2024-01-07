// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemListServiceItemParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemListServiceItemParam struct {
	types.Structure
	Language           string
	Offset             *types.PrimitiveU32
	Size               *types.PrimitiveU32
	IsBalanceAvailable *types.PrimitiveBool
	UniqueID           *types.PrimitiveU32
	Platform           *types.PrimitiveU8 // * Revision 1
}

// ExtractFrom extracts the ServiceItemListServiceItemParam from the given readable
func (serviceItemListServiceItemParam *ServiceItemListServiceItemParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemListServiceItemParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemListServiceItemParam header. %s", err.Error())
	}

	err = serviceItemListServiceItemParam.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Language from stream. %s", err.Error())
	}

	err = serviceItemListServiceItemParam.Offset.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Offset from stream. %s", err.Error())
	}

	err = serviceItemListServiceItemParam.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Size from stream. %s", err.Error())
	}

	err = serviceItemListServiceItemParam.IsBalanceAvailable.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.IsBalanceAvailable from stream. %s", err.Error())
	}

	err = serviceItemListServiceItemParam.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.UniqueID from stream. %s", err.Error())
	}

	if serviceItemListServiceItemParam.StructureVersion >= 1 {
	err = 	serviceItemListServiceItemParam.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// WriteTo writes the ServiceItemListServiceItemParam to the given writable
func (serviceItemListServiceItemParam *ServiceItemListServiceItemParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemListServiceItemParam.Language.WriteTo(contentWritable)
	serviceItemListServiceItemParam.Offset.WriteTo(contentWritable)
	serviceItemListServiceItemParam.Size.WriteTo(contentWritable)
	serviceItemListServiceItemParam.IsBalanceAvailable.WriteTo(contentWritable)
	serviceItemListServiceItemParam.UniqueID.WriteTo(contentWritable)

	if serviceItemListServiceItemParam.StructureVersion >= 1 {
		serviceItemListServiceItemParam.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemListServiceItemParam
func (serviceItemListServiceItemParam *ServiceItemListServiceItemParam) Copy() types.RVType {
	copied := NewServiceItemListServiceItemParam()

	copied.StructureVersion = serviceItemListServiceItemParam.StructureVersion

	copied.Language = serviceItemListServiceItemParam.Language
	copied.Offset = serviceItemListServiceItemParam.Offset
	copied.Size = serviceItemListServiceItemParam.Size
	copied.IsBalanceAvailable = serviceItemListServiceItemParam.IsBalanceAvailable
	copied.UniqueID = serviceItemListServiceItemParam.UniqueID
	copied.Platform = serviceItemListServiceItemParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemListServiceItemParam *ServiceItemListServiceItemParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemListServiceItemParam); !ok {
		return false
	}

	other := o.(*ServiceItemListServiceItemParam)

	if serviceItemListServiceItemParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemListServiceItemParam.Language.Equals(other.Language) {
		return false
	}

	if !serviceItemListServiceItemParam.Offset.Equals(other.Offset) {
		return false
	}

	if !serviceItemListServiceItemParam.Size.Equals(other.Size) {
		return false
	}

	if !serviceItemListServiceItemParam.IsBalanceAvailable.Equals(other.IsBalanceAvailable) {
		return false
	}

	if !serviceItemListServiceItemParam.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !serviceItemListServiceItemParam.Platform.Equals(other.Platform) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemListServiceItemParam *ServiceItemListServiceItemParam) String() string {
	return serviceItemListServiceItemParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemListServiceItemParam *ServiceItemListServiceItemParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemListServiceItemParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemListServiceItemParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemListServiceItemParam.Language))
	b.WriteString(fmt.Sprintf("%sOffset: %d,\n", indentationValues, serviceItemListServiceItemParam.Offset))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, serviceItemListServiceItemParam.Size))
	b.WriteString(fmt.Sprintf("%sIsBalanceAvailable: %t,\n", indentationValues, serviceItemListServiceItemParam.IsBalanceAvailable))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemListServiceItemParam.UniqueID))

	if serviceItemListServiceItemParam.StructureVersion >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemListServiceItemParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemListServiceItemParam returns a new ServiceItemListServiceItemParam
func NewServiceItemListServiceItemParam() *ServiceItemListServiceItemParam {
	return &ServiceItemListServiceItemParam{}
}
