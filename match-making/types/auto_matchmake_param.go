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

// AutoMatchmakeParam holds parameters for a matchmake session
type AutoMatchmakeParam struct {
	types.Structure
	SourceMatchmakeSession   *MatchmakeSession
	AdditionalParticipants   *types.List[*types.PID]
	GIDForParticipationCheck *types.PrimitiveU32
	AutoMatchmakeOption      *types.PrimitiveU32
	JoinMessage              string
	ParticipationCount       *types.PrimitiveU16
	LstSearchCriteria        []*MatchmakeSessionSearchCriteria
	TargetGIDs               *types.List[*types.PrimitiveU32]
}

// ExtractFrom extracts the AutoMatchmakeParam from the given readable
func (autoMatchmakeParam *AutoMatchmakeParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = autoMatchmakeParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read AutoMatchmakeParam header. %s", err.Error())
	}

	err = autoMatchmakeParam.SourceMatchmakeSession.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.SourceMatchmakeSession. %s", err.Error())
	}

	err = autoMatchmakeParam.AdditionalParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.AdditionalParticipants. %s", err.Error())
	}

	err = autoMatchmakeParam.GIDForParticipationCheck.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.GIDForParticipationCheck. %s", err.Error())
	}

	err = autoMatchmakeParam.AutoMatchmakeOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.AutoMatchmakeOption. %s", err.Error())
	}

	err = autoMatchmakeParam.JoinMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.JoinMessage. %s", err.Error())
	}

	err = autoMatchmakeParam.ParticipationCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.ParticipationCount. %s", err.Error())
	}

	lstSearchCriteria, err := nex.StreamReadListStructure(stream, NewMatchmakeSessionSearchCriteria())
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.LstSearchCriteria. %s", err.Error())
	}

	autoMatchmakeParam.LstSearchCriteria = lstSearchCriteria
	err = autoMatchmakeParam.TargetGIDs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.TargetGIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of AutoMatchmakeParam
