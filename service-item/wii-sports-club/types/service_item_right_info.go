// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemRightInfo holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemRightInfo struct {
	nex.Structure
	ReferenceID   string
	AccountRights []*ServiceItemAccountRight
}

// ExtractFromStream extracts a ServiceItemRightInfo structure from a stream
func (serviceItemRightInfo *ServiceItemRightInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemRightInfo.ReferenceID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfo.ReferenceID from stream. %s", err.Error())
	}

	accountRights, err := nex.StreamReadListStructure(stream, NewServiceItemAccountRight())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfo.AccountRights from stream. %s", err.Error())
	}

	serviceItemRightInfo.AccountRights = accountRights

	return nil
}

// Bytes encodes the ServiceItemRightInfo and returns a byte array
func (serviceItemRightInfo *ServiceItemRightInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemRightInfo.ReferenceID)
	stream.WriteListStructure(serviceItemRightInfo.AccountRights)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemRightInfo
func (serviceItemRightInfo *ServiceItemRightInfo) Copy() nex.StructureInterface {
	copied := NewServiceItemRightInfo()

	copied.SetStructureVersion(serviceItemRightInfo.StructureVersion())

	copied.ReferenceID = serviceItemRightInfo.ReferenceID
	copied.AccountRights = make([]*ServiceItemAccountRight, len(serviceItemRightInfo.AccountRights))

	for i := 0; i < len(serviceItemRightInfo.AccountRights); i++ {
		copied.AccountRights[i] = serviceItemRightInfo.AccountRights[i].Copy().(*ServiceItemAccountRight)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRightInfo *ServiceItemRightInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemRightInfo)

	if serviceItemRightInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemRightInfo.ReferenceID != other.ReferenceID {
		return false
	}

	if len(serviceItemRightInfo.AccountRights) != len(other.AccountRights) {
		return false
	}

	for i := 0; i < len(serviceItemRightInfo.AccountRights); i++ {
		if !serviceItemRightInfo.AccountRights[i].Equals(other.AccountRights[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemRightInfo *ServiceItemRightInfo) String() string {
	return serviceItemRightInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemRightInfo *ServiceItemRightInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemRightInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sReferenceID: %q,\n", indentationValues, serviceItemRightInfo.ReferenceID))

	if len(serviceItemRightInfo.AccountRights) == 0 {
		b.WriteString(fmt.Sprintf("%sAccountRights: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sAccountRights: [\n", indentationValues))

		for i := 0; i < len(serviceItemRightInfo.AccountRights); i++ {
			str := serviceItemRightInfo.AccountRights[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemRightInfo.AccountRights)-1 {
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

// NewServiceItemRightInfo returns a new ServiceItemRightInfo
func NewServiceItemRightInfo() *ServiceItemRightInfo {
	return &ServiceItemRightInfo{}
}
