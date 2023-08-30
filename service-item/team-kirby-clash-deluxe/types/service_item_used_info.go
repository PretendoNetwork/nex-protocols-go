// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemUsedInfo holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemUsedInfo struct {
	nex.Structure
	AcquiredCount uint32
	UsedCount     uint32
}

// ExtractFromStream extracts a ServiceItemUsedInfo structure from a stream
func (serviceItemUsedInfo *ServiceItemUsedInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemUsedInfo.AcquiredCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUsedInfo.AcquiredCount from stream. %s", err.Error())
	}

	serviceItemUsedInfo.UsedCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUsedInfo.UsedCount from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemUsedInfo and returns a byte array
func (serviceItemUsedInfo *ServiceItemUsedInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemUsedInfo.AcquiredCount)
	stream.WriteUInt32LE(serviceItemUsedInfo.UsedCount)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemUsedInfo
func (serviceItemUsedInfo *ServiceItemUsedInfo) Copy() nex.StructureInterface {
	copied := NewServiceItemUsedInfo()

	copied.SetStructureVersion(serviceItemUsedInfo.StructureVersion())

	copied.AcquiredCount = serviceItemUsedInfo.AcquiredCount
	copied.UsedCount = serviceItemUsedInfo.UsedCount

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemUsedInfo *ServiceItemUsedInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemUsedInfo)

	if serviceItemUsedInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemUsedInfo.AcquiredCount != other.AcquiredCount {
		return false
	}

	if serviceItemUsedInfo.UsedCount != other.UsedCount {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemUsedInfo *ServiceItemUsedInfo) String() string {
	return serviceItemUsedInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemUsedInfo *ServiceItemUsedInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemUsedInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemUsedInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sAcquiredCount: %d,\n", indentationValues, serviceItemUsedInfo.AcquiredCount))
	b.WriteString(fmt.Sprintf("%sUsedCount: %d,\n", indentationValues, serviceItemUsedInfo.UsedCount))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemUsedInfo returns a new ServiceItemUsedInfo
func NewServiceItemUsedInfo() *ServiceItemUsedInfo {
	return &ServiceItemUsedInfo{}
}
