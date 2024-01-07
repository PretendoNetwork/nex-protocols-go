// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemGetBalanceParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetBalanceParam struct {
	types.Structure
	Language string
	UniqueID *types.PrimitiveU32
	Platform *types.PrimitiveU8 // * Revision 1
}

// ExtractFrom extracts the ServiceItemGetBalanceParam from the given readable
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemGetBalanceParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemGetBalanceParam header. %s", err.Error())
	}

	err = serviceItemGetBalanceParam.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam.Language from stream. %s", err.Error())
	}

	err = serviceItemGetBalanceParam.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam.UniqueID from stream. %s", err.Error())
	}

	if serviceItemGetBalanceParam.StructureVersion >= 1 {
	err = 	serviceItemGetBalanceParam.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// WriteTo writes the ServiceItemGetBalanceParam to the given writable
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemGetBalanceParam.Language.WriteTo(contentWritable)
	serviceItemGetBalanceParam.UniqueID.WriteTo(contentWritable)

	if serviceItemGetBalanceParam.StructureVersion >= 1 {
		serviceItemGetBalanceParam.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemGetBalanceParam
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) Copy() types.RVType {
	copied := NewServiceItemGetBalanceParam()

	copied.StructureVersion = serviceItemGetBalanceParam.StructureVersion

	copied.Language = serviceItemGetBalanceParam.Language
	copied.UniqueID = serviceItemGetBalanceParam.UniqueID
	copied.Platform = serviceItemGetBalanceParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetBalanceParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetBalanceParam)

	if serviceItemGetBalanceParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemGetBalanceParam.Language.Equals(other.Language) {
		return false
	}

	if !serviceItemGetBalanceParam.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !serviceItemGetBalanceParam.Platform.Equals(other.Platform) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) String() string {
	return serviceItemGetBalanceParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetBalanceParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemGetBalanceParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemGetBalanceParam.Language))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemGetBalanceParam.UniqueID))

	if serviceItemGetBalanceParam.StructureVersion >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemGetBalanceParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetBalanceParam returns a new ServiceItemGetBalanceParam
func NewServiceItemGetBalanceParam() *ServiceItemGetBalanceParam {
	return &ServiceItemGetBalanceParam{}
}
