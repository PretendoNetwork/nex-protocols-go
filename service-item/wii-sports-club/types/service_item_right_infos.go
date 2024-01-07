// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemRightInfos holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemRightInfos struct {
	types.Structure
	RightInfos []*ServiceItemRightInfo
}

// ExtractFrom extracts the ServiceItemRightInfos from the given readable
func (serviceItemRightInfos *ServiceItemRightInfos) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemRightInfos.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemRightInfos header. %s", err.Error())
	}

	rightInfos, err := nex.StreamReadListStructure(stream, NewServiceItemRightInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.RightInfos from stream. %s", err.Error())
	}

	serviceItemRightInfos.RightInfos = rightInfos

	return nil
}

// WriteTo writes the ServiceItemRightInfos to the given writable
func (serviceItemRightInfos *ServiceItemRightInfos) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemRightInfos.RightInfos.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemRightInfos.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemRightInfos
func (serviceItemRightInfos *ServiceItemRightInfos) Copy() types.RVType {
	copied := NewServiceItemRightInfos()

	copied.StructureVersion = serviceItemRightInfos.StructureVersion

	copied.RightInfos = make([]*ServiceItemRightInfo, len(serviceItemRightInfos.RightInfos))

	for i := 0; i < len(serviceItemRightInfos.RightInfos); i++ {
		copied.RightInfos[i] = serviceItemRightInfos.RightInfos[i].Copy().(*ServiceItemRightInfo)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRightInfos *ServiceItemRightInfos) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRightInfos); !ok {
		return false
	}

	other := o.(*ServiceItemRightInfos)

	if serviceItemRightInfos.StructureVersion != other.StructureVersion {
		return false
	}

	if len(serviceItemRightInfos.RightInfos) != len(other.RightInfos) {
		return false
	}

	for i := 0; i < len(serviceItemRightInfos.RightInfos); i++ {
		if !serviceItemRightInfos.RightInfos[i].Equals(other.RightInfos[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemRightInfos *ServiceItemRightInfos) String() string {
	return serviceItemRightInfos.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemRightInfos *ServiceItemRightInfos) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightInfos{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemRightInfos.StructureVersion))

	if len(serviceItemRightInfos.RightInfos) == 0 {
		b.WriteString(fmt.Sprintf("%sRightInfos: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sRightInfos: [\n", indentationValues))

		for i := 0; i < len(serviceItemRightInfos.RightInfos); i++ {
			str := serviceItemRightInfos.RightInfos[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemRightInfos.RightInfos)-1 {
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

// NewServiceItemRightInfos returns a new ServiceItemRightInfos
func NewServiceItemRightInfos() *ServiceItemRightInfos {
	return &ServiceItemRightInfos{}
}
