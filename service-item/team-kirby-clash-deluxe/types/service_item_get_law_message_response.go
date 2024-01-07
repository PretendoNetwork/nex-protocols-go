// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemGetLawMessageResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetLawMessageResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullableLawMessage []*ServiceItemLawMessage
}

// ExtractFrom extracts the ServiceItemGetLawMessageResponse from the given readable
func (serviceItemGetLawMessageResponse *ServiceItemGetLawMessageResponse) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemGetLawMessageResponse.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemGetLawMessageResponse header. %s", err.Error())
	}

	nullableLawMessage, err := nex.StreamReadListStructure(stream, NewServiceItemLawMessage())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetLawMessageResponse.NullableLawMessage from stream. %s", err.Error())
	}

	serviceItemGetLawMessageResponse.NullableLawMessage = nullableLawMessage

	return nil
}

// WriteTo writes the ServiceItemGetLawMessageResponse to the given writable
func (serviceItemGetLawMessageResponse *ServiceItemGetLawMessageResponse) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemGetLawMessageResponse.NullableLawMessage.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemGetLawMessageResponse.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemGetLawMessageResponse
func (serviceItemGetLawMessageResponse *ServiceItemGetLawMessageResponse) Copy() types.RVType {
	copied := NewServiceItemGetLawMessageResponse()

	copied.StructureVersion = serviceItemGetLawMessageResponse.StructureVersion

	copied.ServiceItemEShopResponse = serviceItemGetLawMessageResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)

	copied.NullableLawMessage = make([]*ServiceItemLawMessage, len(serviceItemGetLawMessageResponse.NullableLawMessage))

	for i := 0; i < len(serviceItemGetLawMessageResponse.NullableLawMessage); i++ {
		copied.NullableLawMessage[i] = serviceItemGetLawMessageResponse.NullableLawMessage[i].Copy().(*ServiceItemLawMessage)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetLawMessageResponse *ServiceItemGetLawMessageResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetLawMessageResponse); !ok {
		return false
	}

	other := o.(*ServiceItemGetLawMessageResponse)

	if serviceItemGetLawMessageResponse.StructureVersion != other.StructureVersion {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemGetLawMessageResponse.StructureVersion))

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
