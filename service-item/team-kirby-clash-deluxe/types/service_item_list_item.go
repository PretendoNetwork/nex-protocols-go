// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemListItem holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemListItem struct {
	types.Structure
	ItemCode            string
	RegularPrice        *ServiceItemAmount
	TaxExcluded         *types.PrimitiveBool
	InitialPurchaseOnly *types.PrimitiveBool
	Limitation          *ServiceItemLimitation
	Attributes          []*ServiceItemAttribute
}

// ExtractFrom extracts the ServiceItemListItem from the given readable
func (serviceItemListItem *ServiceItemListItem) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemListItem.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemListItem header. %s", err.Error())
	}

	err = serviceItemListItem.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.ItemCode from stream. %s", err.Error())
	}

	err = serviceItemListItem.RegularPrice.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.RegularPrice from stream. %s", err.Error())
	}

	err = serviceItemListItem.TaxExcluded.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.TaxExcluded from stream. %s", err.Error())
	}

	err = serviceItemListItem.InitialPurchaseOnly.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemListItem.InitialPurchaseOnly from stream. %s", err.Error())
	}

	err = serviceItemListItem.Limitation.ExtractFrom(readable)
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

// WriteTo writes the ServiceItemListItem to the given writable
func (serviceItemListItem *ServiceItemListItem) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemListItem.ItemCode.WriteTo(contentWritable)
	serviceItemListItem.RegularPrice.WriteTo(contentWritable)
	serviceItemListItem.TaxExcluded.WriteTo(contentWritable)
	serviceItemListItem.InitialPurchaseOnly.WriteTo(contentWritable)
	serviceItemListItem.Limitation.WriteTo(contentWritable)
	serviceItemListItem.Attributes.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemListItem.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemListItem
func (serviceItemListItem *ServiceItemListItem) Copy() types.RVType {
	copied := NewServiceItemListItem()

	copied.StructureVersion = serviceItemListItem.StructureVersion

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
func (serviceItemListItem *ServiceItemListItem) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemListItem); !ok {
		return false
	}

	other := o.(*ServiceItemListItem)

	if serviceItemListItem.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemListItem.ItemCode.Equals(other.ItemCode) {
		return false
	}

	if !serviceItemListItem.RegularPrice.Equals(other.RegularPrice) {
		return false
	}

	if !serviceItemListItem.TaxExcluded.Equals(other.TaxExcluded) {
		return false
	}

	if !serviceItemListItem.InitialPurchaseOnly.Equals(other.InitialPurchaseOnly) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemListItem.StructureVersion))
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
