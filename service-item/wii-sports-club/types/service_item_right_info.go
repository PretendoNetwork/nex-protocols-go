// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemRightInfo is a type within the ServiceItem protocol
type ServiceItemRightInfo struct {
	types.Structure
	ReferenceID   *types.String
	AccountRights *types.List[*ServiceItemAccountRight]
}

// WriteTo writes the ServiceItemRightInfo to the given writable
func (siri *ServiceItemRightInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siri.ReferenceID.WriteTo(writable)
	siri.AccountRights.WriteTo(writable)

	content := contentWritable.Bytes()

	siri.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemRightInfo from the given readable
func (siri *ServiceItemRightInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = siri.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfo header. %s", err.Error())
	}

	err = siri.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfo.ReferenceID. %s", err.Error())
	}

	err = siri.AccountRights.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfo.AccountRights. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemRightInfo
func (siri *ServiceItemRightInfo) Copy() types.RVType {
	copied := NewServiceItemRightInfo()

	copied.StructureVersion = siri.StructureVersion
	copied.ReferenceID = siri.ReferenceID.Copy().(*types.String)
	copied.AccountRights = siri.AccountRights.Copy().(*types.List[*ServiceItemAccountRight])

	return copied
}

// Equals checks if the given ServiceItemRightInfo contains the same data as the current ServiceItemRightInfo
func (siri *ServiceItemRightInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRightInfo); !ok {
		return false
	}

	other := o.(*ServiceItemRightInfo)

	if siri.StructureVersion != other.StructureVersion {
		return false
	}

	if !siri.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	return siri.AccountRights.Equals(other.AccountRights)
}

// String returns the string representation of the ServiceItemRightInfo
func (siri *ServiceItemRightInfo) String() string {
	return siri.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemRightInfo using the provided indentation level
func (siri *ServiceItemRightInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightInfo{\n")
	b.WriteString(fmt.Sprintf("%sReferenceID: %s,\n", indentationValues, siri.ReferenceID))
	b.WriteString(fmt.Sprintf("%sAccountRights: %s,\n", indentationValues, siri.AccountRights))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRightInfo returns a new ServiceItemRightInfo
func NewServiceItemRightInfo() *ServiceItemRightInfo {
	siri := &ServiceItemRightInfo{
		ReferenceID:   types.NewString(""),
		AccountRights: types.NewList[*ServiceItemAccountRight](),
	}

	siri.AccountRights.Type = NewServiceItemAccountRight()

	return siri
}
