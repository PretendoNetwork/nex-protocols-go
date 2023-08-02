// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemRequestTicketRestorationParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemRequestTicketRestorationParam struct {
	nex.Structure
	TicketType uint32
	NumTicket uint32
}

// ExtractFromStream extracts a ServiceItemRequestTicketRestorationParam structure from a stream
func (serviceItemRequestTicketRestorationParam *ServiceItemRequestTicketRestorationParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemRequestTicketRestorationParam.TicketType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRequestTicketRestorationParam.TicketType from stream. %s", err.Error())
	}

	serviceItemRequestTicketRestorationParam.NumTicket, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRequestTicketRestorationParam.NumTicket from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemRequestTicketRestorationParam and returns a byte array
func (serviceItemRequestTicketRestorationParam *ServiceItemRequestTicketRestorationParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemRequestTicketRestorationParam.TicketType)
	stream.WriteUInt32LE(serviceItemRequestTicketRestorationParam.NumTicket)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemRequestTicketRestorationParam
func (serviceItemRequestTicketRestorationParam *ServiceItemRequestTicketRestorationParam) Copy() nex.StructureInterface {
	copied := NewServiceItemRequestTicketRestorationParam()

	copied.TicketType = serviceItemRequestTicketRestorationParam.TicketType
	copied.NumTicket = serviceItemRequestTicketRestorationParam.NumTicket

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRequestTicketRestorationParam *ServiceItemRequestTicketRestorationParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemRequestTicketRestorationParam)

	if serviceItemRequestTicketRestorationParam.TicketType != other.TicketType {
		return false
	}

	if serviceItemRequestTicketRestorationParam.NumTicket != other.NumTicket {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemRequestTicketRestorationParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sTicketType: %d,\n", indentationValues, serviceItemRequestTicketRestorationParam.TicketType))
	b.WriteString(fmt.Sprintf("%sNumTicket: %d,\n", indentationValues, serviceItemRequestTicketRestorationParam.NumTicket))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRequestTicketRestorationParam returns a new ServiceItemRequestTicketRestorationParam
func NewServiceItemRequestTicketRestorationParam() *ServiceItemRequestTicketRestorationParam {
	return &ServiceItemRequestTicketRestorationParam{}
}
