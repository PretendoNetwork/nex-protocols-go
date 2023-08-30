// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemCatalog holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemCatalog struct {
	nex.Structure
	TotalSize          uint32
	Offset             uint32
	ListItems          []*ServiceItemListItem
	IsBalanceAvailable bool
	Balance            *ServiceItemAmount
}

// ExtractFromStream extracts a ServiceItemCatalog structure from a stream
func (serviceItemCatalog *ServiceItemCatalog) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemCatalog.TotalSize, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemCatalog.TotalSize from stream. %s", err.Error())
	}

	serviceItemCatalog.Offset, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemCatalog.Offset from stream. %s", err.Error())
	}

	listItems, err := stream.ReadListStructure(NewServiceItemListItem())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemCatalog.ListItems from stream. %s", err.Error())
	}

	serviceItemCatalog.ListItems = listItems.([]*ServiceItemListItem)

	serviceItemCatalog.IsBalanceAvailable, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemCatalog.IsBalanceAvailable from stream. %s", err.Error())
	}

	balance, err := stream.ReadStructure(NewServiceItemAmount())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemCatalog.Balance from stream. %s", err.Error())
	}

	serviceItemCatalog.Balance = balance.(*ServiceItemAmount)

	return nil
}

// Bytes encodes the ServiceItemCatalog and returns a byte array
func (serviceItemCatalog *ServiceItemCatalog) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemCatalog.TotalSize)
	stream.WriteUInt32LE(serviceItemCatalog.Offset)
	stream.WriteListStructure(serviceItemCatalog.ListItems)
	stream.WriteBool(serviceItemCatalog.IsBalanceAvailable)
	stream.WriteStructure(serviceItemCatalog.Balance)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemCatalog
func (serviceItemCatalog *ServiceItemCatalog) Copy() nex.StructureInterface {
	copied := NewServiceItemCatalog()

	copied.SetStructureVersion(serviceItemCatalog.StructureVersion())

	copied.TotalSize = serviceItemCatalog.TotalSize
	copied.Offset = serviceItemCatalog.Offset
	copied.ListItems = make([]*ServiceItemListItem, len(serviceItemCatalog.ListItems))

	for i := 0; i < len(serviceItemCatalog.ListItems); i++ {
		copied.ListItems[i] = serviceItemCatalog.ListItems[i].Copy().(*ServiceItemListItem)
	}

	copied.IsBalanceAvailable = serviceItemCatalog.IsBalanceAvailable
	copied.Balance = serviceItemCatalog.Balance.Copy().(*ServiceItemAmount)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemCatalog *ServiceItemCatalog) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemCatalog)

	if serviceItemCatalog.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemCatalog.TotalSize != other.TotalSize {
		return false
	}

	if serviceItemCatalog.Offset != other.Offset {
		return false
	}

	if len(serviceItemCatalog.ListItems) != len(other.ListItems) {
		return false
	}

	for i := 0; i < len(serviceItemCatalog.ListItems); i++ {
		if !serviceItemCatalog.ListItems[i].Equals(other.ListItems[i]) {
			return false
		}
	}

	if serviceItemCatalog.IsBalanceAvailable != other.IsBalanceAvailable {
		return false
	}

	if !serviceItemCatalog.Balance.Equals(other.Balance) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemCatalog *ServiceItemCatalog) String() string {
	return serviceItemCatalog.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemCatalog *ServiceItemCatalog) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemCatalog{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemCatalog.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sTotalSize: %d,\n", indentationValues, serviceItemCatalog.TotalSize))
	b.WriteString(fmt.Sprintf("%sOffset: %d,\n", indentationValues, serviceItemCatalog.Offset))

	if len(serviceItemCatalog.ListItems) == 0 {
		b.WriteString(fmt.Sprintf("%sListItems: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sListItems: [\n", indentationValues))

		for i := 0; i < len(serviceItemCatalog.ListItems); i++ {
			str := serviceItemCatalog.ListItems[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemCatalog.ListItems)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sIsBalanceAvailable: %t,\n", indentationValues, serviceItemCatalog.IsBalanceAvailable))

	if serviceItemCatalog.Balance != nil {
		b.WriteString(fmt.Sprintf("%sBalance: %s\n", indentationValues, serviceItemCatalog.Balance.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sBalance: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemCatalog returns a new ServiceItemCatalog
func NewServiceItemCatalog() *ServiceItemCatalog {
	return &ServiceItemCatalog{}
}
