// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemGetSupportIDParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetSupportIDParam struct {
	types.Structure
	UniqueID *types.PrimitiveU32
	Platform *types.PrimitiveU8
}

// ExtractFrom extracts the ServiceItemGetSupportIDParam from the given readable
func (serviceItemGetSupportIDParam *ServiceItemGetSupportIDParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemGetSupportIDParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemGetSupportIDParam header. %s", err.Error())
	}

	err = serviceItemGetSupportIDParam.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetSupportIDParam.UniqueID from stream. %s", err.Error())
	}

	if serviceItemGetSupportIDParam.StructureVersion >= 1 {
	err = 	serviceItemGetSupportIDParam.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemGetSupportIDParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// WriteTo writes the ServiceItemGetSupportIDParam to the given writable
func (serviceItemGetSupportIDParam *ServiceItemGetSupportIDParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemGetSupportIDParam.UniqueID.WriteTo(contentWritable)

	if serviceItemGetSupportIDParam.StructureVersion >= 1 {
		serviceItemGetSupportIDParam.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemGetSupportIDParam
func (serviceItemGetSupportIDParam *ServiceItemGetSupportIDParam) Copy() types.RVType {
	copied := NewServiceItemGetSupportIDParam()

	copied.StructureVersion = serviceItemGetSupportIDParam.StructureVersion

	copied.UniqueID = serviceItemGetSupportIDParam.UniqueID
	copied.Platform = serviceItemGetSupportIDParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetSupportIDParam *ServiceItemGetSupportIDParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetSupportIDParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetSupportIDParam)

	if serviceItemGetSupportIDParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemGetSupportIDParam.UniqueID.Equals(other.UniqueID) {
		return false
	}

	return serviceItemGetSupportIDParam.Platform == other.Platform
}

// String returns a string representation of the struct
func (serviceItemGetSupportIDParam *ServiceItemGetSupportIDParam) String() string {
	return serviceItemGetSupportIDParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetSupportIDParam *ServiceItemGetSupportIDParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetSupportIDParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemGetSupportIDParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemGetSupportIDParam.UniqueID))

	if serviceItemGetSupportIDParam.StructureVersion >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemGetSupportIDParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetSupportIDParam returns a new ServiceItemGetSupportIDParam
func NewServiceItemGetSupportIDParam() *ServiceItemGetSupportIDParam {
	return &ServiceItemGetSupportIDParam{}
}
