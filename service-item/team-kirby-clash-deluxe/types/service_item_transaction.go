// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemTransaction holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemTransaction struct {
	types.Structure
	TransactionID          string
	ExtTransactionID       string
	Time                   *types.DateTime
	TransactionType        *types.PrimitiveU32
	TransactionDescription string
	TransactionAmount      *ServiceItemAmount
	ItemCode               string
	ReferenceID            string
	Limitation             *ServiceItemLimitation
}

// ExtractFrom extracts the ServiceItemTransaction from the given readable
func (serviceItemTransaction *ServiceItemTransaction) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemTransaction.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemTransaction header. %s", err.Error())
	}

	err = serviceItemTransaction.TransactionID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.TransactionID from stream. %s", err.Error())
	}

	err = serviceItemTransaction.ExtTransactionID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.ExtTransactionID from stream. %s", err.Error())
	}

	err = serviceItemTransaction.Time.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.Time from stream. %s", err.Error())
	}

	err = serviceItemTransaction.TransactionType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.TransactionType from stream. %s", err.Error())
	}

	err = serviceItemTransaction.TransactionDescription.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.TransactionDescription from stream. %s", err.Error())
	}

	err = serviceItemTransaction.TransactionAmount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.TransactionAmount from stream. %s", err.Error())
	}

	err = serviceItemTransaction.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.ItemCode from stream. %s", err.Error())
	}

	err = serviceItemTransaction.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.ReferenceID from stream. %s", err.Error())
	}

	err = serviceItemTransaction.Limitation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.Limitation from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemTransaction to the given writable
func (serviceItemTransaction *ServiceItemTransaction) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemTransaction.TransactionID.WriteTo(contentWritable)
	serviceItemTransaction.ExtTransactionID.WriteTo(contentWritable)
	serviceItemTransaction.Time.WriteTo(contentWritable)
	serviceItemTransaction.TransactionType.WriteTo(contentWritable)
	serviceItemTransaction.TransactionDescription.WriteTo(contentWritable)
	serviceItemTransaction.TransactionAmount.WriteTo(contentWritable)
	serviceItemTransaction.ItemCode.WriteTo(contentWritable)
	serviceItemTransaction.ReferenceID.WriteTo(contentWritable)
	serviceItemTransaction.Limitation.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemTransaction.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemTransaction
func (serviceItemTransaction *ServiceItemTransaction) Copy() types.RVType {
	copied := NewServiceItemTransaction()

	copied.StructureVersion = serviceItemTransaction.StructureVersion

	copied.TransactionID = serviceItemTransaction.TransactionID
	copied.ExtTransactionID = serviceItemTransaction.ExtTransactionID
	copied.Time = serviceItemTransaction.Time.Copy()
	copied.TransactionType = serviceItemTransaction.TransactionType
	copied.TransactionDescription = serviceItemTransaction.TransactionDescription
	copied.TransactionAmount = serviceItemTransaction.TransactionAmount.Copy().(*ServiceItemAmount)
	copied.ItemCode = serviceItemTransaction.ItemCode
	copied.ReferenceID = serviceItemTransaction.ReferenceID
	copied.Limitation = serviceItemTransaction.Limitation.Copy().(*ServiceItemLimitation)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemTransaction *ServiceItemTransaction) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemTransaction); !ok {
		return false
	}

	other := o.(*ServiceItemTransaction)

	if serviceItemTransaction.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemTransaction.TransactionID.Equals(other.TransactionID) {
		return false
	}

	if !serviceItemTransaction.ExtTransactionID.Equals(other.ExtTransactionID) {
		return false
	}

	if !serviceItemTransaction.Time.Equals(other.Time) {
		return false
	}

	if !serviceItemTransaction.TransactionType.Equals(other.TransactionType) {
		return false
	}

	if !serviceItemTransaction.TransactionDescription.Equals(other.TransactionDescription) {
		return false
	}

	if !serviceItemTransaction.TransactionAmount.Equals(other.TransactionAmount) {
		return false
	}

	if !serviceItemTransaction.ItemCode.Equals(other.ItemCode) {
		return false
	}

	if !serviceItemTransaction.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !serviceItemTransaction.Limitation.Equals(other.Limitation) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemTransaction *ServiceItemTransaction) String() string {
	return serviceItemTransaction.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemTransaction *ServiceItemTransaction) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemTransaction{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemTransaction.StructureVersion))
	b.WriteString(fmt.Sprintf("%sTransactionID: %q,\n", indentationValues, serviceItemTransaction.TransactionID))
	b.WriteString(fmt.Sprintf("%sExtTransactionID: %q,\n", indentationValues, serviceItemTransaction.ExtTransactionID))

	if serviceItemTransaction.Time != nil {
		b.WriteString(fmt.Sprintf("%sTime: %s\n", indentationValues, serviceItemTransaction.Time.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sTime: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sTransactionType: %d,\n", indentationValues, serviceItemTransaction.TransactionType))
	b.WriteString(fmt.Sprintf("%sTransactionDescription: %q,\n", indentationValues, serviceItemTransaction.TransactionDescription))

	if serviceItemTransaction.TransactionAmount != nil {
		b.WriteString(fmt.Sprintf("%sTransactionAmount: %s\n", indentationValues, serviceItemTransaction.TransactionAmount.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sTransactionAmount: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sItemCode: %q,\n", indentationValues, serviceItemTransaction.ItemCode))
	b.WriteString(fmt.Sprintf("%sReferenceID: %q,\n", indentationValues, serviceItemTransaction.ReferenceID))

	if serviceItemTransaction.Limitation != nil {
		b.WriteString(fmt.Sprintf("%sLimitation: %s\n", indentationValues, serviceItemTransaction.Limitation.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sLimitation: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemTransaction returns a new ServiceItemTransaction
func NewServiceItemTransaction() *ServiceItemTransaction {
	return &ServiceItemTransaction{}
}
