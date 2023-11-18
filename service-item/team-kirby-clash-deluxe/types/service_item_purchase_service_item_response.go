// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemPurchaseServiceItemResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemPurchaseServiceItemResponse struct {
	nex.Structure
	*ServiceItemEShopResponse
	NullablePurchaseInfo []*ServiceItemPurchaseInfo
}

// ExtractFromStream extracts a ServiceItemPurchaseServiceItemResponse structure from a stream
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	nullablePurchaseInfo, err := nex.StreamReadListStructure(stream, NewServiceItemPurchaseInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemResponse.NullablePurchaseInfo from stream. %s", err.Error())
	}

	serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo = nullablePurchaseInfo

	return nil
}

// Bytes encodes the ServiceItemPurchaseServiceItemResponse and returns a byte array
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListStructure(serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemPurchaseServiceItemResponse
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) Copy() nex.StructureInterface {
	copied := NewServiceItemPurchaseServiceItemResponse()

	copied.SetStructureVersion(serviceItemPurchaseServiceItemResponse.StructureVersion())

	copied.ServiceItemEShopResponse = serviceItemPurchaseServiceItemResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.SetParentType(copied.ServiceItemEShopResponse)

	copied.NullablePurchaseInfo = make([]*ServiceItemPurchaseInfo, len(serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo))

	for i := 0; i < len(serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo); i++ {
		copied.NullablePurchaseInfo[i] = serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo[i].Copy().(*ServiceItemPurchaseInfo)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemPurchaseServiceItemResponse)

	if serviceItemPurchaseServiceItemResponse.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !serviceItemPurchaseServiceItemResponse.ParentType().Equals(other.ParentType()) {
		return false
	}

	if len(serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo) != len(other.NullablePurchaseInfo) {
		return false
	}

	for i := 0; i < len(serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo); i++ {
		if !serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo[i].Equals(other.NullablePurchaseInfo[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) String() string {
	return serviceItemPurchaseServiceItemResponse.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPurchaseServiceItemResponse{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, serviceItemPurchaseServiceItemResponse.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemPurchaseServiceItemResponse.StructureVersion()))

	if len(serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo) == 0 {
		b.WriteString(fmt.Sprintf("%sNullablePurchaseInfo: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sNullablePurchaseInfo: [\n", indentationValues))

		for i := 0; i < len(serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo); i++ {
			str := serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo)-1 {
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

// NewServiceItemPurchaseServiceItemResponse returns a new ServiceItemPurchaseServiceItemResponse
func NewServiceItemPurchaseServiceItemResponse() *ServiceItemPurchaseServiceItemResponse {
	serviceItemPurchaseServiceItemResponse := &ServiceItemPurchaseServiceItemResponse{}

	serviceItemPurchaseServiceItemResponse.ServiceItemEShopResponse = NewServiceItemEShopResponse()
	serviceItemPurchaseServiceItemResponse.SetParentType(serviceItemPurchaseServiceItemResponse.ServiceItemEShopResponse)

	return serviceItemPurchaseServiceItemResponse
}
