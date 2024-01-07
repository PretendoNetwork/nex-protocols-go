// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemRightConsumptionInfo holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemRightConsumptionInfo struct {
	types.Structure
	*ServiceItemRightInfo
	AccountRights []*ServiceItemAccountRightConsumption
}

// ExtractFrom extracts the ServiceItemRightConsumptionInfo from the given readable
func (serviceItemRightConsumptionInfo *ServiceItemRightConsumptionInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemRightConsumptionInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemRightConsumptionInfo header. %s", err.Error())
	}

	accountRights, err := nex.StreamReadListStructure(stream, NewServiceItemAccountRightConsumption())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightConsumptionInfo.AccountRights from stream. %s", err.Error())
	}

	serviceItemRightConsumptionInfo.AccountRights = accountRights

	return nil
}

// WriteTo writes the ServiceItemRightConsumptionInfo to the given writable
func (serviceItemRightConsumptionInfo *ServiceItemRightConsumptionInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemRightConsumptionInfo.AccountRights.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemRightConsumptionInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemRightConsumptionInfo
func (serviceItemRightConsumptionInfo *ServiceItemRightConsumptionInfo) Copy() types.RVType {
	copied := NewServiceItemRightConsumptionInfo()

	copied.StructureVersion = serviceItemRightConsumptionInfo.StructureVersion

	copied.ServiceItemRightInfo = serviceItemRightConsumptionInfo.ServiceItemRightInfo.Copy().(*ServiceItemRightInfo)

	copied.AccountRights = make([]*ServiceItemAccountRightConsumption, len(serviceItemRightConsumptionInfo.AccountRights))

	for i := 0; i < len(serviceItemRightConsumptionInfo.AccountRights); i++ {
		copied.AccountRights[i] = serviceItemRightConsumptionInfo.AccountRights[i].Copy().(*ServiceItemAccountRightConsumption)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRightConsumptionInfo *ServiceItemRightConsumptionInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRightConsumptionInfo); !ok {
		return false
	}

	other := o.(*ServiceItemRightConsumptionInfo)

	if serviceItemRightConsumptionInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemRightConsumptionInfo.ParentType().Equals(other.ParentType()) {
		return false
	}

	if len(serviceItemRightConsumptionInfo.AccountRights) != len(other.AccountRights) {
		return false
	}

	for i := 0; i < len(serviceItemRightConsumptionInfo.AccountRights); i++ {
		if !serviceItemRightConsumptionInfo.AccountRights[i].Equals(other.AccountRights[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemRightConsumptionInfo *ServiceItemRightConsumptionInfo) String() string {
	return serviceItemRightConsumptionInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemRightConsumptionInfo *ServiceItemRightConsumptionInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightConsumptionInfo{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, serviceItemRightConsumptionInfo.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemRightConsumptionInfo.StructureVersion))

	if len(serviceItemRightConsumptionInfo.AccountRights) == 0 {
		b.WriteString(fmt.Sprintf("%sAccountRights: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sAccountRights: [\n", indentationValues))

		for i := 0; i < len(serviceItemRightConsumptionInfo.AccountRights); i++ {
			str := serviceItemRightConsumptionInfo.AccountRights[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemRightConsumptionInfo.AccountRights)-1 {
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

// NewServiceItemRightConsumptionInfo returns a new ServiceItemRightConsumptionInfo
func NewServiceItemRightConsumptionInfo() *ServiceItemRightConsumptionInfo {
	serviceItemRightConsumptionInfo := &ServiceItemRightConsumptionInfo{}

	serviceItemRightConsumptionInfo.ServiceItemRightInfo = NewServiceItemRightInfo()
	serviceItemRightConsumptionInfo.SetParentType(serviceItemRightConsumptionInfo.ServiceItemRightInfo)

	return serviceItemRightConsumptionInfo
}
