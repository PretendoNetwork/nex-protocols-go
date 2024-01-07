// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemAcquireServiceItemResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemAcquireServiceItemResponse struct {
	types.Structure
	LimitationType *types.PrimitiveU32
	AcquiredCount  *types.PrimitiveU32
	UsedCount      *types.PrimitiveU32
	ExpiryDate     *types.PrimitiveU32
	ExpiredCount   *types.PrimitiveU32
	ExpiryCounts   *types.List[*types.PrimitiveU32]
}

// ExtractFrom extracts the ServiceItemAcquireServiceItemResponse from the given readable
func (serviceItemAcquireServiceItemResponse *ServiceItemAcquireServiceItemResponse) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemAcquireServiceItemResponse.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemAcquireServiceItemResponse header. %s", err.Error())
	}

	err = serviceItemAcquireServiceItemResponse.LimitationType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.LimitationType from stream. %s", err.Error())
	}

	err = serviceItemAcquireServiceItemResponse.AcquiredCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.AcquiredCount from stream. %s", err.Error())
	}

	err = serviceItemAcquireServiceItemResponse.UsedCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.UsedCount from stream. %s", err.Error())
	}

	err = serviceItemAcquireServiceItemResponse.ExpiryDate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.ExpiryDate from stream. %s", err.Error())
	}

	err = serviceItemAcquireServiceItemResponse.ExpiredCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.ExpiredCount from stream. %s", err.Error())
	}

	err = serviceItemAcquireServiceItemResponse.ExpiryCounts.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.ExpiryCounts from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemAcquireServiceItemResponse to the given writable
func (serviceItemAcquireServiceItemResponse *ServiceItemAcquireServiceItemResponse) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemAcquireServiceItemResponse.LimitationType.WriteTo(contentWritable)
	serviceItemAcquireServiceItemResponse.AcquiredCount.WriteTo(contentWritable)
	serviceItemAcquireServiceItemResponse.UsedCount.WriteTo(contentWritable)
	serviceItemAcquireServiceItemResponse.ExpiryDate.WriteTo(contentWritable)
	serviceItemAcquireServiceItemResponse.ExpiredCount.WriteTo(contentWritable)
	serviceItemAcquireServiceItemResponse.ExpiryCounts.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemAcquireServiceItemResponse.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemAcquireServiceItemResponse
func (serviceItemAcquireServiceItemResponse *ServiceItemAcquireServiceItemResponse) Copy() types.RVType {
	copied := NewServiceItemAcquireServiceItemResponse()

	copied.StructureVersion = serviceItemAcquireServiceItemResponse.StructureVersion

	copied.LimitationType = serviceItemAcquireServiceItemResponse.LimitationType
	copied.AcquiredCount = serviceItemAcquireServiceItemResponse.AcquiredCount
	copied.UsedCount = serviceItemAcquireServiceItemResponse.UsedCount
	copied.ExpiryDate = serviceItemAcquireServiceItemResponse.ExpiryDate
	copied.ExpiredCount = serviceItemAcquireServiceItemResponse.ExpiredCount
	copied.ExpiryCounts = make(*types.List[*types.PrimitiveU32], len(serviceItemAcquireServiceItemResponse.ExpiryCounts))

	copy(copied.ExpiryCounts, serviceItemAcquireServiceItemResponse.ExpiryCounts)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAcquireServiceItemResponse *ServiceItemAcquireServiceItemResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemAcquireServiceItemResponse); !ok {
		return false
	}

	other := o.(*ServiceItemAcquireServiceItemResponse)

	if serviceItemAcquireServiceItemResponse.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemAcquireServiceItemResponse.LimitationType.Equals(other.LimitationType) {
		return false
	}

	if !serviceItemAcquireServiceItemResponse.AcquiredCount.Equals(other.AcquiredCount) {
		return false
	}

	if !serviceItemAcquireServiceItemResponse.UsedCount.Equals(other.UsedCount) {
		return false
	}

	if !serviceItemAcquireServiceItemResponse.ExpiryDate.Equals(other.ExpiryDate) {
		return false
	}

	if !serviceItemAcquireServiceItemResponse.ExpiredCount.Equals(other.ExpiredCount) {
		return false
	}

	if len(serviceItemAcquireServiceItemResponse.ExpiryCounts) != len(other.ExpiryCounts) {
		return false
	}

	for i := 0; i < len(serviceItemAcquireServiceItemResponse.ExpiryCounts); i++ {
		if serviceItemAcquireServiceItemResponse.ExpiryCounts[i] != other.ExpiryCounts[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemAcquireServiceItemResponse *ServiceItemAcquireServiceItemResponse) String() string {
	return serviceItemAcquireServiceItemResponse.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemAcquireServiceItemResponse *ServiceItemAcquireServiceItemResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAcquireServiceItemResponse{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemAcquireServiceItemResponse.StructureVersion))
	b.WriteString(fmt.Sprintf("%sLimitationType: %d,\n", indentationValues, serviceItemAcquireServiceItemResponse.LimitationType))
	b.WriteString(fmt.Sprintf("%sAcquiredCount: %d,\n", indentationValues, serviceItemAcquireServiceItemResponse.AcquiredCount))
	b.WriteString(fmt.Sprintf("%sUsedCount: %d,\n", indentationValues, serviceItemAcquireServiceItemResponse.UsedCount))
	b.WriteString(fmt.Sprintf("%sExpiryDate: %d,\n", indentationValues, serviceItemAcquireServiceItemResponse.ExpiryDate))
	b.WriteString(fmt.Sprintf("%sExpiredCount: %d,\n", indentationValues, serviceItemAcquireServiceItemResponse.ExpiredCount))
	b.WriteString(fmt.Sprintf("%sExpiryCounts: %v,\n", indentationValues, serviceItemAcquireServiceItemResponse.ExpiryCounts))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAcquireServiceItemResponse returns a new ServiceItemAcquireServiceItemResponse
func NewServiceItemAcquireServiceItemResponse() *ServiceItemAcquireServiceItemResponse {
	return &ServiceItemAcquireServiceItemResponse{}
}
