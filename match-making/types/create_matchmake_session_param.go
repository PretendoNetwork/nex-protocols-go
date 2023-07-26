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

// CreateMatchmakeSessionParam holds parameters for a matchmake session
type CreateMatchmakeSessionParam struct {
	nex.Structure
	SourceMatchmakeSession       *MatchmakeSession
	AdditionalParticipants       []uint32
	GIDForParticipationCheck     uint32
	CreateMatchmakeSessionOption uint32
	JoinMessage                  string
	ParticipationCount           uint16
}

// ExtractFromStream extracts a CreateMatchmakeSessionParam structure from a stream
func (createMatchmakeSessionParam *CreateMatchmakeSessionParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	sourceMatchmakeSession, err := stream.ReadStructure(NewMatchmakeSession())
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.SourceMatchmakeSession. %s", err.Error())
	}

	createMatchmakeSessionParam.SourceMatchmakeSession = sourceMatchmakeSession.(*MatchmakeSession)
	createMatchmakeSessionParam.AdditionalParticipants, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.AdditionalParticipants. %s", err.Error())
	}

	createMatchmakeSessionParam.GIDForParticipationCheck, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.GIDForParticipationCheck. %s", err.Error())
	}

	createMatchmakeSessionParam.CreateMatchmakeSessionOption, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.CreateMatchmakeSessionOption. %s", err.Error())
	}

	createMatchmakeSessionParam.JoinMessage, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.JoinMessage. %s", err.Error())
	}

	createMatchmakeSessionParam.ParticipationCount, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.ParticipationCount. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of CreateMatchmakeSessionParam
func (createMatchmakeSessionParam *CreateMatchmakeSessionParam) Copy() nex.StructureInterface {
	copied := NewCreateMatchmakeSessionParam()

	if createMatchmakeSessionParam.SourceMatchmakeSession != nil {
		copied.SourceMatchmakeSession = createMatchmakeSessionParam.SourceMatchmakeSession.Copy().(*MatchmakeSession)
	}

	copied.AdditionalParticipants = make([]uint32, len(createMatchmakeSessionParam.AdditionalParticipants))

	copy(copied.AdditionalParticipants, createMatchmakeSessionParam.AdditionalParticipants)

	copied.GIDForParticipationCheck = createMatchmakeSessionParam.GIDForParticipationCheck
	copied.CreateMatchmakeSessionOption = createMatchmakeSessionParam.CreateMatchmakeSessionOption
	copied.JoinMessage = createMatchmakeSessionParam.JoinMessage
	copied.ParticipationCount = createMatchmakeSessionParam.ParticipationCount

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (createMatchmakeSessionParam *CreateMatchmakeSessionParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*CreateMatchmakeSessionParam)

	if createMatchmakeSessionParam.SourceMatchmakeSession != nil && other.SourceMatchmakeSession == nil {
		return false
	}

	if createMatchmakeSessionParam.SourceMatchmakeSession == nil && other.SourceMatchmakeSession != nil {
		return false
	}

	if createMatchmakeSessionParam.SourceMatchmakeSession != nil && other.SourceMatchmakeSession != nil {
		if !createMatchmakeSessionParam.SourceMatchmakeSession.Equals(other.SourceMatchmakeSession) {
			return false
		}
	}

	if len(createMatchmakeSessionParam.AdditionalParticipants) != len(other.AdditionalParticipants) {
		return false
	}

	for i := 0; i < len(createMatchmakeSessionParam.AdditionalParticipants); i++ {
		if createMatchmakeSessionParam.AdditionalParticipants[i] != other.AdditionalParticipants[i] {
			return false
		}
	}

	if createMatchmakeSessionParam.GIDForParticipationCheck != other.GIDForParticipationCheck {
		return false
	}

	if createMatchmakeSessionParam.CreateMatchmakeSessionOption != other.CreateMatchmakeSessionOption {
		return false
	}

	if createMatchmakeSessionParam.JoinMessage != other.JoinMessage {
		return false
	}

	if createMatchmakeSessionParam.ParticipationCount != other.ParticipationCount {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (createMatchmakeSessionParam *CreateMatchmakeSessionParam) String() string {
	return createMatchmakeSessionParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (createMatchmakeSessionParam *CreateMatchmakeSessionParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CreateMatchmakeSessionParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, createMatchmakeSessionParam.StructureVersion()))

	if createMatchmakeSessionParam.SourceMatchmakeSession != nil {
		b.WriteString(fmt.Sprintf("%sSourceMatchmakeSession: %s,\n", indentationValues, createMatchmakeSessionParam.SourceMatchmakeSession.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sSourceMatchmakeSession: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sAdditionalParticipants: %v,\n", indentationValues, createMatchmakeSessionParam.AdditionalParticipants))
	b.WriteString(fmt.Sprintf("%sGIDForParticipationCheck: %d,\n", indentationValues, createMatchmakeSessionParam.GIDForParticipationCheck))
	b.WriteString(fmt.Sprintf("%sCreateMatchmakeSessionOption: %d,\n", indentationValues, createMatchmakeSessionParam.CreateMatchmakeSessionOption))
	b.WriteString(fmt.Sprintf("%sJoinMessage: %q,\n", indentationValues, createMatchmakeSessionParam.JoinMessage))
	b.WriteString(fmt.Sprintf("%sParticipationCount: %d\n", indentationValues, createMatchmakeSessionParam.ParticipationCount))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewCreateMatchmakeSessionParam returns a new CreateMatchmakeSessionParam
func NewCreateMatchmakeSessionParam() *CreateMatchmakeSessionParam {
	return &CreateMatchmakeSessionParam{}
}
