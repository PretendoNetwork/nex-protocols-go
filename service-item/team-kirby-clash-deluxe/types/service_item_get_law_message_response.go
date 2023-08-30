// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetLawMessageResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetLawMessageResponse struct {
	nex.Structure
	*ServiceItemEShopResponse
	NullableLawMessage []*ServiceItemLawMessage
}

// ExtractFromStream extracts a ServiceItemGetLawMessageResponse structure from a stream
func (serviceItemGetLawMessageResponse *ServiceItemGetLawMessageResponse) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	nullableLawMessage, err := stream.ReadListStructure(NewServiceItemLawMessage())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetLawMessageResponse.NullableLawMessage from stream. %s", err.Error())
	}

	serviceItemGetLawMessageResponse.NullableLawMessage = nullableLawMessage.([]*ServiceItemLawMessage)

	return nil
}

// Bytes encodes the ServiceItemGetLawMessageResponse and returns a byte array
func (serviceItemGetLawMessageResponse *ServiceItemGetLawMessageResponse) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListStructure(serviceItemGetLawMessageResponse.NullableLawMessage)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetLawMessageResponse
func (serviceItemGetLawMessageResponse *ServiceItemGetLawMessageResponse) Copy() nex.StructureInterface {
	copied := NewServiceItemGetLawMessageResponse()

	copied.SetStructureVersion(serviceItemGetLawMessageResponse.StructureVersion())

	copied.ServiceItemEShopResponse = serviceItemGetLawMessageResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.SetParentType(copied.ServiceItemEShopResponse)

	copied.NullableLawMessage = make([]*ServiceItemLawMessage, len(serviceItemGetLawMessageResponse.NullableLawMessage))

	for i := 0; i < len(serviceItemGetLawMessageResponse.NullableLawMessage); i++ {
		copied.NullableLawMessage[i] = serviceItemGetLawMessageResponse.NullableLawMessage[i].Copy().(*ServiceItemLawMessage)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetLawMessageResponse *ServiceItemGetLawMessageResponse) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetLawMessageResponse)

	if serviceItemGetLawMessageResponse.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !serviceItemGetLawMessageResponse.ParentType().Equals(other.ParentType()) {
		return false
	}

	if len(serviceItemGetLawMessageResponse.NullableLawMessage) != len(other.NullableLawMessage) {
		return false
	}

	for i := 0; i < len(serviceItemGetLawMessageResponse.NullableLawMessage); i++ {
		if !serviceItemGetLawMessageResponse.NullableLawMessage[i].Equals(other.NullableLawMessage[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetLawMessageResponse *ServiceItemGetLawMessageResponse) String() string {
	return serviceItemGetLawMessageResponse.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetLawMessageResponse *ServiceItemGetLawMessageResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetLawMessageResponse{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, serviceItemGetLawMessageResponse.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemGetLawMessageResponse.StructureVersion()))

	if len(serviceItemGetLawMessageResponse.NullableLawMessage) == 0 {
		b.WriteString(fmt.Sprintf("%sNullableLawMessage: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sNullableLawMessage: [\n", indentationValues))

		for i := 0; i < len(serviceItemGetLawMessageResponse.NullableLawMessage); i++ {
			str := serviceItemGetLawMessageResponse.NullableLawMessage[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemGetLawMessageResponse.NullableLawMessage)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetLawMessageResponse returns a new ServiceItemGetLawMessageResponse
func NewServiceItemGetLawMessageResponse() *ServiceItemGetLawMessageResponse {
	serviceItemGetLawMessageResponse := &ServiceItemGetLawMessageResponse{}

	serviceItemGetLawMessageResponse.ServiceItemEShopResponse = NewServiceItemEShopResponse()
	serviceItemGetLawMessageResponse.SetParentType(serviceItemGetLawMessageResponse.ServiceItemEShopResponse)

	return serviceItemGetLawMessageResponse
}
