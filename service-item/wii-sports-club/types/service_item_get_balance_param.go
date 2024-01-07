// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemGetBalanceParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemGetBalanceParam struct {
	types.Structure
	Language string
	TitleID  string
}

// ExtractFrom extracts the ServiceItemGetBalanceParam from the given readable
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemGetBalanceParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemGetBalanceParam header. %s", err.Error())
	}

	err = serviceItemGetBalanceParam.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam.Language from stream. %s", err.Error())
	}

	err = serviceItemGetBalanceParam.TitleID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam.TitleID from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemGetBalanceParam to the given writable
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemGetBalanceParam.Language.WriteTo(contentWritable)
	serviceItemGetBalanceParam.TitleID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemGetBalanceParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemGetBalanceParam
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) Copy() types.RVType {
	copied := NewServiceItemGetBalanceParam()

	copied.StructureVersion = serviceItemGetBalanceParam.StructureVersion

	copied.Language = serviceItemGetBalanceParam.Language
	copied.TitleID = serviceItemGetBalanceParam.TitleID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetBalanceParam *ServiceItemGetBalanceParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetBalanceParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetBalanceParam)

	if serviceItemGetBalanceParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemGetBalanceParam.Language.Equals(other.Language) {
		return false
	}

	if !serviceItemGetBalanceParam.TitleID.Equals(other.TitleID) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemGetBalanceParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemGetBalanceParam.Language))
	b.WriteString(fmt.Sprintf("%sTitleID: %q,\n", indentationValues, serviceItemGetBalanceParam.TitleID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetBalanceParam returns a new ServiceItemGetBalanceParam
func NewServiceItemGetBalanceParam() *ServiceItemGetBalanceParam {
	return &ServiceItemGetBalanceParam{}
}
