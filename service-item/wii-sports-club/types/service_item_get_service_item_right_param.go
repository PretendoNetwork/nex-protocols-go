// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemGetServiceItemRightParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemGetServiceItemRightParam struct {
	types.Structure
	ReferenceID string
	TitleID     string
}

// ExtractFrom extracts the ServiceItemGetServiceItemRightParam from the given readable
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemGetServiceItemRightParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemGetServiceItemRightParam header. %s", err.Error())
	}

	err = serviceItemGetServiceItemRightParam.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.ReferenceID from stream. %s", err.Error())
	}

	err = serviceItemGetServiceItemRightParam.TitleID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.TitleID from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemGetServiceItemRightParam to the given writable
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemGetServiceItemRightParam.ReferenceID.WriteTo(contentWritable)
	serviceItemGetServiceItemRightParam.TitleID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemGetServiceItemRightParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemGetServiceItemRightParam
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) Copy() types.RVType {
	copied := NewServiceItemGetServiceItemRightParam()

	copied.StructureVersion = serviceItemGetServiceItemRightParam.StructureVersion

	copied.ReferenceID = serviceItemGetServiceItemRightParam.ReferenceID
	copied.TitleID = serviceItemGetServiceItemRightParam.TitleID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetServiceItemRightParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetServiceItemRightParam)

	if serviceItemGetServiceItemRightParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemGetServiceItemRightParam.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !serviceItemGetServiceItemRightParam.TitleID.Equals(other.TitleID) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) String() string {
	return serviceItemGetServiceItemRightParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetServiceItemRightParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemGetServiceItemRightParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sReferenceID: %q,\n", indentationValues, serviceItemGetServiceItemRightParam.ReferenceID))
	b.WriteString(fmt.Sprintf("%sTitleID: %q,\n", indentationValues, serviceItemGetServiceItemRightParam.TitleID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetServiceItemRightParam returns a new ServiceItemGetServiceItemRightParam
func NewServiceItemGetServiceItemRightParam() *ServiceItemGetServiceItemRightParam {
	return &ServiceItemGetServiceItemRightParam{}
}
