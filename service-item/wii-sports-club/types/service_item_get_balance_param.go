// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetBalanceParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemGetBalanceParam struct {
	nex.Structure
	Language string
	TitleID  string
}

// ExtractFromStream extracts a ServiceItemGetBalanceParam structure from a stream
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemGetBalanceParam.Language, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam.Language from stream. %s", err.Error())
	}

	serviceItemGetBalanceParam.TitleID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam.TitleID from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemGetBalanceParam and returns a byte array
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemGetBalanceParam.Language)
	stream.WriteString(serviceItemGetBalanceParam.TitleID)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetBalanceParam
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) Copy() nex.StructureInterface {
	copied := NewServiceItemGetBalanceParam()

	copied.SetStructureVersion(serviceItemGetBalanceParam.StructureVersion())

	copied.Language = serviceItemGetBalanceParam.Language
	copied.TitleID = serviceItemGetBalanceParam.TitleID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetBalanceParam)

	if serviceItemGetBalanceParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemGetBalanceParam.Language != other.Language {
		return false
	}

	if serviceItemGetBalanceParam.TitleID != other.TitleID {
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
	b.WriteString(fmt.Sprintf("%sTitleID: %q,\n", indentationValues, serviceItemGetBalanceParam.TitleID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetBalanceParam returns a new ServiceItemGetBalanceParam
func NewServiceItemGetBalanceParam() *ServiceItemGetBalanceParam {
	return &ServiceItemGetBalanceParam{}
}
