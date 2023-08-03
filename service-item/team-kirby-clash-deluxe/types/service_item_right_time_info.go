// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemRightTimeInfo holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemRightTimeInfo struct {
	nex.Structure
	*ServiceItemRightInfo
	AccountRights []*ServiceItemAccountRightTime
}

// ExtractFromStream extracts a ServiceItemRightTimeInfo structure from a stream
func (serviceItemRightTimeInfo *ServiceItemRightTimeInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	accountRights, err := stream.ReadListStructure(NewServiceItemAccountRightTime())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightTimeInfo.AccountRights from stream. %s", err.Error())
	}

	serviceItemRightTimeInfo.AccountRights = accountRights.([]*ServiceItemAccountRightTime)

	return nil
}

// Bytes encodes the ServiceItemRightTimeInfo and returns a byte array
func (serviceItemRightTimeInfo *ServiceItemRightTimeInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListStructure(serviceItemRightTimeInfo.AccountRights)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemRightTimeInfo
func (serviceItemRightTimeInfo *ServiceItemRightTimeInfo) Copy() nex.StructureInterface {
	copied := NewServiceItemRightTimeInfo()

	copied.ServiceItemRightInfo = serviceItemRightTimeInfo.ServiceItemRightInfo.Copy().(*ServiceItemRightInfo)
	copied.SetParentType(copied.ServiceItemRightInfo)

	copied.AccountRights = make([]*ServiceItemAccountRightTime, len(serviceItemRightTimeInfo.AccountRights))

	for i := 0; i < len(serviceItemRightTimeInfo.AccountRights); i++ {
		copied.AccountRights[i] = serviceItemRightTimeInfo.AccountRights[i].Copy().(*ServiceItemAccountRightTime)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRightTimeInfo *ServiceItemRightTimeInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemRightTimeInfo)

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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemRightTimeInfo.StructureVersion()))

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
