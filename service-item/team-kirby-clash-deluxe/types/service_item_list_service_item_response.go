// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemListServiceItemResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemListServiceItemResponse struct {
	nex.Structure
	*ServiceItemEShopResponse
	NullableCatalog []*ServiceItemCatalog
}

// ExtractFromStream extracts a ServiceItemListServiceItemResponse structure from a stream
func (serviceItemListServiceItemResponse *ServiceItemListServiceItemResponse) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	nullableCatalog, err := nex.StreamReadListStructure(stream, NewServiceItemCatalog())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemResponse.NullableCatalog from stream. %s", err.Error())
	}

	serviceItemListServiceItemResponse.NullableCatalog = nullableCatalog

	return nil
}

// Bytes encodes the ServiceItemListServiceItemResponse and returns a byte array
func (serviceItemListServiceItemResponse *ServiceItemListServiceItemResponse) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListStructure(serviceItemListServiceItemResponse.NullableCatalog)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemListServiceItemResponse
func (serviceItemListServiceItemResponse *ServiceItemListServiceItemResponse) Copy() nex.StructureInterface {
	copied := NewServiceItemListServiceItemResponse()

	copied.SetStructureVersion(serviceItemListServiceItemResponse.StructureVersion())

	copied.ServiceItemEShopResponse = serviceItemListServiceItemResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.SetParentType(copied.ServiceItemEShopResponse)

	copied.NullableCatalog = make([]*ServiceItemCatalog, len(serviceItemListServiceItemResponse.NullableCatalog))

	for i := 0; i < len(serviceItemListServiceItemResponse.NullableCatalog); i++ {
		copied.NullableCatalog[i] = serviceItemListServiceItemResponse.NullableCatalog[i].Copy().(*ServiceItemCatalog)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemListServiceItemResponse *ServiceItemListServiceItemResponse) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemListServiceItemResponse)

	if serviceItemListServiceItemResponse.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !serviceItemListServiceItemResponse.ParentType().Equals(other.ParentType()) {
		return false
	}

	if len(serviceItemListServiceItemResponse.NullableCatalog) != len(other.NullableCatalog) {
		return false
	}

	for i := 0; i < len(serviceItemListServiceItemResponse.NullableCatalog); i++ {
		if !serviceItemListServiceItemResponse.NullableCatalog[i].Equals(other.NullableCatalog[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemListServiceItemResponse *ServiceItemListServiceItemResponse) String() string {
	return serviceItemListServiceItemResponse.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemListServiceItemResponse *ServiceItemListServiceItemResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemListServiceItemResponse{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, serviceItemListServiceItemResponse.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemListServiceItemResponse.StructureVersion()))

	if len(serviceItemListServiceItemResponse.NullableCatalog) == 0 {
		b.WriteString(fmt.Sprintf("%sNullableCatalog: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sNullableCatalog: [\n", indentationValues))

		for i := 0; i < len(serviceItemListServiceItemResponse.NullableCatalog); i++ {
			str := serviceItemListServiceItemResponse.NullableCatalog[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemListServiceItemResponse.NullableCatalog)-1 {
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

// NewServiceItemListServiceItemResponse returns a new ServiceItemListServiceItemResponse
func NewServiceItemListServiceItemResponse() *ServiceItemListServiceItemResponse {
	serviceItemListServiceItemResponse := &ServiceItemListServiceItemResponse{}

	serviceItemListServiceItemResponse.ServiceItemEShopResponse = NewServiceItemEShopResponse()
	serviceItemListServiceItemResponse.SetParentType(serviceItemListServiceItemResponse.ServiceItemEShopResponse)

	return serviceItemListServiceItemResponse
}
