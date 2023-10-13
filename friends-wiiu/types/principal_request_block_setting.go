// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// PrincipalRequestBlockSetting contains unknow data
type PrincipalRequestBlockSetting struct {
	nex.Structure
	*nex.Data
	PID       uint32
	IsBlocked bool
}

// Bytes encodes the PrincipalRequestBlockSetting and returns a byte array
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(principalRequestBlockSetting.PID)
	stream.WriteBool(principalRequestBlockSetting.IsBlocked)

	return stream.Bytes()
}

// Copy returns a new copied instance of PrincipalRequestBlockSetting
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) Copy() nex.StructureInterface {
	copied := NewPrincipalRequestBlockSetting()

	copied.SetStructureVersion(principalRequestBlockSetting.StructureVersion())

	if principalRequestBlockSetting.ParentType() != nil {
		copied.Data = principalRequestBlockSetting.ParentType().Copy().(*nex.Data)
	} else {
		copied.Data = nex.NewData()
	}

	copied.PID = principalRequestBlockSetting.PID
	copied.IsBlocked = principalRequestBlockSetting.IsBlocked

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PrincipalRequestBlockSetting)

	if principalRequestBlockSetting.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !principalRequestBlockSetting.ParentType().Equals(other.ParentType()) {
		return false
	}

	if principalRequestBlockSetting.PID != other.PID {
		return false
	}

	if principalRequestBlockSetting.IsBlocked != other.IsBlocked {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, principalRequestBlockSetting.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, principalRequestBlockSetting.PID))
	b.WriteString(fmt.Sprintf("%sIsBlocked: %t\n", indentationValues, principalRequestBlockSetting.IsBlocked))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPrincipalRequestBlockSetting returns a new PrincipalRequestBlockSetting
func NewPrincipalRequestBlockSetting() *PrincipalRequestBlockSetting {
	return &PrincipalRequestBlockSetting{}
}
