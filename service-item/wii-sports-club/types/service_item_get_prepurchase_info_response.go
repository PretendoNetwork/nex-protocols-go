// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemGetPrepurchaseInfoResponse holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemGetPrepurchaseInfoResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullablePrepurchaseInfo []*ServiceItemPrepurchaseInfo
}

// ExtractFrom extracts the ServiceItemGetPrepurchaseInfoResponse from the given readable
func (serviceItemGetPrepurchaseInfoResponse *ServiceItemGetPrepurchaseInfoResponse) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemGetPrepurchaseInfoResponse.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemGetPrepurchaseInfoResponse header. %s", err.Error())
	}

	nullablePrepurchaseInfo, err := nex.StreamReadListStructure(stream, NewServiceItemPrepurchaseInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo from stream. %s", err.Error())
	}

	serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo = nullablePrepurchaseInfo

	return nil
}

// WriteTo writes the ServiceItemGetPrepurchaseInfoResponse to the given writable
func (serviceItemGetPrepurchaseInfoResponse *ServiceItemGetPrepurchaseInfoResponse) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemGetPrepurchaseInfoResponse.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemGetPrepurchaseInfoResponse
func (serviceItemGetPrepurchaseInfoResponse *ServiceItemGetPrepurchaseInfoResponse) Copy() types.RVType {
	copied := NewServiceItemGetPrepurchaseInfoResponse()

	copied.StructureVersion = serviceItemGetPrepurchaseInfoResponse.StructureVersion

	copied.ServiceItemEShopResponse = serviceItemGetPrepurchaseInfoResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)

	copied.NullablePrepurchaseInfo = make([]*ServiceItemPrepurchaseInfo, len(serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo))

	for i := 0; i < len(serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo); i++ {
		copied.NullablePrepurchaseInfo[i] = serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo[i].Copy().(*ServiceItemPrepurchaseInfo)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetPrepurchaseInfoResponse *ServiceItemGetPrepurchaseInfoResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetPrepurchaseInfoResponse); !ok {
		return false
	}

	other := o.(*ServiceItemGetPrepurchaseInfoResponse)

	if serviceItemGetPrepurchaseInfoResponse.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemGetPrepurchaseInfoResponse.ParentType().Equals(other.ParentType()) {
		return false
	}

	if len(serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo) != len(other.NullablePrepurchaseInfo) {
		return false
	}

	for i := 0; i < len(serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo); i++ {
		if !serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo[i].Equals(other.NullablePrepurchaseInfo[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetPrepurchaseInfoResponse *ServiceItemGetPrepurchaseInfoResponse) String() string {
	return serviceItemGetPrepurchaseInfoResponse.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetPrepurchaseInfoResponse *ServiceItemGetPrepurchaseInfoResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetPrepurchaseInfoResponse{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, serviceItemGetPrepurchaseInfoResponse.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemGetPrepurchaseInfoResponse.StructureVersion))

	if len(serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo) == 0 {
		b.WriteString(fmt.Sprintf("%sNullablePrepurchaseInfo: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sNullablePrepurchaseInfo: [\n", indentationValues))

		for i := 0; i < len(serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo); i++ {
			str := serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo)-1 {
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

// NewServiceItemGetPrepurchaseInfoResponse returns a new ServiceItemGetPrepurchaseInfoResponse
func NewServiceItemGetPrepurchaseInfoResponse() *ServiceItemGetPrepurchaseInfoResponse {
	serviceItemGetPrepurchaseInfoResponse := &ServiceItemGetPrepurchaseInfoResponse{}

	serviceItemGetPrepurchaseInfoResponse.ServiceItemEShopResponse = NewServiceItemEShopResponse()
	serviceItemGetPrepurchaseInfoResponse.SetParentType(serviceItemGetPrepurchaseInfoResponse.ServiceItemEShopResponse)

	return serviceItemGetPrepurchaseInfoResponse
}
