// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetServiceItemRightParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemGetServiceItemRightParam struct {
	nex.Structure
	ReferenceID string
	TitleID string
}

// ExtractFromStream extracts a ServiceItemGetServiceItemRightParam structure from a stream
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemGetServiceItemRightParam.ReferenceID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.ReferenceID from stream. %s", err.Error())
	}

	serviceItemGetServiceItemRightParam.TitleID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.TitleID from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemGetServiceItemRightParam and returns a byte array
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemGetServiceItemRightParam.ReferenceID)
	stream.WriteString(serviceItemGetServiceItemRightParam.TitleID)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetServiceItemRightParam
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) Copy() nex.StructureInterface {
	copied := NewServiceItemGetServiceItemRightParam()

	copied.ReferenceID = serviceItemGetServiceItemRightParam.ReferenceID
	copied.TitleID = serviceItemGetServiceItemRightParam.TitleID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetServiceItemRightParam)

	if serviceItemGetServiceItemRightParam.ReferenceID != other.ReferenceID {
		return false
	}

	if serviceItemGetServiceItemRightParam.TitleID != other.TitleID {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemGetServiceItemRightParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sReferenceID: %q,\n", indentationValues, serviceItemGetServiceItemRightParam.ReferenceID))
	b.WriteString(fmt.Sprintf("%sTitleID: %q,\n", indentationValues, serviceItemGetServiceItemRightParam.TitleID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetServiceItemRightParam returns a new ServiceItemGetServiceItemRightParam
func NewServiceItemGetServiceItemRightParam() *ServiceItemGetServiceItemRightParam {
	return &ServiceItemGetServiceItemRightParam{}
}
