// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// UpdateMatchmakeSessionParam holds parameters for a matchmake session
type UpdateMatchmakeSessionParam struct {
	types.Structure
	GID                 *types.PrimitiveU32
	ModificationFlag    *types.PrimitiveU32
	Attributes          *types.List[*types.PrimitiveU32]
	OpenParticipation   *types.PrimitiveBool
	ApplicationBuffer   []byte
	ProgressScore       *types.PrimitiveU8
	MatchmakeParam      *MatchmakeParam
	StartedTime         *types.DateTime
	UserPassword        string
	GameMode            *types.PrimitiveU32
	Description         string
	MinParticipants     *types.PrimitiveU16
	MaxParticipants     *types.PrimitiveU16
	MatchmakeSystemType *types.PrimitiveU32
	ParticipationPolicy *types.PrimitiveU32
	PolicyArgument      *types.PrimitiveU32
	Codeword            string
}

// ExtractFrom extracts the UpdateMatchmakeSessionParam from the given readable
func (updateMatchmakeSessionParam *UpdateMatchmakeSessionParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = updateMatchmakeSessionParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read UpdateMatchmakeSessionParam header. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.GID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.GID. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.ModificationFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.ModificationFlag. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.Attributes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.Attributes. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.OpenParticipation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.OpenParticipation. %s", err.Error())
	}

	updateMatchmakeSessionParam.ApplicationBuffer, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.ApplicationBuffer. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.ProgressScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.ProgressScore. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.MatchmakeParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.MatchmakeParam. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.StartedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.StartedTime. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.UserPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.UserPassword. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.GameMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.GameMode. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.Description.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.Description. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.MinParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.MinParticipants. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.MaxParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.MaxParticipants. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.MatchmakeSystemType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.MatchmakeSystemType. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.ParticipationPolicy.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.ParticipationPolicy. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.PolicyArgument.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.PolicyArgument. %s", err.Error())
	}

	err = updateMatchmakeSessionParam.Codeword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.Codeword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UpdateMatchmakeSessionParam
func (updateMatchmakeSessionParam *UpdateMatchmakeSessionParam) Copy() types.RVType {
	copied := NewUpdateMatchmakeSessionParam()

	copied.StructureVersion = updateMatchmakeSessionParam.StructureVersion

	copied.GID = updateMatchmakeSessionParam.GID
	copied.ModificationFlag = updateMatchmakeSessionParam.ModificationFlag
	copied.Attributes = make(*types.List[*types.PrimitiveU32], len(updateMatchmakeSessionParam.Attributes))

	copy(copied.Attributes, updateMatchmakeSessionParam.Attributes)

	copied.OpenParticipation = updateMatchmakeSessionParam.OpenParticipation
	copied.ApplicationBuffer = make([]byte, len(updateMatchmakeSessionParam.ApplicationBuffer))

	copy(copied.ApplicationBuffer, updateMatchmakeSessionParam.ApplicationBuffer)

	copied.ProgressScore = updateMatchmakeSessionParam.ProgressScore

	copied.MatchmakeParam = updateMatchmakeSessionParam.MatchmakeParam.Copy().(*MatchmakeParam)

	copied.StartedTime = updateMatchmakeSessionParam.StartedTime.Copy()

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
func (updateMatchmakeSessionParam *UpdateMatchmakeSessionParam) Equals(o types.RVType) bool {
	if _, ok := o.(*UpdateMatchmakeSessionParam); !ok {
		return false
	}

	other := o.(*UpdateMatchmakeSessionParam)

	if updateMatchmakeSessionParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !updateMatchmakeSessionParam.GID.Equals(other.GID) {
		return false
	}

	if !updateMatchmakeSessionParam.ModificationFlag.Equals(other.ModificationFlag) {
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

	if !updateMatchmakeSessionParam.OpenParticipation.Equals(other.OpenParticipation) {
		return false
	}

	if !updateMatchmakeSessionParam.ApplicationBuffer.Equals(other.ApplicationBuffer) {
		return false
	}

	if !updateMatchmakeSessionParam.ProgressScore.Equals(other.ProgressScore) {
		return false
	}

	if updateMatchmakeSessionParam.MatchmakeParam != nil && other.MatchmakeParam != nil {
		if updateMatchmakeSessionParam.MatchmakeParam.Equals(other.MatchmakeParam) {
			return false
		}
	}

	if !updateMatchmakeSessionParam.StartedTime.Equals(other.StartedTime) {
		return false
	}

	if !updateMatchmakeSessionParam.UserPassword.Equals(other.UserPassword) {
		return false
	}

	if !updateMatchmakeSessionParam.GameMode.Equals(other.GameMode) {
		return false
	}

	if !updateMatchmakeSessionParam.Description.Equals(other.Description) {
		return false
	}

	if !updateMatchmakeSessionParam.MinParticipants.Equals(other.MinParticipants) {
		return false
	}

	if !updateMatchmakeSessionParam.MaxParticipants.Equals(other.MaxParticipants) {
		return false
	}

	if !updateMatchmakeSessionParam.MatchmakeSystemType.Equals(other.MatchmakeSystemType) {
		return false
	}

	if !updateMatchmakeSessionParam.ParticipationPolicy.Equals(other.ParticipationPolicy) {
		return false
	}

	if !updateMatchmakeSessionParam.PolicyArgument.Equals(other.PolicyArgument) {
		return false
	}

	if !updateMatchmakeSessionParam.Codeword.Equals(other.Codeword) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, updateMatchmakeSessionParam.StructureVersion))
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
