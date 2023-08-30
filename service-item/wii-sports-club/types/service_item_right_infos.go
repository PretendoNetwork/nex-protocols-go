// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemRightInfos holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemRightInfos struct {
	nex.Structure
	RightInfos []*ServiceItemRightInfo
}

// ExtractFromStream extracts a ServiceItemRightInfos structure from a stream
func (serviceItemRightInfos *ServiceItemRightInfos) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	rightInfos, err := stream.ReadListStructure(NewServiceItemRightInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.RightInfos from stream. %s", err.Error())
	}

	serviceItemRightInfos.RightInfos = rightInfos.([]*ServiceItemRightInfo)

	return nil
}

// Bytes encodes the ServiceItemRightInfos and returns a byte array
func (serviceItemRightInfos *ServiceItemRightInfos) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListStructure(serviceItemRightInfos.RightInfos)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemRightInfos
func (serviceItemRightInfos *ServiceItemRightInfos) Copy() nex.StructureInterface {
	copied := NewServiceItemRightInfos()

	copied.SetStructureVersion(serviceItemRightInfos.StructureVersion())

	copied.RightInfos = make([]*ServiceItemRightInfo, len(serviceItemRightInfos.RightInfos))

	for i := 0; i < len(serviceItemRightInfos.RightInfos); i++ {
		copied.RightInfos[i] = serviceItemRightInfos.RightInfos[i].Copy().(*ServiceItemRightInfo)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRightInfos *ServiceItemRightInfos) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemRightInfos)

	if serviceItemRightInfos.StructureVersion() != other.StructureVersion() {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemRightInfos.StructureVersion()))

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
