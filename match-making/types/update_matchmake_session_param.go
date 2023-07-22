// Package match_making_types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package match_making_types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// UpdateMatchmakeSessionParam holds parameters for a matchmake session
type UpdateMatchmakeSessionParam struct {
	nex.Structure
	GID                 uint32
	ModificationFlag    uint32
	Attributes          []uint32
	OpenParticipation   bool
	ApplicationBuffer   []byte
	ProgressScore       uint8
	MatchmakeParam      *MatchmakeParam
	StartedTime         *nex.DateTime
	UserPassword        string
	GameMode            uint32
	Description         string
	MinParticipants     uint16
	MaxParticipants     uint16
	MatchmakeSystemType uint32
	ParticipationPolicy uint32
	PolicyArgument      uint32
	Codeword            string
}

// ExtractFromStream extracts a UpdateMatchmakeSessionParam structure from a stream
func (updateMatchmakeSessionParam *UpdateMatchmakeSessionParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	updateMatchmakeSessionParam.GID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.GID. %s", err.Error())
	}

	updateMatchmakeSessionParam.ModificationFlag, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.ModificationFlag. %s", err.Error())
	}

	updateMatchmakeSessionParam.Attributes, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.Attributes. %s", err.Error())
	}

	updateMatchmakeSessionParam.OpenParticipation, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.OpenParticipation. %s", err.Error())
	}

	updateMatchmakeSessionParam.ApplicationBuffer, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.ApplicationBuffer. %s", err.Error())
	}

	updateMatchmakeSessionParam.ProgressScore, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.ProgressScore. %s", err.Error())
	}

	matchmakeParam, err := stream.ReadStructure(NewMatchmakeParam())
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.MatchmakeParam. %s", err.Error())
	}

	updateMatchmakeSessionParam.MatchmakeParam = matchmakeParam.(*MatchmakeParam)
	updateMatchmakeSessionParam.StartedTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.StartedTime. %s", err.Error())
	}

	updateMatchmakeSessionParam.UserPassword, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.UserPassword. %s", err.Error())
	}

	updateMatchmakeSessionParam.GameMode, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.GameMode. %s", err.Error())
	}

	updateMatchmakeSessionParam.Description, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.Description. %s", err.Error())
	}

	updateMatchmakeSessionParam.MinParticipants, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.MinParticipants. %s", err.Error())
	}

	updateMatchmakeSessionParam.MaxParticipants, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.MaxParticipants. %s", err.Error())
	}

	updateMatchmakeSessionParam.MatchmakeSystemType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.MatchmakeSystemType. %s", err.Error())
	}

	updateMatchmakeSessionParam.ParticipationPolicy, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.ParticipationPolicy. %s", err.Error())
	}

	updateMatchmakeSessionParam.PolicyArgument, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.PolicyArgument. %s", err.Error())
	}

	updateMatchmakeSessionParam.Codeword, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.Codeword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UpdateMatchmakeSessionParam
