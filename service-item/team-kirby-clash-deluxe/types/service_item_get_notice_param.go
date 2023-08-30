// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemGetNoticeParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetNoticeParam struct {
	nex.Structure
	ScheduleType uint32
}

// ExtractFromStream extracts a ServiceItemGetNoticeParam structure from a stream
func (serviceItemGetNoticeParam *ServiceItemGetNoticeParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemGetNoticeParam.ScheduleType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetNoticeParam.ScheduleType from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemGetNoticeParam and returns a byte array
func (serviceItemGetNoticeParam *ServiceItemGetNoticeParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemGetNoticeParam.ScheduleType)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemGetNoticeParam
func (serviceItemGetNoticeParam *ServiceItemGetNoticeParam) Copy() nex.StructureInterface {
	copied := NewServiceItemGetNoticeParam()

	copied.SetStructureVersion(serviceItemGetNoticeParam.StructureVersion())

	copied.ScheduleType = serviceItemGetNoticeParam.ScheduleType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetNoticeParam *ServiceItemGetNoticeParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemGetNoticeParam)

	if serviceItemGetNoticeParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	return serviceItemGetNoticeParam.ScheduleType == other.ScheduleType
}

// String returns a string representation of the struct
func (serviceItemGetNoticeParam *ServiceItemGetNoticeParam) String() string {
	return serviceItemGetNoticeParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetNoticeParam *ServiceItemGetNoticeParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetNoticeParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemGetNoticeParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sScheduleType: %d,\n", indentationValues, serviceItemGetNoticeParam.ScheduleType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetNoticeParam returns a new ServiceItemGetNoticeParam
func NewServiceItemGetNoticeParam() *ServiceItemGetNoticeParam {
	return &ServiceItemGetNoticeParam{}
}
