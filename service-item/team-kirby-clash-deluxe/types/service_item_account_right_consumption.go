// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemAccountRightConsumption holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemAccountRightConsumption struct {
	nex.Structure
	*ServiceItemAccountRight
	UsedCount    uint32
	ExpiredCount uint32
	ExpiryCounts []uint32
}

// ExtractFromStream extracts a ServiceItemAccountRightConsumption structure from a stream
func (serviceItemAccountRightConsumption *ServiceItemAccountRightConsumption) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemAccountRightConsumption.UsedCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRightConsumption.UsedCount from stream. %s", err.Error())
	}

	serviceItemAccountRightConsumption.ExpiredCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRightConsumption.ExpiredCount from stream. %s", err.Error())
	}

	serviceItemAccountRightConsumption.ExpiryCounts, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRightConsumption.ExpiryCounts from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemAccountRightConsumption and returns a byte array
func (serviceItemAccountRightConsumption *ServiceItemAccountRightConsumption) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemAccountRightConsumption.UsedCount)
	stream.WriteUInt32LE(serviceItemAccountRightConsumption.ExpiredCount)
	stream.WriteListUInt32LE(serviceItemAccountRightConsumption.ExpiryCounts)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemAccountRightConsumption
func (serviceItemAccountRightConsumption *ServiceItemAccountRightConsumption) Copy() nex.StructureInterface {
	copied := NewServiceItemAccountRightConsumption()

	copied.SetStructureVersion(serviceItemAccountRightConsumption.StructureVersion())

	copied.ServiceItemAccountRight = serviceItemAccountRightConsumption.ServiceItemAccountRight.Copy().(*ServiceItemAccountRight)
	copied.SetParentType(copied.ServiceItemAccountRight)

	copied.UsedCount = serviceItemAccountRightConsumption.UsedCount
	copied.ExpiredCount = serviceItemAccountRightConsumption.ExpiredCount
	copied.ExpiryCounts = make([]uint32, len(serviceItemAccountRightConsumption.ExpiryCounts))

	copy(copied.ExpiryCounts, serviceItemAccountRightConsumption.ExpiryCounts)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAccountRightConsumption *ServiceItemAccountRightConsumption) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemAccountRightConsumption)

	if serviceItemAccountRightConsumption.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !serviceItemAccountRightConsumption.ParentType().Equals(other.ParentType()) {
		return false
	}

	if serviceItemAccountRightConsumption.UsedCount != other.UsedCount {
		return false
	}

	if serviceItemAccountRightConsumption.ExpiredCount != other.ExpiredCount {
		return false
	}

	if len(serviceItemAccountRightConsumption.ExpiryCounts) != len(other.ExpiryCounts) {
		return false
	}

	for i := 0; i < len(serviceItemAccountRightConsumption.ExpiryCounts); i++ {
		if serviceItemAccountRightConsumption.ExpiryCounts[i] != other.ExpiryCounts[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemAccountRightConsumption *ServiceItemAccountRightConsumption) String() string {
	return serviceItemAccountRightConsumption.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemAccountRightConsumption *ServiceItemAccountRightConsumption) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAccountRightConsumption{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, serviceItemAccountRightConsumption.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemAccountRightConsumption.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUsedCount: %d,\n", indentationValues, serviceItemAccountRightConsumption.UsedCount))
	b.WriteString(fmt.Sprintf("%sExpiredCount: %d,\n", indentationValues, serviceItemAccountRightConsumption.ExpiredCount))
	b.WriteString(fmt.Sprintf("%sExpiryCounts: %v,\n", indentationValues, serviceItemAccountRightConsumption.ExpiryCounts))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAccountRightConsumption returns a new ServiceItemAccountRightConsumption
func NewServiceItemAccountRightConsumption() *ServiceItemAccountRightConsumption {
	serviceItemAccountRightConsumption := &ServiceItemAccountRightConsumption{}

	serviceItemAccountRightConsumption.ServiceItemAccountRight = NewServiceItemAccountRight()
	serviceItemAccountRightConsumption.SetParentType(serviceItemAccountRightConsumption.ServiceItemAccountRight)

	return serviceItemAccountRightConsumption
}
