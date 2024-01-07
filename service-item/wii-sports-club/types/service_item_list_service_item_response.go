// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemListServiceItemResponse holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemListServiceItemResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullableCatalog []*ServiceItemCatalog
}

// ExtractFrom extracts the ServiceItemListServiceItemResponse from the given readable
func (serviceItemListServiceItemResponse *ServiceItemListServiceItemResponse) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemListServiceItemResponse.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemListServiceItemResponse header. %s", err.Error())
	}

	nullableCatalog, err := nex.StreamReadListStructure(stream, NewServiceItemCatalog())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListServiceItemResponse.NullableCatalog from stream. %s", err.Error())
	}

	serviceItemListServiceItemResponse.NullableCatalog = nullableCatalog

	return nil
}

// WriteTo writes the ServiceItemListServiceItemResponse to the given writable
func (serviceItemListServiceItemResponse *ServiceItemListServiceItemResponse) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemListServiceItemResponse.NullableCatalog.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemListServiceItemResponse.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemListServiceItemResponse
func (serviceItemListServiceItemResponse *ServiceItemListServiceItemResponse) Copy() types.RVType {
	copied := NewServiceItemListServiceItemResponse()

	copied.StructureVersion = serviceItemListServiceItemResponse.StructureVersion

	copied.ServiceItemEShopResponse = serviceItemListServiceItemResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)

	copied.NullableCatalog = make([]*ServiceItemCatalog, len(serviceItemListServiceItemResponse.NullableCatalog))

	for i := 0; i < len(serviceItemListServiceItemResponse.NullableCatalog); i++ {
		copied.NullableCatalog[i] = serviceItemListServiceItemResponse.NullableCatalog[i].Copy().(*ServiceItemCatalog)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemListServiceItemResponse *ServiceItemListServiceItemResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemListServiceItemResponse); !ok {
		return false
	}

	other := o.(*ServiceItemListServiceItemResponse)

	if serviceItemListServiceItemResponse.StructureVersion != other.StructureVersion {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemListServiceItemResponse.StructureVersion))

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
