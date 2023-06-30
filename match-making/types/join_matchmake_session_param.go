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

// JoinMatchmakeSessionParam holds parameters for a matchmake session
type JoinMatchmakeSessionParam struct {
	nex.Structure
	GID                          uint32
	AdditionalParticipants       []uint32
	GIDForParticipationCheck     uint32
	JoinMatchmakeSessionOption   uint32
	JoinMatchmakeSessionBehavior uint8
	StrUserPassword              string
	StrSystemPassword            string
	JoinMessage                  string
	ParticipationCount           uint16
	ExtraParticipants            uint16
	BlockListParam               *MatchmakeBlockListParam
}

// ExtractFromStream extracts a JoinMatchmakeSessionParam structure from a stream
func (joinMatchmakeSessionParam *JoinMatchmakeSessionParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	joinMatchmakeSessionParam.GID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.GID. %s", err.Error())
	}

	joinMatchmakeSessionParam.AdditionalParticipants, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.AdditionalParticipants. %s", err.Error())
	}

	joinMatchmakeSessionParam.GIDForParticipationCheck, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.GIDForParticipationCheck. %s", err.Error())
	}

	joinMatchmakeSessionParam.JoinMatchmakeSessionOption, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.JoinMatchmakeSessionOption. %s", err.Error())
	}

	joinMatchmakeSessionParam.JoinMatchmakeSessionBehavior, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.JoinMatchmakeSessionBehavior. %s", err.Error())
	}

	joinMatchmakeSessionParam.StrUserPassword, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.StrUserPassword. %s", err.Error())
	}

	joinMatchmakeSessionParam.StrSystemPassword, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.StrSystemPassword. %s", err.Error())
	}

	joinMatchmakeSessionParam.JoinMessage, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.JoinMessage. %s", err.Error())
	}

	joinMatchmakeSessionParam.ParticipationCount, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.ParticipationCount. %s", err.Error())
	}

	joinMatchmakeSessionParam.ExtraParticipants, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.ExtraParticipants. %s", err.Error())
	}

	blockListParam, err := stream.ReadStructure(NewMatchmakeBlockListParam())
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.BlockListParam. %s", err.Error())
	}

	joinMatchmakeSessionParam.BlockListParam = blockListParam.(*MatchmakeBlockListParam)

	return nil
}

