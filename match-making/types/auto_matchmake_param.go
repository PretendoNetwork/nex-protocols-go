package match_making_types

import (
	"fmt"
	"strings"

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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, autoMatchmakeParam.StructureVersion()))

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
