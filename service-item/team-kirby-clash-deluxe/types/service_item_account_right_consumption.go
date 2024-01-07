// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemAccountRightConsumption holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemAccountRightConsumption struct {
	types.Structure
	*ServiceItemAccountRight
	UsedCount    *types.PrimitiveU32
	ExpiredCount *types.PrimitiveU32
	ExpiryCounts *types.List[*types.PrimitiveU32]
}

// ExtractFrom extracts the ServiceItemAccountRightConsumption from the given readable
func (serviceItemAccountRightConsumption *ServiceItemAccountRightConsumption) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemAccountRightConsumption.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemAccountRightConsumption header. %s", err.Error())
	}

	err = serviceItemAccountRightConsumption.UsedCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRightConsumption.UsedCount from stream. %s", err.Error())
	}

	err = serviceItemAccountRightConsumption.ExpiredCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRightConsumption.ExpiredCount from stream. %s", err.Error())
	}

	err = serviceItemAccountRightConsumption.ExpiryCounts.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRightConsumption.ExpiryCounts from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemAccountRightConsumption to the given writable
func (serviceItemAccountRightConsumption *ServiceItemAccountRightConsumption) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemAccountRightConsumption.UsedCount.WriteTo(contentWritable)
	serviceItemAccountRightConsumption.ExpiredCount.WriteTo(contentWritable)
	serviceItemAccountRightConsumption.ExpiryCounts.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemAccountRightConsumption.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemAccountRightConsumption
func (serviceItemAccountRightConsumption *ServiceItemAccountRightConsumption) Copy() types.RVType {
	copied := NewServiceItemAccountRightConsumption()

	copied.StructureVersion = serviceItemAccountRightConsumption.StructureVersion

	copied.ServiceItemAccountRight = serviceItemAccountRightConsumption.ServiceItemAccountRight.Copy().(*ServiceItemAccountRight)

	copied.UsedCount = serviceItemAccountRightConsumption.UsedCount
	copied.ExpiredCount = serviceItemAccountRightConsumption.ExpiredCount
	copied.ExpiryCounts = make(*types.List[*types.PrimitiveU32], len(serviceItemAccountRightConsumption.ExpiryCounts))

	copy(copied.ExpiryCounts, serviceItemAccountRightConsumption.ExpiryCounts)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAccountRightConsumption *ServiceItemAccountRightConsumption) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemAccountRightConsumption); !ok {
		return false
	}

	other := o.(*ServiceItemAccountRightConsumption)

	if serviceItemAccountRightConsumption.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemAccountRightConsumption.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !serviceItemAccountRightConsumption.UsedCount.Equals(other.UsedCount) {
		return false
	}

	if !serviceItemAccountRightConsumption.ExpiredCount.Equals(other.ExpiredCount) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemAccountRightConsumption.StructureVersion))
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
