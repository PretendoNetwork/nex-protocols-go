// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemRequestTicketRestorationParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemRequestTicketRestorationParam struct {
	types.Structure
	TicketType *types.PrimitiveU32
	NumTicket  *types.PrimitiveU32
}

// ExtractFrom extracts the ServiceItemRequestTicketRestorationParam from the given readable
func (serviceItemRequestTicketRestorationParam *ServiceItemRequestTicketRestorationParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemRequestTicketRestorationParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemRequestTicketRestorationParam header. %s", err.Error())
	}

	err = serviceItemRequestTicketRestorationParam.TicketType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRequestTicketRestorationParam.TicketType from stream. %s", err.Error())
	}

	err = serviceItemRequestTicketRestorationParam.NumTicket.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRequestTicketRestorationParam.NumTicket from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemRequestTicketRestorationParam to the given writable
func (serviceItemRequestTicketRestorationParam *ServiceItemRequestTicketRestorationParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemRequestTicketRestorationParam.TicketType.WriteTo(contentWritable)
	serviceItemRequestTicketRestorationParam.NumTicket.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemRequestTicketRestorationParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemRequestTicketRestorationParam
func (serviceItemRequestTicketRestorationParam *ServiceItemRequestTicketRestorationParam) Copy() types.RVType {
	copied := NewServiceItemRequestTicketRestorationParam()

	copied.StructureVersion = serviceItemRequestTicketRestorationParam.StructureVersion

	copied.TicketType = serviceItemRequestTicketRestorationParam.TicketType
	copied.NumTicket = serviceItemRequestTicketRestorationParam.NumTicket

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRequestTicketRestorationParam *ServiceItemRequestTicketRestorationParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRequestTicketRestorationParam); !ok {
		return false
	}

	other := o.(*ServiceItemRequestTicketRestorationParam)

	if serviceItemRequestTicketRestorationParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemRequestTicketRestorationParam.TicketType.Equals(other.TicketType) {
		return false
	}

	if !serviceItemRequestTicketRestorationParam.NumTicket.Equals(other.NumTicket) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemRequestTicketRestorationParam *ServiceItemRequestTicketRestorationParam) String() string {
	return serviceItemRequestTicketRestorationParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemRequestTicketRestorationParam *ServiceItemRequestTicketRestorationParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRequestTicketRestorationParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemRequestTicketRestorationParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sTicketType: %d,\n", indentationValues, serviceItemRequestTicketRestorationParam.TicketType))
	b.WriteString(fmt.Sprintf("%sNumTicket: %d,\n", indentationValues, serviceItemRequestTicketRestorationParam.NumTicket))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRequestTicketRestorationParam returns a new ServiceItemRequestTicketRestorationParam
func NewServiceItemRequestTicketRestorationParam() *ServiceItemRequestTicketRestorationParam {
	return &ServiceItemRequestTicketRestorationParam{}
}
