// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemTicketInfo holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemTicketInfo struct {
	nex.Structure
	TicketType uint32
	NumTotal   uint32
}

// ExtractFromStream extracts a ServiceItemTicketInfo structure from a stream
func (serviceItemTicketInfo *ServiceItemTicketInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemTicketInfo.TicketType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTicketInfo.TicketType from stream. %s", err.Error())
	}

	serviceItemTicketInfo.NumTotal, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTicketInfo.NumTotal from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemTicketInfo and returns a byte array
func (serviceItemTicketInfo *ServiceItemTicketInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemTicketInfo.TicketType)
	stream.WriteUInt32LE(serviceItemTicketInfo.NumTotal)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemTicketInfo
func (serviceItemTicketInfo *ServiceItemTicketInfo) Copy() nex.StructureInterface {
	copied := NewServiceItemTicketInfo()

	copied.SetStructureVersion(serviceItemTicketInfo.StructureVersion())

	copied.TicketType = serviceItemTicketInfo.TicketType
	copied.NumTotal = serviceItemTicketInfo.NumTotal

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemTicketInfo *ServiceItemTicketInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemTicketInfo)

	if serviceItemTicketInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemTicketInfo.TicketType != other.TicketType {
		return false
	}

	if serviceItemTicketInfo.NumTotal != other.NumTotal {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemTicketInfo *ServiceItemTicketInfo) String() string {
	return serviceItemTicketInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemTicketInfo *ServiceItemTicketInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemTicketInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemTicketInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sTicketType: %d,\n", indentationValues, serviceItemTicketInfo.TicketType))
	b.WriteString(fmt.Sprintf("%sNumTotal: %d,\n", indentationValues, serviceItemTicketInfo.NumTotal))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemTicketInfo returns a new ServiceItemTicketInfo
func NewServiceItemTicketInfo() *ServiceItemTicketInfo {
	return &ServiceItemTicketInfo{}
}
