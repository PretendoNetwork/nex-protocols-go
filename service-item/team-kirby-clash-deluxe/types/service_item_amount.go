// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemAmount holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemAmount struct {
	types.Structure
	FormattedAmount string
	Currency        string
	RawValue        string
}

// ExtractFrom extracts the ServiceItemAmount from the given readable
func (serviceItemAmount *ServiceItemAmount) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemAmount.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemAmount header. %s", err.Error())
	}

	err = serviceItemAmount.FormattedAmount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAmount.FormattedAmount from stream. %s", err.Error())
	}

	err = serviceItemAmount.Currency.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAmount.Currency from stream. %s", err.Error())
	}

	err = serviceItemAmount.RawValue.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAmount.RawValue from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemAmount to the given writable
func (serviceItemAmount *ServiceItemAmount) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemAmount.FormattedAmount.WriteTo(contentWritable)
	serviceItemAmount.Currency.WriteTo(contentWritable)
	serviceItemAmount.RawValue.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemAmount.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemAmount
func (serviceItemAmount *ServiceItemAmount) Copy() types.RVType {
	copied := NewServiceItemAmount()

	copied.StructureVersion = serviceItemAmount.StructureVersion

	copied.FormattedAmount = serviceItemAmount.FormattedAmount
	copied.Currency = serviceItemAmount.Currency
	copied.RawValue = serviceItemAmount.RawValue

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAmount *ServiceItemAmount) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemAmount); !ok {
		return false
	}

	other := o.(*ServiceItemAmount)

	if serviceItemAmount.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemAmount.FormattedAmount.Equals(other.FormattedAmount) {
		return false
	}

	if !serviceItemAmount.Currency.Equals(other.Currency) {
		return false
	}

	if !serviceItemAmount.RawValue.Equals(other.RawValue) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemAmount.StructureVersion))
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
