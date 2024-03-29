// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetSupportIDParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetSupportIDParam struct {
	nex.Structure
	UniqueID uint32
	Platform uint8
}

// ExtractFromStream extracts a ServiceItemGetSupportIDParam structure from a stream
func (serviceItemGetSupportIDParam *ServiceItemGetSupportIDParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemGetSupportIDParam.UniqueID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetSupportIDParam.UniqueID from stream. %s", err.Error())
	}

	if serviceItemGetSupportIDParam.StructureVersion() >= 1 {
		serviceItemGetSupportIDParam.Platform, err = stream.ReadUInt8()
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemGetSupportIDParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// Bytes encodes the ServiceItemGetSupportIDParam and returns a byte array
func (serviceItemGetSupportIDParam *ServiceItemGetSupportIDParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemGetSupportIDParam.UniqueID)

	if serviceItemGetSupportIDParam.StructureVersion() >= 1 {
		stream.WriteUInt8(serviceItemGetSupportIDParam.Platform)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetSupportIDParam
func (serviceItemGetSupportIDParam *ServiceItemGetSupportIDParam) Copy() nex.StructureInterface {
	copied := NewServiceItemGetSupportIDParam()

	copied.SetStructureVersion(serviceItemGetSupportIDParam.StructureVersion())

	copied.UniqueID = serviceItemGetSupportIDParam.UniqueID
	copied.Platform = serviceItemGetSupportIDParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetSupportIDParam *ServiceItemGetSupportIDParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetSupportIDParam)

	if serviceItemGetSupportIDParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemGetSupportIDParam.UniqueID != other.UniqueID {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemGetSupportIDParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemGetSupportIDParam.UniqueID))

	if serviceItemGetSupportIDParam.StructureVersion() >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemGetSupportIDParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetSupportIDParam returns a new ServiceItemGetSupportIDParam
func NewServiceItemGetSupportIDParam() *ServiceItemGetSupportIDParam {
	return &ServiceItemGetSupportIDParam{}
}
