// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemRightInfo holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemRightInfo struct {
	nex.Structure
	ReferenceID     string
	ReferenceIDType uint32
}

// ExtractFromStream extracts a ServiceItemRightInfo structure from a stream
func (serviceItemRightInfo *ServiceItemRightInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemRightInfo.ReferenceID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfo.ReferenceID from stream. %s", err.Error())
	}

	serviceItemRightInfo.ReferenceIDType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfo.ReferenceIDType from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemRightInfo and returns a byte array
func (serviceItemRightInfo *ServiceItemRightInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemRightInfo.ReferenceID)
	stream.WriteUInt32LE(serviceItemRightInfo.ReferenceIDType)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemRightInfo
func (serviceItemRightInfo *ServiceItemRightInfo) Copy() nex.StructureInterface {
	copied := NewServiceItemRightInfo()

	copied.SetStructureVersion(serviceItemRightInfo.StructureVersion())

	copied.ReferenceID = serviceItemRightInfo.ReferenceID
	copied.ReferenceIDType = serviceItemRightInfo.ReferenceIDType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRightInfo *ServiceItemRightInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemRightInfo)

	if serviceItemRightInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemRightInfo.ReferenceID != other.ReferenceID {
		return false
	}

	if serviceItemRightInfo.ReferenceIDType != other.ReferenceIDType {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemRightInfo *ServiceItemRightInfo) String() string {
	return serviceItemRightInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemRightInfo *ServiceItemRightInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemRightInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sReferenceID: %q,\n", indentationValues, serviceItemRightInfo.ReferenceID))
	b.WriteString(fmt.Sprintf("%sReferenceIDType: %d,\n", indentationValues, serviceItemRightInfo.ReferenceIDType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRightInfo returns a new ServiceItemRightInfo
func NewServiceItemRightInfo() *ServiceItemRightInfo {
	return &ServiceItemRightInfo{}
}
