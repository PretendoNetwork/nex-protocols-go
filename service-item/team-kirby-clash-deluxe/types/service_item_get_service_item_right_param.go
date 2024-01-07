// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemGetServiceItemRightParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemGetServiceItemRightParam struct {
	types.Structure
	ReferenceID string
	DeviceID    string
	UniqueID    *types.PrimitiveU32
	ItemGroup   *types.PrimitiveU8
	Platform    *types.PrimitiveU8 // * Revision 1
}

// ExtractFrom extracts the ServiceItemGetServiceItemRightParam from the given readable
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemGetServiceItemRightParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemGetServiceItemRightParam header. %s", err.Error())
	}

	err = serviceItemGetServiceItemRightParam.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.ReferenceID from stream. %s", err.Error())
	}

	err = serviceItemGetServiceItemRightParam.DeviceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.DeviceID from stream. %s", err.Error())
	}

	err = serviceItemGetServiceItemRightParam.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.UniqueID from stream. %s", err.Error())
	}

	err = serviceItemGetServiceItemRightParam.ItemGroup.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.ItemGroup from stream. %s", err.Error())
	}

	if serviceItemGetServiceItemRightParam.StructureVersion >= 1 {
	err = 	serviceItemGetServiceItemRightParam.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemGetServiceItemRightParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// WriteTo writes the ServiceItemGetServiceItemRightParam to the given writable
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemGetServiceItemRightParam.ReferenceID.WriteTo(contentWritable)
	serviceItemGetServiceItemRightParam.DeviceID.WriteTo(contentWritable)
	serviceItemGetServiceItemRightParam.UniqueID.WriteTo(contentWritable)
	serviceItemGetServiceItemRightParam.ItemGroup.WriteTo(contentWritable)

	if serviceItemGetServiceItemRightParam.StructureVersion >= 1 {
		serviceItemGetServiceItemRightParam.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemGetServiceItemRightParam
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) Copy() types.RVType {
	copied := NewServiceItemGetServiceItemRightParam()

	copied.StructureVersion = serviceItemGetServiceItemRightParam.StructureVersion

	copied.ReferenceID = serviceItemGetServiceItemRightParam.ReferenceID
	copied.DeviceID = serviceItemGetServiceItemRightParam.DeviceID
	copied.UniqueID = serviceItemGetServiceItemRightParam.UniqueID
	copied.ItemGroup = serviceItemGetServiceItemRightParam.ItemGroup
	copied.Platform = serviceItemGetServiceItemRightParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetServiceItemRightParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetServiceItemRightParam)

	if serviceItemGetServiceItemRightParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemGetServiceItemRightParam.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !serviceItemGetServiceItemRightParam.DeviceID.Equals(other.DeviceID) {
		return false
	}

	if !serviceItemGetServiceItemRightParam.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !serviceItemGetServiceItemRightParam.ItemGroup.Equals(other.ItemGroup) {
		return false
	}

	if !serviceItemGetServiceItemRightParam.Platform.Equals(other.Platform) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) String() string {
	return serviceItemGetServiceItemRightParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemGetServiceItemRightParam *ServiceItemGetServiceItemRightParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetServiceItemRightParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemGetServiceItemRightParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sReferenceID: %q,\n", indentationValues, serviceItemGetServiceItemRightParam.ReferenceID))
	b.WriteString(fmt.Sprintf("%sDeviceID: %q,\n", indentationValues, serviceItemGetServiceItemRightParam.DeviceID))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemGetServiceItemRightParam.UniqueID))
	b.WriteString(fmt.Sprintf("%sItemGroup: %d,\n", indentationValues, serviceItemGetServiceItemRightParam.ItemGroup))

	if serviceItemGetServiceItemRightParam.StructureVersion >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemGetServiceItemRightParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetServiceItemRightParam returns a new ServiceItemGetServiceItemRightParam
func NewServiceItemGetServiceItemRightParam() *ServiceItemGetServiceItemRightParam {
	return &ServiceItemGetServiceItemRightParam{}
}
