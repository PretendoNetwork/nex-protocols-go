// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemGetBalanceParam is a type within the ServiceItem protocol
type ServiceItemGetBalanceParam struct {
	types.Structure
	Language types.String
	UniqueID types.UInt32
	Platform types.UInt8 // * Revision 1
}

// WriteTo writes the ServiceItemGetBalanceParam to the given writable
func (sigbp ServiceItemGetBalanceParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sigbp.Language.WriteTo(contentWritable)
	sigbp.UniqueID.WriteTo(contentWritable)

	if sigbp.StructureVersion >= 1 {
		sigbp.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	sigbp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemGetBalanceParam from the given readable
func (sigbp *ServiceItemGetBalanceParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = sigbp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam header. %s", err.Error())
	}

	err = sigbp.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam.Language. %s", err.Error())
	}

	err = sigbp.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam.UniqueID. %s", err.Error())
	}

	if sigbp.StructureVersion >= 1 {
		err = sigbp.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemGetBalanceParam.Platform. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemGetBalanceParam
func (sigbp ServiceItemGetBalanceParam) Copy() types.RVType {
	copied := NewServiceItemGetBalanceParam()

	copied.StructureVersion = sigbp.StructureVersion
	copied.Language = sigbp.Language.Copy().(types.String)
	copied.UniqueID = sigbp.UniqueID.Copy().(types.UInt32)
	copied.Platform = sigbp.Platform.Copy().(types.UInt8)

	return copied
}

// Equals checks if the given ServiceItemGetBalanceParam contains the same data as the current ServiceItemGetBalanceParam
func (sigbp ServiceItemGetBalanceParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemGetBalanceParam); !ok {
		return false
	}

	other := o.(*ServiceItemGetBalanceParam)

	if sigbp.StructureVersion != other.StructureVersion {
		return false
	}

	if !sigbp.Language.Equals(other.Language) {
		return false
	}

	if !sigbp.UniqueID.Equals(other.UniqueID) {
		return false
	}

	return sigbp.Platform.Equals(other.Platform)
}

// CopyRef copies the current value of the ServiceItemGetBalanceParam
// and returns a pointer to the new copy
func (sigbp ServiceItemGetBalanceParam) CopyRef() types.RVTypePtr {
	copied := sigbp.Copy().(ServiceItemGetBalanceParam)
	return &copied
}

// Deref takes a pointer to the ServiceItemGetBalanceParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sigbp *ServiceItemGetBalanceParam) Deref() types.RVType {
	return *sigbp
}

// String returns the string representation of the ServiceItemGetBalanceParam
func (sigbp ServiceItemGetBalanceParam) String() string {
	return sigbp.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemGetBalanceParam using the provided indentation level
func (sigbp ServiceItemGetBalanceParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemGetBalanceParam{\n")
	b.WriteString(fmt.Sprintf("%sLanguage: %s,\n", indentationValues, sigbp.Language))
	b.WriteString(fmt.Sprintf("%sUniqueID: %s,\n", indentationValues, sigbp.UniqueID))
	b.WriteString(fmt.Sprintf("%sPlatform: %s,\n", indentationValues, sigbp.Platform))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemGetBalanceParam returns a new ServiceItemGetBalanceParam
func NewServiceItemGetBalanceParam() ServiceItemGetBalanceParam {
	return ServiceItemGetBalanceParam{
		Language: types.NewString(""),
		UniqueID: types.NewUInt32(0),
		Platform: types.NewUInt8(0),
	}

}
