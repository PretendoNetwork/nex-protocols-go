// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemPurchaseServiceItemResponse holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemPurchaseServiceItemResponse struct {
	nex.Structure
	*ServiceItemEShopResponse
	NullablePurchaceInfo []*ServiceItemPurchaceInfo
}

// ExtractFromStream extracts a ServiceItemPurchaseServiceItemResponse structure from a stream
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	nullablePurchaceInfo, err := stream.ReadListStructure(NewServiceItemPurchaceInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemResponse.NullablePurchaceInfo from stream. %s", err.Error())
	}

	serviceItemPurchaseServiceItemResponse.NullablePurchaceInfo = nullablePurchaceInfo.([]*ServiceItemPurchaceInfo)

	return nil
}

// Bytes encodes the ServiceItemPurchaseServiceItemResponse and returns a byte array
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListStructure(serviceItemPurchaseServiceItemResponse.NullablePurchaceInfo)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemPurchaseServiceItemResponse
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) Copy() nex.StructureInterface {
	copied := NewServiceItemPurchaseServiceItemResponse()

	copied.ServiceItemEShopResponse = serviceItemPurchaseServiceItemResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.SetParentType(copied.ServiceItemEShopResponse)

	copied.NullablePurchaceInfo = make([]*ServiceItemPurchaceInfo, len(serviceItemPurchaseServiceItemResponse.NullablePurchaceInfo))

	for i := 0; i < len(serviceItemPurchaseServiceItemResponse.NullablePurchaceInfo); i++ {
		copied.NullablePurchaceInfo[i] = serviceItemPurchaseServiceItemResponse.NullablePurchaceInfo[i].Copy().(*ServiceItemPurchaceInfo)
	}


	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemPurchaseServiceItemResponse)

	if !serviceItemPurchaseServiceItemResponse.ParentType().Equals(other.ParentType()) {
		return false
	}

	if len(serviceItemPurchaseServiceItemResponse.NullablePurchaceInfo) != len(other.NullablePurchaceInfo) {
		return false
	}

	for i := 0; i < len(serviceItemPurchaseServiceItemResponse.NullablePurchaceInfo); i++ {
		if !serviceItemPurchaseServiceItemResponse.NullablePurchaceInfo[i].Equals(other.NullablePurchaceInfo[i]) {
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

	if len(serviceItemPurchaseServiceItemResponse.NullablePurchaceInfo) == 0 {
		b.WriteString(fmt.Sprintf("%sNullablePurchaceInfo: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sNullablePurchaceInfo: [\n", indentationValues))

		for i := 0; i < len(serviceItemPurchaseServiceItemResponse.NullablePurchaceInfo); i++ {
			str := serviceItemPurchaseServiceItemResponse.NullablePurchaceInfo[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemPurchaseServiceItemResponse.NullablePurchaceInfo)-1 {
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
