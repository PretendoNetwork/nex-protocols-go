// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetPrepurchaseInfoResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetPrepurchaseInfoResponse struct {
	nex.Structure
	*ServiceItemEShopResponse
	NullablePrepurchaseInfo []*ServiceItemPrepurchaseInfo
}

// ExtractFromStream extracts a ServiceItemGetPrepurchaseInfoResponse structure from a stream
func (serviceItemGetPrepurchaseInfoResponse *ServiceItemGetPrepurchaseInfoResponse) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	nullablePrepurchaseInfo, err := nex.StreamReadListStructure(stream, NewServiceItemPrepurchaseInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo from stream. %s", err.Error())
	}

	serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo = nullablePrepurchaseInfo

	return nil
}

// Bytes encodes the ServiceItemGetPrepurchaseInfoResponse and returns a byte array
func (serviceItemGetPrepurchaseInfoResponse *ServiceItemGetPrepurchaseInfoResponse) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListStructure(serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetPrepurchaseInfoResponse
func (serviceItemGetPrepurchaseInfoResponse *ServiceItemGetPrepurchaseInfoResponse) Copy() nex.StructureInterface {
	copied := NewServiceItemGetPrepurchaseInfoResponse()

	copied.SetStructureVersion(serviceItemGetPrepurchaseInfoResponse.StructureVersion())

	copied.ServiceItemEShopResponse = serviceItemGetPrepurchaseInfoResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.SetParentType(copied.ServiceItemEShopResponse)

	copied.NullablePrepurchaseInfo = make([]*ServiceItemPrepurchaseInfo, len(serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo))

	for i := 0; i < len(serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo); i++ {
		copied.NullablePrepurchaseInfo[i] = serviceItemGetPrepurchaseInfoResponse.NullablePrepurchaseInfo[i].Copy().(*ServiceItemPrepurchaseInfo)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetPrepurchaseInfoResponse *ServiceItemGetPrepurchaseInfoResponse) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetPrepurchaseInfoResponse)

	if serviceItemGetPrepurchaseInfoResponse.StructureVersion() != other.StructureVersion() {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemGetPrepurchaseInfoResponse.StructureVersion()))

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
