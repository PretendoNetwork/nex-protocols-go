// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemRightTimeInfo holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemRightTimeInfo struct {
	types.Structure
	*ServiceItemRightInfo
	AccountRights []*ServiceItemAccountRightTime
}

// ExtractFrom extracts the ServiceItemRightTimeInfo from the given readable
func (serviceItemRightTimeInfo *ServiceItemRightTimeInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemRightTimeInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemRightTimeInfo header. %s", err.Error())
	}

	accountRights, err := nex.StreamReadListStructure(stream, NewServiceItemAccountRightTime())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightTimeInfo.AccountRights from stream. %s", err.Error())
	}

	serviceItemRightTimeInfo.AccountRights = accountRights

	return nil
}

// WriteTo writes the ServiceItemRightTimeInfo to the given writable
func (serviceItemRightTimeInfo *ServiceItemRightTimeInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemRightTimeInfo.AccountRights.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemRightTimeInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemRightTimeInfo
func (serviceItemRightTimeInfo *ServiceItemRightTimeInfo) Copy() types.RVType {
	copied := NewServiceItemRightTimeInfo()

	copied.StructureVersion = serviceItemRightTimeInfo.StructureVersion

	copied.ServiceItemRightInfo = serviceItemRightTimeInfo.ServiceItemRightInfo.Copy().(*ServiceItemRightInfo)

	copied.AccountRights = make([]*ServiceItemAccountRightTime, len(serviceItemRightTimeInfo.AccountRights))

	for i := 0; i < len(serviceItemRightTimeInfo.AccountRights); i++ {
		copied.AccountRights[i] = serviceItemRightTimeInfo.AccountRights[i].Copy().(*ServiceItemAccountRightTime)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRightTimeInfo *ServiceItemRightTimeInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRightTimeInfo); !ok {
		return false
	}

	other := o.(*ServiceItemRightTimeInfo)

	if serviceItemRightTimeInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemRightTimeInfo.ParentType().Equals(other.ParentType()) {
		return false
	}

	if len(serviceItemRightTimeInfo.AccountRights) != len(other.AccountRights) {
		return false
	}

	for i := 0; i < len(serviceItemRightTimeInfo.AccountRights); i++ {
		if !serviceItemRightTimeInfo.AccountRights[i].Equals(other.AccountRights[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemRightTimeInfo *ServiceItemRightTimeInfo) String() string {
	return serviceItemRightTimeInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemRightTimeInfo *ServiceItemRightTimeInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightTimeInfo{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, serviceItemRightTimeInfo.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemRightTimeInfo.StructureVersion))

	if len(serviceItemRightTimeInfo.AccountRights) == 0 {
		b.WriteString(fmt.Sprintf("%sAccountRights: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sAccountRights: [\n", indentationValues))

		for i := 0; i < len(serviceItemRightTimeInfo.AccountRights); i++ {
			str := serviceItemRightTimeInfo.AccountRights[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemRightTimeInfo.AccountRights)-1 {
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

// NewServiceItemRightTimeInfo returns a new ServiceItemRightTimeInfo
func NewServiceItemRightTimeInfo() *ServiceItemRightTimeInfo {
	serviceItemRightTimeInfo := &ServiceItemRightTimeInfo{}

	serviceItemRightTimeInfo.ServiceItemRightInfo = NewServiceItemRightInfo()
	serviceItemRightTimeInfo.SetParentType(serviceItemRightTimeInfo.ServiceItemRightInfo)

	return serviceItemRightTimeInfo
}
