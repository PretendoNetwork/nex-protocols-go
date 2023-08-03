// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemPrepurchaseRightInfo holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemPrepurchaseRightInfo struct {
	nex.Structure
	LimitationType uint32
	AcquiredCount  uint32
	UsedCount      uint32
	ExpiryDate     uint32
	ExpiredCount   uint32
	ExpiryCounts   []uint32
}

// ExtractFromStream extracts a ServiceItemPrepurchaseRightInfo structure from a stream
func (serviceItemPrepurchaseRightInfo *ServiceItemPrepurchaseRightInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemPrepurchaseRightInfo.LimitationType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.LimitationType from stream. %s", err.Error())
	}

	serviceItemPrepurchaseRightInfo.AcquiredCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.AcquiredCount from stream. %s", err.Error())
	}

	serviceItemPrepurchaseRightInfo.UsedCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.UsedCount from stream. %s", err.Error())
	}

	serviceItemPrepurchaseRightInfo.ExpiryDate, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.ExpiryDate from stream. %s", err.Error())
	}

	serviceItemPrepurchaseRightInfo.ExpiredCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.ExpiredCount from stream. %s", err.Error())
	}

	serviceItemPrepurchaseRightInfo.ExpiryCounts, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseRightInfo.ExpiryCounts from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemPrepurchaseRightInfo and returns a byte array
func (serviceItemPrepurchaseRightInfo *ServiceItemPrepurchaseRightInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemPrepurchaseRightInfo.LimitationType)
	stream.WriteUInt32LE(serviceItemPrepurchaseRightInfo.AcquiredCount)
	stream.WriteUInt32LE(serviceItemPrepurchaseRightInfo.UsedCount)
	stream.WriteUInt32LE(serviceItemPrepurchaseRightInfo.ExpiryDate)
	stream.WriteUInt32LE(serviceItemPrepurchaseRightInfo.ExpiredCount)
	stream.WriteListUInt32LE(serviceItemPrepurchaseRightInfo.ExpiryCounts)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemPrepurchaseRightInfo
func (serviceItemPrepurchaseRightInfo *ServiceItemPrepurchaseRightInfo) Copy() nex.StructureInterface {
	copied := NewServiceItemPrepurchaseRightInfo()

	copied.LimitationType = serviceItemPrepurchaseRightInfo.LimitationType
	copied.AcquiredCount = serviceItemPrepurchaseRightInfo.AcquiredCount
	copied.UsedCount = serviceItemPrepurchaseRightInfo.UsedCount
	copied.ExpiryDate = serviceItemPrepurchaseRightInfo.ExpiryDate
	copied.ExpiredCount = serviceItemPrepurchaseRightInfo.ExpiredCount
	copied.ExpiryCounts = make([]uint32, len(serviceItemPrepurchaseRightInfo.ExpiryCounts))

	copy(copied.ExpiryCounts, serviceItemPrepurchaseRightInfo.ExpiryCounts)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPrepurchaseRightInfo *ServiceItemPrepurchaseRightInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemPrepurchaseRightInfo)

	if serviceItemPrepurchaseRightInfo.LimitationType != other.LimitationType {
		return false
	}

	if serviceItemPrepurchaseRightInfo.AcquiredCount != other.AcquiredCount {
		return false
	}

	if serviceItemPrepurchaseRightInfo.UsedCount != other.UsedCount {
		return false
	}

	if serviceItemPrepurchaseRightInfo.ExpiryDate != other.ExpiryDate {
		return false
	}

	if serviceItemPrepurchaseRightInfo.ExpiredCount != other.ExpiredCount {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemPrepurchaseRightInfo.StructureVersion()))
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
