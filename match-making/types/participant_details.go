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

// ParticipantDetails holds information a participant
type ParticipantDetails struct {
	types.Structure
	IDParticipant  *types.PID
	StrName        string
	StrMessage     string
	UIParticipants *types.PrimitiveU16
}

// ExtractFrom extracts the ParticipantDetails from the given readable
func (participantDetails *ParticipantDetails) ExtractFrom(readable types.Readable) error {
	var err error

	if err = participantDetails.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ParticipantDetails header. %s", err.Error())
	}

	err = participantDetails.IDParticipant.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ParticipantDetails.IDParticipant. %s", err.Error())
	}

	err = participantDetails.StrName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ParticipantDetails.StrName. %s", err.Error())
	}

	err = participantDetails.StrMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ParticipantDetails.StrMessage. %s", err.Error())
	}

	err = participantDetails.UIParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ParticipantDetails.UIParticipants. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ParticipantDetails to the given writable
func (participantDetails *ParticipantDetails) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	participantDetails.IDParticipant.WriteTo(contentWritable)
	participantDetails.StrName.WriteTo(contentWritable)
	participantDetails.StrMessage.WriteTo(contentWritable)
	participantDetails.UIParticipants.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	participantDetails.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ParticipantDetails
func (participantDetails *ParticipantDetails) Copy() types.RVType {
	copied := NewParticipantDetails()

	copied.StructureVersion = participantDetails.StructureVersion

	copied.IDParticipant = participantDetails.IDParticipant.Copy()
	copied.StrName = participantDetails.StrName
	copied.StrMessage = participantDetails.StrMessage
	copied.UIParticipants = participantDetails.UIParticipants

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (participantDetails *ParticipantDetails) Equals(o types.RVType) bool {
	if _, ok := o.(*ParticipantDetails); !ok {
		return false
	}

	other := o.(*ParticipantDetails)

	if participantDetails.StructureVersion != other.StructureVersion {
		return false
	}

	if !participantDetails.IDParticipant.Equals(other.IDParticipant) {
		return false
	}

	if !participantDetails.StrName.Equals(other.StrName) {
		return false
	}

	if !participantDetails.StrMessage.Equals(other.StrMessage) {
		return false
	}

	if !participantDetails.UIParticipants.Equals(other.UIParticipants) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, participantDetails.StructureVersion))
	b.WriteString(fmt.Sprintf("%sIDParticipant: %s,\n", indentationValues, participantDetails.IDParticipant.FormatToString(indentationLevel+1)))
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
