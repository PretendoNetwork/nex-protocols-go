// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetLawMessageParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetLawMessageParam struct {
	nex.Structure
	Language string
	UniqueID uint32
	Platform uint8 // * Revision 1
}

// ExtractFromStream extracts a ServiceItemGetLawMessageParam structure from a stream
func (serviceItemGetLawMessageParam *ServiceItemGetLawMessageParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemGetLawMessageParam.Language, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetLawMessageParam.Language from stream. %s", err.Error())
	}

	serviceItemGetLawMessageParam.UniqueID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetLawMessageParam.UniqueID from stream. %s", err.Error())
	}

	if serviceItemGetLawMessageParam.StructureVersion() >= 1 {
		serviceItemGetLawMessageParam.Platform, err = stream.ReadUInt8()
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemGetLawMessageParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// Bytes encodes the ServiceItemGetLawMessageParam and returns a byte array
func (serviceItemGetLawMessageParam *ServiceItemGetLawMessageParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemGetLawMessageParam.Language)
	stream.WriteUInt32LE(serviceItemGetLawMessageParam.UniqueID)

	if serviceItemGetLawMessageParam.StructureVersion() >= 1 {
		stream.WriteUInt8(serviceItemGetLawMessageParam.Platform)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetLawMessageParam
func (serviceItemGetLawMessageParam *ServiceItemGetLawMessageParam) Copy() nex.StructureInterface {
	copied := NewServiceItemGetLawMessageParam()

	copied.Language = serviceItemGetLawMessageParam.Language
	copied.UniqueID = serviceItemGetLawMessageParam.UniqueID
	copied.Platform = serviceItemGetLawMessageParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetLawMessageParam *ServiceItemGetLawMessageParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetLawMessageParam)

	if serviceItemGetLawMessageParam.Language != other.Language {
		return false
	}

	if serviceItemGetLawMessageParam.UniqueID != other.UniqueID {
		return false
	}

	if serviceItemGetLawMessageParam.Platform != other.Platform {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemGetLawMessageParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemGetLawMessageParam.Language))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemGetLawMessageParam.UniqueID))

	if serviceItemGetLawMessageParam.StructureVersion() >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemGetLawMessageParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetLawMessageParam returns a new ServiceItemGetLawMessageParam
func NewServiceItemGetLawMessageParam() *ServiceItemGetLawMessageParam {
	return &ServiceItemGetLawMessageParam{}
}
