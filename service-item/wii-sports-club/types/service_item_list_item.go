// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemListItem holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemListItem struct {
	nex.Structure
	ItemCode            string
	RegularPrice        *ServiceItemAmount
	TaxExcluded         bool
	InitialPurchaseOnly bool
	Limitation          *ServiceItemLimitation
	Attributes          []*ServiceItemAttribute
}

// ExtractFromStream extracts a ServiceItemListItem structure from a stream
func (serviceItemListItem *ServiceItemListItem) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemListItem.ItemCode, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.ItemCode from stream. %s", err.Error())
	}

	serviceItemListItem.RegularPrice, err = nex.StreamReadStructure(stream, NewServiceItemAmount())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.RegularPrice from stream. %s", err.Error())
	}

	serviceItemListItem.TaxExcluded, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.TaxExcluded from stream. %s", err.Error())
	}

	serviceItemListItem.InitialPurchaseOnly, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.InitialPurchaseOnly from stream. %s", err.Error())
	}

	serviceItemListItem.Limitation, err = nex.StreamReadStructure(stream, NewServiceItemLimitation())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.Limitation from stream. %s", err.Error())
	}

	attributes, err := nex.StreamReadListStructure(stream, NewServiceItemAttribute())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.Attributes from stream. %s", err.Error())
	}

	serviceItemListItem.Attributes = attributes

	return nil
}

// Bytes encodes the ServiceItemListItem and returns a byte array
func (serviceItemListItem *ServiceItemListItem) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemListItem.ItemCode)
	stream.WriteStructure(serviceItemListItem.RegularPrice)
	stream.WriteBool(serviceItemListItem.TaxExcluded)
	stream.WriteBool(serviceItemListItem.InitialPurchaseOnly)
	stream.WriteStructure(serviceItemListItem.Limitation)
	nex.StreamWriteListStructure(stream, serviceItemListItem.Attributes)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemListItem
func (serviceItemListItem *ServiceItemListItem) Copy() nex.StructureInterface {
	copied := NewServiceItemListItem()

	copied.SetStructureVersion(serviceItemListItem.StructureVersion())

	copied.ItemCode = serviceItemListItem.ItemCode
	copied.RegularPrice = serviceItemListItem.RegularPrice.Copy().(*ServiceItemAmount)
	copied.TaxExcluded = serviceItemListItem.TaxExcluded
	copied.InitialPurchaseOnly = serviceItemListItem.InitialPurchaseOnly
	copied.Limitation = serviceItemListItem.Limitation.Copy().(*ServiceItemLimitation)
	copied.Attributes = make([]*ServiceItemAttribute, len(serviceItemListItem.Attributes))

	for i := 0; i < len(serviceItemListItem.Attributes); i++ {
		copied.Attributes[i] = serviceItemListItem.Attributes[i].Copy().(*ServiceItemAttribute)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemListItem *ServiceItemListItem) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemListItem)

	if serviceItemListItem.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemListItem.ItemCode != other.ItemCode {
		return false
	}

	if !serviceItemListItem.RegularPrice.Equals(other.RegularPrice) {
		return false
	}

	if serviceItemListItem.TaxExcluded != other.TaxExcluded {
		return false
	}

	if serviceItemListItem.InitialPurchaseOnly != other.InitialPurchaseOnly {
		return false
	}

	if !serviceItemListItem.Limitation.Equals(other.Limitation) {
		return false
	}

	if len(serviceItemListItem.Attributes) != len(other.Attributes) {
		return false
	}

	for i := 0; i < len(serviceItemListItem.Attributes); i++ {
		if !serviceItemListItem.Attributes[i].Equals(other.Attributes[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemListItem *ServiceItemListItem) String() string {
	return serviceItemListItem.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemListItem *ServiceItemListItem) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemListItem{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemListItem.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sItemCode: %q,\n", indentationValues, serviceItemListItem.ItemCode))

	if serviceItemListItem.RegularPrice != nil {
		b.WriteString(fmt.Sprintf("%sRegularPrice: %s\n", indentationValues, serviceItemListItem.RegularPrice.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sRegularPrice: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sTaxExcluded: %t,\n", indentationValues, serviceItemListItem.TaxExcluded))
	b.WriteString(fmt.Sprintf("%sInitialPurchaseOnly: %t,\n", indentationValues, serviceItemListItem.InitialPurchaseOnly))

	if serviceItemListItem.Limitation != nil {
		b.WriteString(fmt.Sprintf("%sLimitation: %s\n", indentationValues, serviceItemListItem.Limitation.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sLimitation: nil\n", indentationValues))
	}

	if len(serviceItemListItem.Attributes) == 0 {
		b.WriteString(fmt.Sprintf("%sAttributes: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sAttributes: [\n", indentationValues))

		for i := 0; i < len(serviceItemListItem.Attributes); i++ {
			str := serviceItemListItem.Attributes[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemListItem.Attributes)-1 {
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

// NewServiceItemListItem returns a new ServiceItemListItem
func NewServiceItemListItem() *ServiceItemListItem {
	return &ServiceItemListItem{}
}
