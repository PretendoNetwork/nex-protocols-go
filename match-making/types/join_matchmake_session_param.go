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

// JoinMatchmakeSessionParam holds parameters for a matchmake session
type JoinMatchmakeSessionParam struct {
	types.Structure
	GID                          *types.PrimitiveU32
	AdditionalParticipants       *types.List[*types.PID]
	GIDForParticipationCheck     *types.PrimitiveU32
	JoinMatchmakeSessionOption   *types.PrimitiveU32
	JoinMatchmakeSessionBehavior *types.PrimitiveU8
	StrUserPassword              string
	StrSystemPassword            string
	JoinMessage                  string
	ParticipationCount           *types.PrimitiveU16
	ExtraParticipants            *types.PrimitiveU16
	BlockListParam               *MatchmakeBlockListParam // * NEX 4.0+ ? Not seen in Minecraft, which is 3.10.0
}

// ExtractFrom extracts the JoinMatchmakeSessionParam from the given readable
func (joinMatchmakeSessionParam *JoinMatchmakeSessionParam) ExtractFrom(readable types.Readable) error {
	matchmakingVersion := stream.Server.MatchMakingProtocolVersion()

	var err error

	err = joinMatchmakeSessionParam.GID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.GID. %s", err.Error())
	}

	err = joinMatchmakeSessionParam.AdditionalParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.AdditionalParticipants. %s", err.Error())
	}

	err = joinMatchmakeSessionParam.GIDForParticipationCheck.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.GIDForParticipationCheck. %s", err.Error())
	}

	err = joinMatchmakeSessionParam.JoinMatchmakeSessionOption.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.JoinMatchmakeSessionOption. %s", err.Error())
	}

	err = joinMatchmakeSessionParam.JoinMatchmakeSessionBehavior.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.JoinMatchmakeSessionBehavior. %s", err.Error())
	}

	err = joinMatchmakeSessionParam.StrUserPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.StrUserPassword. %s", err.Error())
	}

	err = joinMatchmakeSessionParam.StrSystemPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.StrSystemPassword. %s", err.Error())
	}

	err = joinMatchmakeSessionParam.JoinMessage.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.JoinMessage. %s", err.Error())
	}

	err = joinMatchmakeSessionParam.ParticipationCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.ParticipationCount. %s", err.Error())
	}

	// * From Dani:
	// * - "Just for future reference, Minecraft has structure version 1 on JoinMatchmakeSessionParam"
	// * These fields COULD be different structure versions, not related to NEX updates
	// TODO - Needs more research

	// * Assuming this to be 3.10.0
	// * Not seen in Terraria, which is 3.8.3
	if matchmakingVersion.GreaterOrEqual("3.10.0") {
	err = 	joinMatchmakeSessionParam.ExtraParticipants.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.ExtraParticipants. %s", err.Error())
		}
	}

	// * Assuming this to be 4.0.0
	// * Not seen in Minecraft, which is 3.10.0
	if matchmakingVersion.GreaterOrEqual("4.0.0") {
		blockListParam, err := nex.StreamReadStructure(stream, NewMatchmakeBlockListParam())
		if err != nil {
			return fmt.Errorf("Failed to extract JoinMatchmakeSessionParam.BlockListParam. %s", err.Error())
		}

		joinMatchmakeSessionParam.BlockListParam = blockListParam
	}

	return nil
}

// Copy returns a new copied instance of JoinMatchmakeSessionParam
func (joinMatchmakeSessionParam *JoinMatchmakeSessionParam) Copy() types.RVType {
	copied := NewJoinMatchmakeSessionParam()

	copied.StructureVersion = joinMatchmakeSessionParam.StructureVersion

	copied.GID = joinMatchmakeSessionParam.GID
	copied.AdditionalParticipants = make(*types.List[*types.PID], len(joinMatchmakeSessionParam.AdditionalParticipants))

	for i := 0; i < len(joinMatchmakeSessionParam.AdditionalParticipants); i++ {
		copied.AdditionalParticipants[i] = joinMatchmakeSessionParam.AdditionalParticipants[i].Copy()
	}

	copied.GIDForParticipationCheck = joinMatchmakeSessionParam.GIDForParticipationCheck
	copied.JoinMatchmakeSessionOption = joinMatchmakeSessionParam.JoinMatchmakeSessionOption
	copied.JoinMatchmakeSessionBehavior = joinMatchmakeSessionParam.JoinMatchmakeSessionBehavior
	copied.StrUserPassword = joinMatchmakeSessionParam.StrUserPassword
	copied.StrSystemPassword = joinMatchmakeSessionParam.StrSystemPassword
	copied.JoinMessage = joinMatchmakeSessionParam.JoinMessage
	copied.ParticipationCount = joinMatchmakeSessionParam.ParticipationCount
	copied.ExtraParticipants = joinMatchmakeSessionParam.ExtraParticipants

	copied.BlockListParam = joinMatchmakeSessionParam.BlockListParam.Copy().(*MatchmakeBlockListParam)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (joinMatchmakeSessionParam *JoinMatchmakeSessionParam) Equals(o types.RVType) bool {
	if _, ok := o.(*JoinMatchmakeSessionParam); !ok {
		return false
	}

	other := o.(*JoinMatchmakeSessionParam)

	if joinMatchmakeSessionParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !joinMatchmakeSessionParam.GID.Equals(other.GID) {
		return false
	}

	if len(joinMatchmakeSessionParam.AdditionalParticipants) != len(other.AdditionalParticipants) {
		return false
	}

	for i := 0; i < len(joinMatchmakeSessionParam.AdditionalParticipants); i++ {
		if !joinMatchmakeSessionParam.AdditionalParticipants[i].Equals(other.AdditionalParticipants[i]) {
			return false
		}
	}

	if !joinMatchmakeSessionParam.GIDForParticipationCheck.Equals(other.GIDForParticipationCheck) {
		return false
	}

	if !joinMatchmakeSessionParam.JoinMatchmakeSessionOption.Equals(other.JoinMatchmakeSessionOption) {
		return false
	}

	if !joinMatchmakeSessionParam.JoinMatchmakeSessionBehavior.Equals(other.JoinMatchmakeSessionBehavior) {
		return false
	}

	if !joinMatchmakeSessionParam.StrUserPassword.Equals(other.StrUserPassword) {
		return false
	}

	if !joinMatchmakeSessionParam.StrSystemPassword.Equals(other.StrSystemPassword) {
		return false
	}

	if !joinMatchmakeSessionParam.JoinMessage.Equals(other.JoinMessage) {
		return false
	}

	if !joinMatchmakeSessionParam.ParticipationCount.Equals(other.ParticipationCount) {
		return false
	}

	if !joinMatchmakeSessionParam.ExtraParticipants.Equals(other.ExtraParticipants) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, joinMatchmakeSessionParam.StructureVersion))
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
