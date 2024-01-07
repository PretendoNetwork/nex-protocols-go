// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemPurchaseHistory holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemPurchaseHistory struct {
	types.Structure
	TotalSize    *types.PrimitiveU32
	Offset       *types.PrimitiveU32
	Transactions []*ServiceItemTransaction
}

// ExtractFrom extracts the ServiceItemPurchaseHistory from the given readable
func (serviceItemPurchaseHistory *ServiceItemPurchaseHistory) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemPurchaseHistory.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemPurchaseHistory header. %s", err.Error())
	}

	err = serviceItemPurchaseHistory.TotalSize.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseHistory.TotalSize from stream. %s", err.Error())
	}

	err = serviceItemPurchaseHistory.Offset.ExtractFrom(readable)
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

// WriteTo writes the ServiceItemPurchaseHistory to the given writable
func (serviceItemPurchaseHistory *ServiceItemPurchaseHistory) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemPurchaseHistory.TotalSize.WriteTo(contentWritable)
	serviceItemPurchaseHistory.Offset.WriteTo(contentWritable)
	serviceItemPurchaseHistory.Transactions.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemPurchaseHistory.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemPurchaseHistory
func (serviceItemPurchaseHistory *ServiceItemPurchaseHistory) Copy() types.RVType {
	copied := NewServiceItemPurchaseHistory()

	copied.StructureVersion = serviceItemPurchaseHistory.StructureVersion

	copied.TotalSize = serviceItemPurchaseHistory.TotalSize
	copied.Offset = serviceItemPurchaseHistory.Offset
	copied.Transactions = make([]*ServiceItemTransaction, len(serviceItemPurchaseHistory.Transactions))

	for i := 0; i < len(serviceItemPurchaseHistory.Transactions); i++ {
		copied.Transactions[i] = serviceItemPurchaseHistory.Transactions[i].Copy().(*ServiceItemTransaction)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPurchaseHistory *ServiceItemPurchaseHistory) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPurchaseHistory); !ok {
		return false
	}

	other := o.(*ServiceItemPurchaseHistory)

	if serviceItemPurchaseHistory.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemPurchaseHistory.TotalSize.Equals(other.TotalSize) {
		return false
	}

	if !serviceItemPurchaseHistory.Offset.Equals(other.Offset) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemPurchaseHistory.StructureVersion))
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
