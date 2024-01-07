// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemGetLawMessageParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetLawMessageParam struct {
	types.Structure
	Language string
	UniqueID *types.PrimitiveU32
	Platform *types.PrimitiveU8 // * Revision 1
}

// ExtractFrom extracts the ServiceItemGetLawMessageParam from the given readable
func (serviceItemGetLawMessageParam *ServiceItemGetLawMessageParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemGetLawMessageParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemGetLawMessageParam header. %s", err.Error())
	}

	err = serviceItemGetLawMessageParam.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetLawMessageParam.Language from stream. %s", err.Error())
	}

	err = serviceItemGetLawMessageParam.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetLawMessageParam.UniqueID from stream. %s", err.Error())
	}

	if serviceItemGetLawMessageParam.StructureVersion >= 1 {
	err = 	serviceItemGetLawMessageParam.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemGetLawMessageParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// WriteTo writes the ServiceItemGetLawMessageParam to the given writable
func (serviceItemGetLawMessageParam *ServiceItemGetLawMessageParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemGetLawMessageParam.Language.WriteTo(contentWritable)
	serviceItemGetLawMessageParam.UniqueID.WriteTo(contentWritable)

	if serviceItemGetLawMessageParam.StructureVersion >= 1 {
		serviceItemGetLawMessageParam.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemGetLawMessageParam
func (serviceItemGetLawMessageParam *ServiceItemGetLawMessageParam) Copy() types.RVType {
	copied := NewServiceItemGetLawMessageParam()

	copied.StructureVersion = serviceItemGetLawMessageParam.StructureVersion

	copied.Language = serviceItemGetLawMessageParam.Language
	copied.UniqueID = serviceItemGetLawMessageParam.UniqueID
	copied.Platform = serviceItemGetLawMessageParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetLawMessageParam *ServiceItemGetLawMessageParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetLawMessageParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetLawMessageParam)

	if serviceItemGetLawMessageParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemGetLawMessageParam.Language.Equals(other.Language) {
		return false
	}

	if !serviceItemGetLawMessageParam.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !serviceItemGetLawMessageParam.Platform.Equals(other.Platform) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetLawMessageParam *ServiceItemGetLawMessageParam) String() string {
	return serviceItemGetLawMessageParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetLawMessageParam *ServiceItemGetLawMessageParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetLawMessageParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemGetLawMessageParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemGetLawMessageParam.Language))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemGetLawMessageParam.UniqueID))

	if serviceItemGetLawMessageParam.StructureVersion >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemGetLawMessageParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetLawMessageParam returns a new ServiceItemGetLawMessageParam
func NewServiceItemGetLawMessageParam() *ServiceItemGetLawMessageParam {
	return &ServiceItemGetLawMessageParam{}
}
