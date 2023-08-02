// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemUserInfo holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemUserInfo struct {
	nex.Structure
	NumTotalEntryTicket uint32
	ApplicationBuffer []byte
}

// ExtractFromStream extracts a ServiceItemUserInfo structure from a stream
func (serviceItemUserInfo *ServiceItemUserInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemUserInfo.NumTotalEntryTicket, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUserInfo.NumTotalEntryTicket from stream. %s", err.Error())
	}

	serviceItemUserInfo.ApplicationBuffer, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUserInfo.ApplicationBuffer from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemUserInfo and returns a byte array
func (serviceItemUserInfo *ServiceItemUserInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemUserInfo.NumTotalEntryTicket)
	stream.WriteQBuffer(serviceItemUserInfo.ApplicationBuffer)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemUserInfo
func (serviceItemUserInfo *ServiceItemUserInfo) Copy() nex.StructureInterface {
	copied := NewServiceItemUserInfo()

	copied.NumTotalEntryTicket = serviceItemUserInfo.NumTotalEntryTicket
	copied.ApplicationBuffer = serviceItemUserInfo.ApplicationBuffer

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemUserInfo *ServiceItemUserInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemUserInfo)

	if serviceItemUserInfo.NumTotalEntryTicket != other.NumTotalEntryTicket {
		return false
	}

	if !bytes.Equal(serviceItemUserInfo.ApplicationBuffer, other.ApplicationBuffer) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemUserInfo *ServiceItemUserInfo) String() string {
	return serviceItemUserInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemUserInfo *ServiceItemUserInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemUserInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemUserInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sNumTotalEntryTicket: %d,\n", indentationValues, serviceItemUserInfo.NumTotalEntryTicket))
	b.WriteString(fmt.Sprintf("%sApplicationBuffer: %x,\n", indentationValues, serviceItemUserInfo.ApplicationBuffer))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemUserInfo returns a new ServiceItemUserInfo
func NewServiceItemUserInfo() *ServiceItemUserInfo {
	return &ServiceItemUserInfo{}
}
