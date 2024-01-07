// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemRightInfo holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemRightInfo struct {
	types.Structure
	ReferenceID     string
	ReferenceIDType *types.PrimitiveU32
}

// ExtractFrom extracts the ServiceItemRightInfo from the given readable
func (serviceItemRightInfo *ServiceItemRightInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemRightInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemRightInfo header. %s", err.Error())
	}

	err = serviceItemRightInfo.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfo.ReferenceID from stream. %s", err.Error())
	}

	err = serviceItemRightInfo.ReferenceIDType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfo.ReferenceIDType from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemRightInfo to the given writable
func (serviceItemRightInfo *ServiceItemRightInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemRightInfo.ReferenceID.WriteTo(contentWritable)
	serviceItemRightInfo.ReferenceIDType.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemRightInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemRightInfo
func (serviceItemRightInfo *ServiceItemRightInfo) Copy() types.RVType {
	copied := NewServiceItemRightInfo()

	copied.StructureVersion = serviceItemRightInfo.StructureVersion

	copied.ReferenceID = serviceItemRightInfo.ReferenceID
	copied.ReferenceIDType = serviceItemRightInfo.ReferenceIDType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRightInfo *ServiceItemRightInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRightInfo); !ok {
		return false
	}

	other := o.(*ServiceItemRightInfo)

	if serviceItemRightInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemRightInfo.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !serviceItemRightInfo.ReferenceIDType.Equals(other.ReferenceIDType) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemRightInfo *ServiceItemRightInfo) String() string {
	return serviceItemRightInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemRightInfo *ServiceItemRightInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemRightInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sReferenceID: %q,\n", indentationValues, serviceItemRightInfo.ReferenceID))
	b.WriteString(fmt.Sprintf("%sReferenceIDType: %d,\n", indentationValues, serviceItemRightInfo.ReferenceIDType))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRightInfo returns a new ServiceItemRightInfo
func NewServiceItemRightInfo() *ServiceItemRightInfo {
	return &ServiceItemRightInfo{}
}
