// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemGetBalanceResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetBalanceResponse struct {
	types.Structure
	*ServiceItemEShopResponse
	NullableBalance []*ServiceItemAmount
}

// ExtractFrom extracts the ServiceItemGetBalanceResponse from the given readable
func (serviceItemGetBalanceResponse *ServiceItemGetBalanceResponse) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemGetBalanceResponse.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemGetBalanceResponse header. %s", err.Error())
	}

	nullableBalance, err := nex.StreamReadListStructure(stream, NewServiceItemAmount())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceResponse.NullableBalance from stream. %s", err.Error())
	}

	serviceItemGetBalanceResponse.NullableBalance = nullableBalance

	return nil
}

// WriteTo writes the ServiceItemGetBalanceResponse to the given writable
func (serviceItemGetBalanceResponse *ServiceItemGetBalanceResponse) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemGetBalanceResponse.NullableBalance.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemGetBalanceResponse.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemGetBalanceResponse
func (serviceItemGetBalanceResponse *ServiceItemGetBalanceResponse) Copy() types.RVType {
	copied := NewServiceItemGetBalanceResponse()

	copied.StructureVersion = serviceItemGetBalanceResponse.StructureVersion

	copied.ServiceItemEShopResponse = serviceItemGetBalanceResponse.ServiceItemEShopResponse.Copy().(*ServiceItemEShopResponse)

	copied.NullableBalance = make([]*ServiceItemAmount, len(serviceItemGetBalanceResponse.NullableBalance))

	for i := 0; i < len(serviceItemGetBalanceResponse.NullableBalance); i++ {
		copied.NullableBalance[i] = serviceItemGetBalanceResponse.NullableBalance[i].Copy().(*ServiceItemAmount)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetBalanceResponse *ServiceItemGetBalanceResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetBalanceResponse); !ok {
		return false
	}

	other := o.(*ServiceItemGetBalanceResponse)

	if serviceItemGetBalanceResponse.StructureVersion != other.StructureVersion {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemGetBalanceResponse.StructureVersion))

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
