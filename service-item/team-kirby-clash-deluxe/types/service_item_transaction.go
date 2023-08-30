// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemTransaction holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemTransaction struct {
	nex.Structure
	TransactionID          string
	ExtTransactionID       string
	Time                   *nex.DateTime
	TransactionType        uint32
	TransactionDescription string
	TransactionAmount      *ServiceItemAmount
	ItemCode               string
	ReferenceID            string
	Limitation             *ServiceItemLimitation
}

// ExtractFromStream extracts a ServiceItemTransaction structure from a stream
func (serviceItemTransaction *ServiceItemTransaction) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemTransaction.TransactionID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.TransactionID from stream. %s", err.Error())
	}

	serviceItemTransaction.ExtTransactionID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.ExtTransactionID from stream. %s", err.Error())
	}

	serviceItemTransaction.Time, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.Time from stream. %s", err.Error())
	}

	serviceItemTransaction.TransactionType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.TransactionType from stream. %s", err.Error())
	}

	serviceItemTransaction.TransactionDescription, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.TransactionDescription from stream. %s", err.Error())
	}

	transactionAmount, err := stream.ReadStructure(NewServiceItemAmount())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.TransactionAmount from stream. %s", err.Error())
	}

	serviceItemTransaction.TransactionAmount = transactionAmount.(*ServiceItemAmount)

	serviceItemTransaction.ItemCode, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.ItemCode from stream. %s", err.Error())
	}

	serviceItemTransaction.ReferenceID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.ReferenceID from stream. %s", err.Error())
	}

	limitation, err := stream.ReadStructure(NewServiceItemLimitation())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTransaction.Limitation from stream. %s", err.Error())
	}

	serviceItemTransaction.Limitation = limitation.(*ServiceItemLimitation)

	return nil
}

// Bytes encodes the ServiceItemTransaction and returns a byte array
func (serviceItemTransaction *ServiceItemTransaction) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemTransaction.TransactionID)
	stream.WriteString(serviceItemTransaction.ExtTransactionID)
	stream.WriteDateTime(serviceItemTransaction.Time)
	stream.WriteUInt32LE(serviceItemTransaction.TransactionType)
	stream.WriteString(serviceItemTransaction.TransactionDescription)
	stream.WriteStructure(serviceItemTransaction.TransactionAmount)
	stream.WriteString(serviceItemTransaction.ItemCode)
	stream.WriteString(serviceItemTransaction.ReferenceID)
	stream.WriteStructure(serviceItemTransaction.Limitation)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemTransaction
func (serviceItemTransaction *ServiceItemTransaction) Copy() nex.StructureInterface {
	copied := NewServiceItemTransaction()

	copied.SetStructureVersion(serviceItemTransaction.StructureVersion())

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
func (serviceItemTransaction *ServiceItemTransaction) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemTransaction)

	if serviceItemTransaction.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemTransaction.TransactionID != other.TransactionID {
		return false
	}

	if serviceItemTransaction.ExtTransactionID != other.ExtTransactionID {
		return false
	}

	if !serviceItemTransaction.Time.Equals(other.Time) {
		return false
	}

	if serviceItemTransaction.TransactionType != other.TransactionType {
		return false
	}

	if serviceItemTransaction.TransactionDescription != other.TransactionDescription {
		return false
	}

	if !serviceItemTransaction.TransactionAmount.Equals(other.TransactionAmount) {
		return false
	}

	if serviceItemTransaction.ItemCode != other.ItemCode {
		return false
	}

	if serviceItemTransaction.ReferenceID != other.ReferenceID {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemTransaction.StructureVersion()))
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
