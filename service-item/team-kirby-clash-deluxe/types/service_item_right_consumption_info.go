// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemRightConsumptionInfo holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemRightConsumptionInfo struct {
	nex.Structure
	*ServiceItemRightInfo
	AccountRights []*ServiceItemAccountRightConsumption
}

// ExtractFromStream extracts a ServiceItemRightConsumptionInfo structure from a stream
func (serviceItemRightConsumptionInfo *ServiceItemRightConsumptionInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	accountRights, err := stream.ReadListStructure(NewServiceItemAccountRightConsumption())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightConsumptionInfo.AccountRights from stream. %s", err.Error())
	}

	serviceItemRightConsumptionInfo.AccountRights = accountRights.([]*ServiceItemAccountRightConsumption)

	return nil
}

// Bytes encodes the ServiceItemRightConsumptionInfo and returns a byte array
func (serviceItemRightConsumptionInfo *ServiceItemRightConsumptionInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListStructure(serviceItemRightConsumptionInfo.AccountRights)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemRightConsumptionInfo
func (serviceItemRightConsumptionInfo *ServiceItemRightConsumptionInfo) Copy() nex.StructureInterface {
	copied := NewServiceItemRightConsumptionInfo()

	copied.SetStructureVersion(serviceItemRightConsumptionInfo.StructureVersion())

	copied.ServiceItemRightInfo = serviceItemRightConsumptionInfo.ServiceItemRightInfo.Copy().(*ServiceItemRightInfo)
	copied.SetParentType(copied.ServiceItemRightInfo)

	copied.AccountRights = make([]*ServiceItemAccountRightConsumption, len(serviceItemRightConsumptionInfo.AccountRights))

	for i := 0; i < len(serviceItemRightConsumptionInfo.AccountRights); i++ {
		copied.AccountRights[i] = serviceItemRightConsumptionInfo.AccountRights[i].Copy().(*ServiceItemAccountRightConsumption)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRightConsumptionInfo *ServiceItemRightConsumptionInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemRightConsumptionInfo)

	if serviceItemRightConsumptionInfo.StructureVersion() != other.StructureVersion() {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemRightConsumptionInfo.StructureVersion()))

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
