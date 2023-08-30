// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetPurchaseHistoryParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetPurchaseHistoryParam struct {
	nex.Structure
	Language string
	Offset   uint32
	Size     uint32
	UniqueID uint32
	Platform uint8
}

// ExtractFromStream extracts a ServiceItemGetPurchaseHistoryParam structure from a stream
func (serviceItemGetPurchaseHistoryParam *ServiceItemGetPurchaseHistoryParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemGetPurchaseHistoryParam.Language, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.Language from stream. %s", err.Error())
	}

	serviceItemGetPurchaseHistoryParam.Offset, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.Offset from stream. %s", err.Error())
	}

	serviceItemGetPurchaseHistoryParam.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.Size from stream. %s", err.Error())
	}

	serviceItemGetPurchaseHistoryParam.UniqueID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.UniqueID from stream. %s", err.Error())
	}

	if serviceItemGetPurchaseHistoryParam.StructureVersion() >= 1 {
		serviceItemGetPurchaseHistoryParam.Platform, err = stream.ReadUInt8()
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// Bytes encodes the ServiceItemGetPurchaseHistoryParam and returns a byte array
func (serviceItemGetPurchaseHistoryParam *ServiceItemGetPurchaseHistoryParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemGetPurchaseHistoryParam.Language)
	stream.WriteUInt32LE(serviceItemGetPurchaseHistoryParam.Offset)
	stream.WriteUInt32LE(serviceItemGetPurchaseHistoryParam.Size)
	stream.WriteUInt32LE(serviceItemGetPurchaseHistoryParam.UniqueID)

	if serviceItemGetPurchaseHistoryParam.StructureVersion() >= 1 {
		stream.WriteUInt8(serviceItemGetPurchaseHistoryParam.Platform)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetPurchaseHistoryParam
func (serviceItemGetPurchaseHistoryParam *ServiceItemGetPurchaseHistoryParam) Copy() nex.StructureInterface {
	copied := NewServiceItemGetPurchaseHistoryParam()

	copied.SetStructureVersion(serviceItemGetPurchaseHistoryParam.StructureVersion())

	copied.Language = serviceItemGetPurchaseHistoryParam.Language
	copied.Offset = serviceItemGetPurchaseHistoryParam.Offset
	copied.Size = serviceItemGetPurchaseHistoryParam.Size
	copied.UniqueID = serviceItemGetPurchaseHistoryParam.UniqueID
	copied.Platform = serviceItemGetPurchaseHistoryParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetPurchaseHistoryParam *ServiceItemGetPurchaseHistoryParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetPurchaseHistoryParam)

	if serviceItemGetPurchaseHistoryParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemGetPurchaseHistoryParam.Language != other.Language {
		return false
	}

	if serviceItemGetPurchaseHistoryParam.Offset != other.Offset {
		return false
	}

	if serviceItemGetPurchaseHistoryParam.Size != other.Size {
		return false
	}

	if serviceItemGetPurchaseHistoryParam.UniqueID != other.UniqueID {
		return false
	}

	if serviceItemGetPurchaseHistoryParam.Platform != other.Platform {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetPurchaseHistoryParam *ServiceItemGetPurchaseHistoryParam) String() string {
	return serviceItemGetPurchaseHistoryParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetPurchaseHistoryParam *ServiceItemGetPurchaseHistoryParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetPurchaseHistoryParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemGetPurchaseHistoryParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemGetPurchaseHistoryParam.Language))
	b.WriteString(fmt.Sprintf("%sOffset: %d,\n", indentationValues, serviceItemGetPurchaseHistoryParam.Offset))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, serviceItemGetPurchaseHistoryParam.Size))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemGetPurchaseHistoryParam.UniqueID))

	if serviceItemGetPurchaseHistoryParam.StructureVersion() >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemGetPurchaseHistoryParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetPurchaseHistoryParam returns a new ServiceItemGetPurchaseHistoryParam
func NewServiceItemGetPurchaseHistoryParam() *ServiceItemGetPurchaseHistoryParam {
	return &ServiceItemGetPurchaseHistoryParam{}
}
