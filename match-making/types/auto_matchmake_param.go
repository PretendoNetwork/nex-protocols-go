package match_making_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// AutoMatchmakeParam holds parameters for a matchmake session
type AutoMatchmakeParam struct {
	nex.Structure
	SourceMatchmakeSession   *MatchmakeSession
	AdditionalParticipants   []uint32
	GIDForParticipationCheck uint32
	AutoMatchmakeOption      uint32
	JoinMessage              string
	ParticipationCount       uint16
	LstSearchCriteria        []*MatchmakeSessionSearchCriteria
	TargetGIDs               []uint32
}

// ExtractFromStream extracts a AutoMatchmakeParam structure from a stream
func (autoMatchmakeParam *AutoMatchmakeParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	sourceMatchmakeSession, err := stream.ReadStructure(NewMatchmakeSession())
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.SourceMatchmakeSession. %s", err.Error())
	}

	autoMatchmakeParam.SourceMatchmakeSession = sourceMatchmakeSession.(*MatchmakeSession)
	autoMatchmakeParam.AdditionalParticipants, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.AdditionalParticipants. %s", err.Error())
	}

	autoMatchmakeParam.GIDForParticipationCheck, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.GIDForParticipationCheck. %s", err.Error())
	}

	autoMatchmakeParam.AutoMatchmakeOption, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.AutoMatchmakeOption. %s", err.Error())
	}

	autoMatchmakeParam.JoinMessage, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.JoinMessage. %s", err.Error())
	}

	autoMatchmakeParam.ParticipationCount, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.ParticipationCount. %s", err.Error())
	}

	lstSearchCriteria, err := stream.ReadListStructure(NewMatchmakeSessionSearchCriteria())
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.LstSearchCriteria. %s", err.Error())
	}

	autoMatchmakeParam.LstSearchCriteria = lstSearchCriteria.([]*MatchmakeSessionSearchCriteria)
	autoMatchmakeParam.TargetGIDs, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract AutoMatchmakeParam.TargetGIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of AutoMatchmakeParam
func (autoMatchmakeParam *AutoMatchmakeParam) Copy() nex.StructureInterface {
	copied := NewAutoMatchmakeParam()

	if autoMatchmakeParam.SourceMatchmakeSession != nil {
		copied.SourceMatchmakeSession = autoMatchmakeParam.SourceMatchmakeSession.Copy().(*MatchmakeSession)
	}

	copied.AdditionalParticipants = make([]uint32, len(autoMatchmakeParam.AdditionalParticipants))

	copy(copied.AdditionalParticipants, autoMatchmakeParam.AdditionalParticipants)

	copied.GIDForParticipationCheck = autoMatchmakeParam.GIDForParticipationCheck
	copied.AutoMatchmakeOption = autoMatchmakeParam.AutoMatchmakeOption
	copied.JoinMessage = autoMatchmakeParam.JoinMessage
	copied.ParticipationCount = autoMatchmakeParam.ParticipationCount
	copied.LstSearchCriteria = make([]*MatchmakeSessionSearchCriteria, len(autoMatchmakeParam.LstSearchCriteria))

	for i := 0; i < len(autoMatchmakeParam.LstSearchCriteria); i++ {
		copied.LstSearchCriteria[i] = autoMatchmakeParam.LstSearchCriteria[i].Copy().(*MatchmakeSessionSearchCriteria)
	}

	copied.TargetGIDs = make([]uint32, len(autoMatchmakeParam.TargetGIDs))

	copy(copied.TargetGIDs, autoMatchmakeParam.TargetGIDs)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (autoMatchmakeParam *AutoMatchmakeParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*AutoMatchmakeParam)

	if autoMatchmakeParam.SourceMatchmakeSession != nil && other.SourceMatchmakeSession == nil {
		return false
	}

	if autoMatchmakeParam.SourceMatchmakeSession == nil && other.SourceMatchmakeSession != nil {
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
		if autoMatchmakeParam.AdditionalParticipants[i] != other.AdditionalParticipants[i] {
			return false
		}
	}

	if autoMatchmakeParam.GIDForParticipationCheck != other.GIDForParticipationCheck {
		return false
	}

	if autoMatchmakeParam.AutoMatchmakeOption != other.AutoMatchmakeOption {
		return false
	}

	if autoMatchmakeParam.JoinMessage != other.JoinMessage {
		return false
	}

	if autoMatchmakeParam.ParticipationCount != other.ParticipationCount {
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

// NewAutoMatchmakeParam returns a new AutoMatchmakeParam
func NewAutoMatchmakeParam() *AutoMatchmakeParam {
	return &AutoMatchmakeParam{}
}
