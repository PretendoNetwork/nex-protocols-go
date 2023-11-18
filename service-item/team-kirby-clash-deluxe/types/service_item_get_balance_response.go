// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetBalanceResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetBalanceResponse struct {
	nex.Structure
	*ServiceItemEShopResponse
	NullableBalance []*ServiceItemAmount
}

// ExtractFromStream extracts a ServiceItemGetBalanceResponse structure from a stream
func (serviceItemGetBalanceResponse *ServiceItemGetBalanceResponse) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	nullableBalance, err := nex.StreamReadListStructure(stream, NewServiceItemAmount())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceResponse.NullableBalance from stream. %s", err.Error())
	}

	serviceItemGetBalanceResponse.NullableBalance = nullableBalance

	return nil
}

// Bytes encodes the ServiceItemGetBalanceResponse and returns a byte array
func (serviceItemGetBalanceResponse *ServiceItemGetBalanceResponse) Bytes(stream *nex.StreamOut) []byte {
	nex.StreamWriteListStructure(stream, serviceItemGetBalanceResponse.NullableBalance)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetBalanceResponse
func (serviceItemGetBalanceResponse *ServiceItemGetBalanceResponse) Copy() nex.StructureInterface {
	copied := NewServiceItemGetBalanceResponse()

	copied.SetStructureVersion(serviceItemGetBalanceResponse.StructureVersion())

	copied.ServiceItemEShopResponse = serviceItemGetBalanceResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)
	copied.SetParentType(copied.ServiceItemEShopResponse)

	copied.NullableBalance = make([]*ServiceItemAmount, len(serviceItemGetBalanceResponse.NullableBalance))

	for i := 0; i < len(serviceItemGetBalanceResponse.NullableBalance); i++ {
		copied.NullableBalance[i] = serviceItemGetBalanceResponse.NullableBalance[i].Copy().(*ServiceItemAmount)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetBalanceResponse *ServiceItemGetBalanceResponse) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetBalanceResponse)

	if serviceItemGetBalanceResponse.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !serviceItemGetBalanceResponse.ParentType().Equals(other.ParentType()) {
		return false
	}

	if len(serviceItemGetBalanceResponse.NullableBalance) != len(other.NullableBalance) {
		return false
	}

	for i := 0; i < len(serviceItemGetBalanceResponse.NullableBalance); i++ {
		if !serviceItemGetBalanceResponse.NullableBalance[i].Equals(other.NullableBalance[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetBalanceResponse *ServiceItemGetBalanceResponse) String() string {
	return serviceItemGetBalanceResponse.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetBalanceResponse *ServiceItemGetBalanceResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetBalanceResponse{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, serviceItemGetBalanceResponse.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemGetBalanceResponse.StructureVersion()))

	if len(serviceItemGetBalanceResponse.NullableBalance) == 0 {
		b.WriteString(fmt.Sprintf("%sNullableBalance: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sNullableBalance: [\n", indentationValues))

		for i := 0; i < len(serviceItemGetBalanceResponse.NullableBalance); i++ {
			str := serviceItemGetBalanceResponse.NullableBalance[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemGetBalanceResponse.NullableBalance)-1 {
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

// NewServiceItemGetBalanceResponse returns a new ServiceItemGetBalanceResponse
func NewServiceItemGetBalanceResponse() *ServiceItemGetBalanceResponse {
	serviceItemGetBalanceResponse := &ServiceItemGetBalanceResponse{}

	serviceItemGetBalanceResponse.ServiceItemEShopResponse = NewServiceItemEShopResponse()
	serviceItemGetBalanceResponse.SetParentType(serviceItemGetBalanceResponse.ServiceItemEShopResponse)

	return serviceItemGetBalanceResponse
}
