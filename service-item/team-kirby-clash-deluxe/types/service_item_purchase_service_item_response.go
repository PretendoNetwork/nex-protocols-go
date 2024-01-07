// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemPurchaseServiceItemResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemPurchaseServiceItemResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullablePurchaseInfo []*ServiceItemPurchaseInfo
}

// ExtractFrom extracts the ServiceItemPurchaseServiceItemResponse from the given readable
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemPurchaseServiceItemResponse.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemPurchaseServiceItemResponse header. %s", err.Error())
	}

	nullablePurchaseInfo, err := nex.StreamReadListStructure(stream, NewServiceItemPurchaseInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemResponse.NullablePurchaseInfo from stream. %s", err.Error())
	}

	serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo = nullablePurchaseInfo

	return nil
}

// WriteTo writes the ServiceItemPurchaseServiceItemResponse to the given writable
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemPurchaseServiceItemResponse.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemPurchaseServiceItemResponse
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) Copy() types.RVType {
	copied := NewServiceItemPurchaseServiceItemResponse()

	copied.StructureVersion = serviceItemPurchaseServiceItemResponse.StructureVersion

	copied.ServiceItemEShopResponse = serviceItemPurchaseServiceItemResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)

	copied.NullablePurchaseInfo = make([]*ServiceItemPurchaseInfo, len(serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo))

	for i := 0; i < len(serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo); i++ {
		copied.NullablePurchaseInfo[i] = serviceItemPurchaseServiceItemResponse.NullablePurchaseInfo[i].Copy().(*ServiceItemPurchaseInfo)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPurchaseServiceItemResponse *ServiceItemPurchaseServiceItemResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPurchaseServiceItemResponse); !ok {
		return false
	}

	other := o.(*ServiceItemPurchaseServiceItemResponse)

	if serviceItemPurchaseServiceItemResponse.StructureVersion != other.StructureVersion {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemPurchaseServiceItemResponse.StructureVersion))

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
