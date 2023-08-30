// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemAcquireServiceItemResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemAcquireServiceItemResponse struct {
	nex.Structure
	LimitationType uint32
	AcquiredCount  uint32
	UsedCount      uint32
	ExpiryDate     uint32
	ExpiredCount   uint32
	ExpiryCounts   []uint32
}

// ExtractFromStream extracts a ServiceItemAcquireServiceItemResponse structure from a stream
func (serviceItemAcquireServiceItemResponse *ServiceItemAcquireServiceItemResponse) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemAcquireServiceItemResponse.LimitationType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.LimitationType from stream. %s", err.Error())
	}

	serviceItemAcquireServiceItemResponse.AcquiredCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.AcquiredCount from stream. %s", err.Error())
	}

	serviceItemAcquireServiceItemResponse.UsedCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.UsedCount from stream. %s", err.Error())
	}

	serviceItemAcquireServiceItemResponse.ExpiryDate, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.ExpiryDate from stream. %s", err.Error())
	}

	serviceItemAcquireServiceItemResponse.ExpiredCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.ExpiredCount from stream. %s", err.Error())
	}

	serviceItemAcquireServiceItemResponse.ExpiryCounts, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAcquireServiceItemResponse.ExpiryCounts from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemAcquireServiceItemResponse and returns a byte array
func (serviceItemAcquireServiceItemResponse *ServiceItemAcquireServiceItemResponse) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemAcquireServiceItemResponse.LimitationType)
	stream.WriteUInt32LE(serviceItemAcquireServiceItemResponse.AcquiredCount)
	stream.WriteUInt32LE(serviceItemAcquireServiceItemResponse.UsedCount)
	stream.WriteUInt32LE(serviceItemAcquireServiceItemResponse.ExpiryDate)
	stream.WriteUInt32LE(serviceItemAcquireServiceItemResponse.ExpiredCount)
	stream.WriteListUInt32LE(serviceItemAcquireServiceItemResponse.ExpiryCounts)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemAcquireServiceItemResponse
func (serviceItemAcquireServiceItemResponse *ServiceItemAcquireServiceItemResponse) Copy() nex.StructureInterface {
	copied := NewServiceItemAcquireServiceItemResponse()

	copied.SetStructureVersion(serviceItemAcquireServiceItemResponse.StructureVersion())

	copied.LimitationType = serviceItemAcquireServiceItemResponse.LimitationType
	copied.AcquiredCount = serviceItemAcquireServiceItemResponse.AcquiredCount
	copied.UsedCount = serviceItemAcquireServiceItemResponse.UsedCount
	copied.ExpiryDate = serviceItemAcquireServiceItemResponse.ExpiryDate
	copied.ExpiredCount = serviceItemAcquireServiceItemResponse.ExpiredCount
	copied.ExpiryCounts = make([]uint32, len(serviceItemAcquireServiceItemResponse.ExpiryCounts))

	copy(copied.ExpiryCounts, serviceItemAcquireServiceItemResponse.ExpiryCounts)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemAcquireServiceItemResponse *ServiceItemAcquireServiceItemResponse) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemAcquireServiceItemResponse)

	if serviceItemAcquireServiceItemResponse.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemAcquireServiceItemResponse.LimitationType != other.LimitationType {
		return false
	}

	if serviceItemAcquireServiceItemResponse.AcquiredCount != other.AcquiredCount {
		return false
	}

	if serviceItemAcquireServiceItemResponse.UsedCount != other.UsedCount {
		return false
	}

	if serviceItemAcquireServiceItemResponse.ExpiryDate != other.ExpiryDate {
		return false
	}

	if serviceItemAcquireServiceItemResponse.ExpiredCount != other.ExpiredCount {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemAcquireServiceItemResponse.StructureVersion()))
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
