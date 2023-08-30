// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemLawMessage holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemLawMessage struct {
	nex.Structure
	IsMessageRequired bool
	LawMessage        string
}

// ExtractFromStream extracts a ServiceItemLawMessage structure from a stream
func (serviceItemLawMessage *ServiceItemLawMessage) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemLawMessage.IsMessageRequired, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLawMessage.IsMessageRequired from stream. %s", err.Error())
	}

	serviceItemLawMessage.LawMessage, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemLawMessage.LawMessage from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemLawMessage and returns a byte array
func (serviceItemLawMessage *ServiceItemLawMessage) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteBool(serviceItemLawMessage.IsMessageRequired)
	stream.WriteString(serviceItemLawMessage.LawMessage)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemLawMessage
func (serviceItemLawMessage *ServiceItemLawMessage) Copy() nex.StructureInterface {
	copied := NewServiceItemLawMessage()

	copied.SetStructureVersion(serviceItemLawMessage.StructureVersion())

	copied.IsMessageRequired = serviceItemLawMessage.IsMessageRequired
	copied.LawMessage = serviceItemLawMessage.LawMessage

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemLawMessage *ServiceItemLawMessage) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemLawMessage)

	if serviceItemLawMessage.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemLawMessage.IsMessageRequired != other.IsMessageRequired {
		return false
	}

	if serviceItemLawMessage.LawMessage != other.LawMessage {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemLawMessage.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sIsMessageRequired: %t,\n", indentationValues, serviceItemLawMessage.IsMessageRequired))
	b.WriteString(fmt.Sprintf("%sLawMessage: %q,\n", indentationValues, serviceItemLawMessage.LawMessage))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemLawMessage returns a new ServiceItemLawMessage
func NewServiceItemLawMessage() *ServiceItemLawMessage {
	return &ServiceItemLawMessage{}
}
