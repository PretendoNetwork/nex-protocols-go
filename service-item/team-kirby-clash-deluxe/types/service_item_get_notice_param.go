// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemGetNoticeParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetNoticeParam struct {
	types.Structure
	ScheduleType *types.PrimitiveU32
}

// ExtractFrom extracts the ServiceItemGetNoticeParam from the given readable
func (serviceItemGetNoticeParam *ServiceItemGetNoticeParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemGetNoticeParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemGetNoticeParam header. %s", err.Error())
	}

	err = serviceItemGetNoticeParam.ScheduleType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetNoticeParam.ScheduleType from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemGetNoticeParam to the given writable
func (serviceItemGetNoticeParam *ServiceItemGetNoticeParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemGetNoticeParam.ScheduleType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemGetNoticeParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemGetNoticeParam
func (serviceItemGetNoticeParam *ServiceItemGetNoticeParam) Copy() types.RVType {
	copied := NewServiceItemGetNoticeParam()

	copied.StructureVersion = serviceItemGetNoticeParam.StructureVersion

	copied.ScheduleType = serviceItemGetNoticeParam.ScheduleType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetNoticeParam *ServiceItemGetNoticeParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetNoticeParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetNoticeParam)

	if serviceItemGetNoticeParam.StructureVersion != other.StructureVersion {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemGetNoticeParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sScheduleType: %d,\n", indentationValues, serviceItemGetNoticeParam.ScheduleType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetNoticeParam returns a new ServiceItemGetNoticeParam
func NewServiceItemGetNoticeParam() *ServiceItemGetNoticeParam {
	return &ServiceItemGetNoticeParam{}
}
