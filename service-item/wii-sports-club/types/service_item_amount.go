// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemAmount holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemAmount struct {
	nex.Structure
	FormattedAmount string
	Currency string
	RawValue string
}

// ExtractFromStream extracts a ServiceItemAmount structure from a stream
func (serviceItemAmount *ServiceItemAmount) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemAmount.FormattedAmount, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAmount.FormattedAmount from stream. %s", err.Error())
	}

	serviceItemAmount.Currency, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAmount.Currency from stream. %s", err.Error())
	}

	serviceItemAmount.RawValue, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAmount.RawValue from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemAmount and returns a byte array
func (serviceItemAmount *ServiceItemAmount) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemAmount.FormattedAmount)
	stream.WriteString(serviceItemAmount.Currency)
	stream.WriteString(serviceItemAmount.RawValue)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemAmount
func (serviceItemAmount *ServiceItemAmount) Copy() nex.StructureInterface {
	copied := NewServiceItemAmount()

	copied.FormattedAmount = serviceItemAmount.FormattedAmount
	copied.Currency = serviceItemAmount.Currency
	copied.RawValue = serviceItemAmount.RawValue

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAmount *ServiceItemAmount) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemAmount)

	if serviceItemAmount.FormattedAmount != other.FormattedAmount {
		return false
	}

	if serviceItemAmount.Currency != other.Currency {
		return false
	}

	if serviceItemAmount.RawValue != other.RawValue {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemAmount *ServiceItemAmount) String() string {
	return serviceItemAmount.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemAmount *ServiceItemAmount) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAmount{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemAmount.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sFormattedAmount: %q,\n", indentationValues, serviceItemAmount.FormattedAmount))
	b.WriteString(fmt.Sprintf("%sCurrency: %q,\n", indentationValues, serviceItemAmount.Currency))
	b.WriteString(fmt.Sprintf("%sRawValue: %q,\n", indentationValues, serviceItemAmount.RawValue))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAmount returns a new ServiceItemAmount
func NewServiceItemAmount() *ServiceItemAmount {
	return &ServiceItemAmount{}
}