func (autoMatchmakeParam *AutoMatchmakeParam) Copy() types.RVType {
	copied := NewAutoMatchmakeParam()

	copied.StructureVersion = autoMatchmakeParam.StructureVersion

	copied.SourceMatchmakeSession = autoMatchmakeParam.SourceMatchmakeSession.Copy().(*MatchmakeSession)

	copied.AdditionalParticipants = make(*types.List[*types.PID], len(autoMatchmakeParam.AdditionalParticipants))

	for i := 0; i < len(autoMatchmakeParam.AdditionalParticipants); i++ {
		copied.AdditionalParticipants[i] = autoMatchmakeParam.AdditionalParticipants[i].Copy()
	}

	copied.GIDForParticipationCheck = autoMatchmakeParam.GIDForParticipationCheck
	copied.AutoMatchmakeOption = autoMatchmakeParam.AutoMatchmakeOption
	copied.JoinMessage = autoMatchmakeParam.JoinMessage
	copied.ParticipationCount = autoMatchmakeParam.ParticipationCount
	copied.LstSearchCriteria = make([]*MatchmakeSessionSearchCriteria, len(autoMatchmakeParam.LstSearchCriteria))

	for i := 0; i < len(autoMatchmakeParam.LstSearchCriteria); i++ {
		copied.LstSearchCriteria[i] = autoMatchmakeParam.LstSearchCriteria[i].Copy().(*MatchmakeSessionSearchCriteria)
	}

	copied.TargetGIDs = make(*types.List[*types.PrimitiveU32], len(autoMatchmakeParam.TargetGIDs))

	copy(copied.TargetGIDs, autoMatchmakeParam.TargetGIDs)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (autoMatchmakeParam *AutoMatchmakeParam) Equals(o types.RVType) bool {
	if _, ok := o.(*AutoMatchmakeParam); !ok {
		return false
	}

	other := o.(*AutoMatchmakeParam)

	if autoMatchmakeParam.StructureVersion != other.StructureVersion {
		return false
	}

	if autoMatchmakeParam.SourceMatchmakeSession != nil && other.SourceMatchmakeSession != nil {
		if autoMatchmakeParam.SourceMatchmakeSession.Equals(other.SourceMatchmakeSession) {
			return false
		}
	}

	if len(autoMatchmakeParam.AdditionalParticipants) != len(other.AdditionalParticipants) {
		return false
	}

	for i := 0; i < len(autoMatchmakeParam.AdditionalParticipants); i++ {
		if !autoMatchmakeParam.AdditionalParticipants[i].Equals(other.AdditionalParticipants[i]) {
			return false
		}
	}

	if !autoMatchmakeParam.GIDForParticipationCheck.Equals(other.GIDForParticipationCheck) {
		return false
	}

	if !autoMatchmakeParam.AutoMatchmakeOption.Equals(other.AutoMatchmakeOption) {
		return false
	}

	if !autoMatchmakeParam.JoinMessage.Equals(other.JoinMessage) {
		return false
	}

	if !autoMatchmakeParam.ParticipationCount.Equals(other.ParticipationCount) {
		return false
	}

	if len(autoMatchmakeParam.LstSearchCriteria) != len(other.LstSearchCriteria) {
		return false
	}

	for i := 0; i < len(autoMatchmakeParam.LstSearchCriteria); i++ {
		if !autoMatchmakeParam.LstSearchCriteria[i].Equals(other.LstSearchCriteria[i]) {
			return false
		}
	}

	if len(autoMatchmakeParam.TargetGIDs) != len(other.TargetGIDs) {
		return false
	}

	for i := 0; i < len(autoMatchmakeParam.TargetGIDs); i++ {
		if autoMatchmakeParam.TargetGIDs[i] != other.TargetGIDs[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (autoMatchmakeParam *AutoMatchmakeParam) String() string {
	return autoMatchmakeParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (autoMatchmakeParam *AutoMatchmakeParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("AutoMatchmakeParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, autoMatchmakeParam.StructureVersion))

	if autoMatchmakeParam.SourceMatchmakeSession != nil {
		b.WriteString(fmt.Sprintf("%sSourceMatchmakeSession: %s,\n", indentationValues, autoMatchmakeParam.SourceMatchmakeSession.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sSourceMatchmakeSession: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sAdditionalParticipants: %v,\n", indentationValues, autoMatchmakeParam.AdditionalParticipants))
	b.WriteString(fmt.Sprintf("%sGIDForParticipationCheck: %d,\n", indentationValues, autoMatchmakeParam.GIDForParticipationCheck))
	b.WriteString(fmt.Sprintf("%sAutoMatchmakeOption: %d,\n", indentationValues, autoMatchmakeParam.AutoMatchmakeOption))
	b.WriteString(fmt.Sprintf("%sJoinMessage: %q,\n", indentationValues, autoMatchmakeParam.JoinMessage))
	b.WriteString(fmt.Sprintf("%sParticipationCount: %d,\n", indentationValues, autoMatchmakeParam.ParticipationCount))

	if len(autoMatchmakeParam.LstSearchCriteria) == 0 {
		b.WriteString(fmt.Sprintf("%sLstSearchCriteria: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sLstSearchCriteria: [\n", indentationValues))

		for i := 0; i < len(autoMatchmakeParam.LstSearchCriteria); i++ {
			str := autoMatchmakeParam.LstSearchCriteria[i].FormatToString(indentationLevel + 2)
			if i == len(autoMatchmakeParam.LstSearchCriteria)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sTargetGIDs: %v\n", indentationValues, autoMatchmakeParam.TargetGIDs))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewAutoMatchmakeParam returns a new AutoMatchmakeParam
func NewAutoMatchmakeParam() *AutoMatchmakeParam {
	return &AutoMatchmakeParam{}
}
