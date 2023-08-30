// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetServiceItemRightResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetServiceItemRightResponse struct {
	nex.Structure
	*ServiceItemEShopResponse
	NullableRightInfos []*ServiceItemRightInfos
}

// ExtractFromStream extracts a ServiceItemGetServiceItemRightResponse structure from a stream
func (serviceItemGetServiceItemRightResponse *ServiceItemGetServiceItemRightResponse) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	nullableRightInfos, err := stream.ReadListStructure(NewServiceItemRightInfos())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightResponse.NullableRightInfos from stream. %s", err.Error())
	}

	serviceItemGetServiceItemRightResponse.NullableRightInfos = nullableRightInfos.([]*ServiceItemRightInfos)

	return nil
}

// Bytes encodes the ServiceItemGetServiceItemRightResponse and returns a byte array
func (serviceItemGetServiceItemRightResponse *ServiceItemGetServiceItemRightResponse) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListStructure(serviceItemGetServiceItemRightResponse.NullableRightInfos)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetServiceItemRightResponse
func (serviceItemGetServiceItemRightResponse *ServiceItemGetServiceItemRightResponse) Copy() nex.StructureInterface {
	copied := NewServiceItemGetServiceItemRightResponse()

	copied.SetStructureVersion(serviceItemGetServiceItemRightResponse.StructureVersion())

	copied.ServiceItemEShopResponse = serviceItemGetServiceItemRightResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.SetParentType(copied.ServiceItemEShopResponse)

	copied.NullableRightInfos = make([]*ServiceItemRightInfos, len(serviceItemGetServiceItemRightResponse.NullableRightInfos))

	for i := 0; i < len(serviceItemGetServiceItemRightResponse.NullableRightInfos); i++ {
		copied.NullableRightInfos[i] = serviceItemGetServiceItemRightResponse.NullableRightInfos[i].Copy().(*ServiceItemRightInfos)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetServiceItemRightResponse *ServiceItemGetServiceItemRightResponse) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetServiceItemRightResponse)

	if serviceItemGetServiceItemRightResponse.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !serviceItemGetServiceItemRightResponse.ParentType().Equals(other.ParentType()) {
		return false
	}

	if len(serviceItemGetServiceItemRightResponse.NullableRightInfos) != len(other.NullableRightInfos) {
		return false
	}

	for i := 0; i < len(serviceItemGetServiceItemRightResponse.NullableRightInfos); i++ {
		if !serviceItemGetServiceItemRightResponse.NullableRightInfos[i].Equals(other.NullableRightInfos[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetServiceItemRightResponse *ServiceItemGetServiceItemRightResponse) String() string {
	return serviceItemGetServiceItemRightResponse.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetServiceItemRightResponse *ServiceItemGetServiceItemRightResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetServiceItemRightResponse{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, serviceItemGetServiceItemRightResponse.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemGetServiceItemRightResponse.StructureVersion()))

	if len(serviceItemGetServiceItemRightResponse.NullableRightInfos) == 0 {
		b.WriteString(fmt.Sprintf("%sNullableRightInfos: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sNullableRightInfos: [\n", indentationValues))

		for i := 0; i < len(serviceItemGetServiceItemRightResponse.NullableRightInfos); i++ {
			str := serviceItemGetServiceItemRightResponse.NullableRightInfos[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemGetServiceItemRightResponse.NullableRightInfos)-1 {
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

// NewServiceItemGetServiceItemRightResponse returns a new ServiceItemGetServiceItemRightResponse
func NewServiceItemGetServiceItemRightResponse() *ServiceItemGetServiceItemRightResponse {
	serviceItemGetServiceItemRightResponse := &ServiceItemGetServiceItemRightResponse{}

	serviceItemGetServiceItemRightResponse.ServiceItemEShopResponse = NewServiceItemEShopResponse()
	serviceItemGetServiceItemRightResponse.SetParentType(serviceItemGetServiceItemRightResponse.ServiceItemEShopResponse)

	return serviceItemGetServiceItemRightResponse
}