// Copy returns a new copied instance of JoinMatchmakeSessionParam
func (joinMatchmakeSessionParam *JoinMatchmakeSessionParam) Copy() nex.StructureInterface {
	copied := NewJoinMatchmakeSessionParam()

	copied.GID = joinMatchmakeSessionParam.GID
	copied.AdditionalParticipants = make([]uint32, len(joinMatchmakeSessionParam.AdditionalParticipants))

	copy(copied.AdditionalParticipants, joinMatchmakeSessionParam.AdditionalParticipants)

	copied.GIDForParticipationCheck = joinMatchmakeSessionParam.GIDForParticipationCheck
	copied.JoinMatchmakeSessionOption = joinMatchmakeSessionParam.JoinMatchmakeSessionOption
	copied.JoinMatchmakeSessionBehavior = joinMatchmakeSessionParam.JoinMatchmakeSessionBehavior
	copied.StrUserPassword = joinMatchmakeSessionParam.StrUserPassword
	copied.StrSystemPassword = joinMatchmakeSessionParam.StrSystemPassword
	copied.JoinMessage = joinMatchmakeSessionParam.JoinMessage
	copied.ParticipationCount = joinMatchmakeSessionParam.ParticipationCount
	copied.ExtraParticipants = joinMatchmakeSessionParam.ExtraParticipants

	if joinMatchmakeSessionParam.BlockListParam != nil {
		copied.BlockListParam = joinMatchmakeSessionParam.BlockListParam.Copy().(*MatchmakeBlockListParam)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (joinMatchmakeSessionParam *JoinMatchmakeSessionParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*JoinMatchmakeSessionParam)

	if joinMatchmakeSessionParam.GID != other.GID {
		return false
	}

	if len(joinMatchmakeSessionParam.AdditionalParticipants) != len(other.AdditionalParticipants) {
		return false
	}

	for i := 0; i < len(joinMatchmakeSessionParam.AdditionalParticipants); i++ {
		if joinMatchmakeSessionParam.AdditionalParticipants[i] != other.AdditionalParticipants[i] {
			return false
		}
	}

	if joinMatchmakeSessionParam.GIDForParticipationCheck != other.GIDForParticipationCheck {
		return false
	}

	if joinMatchmakeSessionParam.JoinMatchmakeSessionOption != other.JoinMatchmakeSessionOption {
		return false
	}

	if joinMatchmakeSessionParam.JoinMatchmakeSessionBehavior != other.JoinMatchmakeSessionBehavior {
		return false
	}

	if joinMatchmakeSessionParam.StrUserPassword != other.StrUserPassword {
		return false
	}

	if joinMatchmakeSessionParam.StrSystemPassword != other.StrSystemPassword {
		return false
	}

	if joinMatchmakeSessionParam.JoinMessage != other.JoinMessage {
		return false
	}

	if joinMatchmakeSessionParam.ParticipationCount != other.ParticipationCount {
		return false
	}

	if joinMatchmakeSessionParam.ExtraParticipants != other.ExtraParticipants {
		return false
	}

	if joinMatchmakeSessionParam.BlockListParam != nil {
		return joinMatchmakeSessionParam.BlockListParam.Equals(other.BlockListParam)
	}

	return true
}

// String returns a string representation of the struct
func (joinMatchmakeSessionParam *JoinMatchmakeSessionParam) String() string {
	return joinMatchmakeSessionParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (joinMatchmakeSessionParam *JoinMatchmakeSessionParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("JoinMatchmakeSessionParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, joinMatchmakeSessionParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sGID: %d,\n", indentationValues, joinMatchmakeSessionParam.GID))
	b.WriteString(fmt.Sprintf("%sAdditionalParticipants: %v,\n", indentationValues, joinMatchmakeSessionParam.AdditionalParticipants))
	b.WriteString(fmt.Sprintf("%sGIDForParticipationCheck: %d,\n", indentationValues, joinMatchmakeSessionParam.GIDForParticipationCheck))
	b.WriteString(fmt.Sprintf("%sJoinMatchmakeSessionOption: %d,\n", indentationValues, joinMatchmakeSessionParam.JoinMatchmakeSessionOption))
	b.WriteString(fmt.Sprintf("%sJoinMatchmakeSessionBehavior: %d,\n", indentationValues, joinMatchmakeSessionParam.JoinMatchmakeSessionBehavior))
	b.WriteString(fmt.Sprintf("%sStrUserPassword: %q,\n", indentationValues, joinMatchmakeSessionParam.StrUserPassword))
	b.WriteString(fmt.Sprintf("%sStrSystemPassword: %q,\n", indentationValues, joinMatchmakeSessionParam.StrSystemPassword))
	b.WriteString(fmt.Sprintf("%sJoinMessage: %q,\n", indentationValues, joinMatchmakeSessionParam.JoinMessage))
	b.WriteString(fmt.Sprintf("%sParticipationCount: %d,\n", indentationValues, joinMatchmakeSessionParam.ParticipationCount))
	b.WriteString(fmt.Sprintf("%sExtraParticipants: %d,\n", indentationValues, joinMatchmakeSessionParam.ExtraParticipants))

	if joinMatchmakeSessionParam.BlockListParam != nil {
		b.WriteString(fmt.Sprintf("%sBlockListParam: %s\n", indentationValues, joinMatchmakeSessionParam.BlockListParam.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sBlockListParam: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewJoinMatchmakeSessionParam returns a new JoinMatchmakeSessionParam
func NewJoinMatchmakeSessionParam() *JoinMatchmakeSessionParam {
	return &JoinMatchmakeSessionParam{}
}
