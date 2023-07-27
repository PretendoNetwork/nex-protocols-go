// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Invitation holds an invitation a Gathering
type Invitation struct {
	nex.Structure
	IDGathering uint32
	IDGuest     uint32
	StrMessage  string
}

// ExtractFromStream extracts a Invitation structure from a stream
func (invitation *Invitation) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	invitation.IDGathering, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Invitation.IDGathering. %s", err.Error())
	}

	invitation.IDGuest, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Invitation.IDGuest. %s", err.Error())
	}

	invitation.StrMessage, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract Invitation.StrMessage. %s", err.Error())
	}

	return nil
}

// Bytes encodes the Invitation and returns a byte array
func (invitation *Invitation) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(invitation.IDGathering)
	stream.WriteUInt32LE(invitation.IDGuest)
	stream.WriteString(invitation.StrMessage)

	return stream.Bytes()
}

// Copy returns a new copied instance of Invitation
func (invitation *Invitation) Copy() nex.StructureInterface {
	copied := NewInvitation()

	copied.IDGathering = invitation.IDGathering
	copied.IDGuest = invitation.IDGuest
	copied.StrMessage = invitation.StrMessage

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (invitation *Invitation) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Invitation)

	if invitation.IDGathering != other.IDGathering {
		return false
	}

	if invitation.IDGuest != other.IDGuest {
		return false
	}

	if invitation.StrMessage != other.StrMessage {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, invitation.StructureVersion()))
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
