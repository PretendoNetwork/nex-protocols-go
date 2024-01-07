// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemLawMessage holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemLawMessage struct {
	types.Structure
	IsMessageRequired *types.PrimitiveBool
	LawMessage        string
}

// ExtractFrom extracts the ServiceItemLawMessage from the given readable
func (serviceItemLawMessage *ServiceItemLawMessage) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemLawMessage.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemLawMessage header. %s", err.Error())
	}

	err = serviceItemLawMessage.IsMessageRequired.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLawMessage.IsMessageRequired from stream. %s", err.Error())
	}

	err = serviceItemLawMessage.LawMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLawMessage.LawMessage from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemLawMessage to the given writable
func (serviceItemLawMessage *ServiceItemLawMessage) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemLawMessage.IsMessageRequired.WriteTo(contentWritable)
	serviceItemLawMessage.LawMessage.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemLawMessage.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemLawMessage
func (serviceItemLawMessage *ServiceItemLawMessage) Copy() types.RVType {
	copied := NewServiceItemLawMessage()

	copied.StructureVersion = serviceItemLawMessage.StructureVersion

	copied.IsMessageRequired = serviceItemLawMessage.IsMessageRequired
	copied.LawMessage = serviceItemLawMessage.LawMessage

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemLawMessage *ServiceItemLawMessage) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemLawMessage); !ok {
		return false
	}

	other := o.(*ServiceItemLawMessage)

	if serviceItemLawMessage.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemLawMessage.IsMessageRequired.Equals(other.IsMessageRequired) {
		return false
	}

	if !serviceItemLawMessage.LawMessage.Equals(other.LawMessage) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemLawMessage *ServiceItemLawMessage) String() string {
	return serviceItemLawMessage.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemLawMessage *ServiceItemLawMessage) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemLawMessage{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemLawMessage.StructureVersion))
	b.WriteString(fmt.Sprintf("%sIsMessageRequired: %t,\n", indentationValues, serviceItemLawMessage.IsMessageRequired))
	b.WriteString(fmt.Sprintf("%sLawMessage: %q,\n", indentationValues, serviceItemLawMessage.LawMessage))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemLawMessage returns a new ServiceItemLawMessage
func NewServiceItemLawMessage() *ServiceItemLawMessage {
	return &ServiceItemLawMessage{}
}
