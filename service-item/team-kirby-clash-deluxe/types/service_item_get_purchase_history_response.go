// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetPurchaseHistoryResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetPurchaseHistoryResponse struct {
	nex.Structure
	*ServiceItemEShopResponse
	NullablePurchaseHistory []*ServiceItemPurchaseHistory
}

// ExtractFromStream extracts a ServiceItemGetPurchaseHistoryResponse structure from a stream
func (serviceItemGetPurchaseHistoryResponse *ServiceItemGetPurchaseHistoryResponse) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	nullablePurchaseHistory, err := nex.StreamReadListStructure(stream, NewServiceItemPurchaseHistory())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPurchaseHistoryResponse.NullablePurchaseHistory from stream. %s", err.Error())
	}

	serviceItemGetPurchaseHistoryResponse.NullablePurchaseHistory = nullablePurchaseHistory

	return nil
}

// Bytes encodes the ServiceItemGetPurchaseHistoryResponse and returns a byte array
func (serviceItemGetPurchaseHistoryResponse *ServiceItemGetPurchaseHistoryResponse) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListStructure(serviceItemGetPurchaseHistoryResponse.NullablePurchaseHistory)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetPurchaseHistoryResponse
func (serviceItemGetPurchaseHistoryResponse *ServiceItemGetPurchaseHistoryResponse) Copy() nex.StructureInterface {
	copied := NewServiceItemGetPurchaseHistoryResponse()

	copied.SetStructureVersion(serviceItemGetPurchaseHistoryResponse.StructureVersion())

	copied.ServiceItemEShopResponse = serviceItemGetPurchaseHistoryResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.SetParentType(copied.ServiceItemEShopResponse)

	copied.NullablePurchaseHistory = make([]*ServiceItemPurchaseHistory, len(serviceItemGetPurchaseHistoryResponse.NullablePurchaseHistory))

	for i := 0; i < len(serviceItemGetPurchaseHistoryResponse.NullablePurchaseHistory); i++ {
		copied.NullablePurchaseHistory[i] = serviceItemGetPurchaseHistoryResponse.NullablePurchaseHistory[i].Copy().(*ServiceItemPurchaseHistory)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetPurchaseHistoryResponse *ServiceItemGetPurchaseHistoryResponse) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetPurchaseHistoryResponse)

	if serviceItemGetPurchaseHistoryResponse.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !serviceItemGetPurchaseHistoryResponse.ParentType().Equals(other.ParentType()) {
		return false
	}

	if len(serviceItemGetPurchaseHistoryResponse.NullablePurchaseHistory) != len(other.NullablePurchaseHistory) {
		return false
	}

	for i := 0; i < len(serviceItemGetPurchaseHistoryResponse.NullablePurchaseHistory); i++ {
		if !serviceItemGetPurchaseHistoryResponse.NullablePurchaseHistory[i].Equals(other.NullablePurchaseHistory[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetPurchaseHistoryResponse *ServiceItemGetPurchaseHistoryResponse) String() string {
	return serviceItemGetPurchaseHistoryResponse.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetPurchaseHistoryResponse *ServiceItemGetPurchaseHistoryResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetPurchaseHistoryResponse{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, serviceItemGetPurchaseHistoryResponse.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemGetPurchaseHistoryResponse.StructureVersion()))

	if len(serviceItemGetPurchaseHistoryResponse.NullablePurchaseHistory) == 0 {
		b.WriteString(fmt.Sprintf("%sNullablePurchaseHistory: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sNullablePurchaseHistory: [\n", indentationValues))

		for i := 0; i < len(serviceItemGetPurchaseHistoryResponse.NullablePurchaseHistory); i++ {
			str := serviceItemGetPurchaseHistoryResponse.NullablePurchaseHistory[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemGetPurchaseHistoryResponse.NullablePurchaseHistory)-1 {
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

// NewServiceItemGetPurchaseHistoryResponse returns a new ServiceItemGetPurchaseHistoryResponse
func NewServiceItemGetPurchaseHistoryResponse() *ServiceItemGetPurchaseHistoryResponse {
	serviceItemGetPurchaseHistoryResponse := &ServiceItemGetPurchaseHistoryResponse{}

	serviceItemGetPurchaseHistoryResponse.ServiceItemEShopResponse = NewServiceItemEShopResponse()
	serviceItemGetPurchaseHistoryResponse.SetParentType(serviceItemGetPurchaseHistoryResponse.ServiceItemEShopResponse)

	return serviceItemGetPurchaseHistoryResponse
}
