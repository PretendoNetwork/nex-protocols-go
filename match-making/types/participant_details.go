// Package match_making_types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package match_making_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ParticipantDetails holds information a participant
type ParticipantDetails struct {
	nex.Structure
	IDParticipant  uint32
	StrName        string
	StrMessage     string
	UIParticipants uint16
}

// ExtractFromStream extracts a ParticipantDetails structure from a stream
func (participantDetails *ParticipantDetails) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	participantDetails.IDParticipant, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ParticipantDetails.IDParticipant. %s", err.Error())
	}

	participantDetails.StrName, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ParticipantDetails.StrName. %s", err.Error())
	}

	participantDetails.StrMessage, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ParticipantDetails.StrMessage. %s", err.Error())
	}

	participantDetails.UIParticipants, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ParticipantDetails.UIParticipants. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ParticipantDetails and returns a byte array
func (participantDetails *ParticipantDetails) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(participantDetails.IDParticipant)
	stream.WriteString(participantDetails.StrName)
	stream.WriteString(participantDetails.StrMessage)
	stream.WriteUInt16LE(participantDetails.UIParticipants)

	return stream.Bytes()
}

// Copy returns a new copied instance of ParticipantDetails
func (participantDetails *ParticipantDetails) Copy() nex.StructureInterface {
	copied := NewParticipantDetails()

	copied.IDParticipant = participantDetails.IDParticipant
	copied.StrName = participantDetails.StrName
	copied.StrMessage = participantDetails.StrMessage
	copied.UIParticipants = participantDetails.UIParticipants

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (participantDetails *ParticipantDetails) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ParticipantDetails)

	if participantDetails.IDParticipant != other.IDParticipant {
		return false
	}

	if participantDetails.StrName != other.StrName {
		return false
	}

	if participantDetails.StrMessage != other.StrMessage {
		return false
	}

	if participantDetails.UIParticipants != other.UIParticipants {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (participantDetails *ParticipantDetails) String() string {
	return participantDetails.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (participantDetails *ParticipantDetails) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ParticipantDetails{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, participantDetails.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sIDParticipant: %d,\n", indentationValues, participantDetails.IDParticipant))
	b.WriteString(fmt.Sprintf("%sStrName: %q,\n", indentationValues, participantDetails.StrName))
	b.WriteString(fmt.Sprintf("%sStrMessage: %q,\n", indentationValues, participantDetails.StrMessage))
	b.WriteString(fmt.Sprintf("%sUIParticipants: %d\n", indentationValues, participantDetails.UIParticipants))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewParticipantDetails returns a new ParticipantDetails
func NewParticipantDetails() *ParticipantDetails {
	return &ParticipantDetails{}
}
