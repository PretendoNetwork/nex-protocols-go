// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetPrepurchaseInfoParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemGetPrepurchaseInfoParam struct {
	nex.Structure
	ItemCode string
	Language string
	TitleID  string
}

// ExtractFromStream extracts a ServiceItemGetPrepurchaseInfoParam structure from a stream
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemGetPrepurchaseInfoParam.ItemCode, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.ItemCode from stream. %s", err.Error())
	}

	serviceItemGetPrepurchaseInfoParam.Language, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.Language from stream. %s", err.Error())
	}

	serviceItemGetPrepurchaseInfoParam.TitleID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoParam.TitleID from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemGetPrepurchaseInfoParam and returns a byte array
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemGetPrepurchaseInfoParam.ItemCode)
	stream.WriteString(serviceItemGetPrepurchaseInfoParam.Language)
	stream.WriteString(serviceItemGetPrepurchaseInfoParam.TitleID)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetPrepurchaseInfoParam
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) Copy() nex.StructureInterface {
	copied := NewServiceItemGetPrepurchaseInfoParam()

	copied.SetStructureVersion(serviceItemGetPrepurchaseInfoParam.StructureVersion())

	copied.ItemCode = serviceItemGetPrepurchaseInfoParam.ItemCode
	copied.Language = serviceItemGetPrepurchaseInfoParam.Language
	copied.TitleID = serviceItemGetPrepurchaseInfoParam.TitleID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetPrepurchaseInfoParam *ServiceItemGetPrepurchaseInfoParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetPrepurchaseInfoParam)

	if serviceItemGetPrepurchaseInfoParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemGetPrepurchaseInfoParam.ItemCode != other.ItemCode {
		return false
	}

	if serviceItemGetPrepurchaseInfoParam.Language != other.Language {
		return false
	}

	if serviceItemGetPrepurchaseInfoParam.TitleID != other.TitleID {
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
	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemGetPrepurchaseInfoParam.Language))
	b.WriteString(fmt.Sprintf("%sTitleID: %q,\n", indentationValues, serviceItemGetPrepurchaseInfoParam.TitleID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetPrepurchaseInfoParam returns a new ServiceItemGetPrepurchaseInfoParam
func NewServiceItemGetPrepurchaseInfoParam() *ServiceItemGetPrepurchaseInfoParam {
	return &ServiceItemGetPrepurchaseInfoParam{}
}
