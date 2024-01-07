// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemPrepurchaseRightInfo holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemPrepurchaseRightInfo struct {
	types.Structure
	LimitationType *types.PrimitiveU32
	AcquiredCount  *types.PrimitiveU32
	UsedCount      *types.PrimitiveU32
	ExpiryDate     *types.PrimitiveU32
	ExpiredCount   *types.PrimitiveU32
	ExpiryCounts   *types.List[*types.PrimitiveU32]
}

// ExtractFrom extracts the ServiceItemPrepurchaseRightInfo from the given readable
func (serviceItemPrepurchaseRightInfo *ServiceItemPrepurchaseRightInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemPrepurchaseRightInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemPrepurchaseRightInfo header. %s", err.Error())
	}

	err = serviceItemPrepurchaseRightInfo.LimitationType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.LimitationType from stream. %s", err.Error())
	}

	err = serviceItemPrepurchaseRightInfo.AcquiredCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.AcquiredCount from stream. %s", err.Error())
	}

	err = serviceItemPrepurchaseRightInfo.UsedCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.UsedCount from stream. %s", err.Error())
	}

	err = serviceItemPrepurchaseRightInfo.ExpiryDate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.ExpiryDate from stream. %s", err.Error())
	}

	err = serviceItemPrepurchaseRightInfo.ExpiredCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.ExpiredCount from stream. %s", err.Error())
	}

	err = serviceItemPrepurchaseRightInfo.ExpiryCounts.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.ExpiryCounts from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemPrepurchaseRightInfo to the given writable
func (serviceItemPrepurchaseRightInfo *ServiceItemPrepurchaseRightInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemPrepurchaseRightInfo.LimitationType.WriteTo(contentWritable)
	serviceItemPrepurchaseRightInfo.AcquiredCount.WriteTo(contentWritable)
	serviceItemPrepurchaseRightInfo.UsedCount.WriteTo(contentWritable)
	serviceItemPrepurchaseRightInfo.ExpiryDate.WriteTo(contentWritable)
	serviceItemPrepurchaseRightInfo.ExpiredCount.WriteTo(contentWritable)
	serviceItemPrepurchaseRightInfo.ExpiryCounts.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemPrepurchaseRightInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemPrepurchaseRightInfo
func (serviceItemPrepurchaseRightInfo *ServiceItemPrepurchaseRightInfo) Copy() types.RVType {
	copied := NewServiceItemPrepurchaseRightInfo()

	copied.StructureVersion = serviceItemPrepurchaseRightInfo.StructureVersion

	copied.LimitationType = serviceItemPrepurchaseRightInfo.LimitationType
	copied.AcquiredCount = serviceItemPrepurchaseRightInfo.AcquiredCount
	copied.UsedCount = serviceItemPrepurchaseRightInfo.UsedCount
	copied.ExpiryDate = serviceItemPrepurchaseRightInfo.ExpiryDate
	copied.ExpiredCount = serviceItemPrepurchaseRightInfo.ExpiredCount
	copied.ExpiryCounts = make(*types.List[*types.PrimitiveU32], len(serviceItemPrepurchaseRightInfo.ExpiryCounts))

	copy(copied.ExpiryCounts, serviceItemPrepurchaseRightInfo.ExpiryCounts)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPrepurchaseRightInfo *ServiceItemPrepurchaseRightInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPrepurchaseRightInfo); !ok {
		return false
	}

	other := o.(*ServiceItemPrepurchaseRightInfo)

	if serviceItemPrepurchaseRightInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemPrepurchaseRightInfo.LimitationType.Equals(other.LimitationType) {
		return false
	}

	if !serviceItemPrepurchaseRightInfo.AcquiredCount.Equals(other.AcquiredCount) {
		return false
	}

	if !serviceItemPrepurchaseRightInfo.UsedCount.Equals(other.UsedCount) {
		return false
	}

	if !serviceItemPrepurchaseRightInfo.ExpiryDate.Equals(other.ExpiryDate) {
		return false
	}

	if !serviceItemPrepurchaseRightInfo.ExpiredCount.Equals(other.ExpiredCount) {
		return false
	}

	if len(serviceItemPrepurchaseRightInfo.ExpiryCounts) != len(other.ExpiryCounts) {
		return false
	}

	for i := 0; i < len(serviceItemPrepurchaseRightInfo.ExpiryCounts); i++ {
		if serviceItemPrepurchaseRightInfo.ExpiryCounts[i] != other.ExpiryCounts[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemPrepurchaseRightInfo *ServiceItemPrepurchaseRightInfo) String() string {
	return serviceItemPrepurchaseRightInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemPrepurchaseRightInfo *ServiceItemPrepurchaseRightInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPrepurchaseRightInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemPrepurchaseRightInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sLimitationType: %d,\n", indentationValues, serviceItemPrepurchaseRightInfo.LimitationType))
	b.WriteString(fmt.Sprintf("%sAcquiredCount: %d,\n", indentationValues, serviceItemPrepurchaseRightInfo.AcquiredCount))
	b.WriteString(fmt.Sprintf("%sUsedCount: %d,\n", indentationValues, serviceItemPrepurchaseRightInfo.UsedCount))
	b.WriteString(fmt.Sprintf("%sExpiryDate: %d,\n", indentationValues, serviceItemPrepurchaseRightInfo.ExpiryDate))
	b.WriteString(fmt.Sprintf("%sExpiredCount: %d,\n", indentationValues, serviceItemPrepurchaseRightInfo.ExpiredCount))
	b.WriteString(fmt.Sprintf("%sExpiryCounts: %v,\n", indentationValues, serviceItemPrepurchaseRightInfo.ExpiryCounts))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPrepurchaseRightInfo returns a new ServiceItemPrepurchaseRightInfo
func NewServiceItemPrepurchaseRightInfo() *ServiceItemPrepurchaseRightInfo {
	return &ServiceItemPrepurchaseRightInfo{}
}
