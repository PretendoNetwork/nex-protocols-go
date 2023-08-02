// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetPrepurchaseInfoParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetPrepurchaseInfoParam struct {
	nex.Structure
	ItemCode string
	ReferenceID string
	Limitation *ServiceItemLimitation
	Language string
	UniqueID uint32
}

// ExtractFromStream extracts a ServiceItemGetPrepurchaseInfoParam structure from a stream
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemGetPrepurchaseInfoParam.ItemCode, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.ItemCode from stream. %s", err.Error())
	}

	serviceItemGetPrepurchaseInfoParam.ReferenceID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.ReferenceID from stream. %s", err.Error())
	}

	limitation, err := stream.ReadStructure(NewServiceItemLimitation())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.Limitation from stream. %s", err.Error())
	}

	serviceItemGetPrepurchaseInfoParam.Limitation = limitation.(*ServiceItemLimitation)

	serviceItemGetPrepurchaseInfoParam.Language, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.Language from stream. %s", err.Error())
	}

	serviceItemGetPrepurchaseInfoParam.UniqueID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.UniqueID from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemGetPrepurchaseInfoParam and returns a byte array
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemGetPrepurchaseInfoParam.ItemCode)
	stream.WriteString(serviceItemGetPrepurchaseInfoParam.ReferenceID)
	stream.WriteStructure(serviceItemGetPrepurchaseInfoParam.Limitation)
	stream.WriteString(serviceItemGetPrepurchaseInfoParam.Language)
	stream.WriteUInt32LE(serviceItemGetPrepurchaseInfoParam.UniqueID)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetPrepurchaseInfoParam
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) Copy() nex.StructureInterface {
	copied := NewServiceItemGetPrepurchaseInfoParam()

	copied.ItemCode = serviceItemGetPrepurchaseInfoParam.ItemCode
	copied.ReferenceID = serviceItemGetPrepurchaseInfoParam.ReferenceID
	copied.Limitation = serviceItemGetPrepurchaseInfoParam.Limitation.Copy().(*ServiceItemLimitation)
	copied.Language = serviceItemGetPrepurchaseInfoParam.Language
	copied.UniqueID = serviceItemGetPrepurchaseInfoParam.UniqueID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetPrepurchaseInfoParam)

	if serviceItemGetPrepurchaseInfoParam.ItemCode != other.ItemCode {
		return false
	}

	if serviceItemGetPrepurchaseInfoParam.ReferenceID != other.ReferenceID {
		return false
	}

	if !serviceItemGetPrepurchaseInfoParam.Limitation.Equals(other.Limitation) {
		return false
	}

	if serviceItemGetPrepurchaseInfoParam.Language != other.Language {
		return false
	}

	if serviceItemGetPrepurchaseInfoParam.UniqueID != other.UniqueID {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) String() string {
	return serviceItemGetPrepurchaseInfoParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetPrepurchaseInfoParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemGetPrepurchaseInfoParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sItemCode: %q,\n", indentationValues, serviceItemGetPrepurchaseInfoParam.ItemCode))
	b.WriteString(fmt.Sprintf("%sReferenceID: %q,\n", indentationValues, serviceItemGetPrepurchaseInfoParam.ReferenceID))

	if serviceItemGetPrepurchaseInfoParam.Limitation != nil {
		b.WriteString(fmt.Sprintf("%sLimitation: %s\n", indentationValues, serviceItemGetPrepurchaseInfoParam.Limitation.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sLimitation: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemGetPrepurchaseInfoParam.Language))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemGetPrepurchaseInfoParam.UniqueID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetPrepurchaseInfoParam returns a new ServiceItemGetPrepurchaseInfoParam
func NewServiceItemGetPrepurchaseInfoParam() *ServiceItemGetPrepurchaseInfoParam {
	return &ServiceItemGetPrepurchaseInfoParam{}
}
