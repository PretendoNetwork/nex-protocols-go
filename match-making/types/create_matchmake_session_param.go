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

// CreateMatchmakeSessionParam holds parameters for a matchmake session
type CreateMatchmakeSessionParam struct {
	types.Structure
	SourceMatchmakeSession       *MatchmakeSession
	AdditionalParticipants       *types.List[*types.PID]
	GIDForParticipationCheck     *types.PrimitiveU32
	CreateMatchmakeSessionOption *types.PrimitiveU32
	JoinMessage                  string
	ParticipationCount           *types.PrimitiveU16
}

// ExtractFrom extracts the CreateMatchmakeSessionParam from the given readable
func (createMatchmakeSessionParam *CreateMatchmakeSessionParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = createMatchmakeSessionParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read CreateMatchmakeSessionParam header. %s", err.Error())
	}

	err = createMatchmakeSessionParam.SourceMatchmakeSession.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.SourceMatchmakeSession. %s", err.Error())
	}

	err = createMatchmakeSessionParam.AdditionalParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.AdditionalParticipants. %s", err.Error())
	}

	err = createMatchmakeSessionParam.GIDForParticipationCheck.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.GIDForParticipationCheck. %s", err.Error())
	}

	err = createMatchmakeSessionParam.CreateMatchmakeSessionOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.CreateMatchmakeSessionOption. %s", err.Error())
	}

	err = createMatchmakeSessionParam.JoinMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.JoinMessage. %s", err.Error())
	}

	err = createMatchmakeSessionParam.ParticipationCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CreateMatchmakeSessionParam.ParticipationCount. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of CreateMatchmakeSessionParam
func (createMatchmakeSessionParam *CreateMatchmakeSessionParam) Copy() types.RVType {
	copied := NewCreateMatchmakeSessionParam()

	copied.StructureVersion = createMatchmakeSessionParam.StructureVersion

	copied.SourceMatchmakeSession = createMatchmakeSessionParam.SourceMatchmakeSession.Copy().(*MatchmakeSession)

	copied.AdditionalParticipants = make(*types.List[*types.PID], len(createMatchmakeSessionParam.AdditionalParticipants))

	for i := 0; i < len(createMatchmakeSessionParam.AdditionalParticipants); i++ {
		copied.AdditionalParticipants[i] = createMatchmakeSessionParam.AdditionalParticipants[i].Copy()
	}

	copied.GIDForParticipationCheck = createMatchmakeSessionParam.GIDForParticipationCheck
	copied.CreateMatchmakeSessionOption = createMatchmakeSessionParam.CreateMatchmakeSessionOption
	copied.JoinMessage = createMatchmakeSessionParam.JoinMessage
	copied.ParticipationCount = createMatchmakeSessionParam.ParticipationCount

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (createMatchmakeSessionParam *CreateMatchmakeSessionParam) Equals(o types.RVType) bool {
	if _, ok := o.(*CreateMatchmakeSessionParam); !ok {
		return false
	}

	other := o.(*CreateMatchmakeSessionParam)

	if createMatchmakeSessionParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !createMatchmakeSessionParam.SourceMatchmakeSession.Equals(other.SourceMatchmakeSession) {
		return false
	}

	if len(createMatchmakeSessionParam.AdditionalParticipants) != len(other.AdditionalParticipants) {
		return false
	}

	for i := 0; i < len(createMatchmakeSessionParam.AdditionalParticipants); i++ {
		if !createMatchmakeSessionParam.AdditionalParticipants[i].Equals(other.AdditionalParticipants[i]) {
			return false
		}
	}

	if !createMatchmakeSessionParam.GIDForParticipationCheck.Equals(other.GIDForParticipationCheck) {
		return false
	}

	if !createMatchmakeSessionParam.CreateMatchmakeSessionOption.Equals(other.CreateMatchmakeSessionOption) {
		return false
	}

	if !createMatchmakeSessionParam.JoinMessage.Equals(other.JoinMessage) {
		return false
	}

	if !createMatchmakeSessionParam.ParticipationCount.Equals(other.ParticipationCount) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, createMatchmakeSessionParam.StructureVersion))

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
