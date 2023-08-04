// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetBalanceParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetBalanceParam struct {
	nex.Structure
	Language string
	UniqueID uint32
	Platform uint8 // * Revision 1
}

// ExtractFromStream extracts a ServiceItemGetBalanceParam structure from a stream
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemGetBalanceParam.Language, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam.Language from stream. %s", err.Error())
	}

	serviceItemGetBalanceParam.UniqueID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam.UniqueID from stream. %s", err.Error())
	}

	if serviceItemGetBalanceParam.StructureVersion() >= 1 {
		serviceItemGetBalanceParam.Platform, err = stream.ReadUInt8()
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// Bytes encodes the ServiceItemGetBalanceParam and returns a byte array
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemGetBalanceParam.Language)
	stream.WriteUInt32LE(serviceItemGetBalanceParam.UniqueID)

	if serviceItemGetBalanceParam.StructureVersion() >= 1 {
		stream.WriteUInt8(serviceItemGetBalanceParam.Platform)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetBalanceParam
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) Copy() nex.StructureInterface {
	copied := NewServiceItemGetBalanceParam()

	copied.Language = serviceItemGetBalanceParam.Language
	copied.UniqueID = serviceItemGetBalanceParam.UniqueID
	copied.Platform = serviceItemGetBalanceParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetBalanceParam)

	if serviceItemGetBalanceParam.Language != other.Language {
		return false
	}

	if serviceItemGetBalanceParam.UniqueID != other.UniqueID {
		return false
	}

	if serviceItemGetBalanceParam.Platform != other.Platform {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) String() string {
	return serviceItemGetBalanceParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetBalanceParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemGetBalanceParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemGetBalanceParam.Language))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemGetBalanceParam.UniqueID))

	if serviceItemGetBalanceParam.StructureVersion() >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemGetBalanceParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetBalanceParam returns a new ServiceItemGetBalanceParam
func NewServiceItemGetBalanceParam() *ServiceItemGetBalanceParam {
	return &ServiceItemGetBalanceParam{}
}
