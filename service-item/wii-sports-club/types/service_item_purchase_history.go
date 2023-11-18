// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemPurchaseHistory holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemPurchaseHistory struct {
	nex.Structure
	TotalSize    uint32
	Offset       uint32
	Transactions []*ServiceItemTransaction
}

// ExtractFromStream extracts a ServiceItemPurchaseHistory structure from a stream
func (serviceItemPurchaseHistory *ServiceItemPurchaseHistory) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemPurchaseHistory.TotalSize, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseHistory.TotalSize from stream. %s", err.Error())
	}

	serviceItemPurchaseHistory.Offset, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseHistory.Offset from stream. %s", err.Error())
	}

	transactions, err := nex.StreamReadListStructure(stream, NewServiceItemTransaction())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseHistory.Transactions from stream. %s", err.Error())
	}

	serviceItemPurchaseHistory.Transactions = transactions

	return nil
}

// Bytes encodes the ServiceItemPurchaseHistory and returns a byte array
func (serviceItemPurchaseHistory *ServiceItemPurchaseHistory) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemPurchaseHistory.TotalSize)
	stream.WriteUInt32LE(serviceItemPurchaseHistory.Offset)
	stream.WriteListStructure(serviceItemPurchaseHistory.Transactions)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemPurchaseHistory
func (serviceItemPurchaseHistory *ServiceItemPurchaseHistory) Copy() nex.StructureInterface {
	copied := NewServiceItemPurchaseHistory()

	copied.SetStructureVersion(serviceItemPurchaseHistory.StructureVersion())

	copied.TotalSize = serviceItemPurchaseHistory.TotalSize
	copied.Offset = serviceItemPurchaseHistory.Offset
	copied.Transactions = make([]*ServiceItemTransaction, len(serviceItemPurchaseHistory.Transactions))

	for i := 0; i < len(serviceItemPurchaseHistory.Transactions); i++ {
		copied.Transactions[i] = serviceItemPurchaseHistory.Transactions[i].Copy().(*ServiceItemTransaction)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPurchaseHistory *ServiceItemPurchaseHistory) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemPurchaseHistory)

	if serviceItemPurchaseHistory.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemPurchaseHistory.TotalSize != other.TotalSize {
		return false
	}

	if serviceItemPurchaseHistory.Offset != other.Offset {
		return false
	}

	if len(serviceItemPurchaseHistory.Transactions) != len(other.Transactions) {
		return false
	}

	for i := 0; i < len(serviceItemPurchaseHistory.Transactions); i++ {
		if !serviceItemPurchaseHistory.Transactions[i].Equals(other.Transactions[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemPurchaseHistory *ServiceItemPurchaseHistory) String() string {
	return serviceItemPurchaseHistory.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemPurchaseHistory *ServiceItemPurchaseHistory) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPurchaseHistory{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemPurchaseHistory.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sTotalSize: %d,\n", indentationValues, serviceItemPurchaseHistory.TotalSize))
	b.WriteString(fmt.Sprintf("%sOffset: %d,\n", indentationValues, serviceItemPurchaseHistory.Offset))

	if len(serviceItemPurchaseHistory.Transactions) == 0 {
		b.WriteString(fmt.Sprintf("%sTransactions: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sTransactions: [\n", indentationValues))

		for i := 0; i < len(serviceItemPurchaseHistory.Transactions); i++ {
			str := serviceItemPurchaseHistory.Transactions[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemPurchaseHistory.Transactions)-1 {
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

// NewServiceItemPurchaseHistory returns a new ServiceItemPurchaseHistory
func NewServiceItemPurchaseHistory() *ServiceItemPurchaseHistory {
	return &ServiceItemPurchaseHistory{}
}
