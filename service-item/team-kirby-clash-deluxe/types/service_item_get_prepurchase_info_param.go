// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemGetPrepurchaseInfoParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetPrepurchaseInfoParam struct {
	types.Structure
	ItemCode    string
	ReferenceID string
	Limitation  *ServiceItemLimitation
	Language    string
	UniqueID    *types.PrimitiveU32
}

// ExtractFrom extracts the ServiceItemGetPrepurchaseInfoParam from the given readable
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemGetPrepurchaseInfoParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemGetPrepurchaseInfoParam header. %s", err.Error())
	}

	err = serviceItemGetPrepurchaseInfoParam.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.ItemCode from stream. %s", err.Error())
	}

	err = serviceItemGetPrepurchaseInfoParam.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.ReferenceID from stream. %s", err.Error())
	}

	err = serviceItemGetPrepurchaseInfoParam.Limitation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.Limitation from stream. %s", err.Error())
	}

	err = serviceItemGetPrepurchaseInfoParam.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.Language from stream. %s", err.Error())
	}

	err = serviceItemGetPrepurchaseInfoParam.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.UniqueID from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemGetPrepurchaseInfoParam to the given writable
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemGetPrepurchaseInfoParam.ItemCode.WriteTo(contentWritable)
	serviceItemGetPrepurchaseInfoParam.ReferenceID.WriteTo(contentWritable)
	serviceItemGetPrepurchaseInfoParam.Limitation.WriteTo(contentWritable)
	serviceItemGetPrepurchaseInfoParam.Language.WriteTo(contentWritable)
	serviceItemGetPrepurchaseInfoParam.UniqueID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemGetPrepurchaseInfoParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemGetPrepurchaseInfoParam
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) Copy() types.RVType {
	copied := NewServiceItemGetPrepurchaseInfoParam()

	copied.StructureVersion = serviceItemGetPrepurchaseInfoParam.StructureVersion

	copied.ItemCode = serviceItemGetPrepurchaseInfoParam.ItemCode
	copied.ReferenceID = serviceItemGetPrepurchaseInfoParam.ReferenceID
	copied.Limitation = serviceItemGetPrepurchaseInfoParam.Limitation.Copy().(*ServiceItemLimitation)
	copied.Language = serviceItemGetPrepurchaseInfoParam.Language
	copied.UniqueID = serviceItemGetPrepurchaseInfoParam.UniqueID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetPrepurchaseInfoParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetPrepurchaseInfoParam)

	if serviceItemGetPrepurchaseInfoParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemGetPrepurchaseInfoParam.ItemCode.Equals(other.ItemCode) {
		return false
	}

	if !serviceItemGetPrepurchaseInfoParam.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !serviceItemGetPrepurchaseInfoParam.Limitation.Equals(other.Limitation) {
		return false
	}

	if !serviceItemGetPrepurchaseInfoParam.Language.Equals(other.Language) {
		return false
	}

	if !serviceItemGetPrepurchaseInfoParam.UniqueID.Equals(other.UniqueID) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) String() string {
	return serviceItemGetPrepurchaseInfoParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetPrepurchaseInfoParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemGetPrepurchaseInfoParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sItemCode: %q,\n", indentationValues, serviceItemGetPrepurchaseInfoParam.ItemCode))
	b.WriteString(fmt.Sprintf("%sReferenceID: %q,\n", indentationValues, serviceItemGetPrepurchaseInfoParam.ReferenceID))

	if serviceItemGetPrepurchaseInfoParam.Limitation != nil {
		b.WriteString(fmt.Sprintf("%sLimitation: %s\n", indentationValues, serviceItemGetPrepurchaseInfoParam.Limitation.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sLimitation: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemGetPrepurchaseInfoParam.Language))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemGetPrepurchaseInfoParam.UniqueID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetPrepurchaseInfoParam returns a new ServiceItemGetPrepurchaseInfoParam
func NewServiceItemGetPrepurchaseInfoParam() *ServiceItemGetPrepurchaseInfoParam {
	return &ServiceItemGetPrepurchaseInfoParam{}
}
