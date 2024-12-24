// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemGetSupportIDParam is a type within the ServiceItem protocol
type ServiceItemGetSupportIDParam struct {
	types.Structure
	UniqueID types.UInt32
	Platform types.UInt8
}

// WriteTo writes the ServiceItemGetSupportIDParam to the given writable
func (sigsidp ServiceItemGetSupportIDParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sigsidp.UniqueID.WriteTo(contentWritable)
	sigsidp.Platform.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sigsidp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemGetSupportIDParam from the given readable
func (sigsidp *ServiceItemGetSupportIDParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = sigsidp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetSupportIDParam header. %s", err.Error())
	}

	err = sigsidp.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetSupportIDParam.UniqueID. %s", err.Error())
	}

	err = sigsidp.Platform.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetSupportIDParam.Platform. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemGetSupportIDParam
func (sigsidp ServiceItemGetSupportIDParam) Copy() types.RVType {
	copied := NewServiceItemGetSupportIDParam()

	copied.StructureVersion = sigsidp.StructureVersion
	copied.UniqueID = sigsidp.UniqueID.Copy().(types.UInt32)
	copied.Platform = sigsidp.Platform.Copy().(types.UInt8)

	return copied
}

// Equals checks if the given ServiceItemGetSupportIDParam contains the same data as the current ServiceItemGetSupportIDParam
func (sigsidp ServiceItemGetSupportIDParam) Equals(o types.RVType) bool {
	if _, ok := o.(ServiceItemGetSupportIDParam); !ok {
		return false
	}

	other := o.(ServiceItemGetSupportIDParam)

	if sigsidp.StructureVersion != other.StructureVersion {
		return false
	}

	if !sigsidp.UniqueID.Equals(other.UniqueID) {
		return false
	}

	return sigsidp.Platform.Equals(other.Platform)
}

// CopyRef copies the current value of the ServiceItemGetSupportIDParam
// and returns a pointer to the new copy
func (sigsidp ServiceItemGetSupportIDParam) CopyRef() types.RVTypePtr {
	copied := sigsidp.Copy().(ServiceItemGetSupportIDParam)
	return &copied
}

// Deref takes a pointer to the ServiceItemGetSupportIDParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sigsidp *ServiceItemGetSupportIDParam) Deref() types.RVType {
	return *sigsidp
}

// String returns the string representation of the ServiceItemGetSupportIDParam
func (sigsidp ServiceItemGetSupportIDParam) String() string {
	return sigsidp.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemGetSupportIDParam using the provided indentation level
func (sigsidp ServiceItemGetSupportIDParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetSupportIDParam{\n")
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, sigsidp.UniqueID))
	b.WriteString(fmt.Sprintf("%sPlatform: %s,\n", indentationValues, sigsidp.Platform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetSupportIDParam returns a new ServiceItemGetSupportIDParam
func NewServiceItemGetSupportIDParam() ServiceItemGetSupportIDParam {
	return ServiceItemGetSupportIDParam{
		UniqueID: types.NewUInt32(0),
		Platform: types.NewUInt8(0),
	}

}
