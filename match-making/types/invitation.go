// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Invitation holds an invitation a Gathering
type Invitation struct {
	types.Structure
	IDGathering *types.PrimitiveU32
	IDGuest     *types.PrimitiveU32
	StrMessage  string
}

// ExtractFrom extracts the Invitation from the given readable
func (invitation *Invitation) ExtractFrom(readable types.Readable) error {
	var err error

	if err = invitation.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Invitation header. %s", err.Error())
	}

	err = invitation.IDGathering.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Invitation.IDGathering. %s", err.Error())
	}

	err = invitation.IDGuest.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Invitation.IDGuest. %s", err.Error())
	}

	err = invitation.StrMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Invitation.StrMessage. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Invitation to the given writable
func (invitation *Invitation) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	invitation.IDGathering.WriteTo(contentWritable)
	invitation.IDGuest.WriteTo(contentWritable)
	invitation.StrMessage.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	invitation.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Invitation
func (invitation *Invitation) Copy() types.RVType {
	copied := NewInvitation()

	copied.StructureVersion = invitation.StructureVersion

	copied.IDGathering = invitation.IDGathering
	copied.IDGuest = invitation.IDGuest
	copied.StrMessage = invitation.StrMessage

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (invitation *Invitation) Equals(o types.RVType) bool {
	if _, ok := o.(*Invitation); !ok {
		return false
	}

	other := o.(*Invitation)

	if invitation.StructureVersion != other.StructureVersion {
		return false
	}

	if !invitation.IDGathering.Equals(other.IDGathering) {
		return false
	}

	if !invitation.IDGuest.Equals(other.IDGuest) {
		return false
	}

	if !invitation.StrMessage.Equals(other.StrMessage) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (invitation *Invitation) String() string {
	return invitation.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (invitation *Invitation) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Invitation{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, invitation.StructureVersion))
	b.WriteString(fmt.Sprintf("%sIDGathering: %d,\n", indentationValues, invitation.IDGathering))
	b.WriteString(fmt.Sprintf("%sIDGuest: %d,\n", indentationValues, invitation.IDGuest))
	b.WriteString(fmt.Sprintf("%sStrMessage: %q\n", indentationValues, invitation.StrMessage))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewInvitation returns a new Invitation
func NewInvitation() *Invitation {
	return &Invitation{}
}
