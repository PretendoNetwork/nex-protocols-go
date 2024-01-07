// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// PrincipalRequestBlockSetting contains unknow data
type PrincipalRequestBlockSetting struct {
	types.Structure
	*types.Data
	PID       *types.PrimitiveU32
	IsBlocked *types.PrimitiveBool
}

// WriteTo writes the PrincipalRequestBlockSetting to the given writable
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	principalRequestBlockSetting.PID.WriteTo(contentWritable)
	principalRequestBlockSetting.IsBlocked.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	principalRequestBlockSetting.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of PrincipalRequestBlockSetting
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) Copy() types.RVType {
	copied := NewPrincipalRequestBlockSetting()

	copied.StructureVersion = principalRequestBlockSetting.StructureVersion

	copied.Data = principalRequestBlockSetting.Data.Copy().(*types.Data)

	copied.PID = principalRequestBlockSetting.PID
	copied.IsBlocked = principalRequestBlockSetting.IsBlocked

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) Equals(o types.RVType) bool {
	if _, ok := o.(*PrincipalRequestBlockSetting); !ok {
		return false
	}

	other := o.(*PrincipalRequestBlockSetting)

	if principalRequestBlockSetting.StructureVersion != other.StructureVersion {
		return false
	}

	if !principalRequestBlockSetting.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !principalRequestBlockSetting.PID.Equals(other.PID) {
		return false
	}

	if !principalRequestBlockSetting.IsBlocked.Equals(other.IsBlocked) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) String() string {
	return principalRequestBlockSetting.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PrincipalRequestBlockSetting{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, principalRequestBlockSetting.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, principalRequestBlockSetting.PID))
	b.WriteString(fmt.Sprintf("%sIsBlocked: %t\n", indentationValues, principalRequestBlockSetting.IsBlocked))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPrincipalRequestBlockSetting returns a new PrincipalRequestBlockSetting
func NewPrincipalRequestBlockSetting() *PrincipalRequestBlockSetting {
	return &PrincipalRequestBlockSetting{}
}
