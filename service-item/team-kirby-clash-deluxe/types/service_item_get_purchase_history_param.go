// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemGetPurchaseHistoryParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetPurchaseHistoryParam struct {
	types.Structure
	Language string
	Offset   *types.PrimitiveU32
	Size     *types.PrimitiveU32
	UniqueID *types.PrimitiveU32
	Platform *types.PrimitiveU8
}

// ExtractFrom extracts the ServiceItemGetPurchaseHistoryParam from the given readable
func (serviceItemGetPurchaseHistoryParam *ServiceItemGetPurchaseHistoryParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemGetPurchaseHistoryParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemGetPurchaseHistoryParam header. %s", err.Error())
	}

	err = serviceItemGetPurchaseHistoryParam.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.Language from stream. %s", err.Error())
	}

	err = serviceItemGetPurchaseHistoryParam.Offset.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.Offset from stream. %s", err.Error())
	}

	err = serviceItemGetPurchaseHistoryParam.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.Size from stream. %s", err.Error())
	}

	err = serviceItemGetPurchaseHistoryParam.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.UniqueID from stream. %s", err.Error())
	}

	if serviceItemGetPurchaseHistoryParam.StructureVersion >= 1 {
	err = 	serviceItemGetPurchaseHistoryParam.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// WriteTo writes the ServiceItemGetPurchaseHistoryParam to the given writable
func (serviceItemGetPurchaseHistoryParam *ServiceItemGetPurchaseHistoryParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemGetPurchaseHistoryParam.Language.WriteTo(contentWritable)
	serviceItemGetPurchaseHistoryParam.Offset.WriteTo(contentWritable)
	serviceItemGetPurchaseHistoryParam.Size.WriteTo(contentWritable)
	serviceItemGetPurchaseHistoryParam.UniqueID.WriteTo(contentWritable)

	if serviceItemGetPurchaseHistoryParam.StructureVersion >= 1 {
		serviceItemGetPurchaseHistoryParam.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemGetPurchaseHistoryParam
func (serviceItemGetPurchaseHistoryParam *ServiceItemGetPurchaseHistoryParam) Copy() types.RVType {
	copied := NewServiceItemGetPurchaseHistoryParam()

	copied.StructureVersion = serviceItemGetPurchaseHistoryParam.StructureVersion

	copied.Language = serviceItemGetPurchaseHistoryParam.Language
	copied.Offset = serviceItemGetPurchaseHistoryParam.Offset
	copied.Size = serviceItemGetPurchaseHistoryParam.Size
	copied.UniqueID = serviceItemGetPurchaseHistoryParam.UniqueID
	copied.Platform = serviceItemGetPurchaseHistoryParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetPurchaseHistoryParam *ServiceItemGetPurchaseHistoryParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetPurchaseHistoryParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetPurchaseHistoryParam)

	if serviceItemGetPurchaseHistoryParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemGetPurchaseHistoryParam.Language.Equals(other.Language) {
		return false
	}

	if !serviceItemGetPurchaseHistoryParam.Offset.Equals(other.Offset) {
		return false
	}

	if !serviceItemGetPurchaseHistoryParam.Size.Equals(other.Size) {
		return false
	}

	if !serviceItemGetPurchaseHistoryParam.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !serviceItemGetPurchaseHistoryParam.Platform.Equals(other.Platform) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetPurchaseHistoryParam *ServiceItemGetPurchaseHistoryParam) String() string {
	return serviceItemGetPurchaseHistoryParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetPurchaseHistoryParam *ServiceItemGetPurchaseHistoryParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetPurchaseHistoryParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemGetPurchaseHistoryParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemGetPurchaseHistoryParam.Language))
	b.WriteString(fmt.Sprintf("%sOffset: %d,\n", indentationValues, serviceItemGetPurchaseHistoryParam.Offset))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, serviceItemGetPurchaseHistoryParam.Size))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemGetPurchaseHistoryParam.UniqueID))

	if serviceItemGetPurchaseHistoryParam.StructureVersion >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemGetPurchaseHistoryParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetPurchaseHistoryParam returns a new ServiceItemGetPurchaseHistoryParam
func NewServiceItemGetPurchaseHistoryParam() *ServiceItemGetPurchaseHistoryParam {
	return &ServiceItemGetPurchaseHistoryParam{}
}
