// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemUseServiceItemResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemUseServiceItemResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullableUsedInfo []*ServiceItemUsedInfo
}

// ExtractFrom extracts the ServiceItemUseServiceItemResponse from the given readable
func (serviceItemUseServiceItemResponse *ServiceItemUseServiceItemResponse) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemUseServiceItemResponse.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemUseServiceItemResponse header. %s", err.Error())
	}

	nullableUsedInfo, err := nex.StreamReadListStructure(stream, NewServiceItemUsedInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUseServiceItemResponse.NullableUsedInfo from stream. %s", err.Error())
	}

	serviceItemUseServiceItemResponse.NullableUsedInfo = nullableUsedInfo

	return nil
}

// WriteTo writes the ServiceItemUseServiceItemResponse to the given writable
func (serviceItemUseServiceItemResponse *ServiceItemUseServiceItemResponse) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemUseServiceItemResponse.NullableUsedInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemUseServiceItemResponse.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemUseServiceItemResponse
func (serviceItemUseServiceItemResponse *ServiceItemUseServiceItemResponse) Copy() types.RVType {
	copied := NewServiceItemUseServiceItemResponse()

	copied.StructureVersion = serviceItemUseServiceItemResponse.StructureVersion

	copied.ServiceItemEShopResponse = serviceItemUseServiceItemResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)

	copied.NullableUsedInfo = make([]*ServiceItemUsedInfo, len(serviceItemUseServiceItemResponse.NullableUsedInfo))

	for i := 0; i < len(serviceItemUseServiceItemResponse.NullableUsedInfo); i++ {
		copied.NullableUsedInfo[i] = serviceItemUseServiceItemResponse.NullableUsedInfo[i].Copy().(*ServiceItemUsedInfo)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemUseServiceItemResponse *ServiceItemUseServiceItemResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemUseServiceItemResponse); !ok {
		return false
	}

	other := o.(*ServiceItemUseServiceItemResponse)

	if serviceItemUseServiceItemResponse.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemUseServiceItemResponse.ParentType().Equals(other.ParentType()) {
		return false
	}

	if len(serviceItemUseServiceItemResponse.NullableUsedInfo) != len(other.NullableUsedInfo) {
		return false
	}

	for i := 0; i < len(serviceItemUseServiceItemResponse.NullableUsedInfo); i++ {
		if !serviceItemUseServiceItemResponse.NullableUsedInfo[i].Equals(other.NullableUsedInfo[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemUseServiceItemResponse *ServiceItemUseServiceItemResponse) String() string {
	return serviceItemUseServiceItemResponse.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemUseServiceItemResponse *ServiceItemUseServiceItemResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemUseServiceItemResponse{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, serviceItemUseServiceItemResponse.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemUseServiceItemResponse.StructureVersion))

	if len(serviceItemUseServiceItemResponse.NullableUsedInfo) == 0 {
		b.WriteString(fmt.Sprintf("%sNullableUsedInfo: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sNullableUsedInfo: [\n", indentationValues))

		for i := 0; i < len(serviceItemUseServiceItemResponse.NullableUsedInfo); i++ {
			str := serviceItemUseServiceItemResponse.NullableUsedInfo[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemUseServiceItemResponse.NullableUsedInfo)-1 {
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

// NewServiceItemUseServiceItemResponse returns a new ServiceItemUseServiceItemResponse
func NewServiceItemUseServiceItemResponse() *ServiceItemUseServiceItemResponse {
	serviceItemUseServiceItemResponse := &ServiceItemUseServiceItemResponse{}

	serviceItemUseServiceItemResponse.ServiceItemEShopResponse = NewServiceItemEShopResponse()
	serviceItemUseServiceItemResponse.SetParentType(serviceItemUseServiceItemResponse.ServiceItemEShopResponse)

	return serviceItemUseServiceItemResponse
}
