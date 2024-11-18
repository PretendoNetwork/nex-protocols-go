// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemAcquireServiceItemResponse is a type within the ServiceItem protocol
type ServiceItemAcquireServiceItemResponse struct {
	types.Structure
	LimitationType types.UInt32
	AcquiredCount  types.UInt32
	UsedCount      types.UInt32
	ExpiryDate     types.UInt32
	ExpiredCount   types.UInt32
	ExpiryCounts   types.List[types.UInt32]
}

// WriteTo writes the ServiceItemAcquireServiceItemResponse to the given writable
func (siasir ServiceItemAcquireServiceItemResponse) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siasir.LimitationType.WriteTo(contentWritable)
	siasir.AcquiredCount.WriteTo(contentWritable)
	siasir.UsedCount.WriteTo(contentWritable)
	siasir.ExpiryDate.WriteTo(contentWritable)
	siasir.ExpiredCount.WriteTo(contentWritable)
	siasir.ExpiryCounts.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siasir.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemAcquireServiceItemResponse from the given readable
func (siasir *ServiceItemAcquireServiceItemResponse) ExtractFrom(readable types.Readable) error {
	var err error

	err = siasir.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse header. %s", err.Error())
	}

	err = siasir.LimitationType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.LimitationType. %s", err.Error())
	}

	err = siasir.AcquiredCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.AcquiredCount. %s", err.Error())
	}

	err = siasir.UsedCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.UsedCount. %s", err.Error())
	}

	err = siasir.ExpiryDate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.ExpiryDate. %s", err.Error())
	}

	err = siasir.ExpiredCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.ExpiredCount. %s", err.Error())
	}

	err = siasir.ExpiryCounts.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.ExpiryCounts. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemAcquireServiceItemResponse
func (siasir ServiceItemAcquireServiceItemResponse) Copy() types.RVType {
	copied := NewServiceItemAcquireServiceItemResponse()

	copied.StructureVersion = siasir.StructureVersion
	copied.LimitationType = siasir.LimitationType.Copy().(types.UInt32)
	copied.AcquiredCount = siasir.AcquiredCount.Copy().(types.UInt32)
	copied.UsedCount = siasir.UsedCount.Copy().(types.UInt32)
	copied.ExpiryDate = siasir.ExpiryDate.Copy().(types.UInt32)
	copied.ExpiredCount = siasir.ExpiredCount.Copy().(types.UInt32)
	copied.ExpiryCounts = siasir.ExpiryCounts.Copy().(types.List[types.UInt32])

	return copied
}

// Equals checks if the given ServiceItemAcquireServiceItemResponse contains the same data as the current ServiceItemAcquireServiceItemResponse
func (siasir ServiceItemAcquireServiceItemResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemAcquireServiceItemResponse); !ok {
		return false
	}

	other := o.(*ServiceItemAcquireServiceItemResponse)

	if siasir.StructureVersion != other.StructureVersion {
		return false
	}

	if !siasir.LimitationType.Equals(other.LimitationType) {
		return false
	}

	if !siasir.AcquiredCount.Equals(other.AcquiredCount) {
		return false
	}

	if !siasir.UsedCount.Equals(other.UsedCount) {
		return false
	}

	if !siasir.ExpiryDate.Equals(other.ExpiryDate) {
		return false
	}

	if !siasir.ExpiredCount.Equals(other.ExpiredCount) {
		return false
	}

	return siasir.ExpiryCounts.Equals(other.ExpiryCounts)
}

// CopyRef copies the current value of the ServiceItemAcquireServiceItemResponse
// and returns a pointer to the new copy
func (siasir ServiceItemAcquireServiceItemResponse) CopyRef() types.RVTypePtr {
	copied := siasir.Copy().(ServiceItemAcquireServiceItemResponse)
	return &copied
}

// Deref takes a pointer to the ServiceItemAcquireServiceItemResponse
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (siasir *ServiceItemAcquireServiceItemResponse) Deref() types.RVType {
	return *siasir
}

// String returns the string representation of the ServiceItemAcquireServiceItemResponse
func (siasir ServiceItemAcquireServiceItemResponse) String() string {
	return siasir.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemAcquireServiceItemResponse using the provided indentation level
func (siasir ServiceItemAcquireServiceItemResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAcquireServiceItemResponse{\n")
	b.WriteString(fmt.Sprintf("%sLimitationType: %s,\n", indentationValues, siasir.LimitationType))
	b.WriteString(fmt.Sprintf("%sAcquiredCount: %s,\n", indentationValues, siasir.AcquiredCount))
	b.WriteString(fmt.Sprintf("%sUsedCount: %s,\n", indentationValues, siasir.UsedCount))
	b.WriteString(fmt.Sprintf("%sExpiryDate: %s,\n", indentationValues, siasir.ExpiryDate))
	b.WriteString(fmt.Sprintf("%sExpiredCount: %s,\n", indentationValues, siasir.ExpiredCount))
	b.WriteString(fmt.Sprintf("%sExpiryCounts: %s,\n", indentationValues, siasir.ExpiryCounts))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAcquireServiceItemResponse returns a new ServiceItemAcquireServiceItemResponse
func NewServiceItemAcquireServiceItemResponse() ServiceItemAcquireServiceItemResponse {
	return ServiceItemAcquireServiceItemResponse{
		LimitationType: types.NewUInt32(0),
		AcquiredCount:  types.NewUInt32(0),
		UsedCount:      types.NewUInt32(0),
		ExpiryDate:     types.NewUInt32(0),
		ExpiredCount:   types.NewUInt32(0),
		ExpiryCounts:   types.NewList[types.UInt32](),
	}

}