func (updateMatchmakeSessionParam *UpdateMatchmakeSessionParam) Copy() nex.StructureInterface {
	copied := NewUpdateMatchmakeSessionParam()

	copied.GID = updateMatchmakeSessionParam.GID
	copied.ModificationFlag = updateMatchmakeSessionParam.ModificationFlag
	copied.Attributes = make([]uint32, len(updateMatchmakeSessionParam.Attributes))

	copy(copied.Attributes, updateMatchmakeSessionParam.Attributes)

	copied.OpenParticipation = updateMatchmakeSessionParam.OpenParticipation
	copied.ApplicationBuffer = make([]byte, len(updateMatchmakeSessionParam.ApplicationBuffer))

	copy(copied.ApplicationBuffer, updateMatchmakeSessionParam.ApplicationBuffer)

	copied.ProgressScore = updateMatchmakeSessionParam.ProgressScore

	if updateMatchmakeSessionParam.MatchmakeParam != nil {
		copied.MatchmakeParam = updateMatchmakeSessionParam.MatchmakeParam.Copy().(*MatchmakeParam)
	}

	if updateMatchmakeSessionParam.StartedTime != nil {
		copied.StartedTime = updateMatchmakeSessionParam.StartedTime.Copy()
	}

	copied.UserPassword = updateMatchmakeSessionParam.UserPassword
	copied.GameMode = updateMatchmakeSessionParam.GameMode
	copied.Description = updateMatchmakeSessionParam.Description
	copied.MinParticipants = updateMatchmakeSessionParam.MinParticipants
	copied.MaxParticipants = updateMatchmakeSessionParam.MaxParticipants
	copied.MatchmakeSystemType = updateMatchmakeSessionParam.MatchmakeSystemType
	copied.ParticipationPolicy = updateMatchmakeSessionParam.ParticipationPolicy
	copied.PolicyArgument = updateMatchmakeSessionParam.PolicyArgument
	copied.Codeword = updateMatchmakeSessionParam.Codeword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (updateMatchmakeSessionParam *UpdateMatchmakeSessionParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*UpdateMatchmakeSessionParam)

	if updateMatchmakeSessionParam.GID != other.GID {
		return false
	}

	if updateMatchmakeSessionParam.ModificationFlag != other.ModificationFlag {
		return false
	}

	if len(updateMatchmakeSessionParam.Attributes) != len(other.Attributes) {
		return false
	}

	for i := 0; i < len(updateMatchmakeSessionParam.Attributes); i++ {
		if updateMatchmakeSessionParam.Attributes[i] != other.Attributes[i] {
			return false
		}
	}

	if updateMatchmakeSessionParam.OpenParticipation != other.OpenParticipation {
		return false
	}

	if !bytes.Equal(updateMatchmakeSessionParam.ApplicationBuffer, other.ApplicationBuffer) {
		return false
	}

	if updateMatchmakeSessionParam.ProgressScore != other.ProgressScore {
		return false
	}

	if updateMatchmakeSessionParam.MatchmakeParam != nil && other.MatchmakeParam == nil {
		return false
	}

	if updateMatchmakeSessionParam.MatchmakeParam == nil && other.MatchmakeParam != nil {
		return false
	}

	if updateMatchmakeSessionParam.MatchmakeParam != nil && other.MatchmakeParam != nil {
		if updateMatchmakeSessionParam.MatchmakeParam.Equals(other.MatchmakeParam) {
			return false
		}
	}

	if updateMatchmakeSessionParam.StartedTime != other.StartedTime {
		return false
	}

	if updateMatchmakeSessionParam.UserPassword != other.UserPassword {
		return false
	}

	if updateMatchmakeSessionParam.GameMode != other.GameMode {
		return false
	}

	if updateMatchmakeSessionParam.Description != other.Description {
		return false
	}

	if updateMatchmakeSessionParam.MinParticipants != other.MinParticipants {
		return false
	}

	if updateMatchmakeSessionParam.MaxParticipants != other.MaxParticipants {
		return false
	}

	if updateMatchmakeSessionParam.MatchmakeSystemType != other.MatchmakeSystemType {
		return false
	}

	if updateMatchmakeSessionParam.ParticipationPolicy != other.ParticipationPolicy {
		return false
	}

	if updateMatchmakeSessionParam.PolicyArgument != other.PolicyArgument {
		return false
	}

	if updateMatchmakeSessionParam.Codeword != other.Codeword {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (updateMatchmakeSessionParam *UpdateMatchmakeSessionParam) String() string {
	return updateMatchmakeSessionParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (updateMatchmakeSessionParam *UpdateMatchmakeSessionParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UpdateMatchmakeSessionParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, updateMatchmakeSessionParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sGID: %d,\n", indentationValues, updateMatchmakeSessionParam.GID))
	b.WriteString(fmt.Sprintf("%sModificationFlag: %d,\n", indentationValues, updateMatchmakeSessionParam.ModificationFlag))
	b.WriteString(fmt.Sprintf("%sAttributes: %v,\n", indentationValues, updateMatchmakeSessionParam.Attributes))
	b.WriteString(fmt.Sprintf("%sOpenParticipation: %t,\n", indentationValues, updateMatchmakeSessionParam.OpenParticipation))
	b.WriteString(fmt.Sprintf("%sApplicationBuffer: %x,\n", indentationValues, updateMatchmakeSessionParam.ApplicationBuffer))
	b.WriteString(fmt.Sprintf("%sProgressScore: %d,\n", indentationValues, updateMatchmakeSessionParam.ProgressScore))

	if updateMatchmakeSessionParam.MatchmakeParam != nil {
		b.WriteString(fmt.Sprintf("%sMatchmakeParam: %s,\n", indentationValues, updateMatchmakeSessionParam.MatchmakeParam.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sMatchmakeParam: nil,\n", indentationValues))
	}

	if updateMatchmakeSessionParam.StartedTime != nil {
		b.WriteString(fmt.Sprintf("%sStartedTime: %s,\n", indentationValues, updateMatchmakeSessionParam.StartedTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sStartedTime: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sUserPassword: %q,\n", indentationValues, updateMatchmakeSessionParam.UserPassword))
	b.WriteString(fmt.Sprintf("%sGameMode: %d,\n", indentationValues, updateMatchmakeSessionParam.GameMode))
	b.WriteString(fmt.Sprintf("%sDescription: %q,\n", indentationValues, updateMatchmakeSessionParam.Description))
	b.WriteString(fmt.Sprintf("%sMinParticipants: %d,\n", indentationValues, updateMatchmakeSessionParam.MinParticipants))
	b.WriteString(fmt.Sprintf("%sMaxParticipants: %d,\n", indentationValues, updateMatchmakeSessionParam.MaxParticipants))
	b.WriteString(fmt.Sprintf("%sMatchmakeSystemType: %d,\n", indentationValues, updateMatchmakeSessionParam.MatchmakeSystemType))
	b.WriteString(fmt.Sprintf("%sParticipationPolicy: %d,\n", indentationValues, updateMatchmakeSessionParam.ParticipationPolicy))
	b.WriteString(fmt.Sprintf("%sPolicyArgument: %d,\n", indentationValues, updateMatchmakeSessionParam.PolicyArgument))
	b.WriteString(fmt.Sprintf("%sCodeword: %q\n", indentationValues, updateMatchmakeSessionParam.Codeword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewUpdateMatchmakeSessionParam returns a new UpdateMatchmakeSessionParam
func NewUpdateMatchmakeSessionParam() *UpdateMatchmakeSessionParam {
	return &UpdateMatchmakeSessionParam{}
}
