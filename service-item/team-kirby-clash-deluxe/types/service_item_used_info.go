// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemUsedInfo holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemUsedInfo struct {
	types.Structure
	AcquiredCount *types.PrimitiveU32
	UsedCount     *types.PrimitiveU32
}

// ExtractFrom extracts the ServiceItemUsedInfo from the given readable
func (serviceItemUsedInfo *ServiceItemUsedInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemUsedInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemUsedInfo header. %s", err.Error())
	}

	err = serviceItemUsedInfo.AcquiredCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUsedInfo.AcquiredCount from stream. %s", err.Error())
	}

	err = serviceItemUsedInfo.UsedCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUsedInfo.UsedCount from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemUsedInfo to the given writable
func (serviceItemUsedInfo *ServiceItemUsedInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemUsedInfo.AcquiredCount.WriteTo(contentWritable)
	serviceItemUsedInfo.UsedCount.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemUsedInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemUsedInfo
func (serviceItemUsedInfo *ServiceItemUsedInfo) Copy() types.RVType {
	copied := NewServiceItemUsedInfo()

	copied.StructureVersion = serviceItemUsedInfo.StructureVersion

	copied.AcquiredCount = serviceItemUsedInfo.AcquiredCount
	copied.UsedCount = serviceItemUsedInfo.UsedCount

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemUsedInfo *ServiceItemUsedInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemUsedInfo); !ok {
		return false
	}

	other := o.(*ServiceItemUsedInfo)

	if serviceItemUsedInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemUsedInfo.AcquiredCount.Equals(other.AcquiredCount) {
		return false
	}

	if !serviceItemUsedInfo.UsedCount.Equals(other.UsedCount) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemUsedInfo *ServiceItemUsedInfo) String() string {
	return serviceItemUsedInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemUsedInfo *ServiceItemUsedInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemUsedInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemUsedInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sAcquiredCount: %d,\n", indentationValues, serviceItemUsedInfo.AcquiredCount))
	b.WriteString(fmt.Sprintf("%sUsedCount: %d,\n", indentationValues, serviceItemUsedInfo.UsedCount))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemUsedInfo returns a new ServiceItemUsedInfo
func NewServiceItemUsedInfo() *ServiceItemUsedInfo {
	return &ServiceItemUsedInfo{}
}
