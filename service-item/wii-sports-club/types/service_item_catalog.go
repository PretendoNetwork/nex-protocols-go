// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemCatalog holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemCatalog struct {
	types.Structure
	TotalSize *types.PrimitiveU32
	Offset    *types.PrimitiveU32
	ListItems []*ServiceItemListItem
}

// ExtractFrom extracts the ServiceItemCatalog from the given readable
func (serviceItemCatalog *ServiceItemCatalog) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemCatalog.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemCatalog header. %s", err.Error())
	}

	err = serviceItemCatalog.TotalSize.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemCatalog.TotalSize from stream. %s", err.Error())
	}

	err = serviceItemCatalog.Offset.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemCatalog.Offset from stream. %s", err.Error())
	}

	listItems, err := nex.StreamReadListStructure(stream, NewServiceItemListItem())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemCatalog.ListItems from stream. %s", err.Error())
	}

	serviceItemCatalog.ListItems = listItems

	return nil
}

// WriteTo writes the ServiceItemCatalog to the given writable
func (serviceItemCatalog *ServiceItemCatalog) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemCatalog.TotalSize.WriteTo(contentWritable)
	serviceItemCatalog.Offset.WriteTo(contentWritable)
	serviceItemCatalog.ListItems.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemCatalog.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemCatalog
func (serviceItemCatalog *ServiceItemCatalog) Copy() types.RVType {
	copied := NewServiceItemCatalog()

	copied.StructureVersion = serviceItemCatalog.StructureVersion

	copied.TotalSize = serviceItemCatalog.TotalSize
	copied.Offset = serviceItemCatalog.Offset
	copied.ListItems = make([]*ServiceItemListItem, len(serviceItemCatalog.ListItems))

	for i := 0; i < len(serviceItemCatalog.ListItems); i++ {
		copied.ListItems[i] = serviceItemCatalog.ListItems[i].Copy().(*ServiceItemListItem)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemCatalog *ServiceItemCatalog) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemCatalog); !ok {
		return false
	}

	other := o.(*ServiceItemCatalog)

	if serviceItemCatalog.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemCatalog.TotalSize.Equals(other.TotalSize) {
		return false
	}

	if !serviceItemCatalog.Offset.Equals(other.Offset) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemCatalog.StructureVersion))
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

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemCatalog returns a new ServiceItemCatalog
func NewServiceItemCatalog() *ServiceItemCatalog {
	return &ServiceItemCatalog{}
}
