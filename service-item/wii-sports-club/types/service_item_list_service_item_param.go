// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemListServiceItemParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemListServiceItemParam struct {
	nex.Structure
	Language string
	Offset   uint32
	Size     uint32
	TitleID  string
}

// ExtractFromStream extracts a ServiceItemListServiceItemParam structure from a stream
func (serviceItemListServiceItemParam *ServiceItemListServiceItemParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemListServiceItemParam.Language, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Language from stream. %s", err.Error())
	}

	serviceItemListServiceItemParam.Offset, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Offset from stream. %s", err.Error())
	}

	serviceItemListServiceItemParam.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.Size from stream. %s", err.Error())
	}

	serviceItemListServiceItemParam.TitleID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemParam.TitleID from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemListServiceItemParam and returns a byte array
func (serviceItemListServiceItemParam *ServiceItemListServiceItemParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemListServiceItemParam.Language)
	stream.WriteUInt32LE(serviceItemListServiceItemParam.Offset)
	stream.WriteUInt32LE(serviceItemListServiceItemParam.Size)
	stream.WriteString(serviceItemListServiceItemParam.TitleID)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemListServiceItemParam
func (serviceItemListServiceItemParam *ServiceItemListServiceItemParam) Copy() nex.StructureInterface {
	copied := NewServiceItemListServiceItemParam()

	copied.SetStructureVersion(serviceItemListServiceItemParam.StructureVersion())

	copied.Language = serviceItemListServiceItemParam.Language
	copied.Offset = serviceItemListServiceItemParam.Offset
	copied.Size = serviceItemListServiceItemParam.Size
	copied.TitleID = serviceItemListServiceItemParam.TitleID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemListServiceItemParam *ServiceItemListServiceItemParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemListServiceItemParam)

	if serviceItemListServiceItemParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemListServiceItemParam.Language != other.Language {
		return false
	}

	if serviceItemListServiceItemParam.Offset != other.Offset {
		return false
	}

	if serviceItemListServiceItemParam.Size != other.Size {
		return false
	}

	if serviceItemListServiceItemParam.TitleID != other.TitleID {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemListServiceItemParam *ServiceItemListServiceItemParam) String() string {
	return serviceItemListServiceItemParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemListServiceItemParam *ServiceItemListServiceItemParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemListServiceItemParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemListServiceItemParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemListServiceItemParam.Language))
	b.WriteString(fmt.Sprintf("%sOffset: %d,\n", indentationValues, serviceItemListServiceItemParam.Offset))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, serviceItemListServiceItemParam.Size))
	b.WriteString(fmt.Sprintf("%sTitleID: %q,\n", indentationValues, serviceItemListServiceItemParam.TitleID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemListServiceItemParam returns a new ServiceItemListServiceItemParam
func NewServiceItemListServiceItemParam() *ServiceItemListServiceItemParam {
	return &ServiceItemListServiceItemParam{}
}
