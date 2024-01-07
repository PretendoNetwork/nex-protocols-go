// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemRightInfo holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemRightInfo struct {
	types.Structure
	ReferenceID   string
	AccountRights []*ServiceItemAccountRight
}

// ExtractFrom extracts the ServiceItemRightInfo from the given readable
func (serviceItemRightInfo *ServiceItemRightInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemRightInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemRightInfo header. %s", err.Error())
	}

	err = serviceItemRightInfo.ReferenceID.ExtractFrom(readable)
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

// WriteTo writes the ServiceItemRightInfo to the given writable
func (serviceItemRightInfo *ServiceItemRightInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemRightInfo.ReferenceID.WriteTo(contentWritable)
	serviceItemRightInfo.AccountRights.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemRightInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemRightInfo
func (serviceItemRightInfo *ServiceItemRightInfo) Copy() types.RVType {
	copied := NewServiceItemRightInfo()

	copied.StructureVersion = serviceItemRightInfo.StructureVersion

	copied.ReferenceID = serviceItemRightInfo.ReferenceID
	copied.AccountRights = make([]*ServiceItemAccountRight, len(serviceItemRightInfo.AccountRights))

	for i := 0; i < len(serviceItemRightInfo.AccountRights); i++ {
		copied.AccountRights[i] = serviceItemRightInfo.AccountRights[i].Copy().(*ServiceItemAccountRight)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRightInfo *ServiceItemRightInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRightInfo); !ok {
		return false
	}

	other := o.(*ServiceItemRightInfo)

	if serviceItemRightInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemRightInfo.ReferenceID.Equals(other.ReferenceID) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemRightInfo.StructureVersion))
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
