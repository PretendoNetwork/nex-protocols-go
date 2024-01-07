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

// MatchmakeSession holds information about a matchmake session
type MatchmakeSession struct {
	types.Structure
	*Gathering
	GameMode              *types.PrimitiveU32
	Attributes            *types.List[*types.PrimitiveU32]
	OpenParticipation     *types.PrimitiveBool
	MatchmakeSystemType   *types.PrimitiveU32
	ApplicationData       []byte
	ParticipationCount    *types.PrimitiveU32
	ProgressScore         *types.PrimitiveU8           // NEX v3.4.0+
	SessionKey            []byte          // NEX v3.0.0+
	Option                *types.PrimitiveU32          // NEX v3.5.0+
	MatchmakeParam        *MatchmakeParam // NEX v3.6.0+
	StartedTime           *types.DateTime   // NEX v3.6.0+
	UserPassword          string          // NEX v3.7.0+
	ReferGID              *types.PrimitiveU32          // NEX v3.8.0+
	UserPasswordEnabled   *types.PrimitiveBool            // NEX v3.8.0+
	SystemPasswordEnabled *types.PrimitiveBool            // NEX v3.8.0+
	CodeWord              string          // NEX v4.0.0+
}

// ExtractFrom extracts the MatchmakeSession from the given readable
func (matchmakeSession *MatchmakeSession) ExtractFrom(readable types.Readable) error {
	matchmakingVersion := stream.Server.MatchMakingProtocolVersion()

	var err error

	err = matchmakeSession.GameMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.GameMode. %s", err.Error())
	}

	err = matchmakeSession.Attributes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.Attributes. %s", err.Error())
	}

	err = matchmakeSession.OpenParticipation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.OpenParticipation. %s", err.Error())
	}

	err = matchmakeSession.MatchmakeSystemType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.MatchmakeSystemType. %s", err.Error())
	}

	matchmakeSession.ApplicationData, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.ApplicationData. %s", err.Error())
	}

	err = matchmakeSession.ParticipationCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.ParticipationCount. %s", err.Error())
	}

	if matchmakingVersion.GreaterOrEqual("3.4.0") {
	err = 	matchmakeSession.ProgressScore.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.ProgressScore. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("3.0.0") {
		matchmakeSession.SessionKey, err = stream.ReadBuffer()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.SessionKey. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("3.5.0") {
	err = 	matchmakeSession.Option.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.Option. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("3.6.0") {
		matchmakeParam, err := nex.StreamReadStructure(stream, NewMatchmakeParam())
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.MatchmakeParam. %s", err.Error())
		}

		matchmakeSession.MatchmakeParam = matchmakeParam
	err = 	matchmakeSession.StartedTime.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.StartedTime. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("3.7.0") {
	err = 	matchmakeSession.UserPassword.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.UserPassword. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("3.8.0") {
	err = 	matchmakeSession.ReferGID.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.ReferGID. %s", err.Error())
		}

	err = 	matchmakeSession.UserPasswordEnabled.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.UserPasswordEnabled. %s", err.Error())
		}

	err = 	matchmakeSession.SystemPasswordEnabled.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.SystemPasswordEnabled. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("4.0.0") {
	err = 	matchmakeSession.CodeWord.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.CodeWord. %s", err.Error())
		}
	}

	return nil
}

// Bytes extracts a MatchmakeSession structure from a stream
func (matchmakeSession *MatchmakeSession) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	matchmakingVersion := stream.Server.MatchMakingProtocolVersion()

	matchmakeSession.GameMode.WriteTo(contentWritable)
	matchmakeSession.Attributes.WriteTo(contentWritable)
	matchmakeSession.OpenParticipation.WriteTo(contentWritable)
	matchmakeSession.MatchmakeSystemType.WriteTo(contentWritable)
	stream.WriteBuffer(matchmakeSession.ApplicationData)

	matchmakeSession.ParticipationCount.WriteTo(contentWritable)

	if matchmakingVersion.GreaterOrEqual("3.4.0") {
		matchmakeSession.ProgressScore.WriteTo(contentWritable)
	}

	if matchmakingVersion.GreaterOrEqual("3.0.0") {
		stream.WriteBuffer(matchmakeSession.SessionKey)
	}

	if matchmakingVersion.GreaterOrEqual("3.5.0") {
		matchmakeSession.Option.WriteTo(contentWritable)
	}

	if matchmakingVersion.GreaterOrEqual("3.6.0") {
		matchmakeSession.MatchmakeParam.WriteTo(contentWritable)
		matchmakeSession.StartedTime.WriteTo(contentWritable)
	}

	if matchmakingVersion.GreaterOrEqual("3.7.0") {
		matchmakeSession.UserPassword.WriteTo(contentWritable)
	}

	if matchmakingVersion.GreaterOrEqual("3.7.0") {
		matchmakeSession.ReferGID.WriteTo(contentWritable)
		matchmakeSession.UserPasswordEnabled.WriteTo(contentWritable)
		matchmakeSession.SystemPasswordEnabled.WriteTo(contentWritable)
	}

	if matchmakingVersion.GreaterOrEqual("4.0.0") {
		matchmakeSession.CodeWord.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of MatchmakeSession
func (matchmakeSession *MatchmakeSession) Copy() types.RVType {
	copied := NewMatchmakeSession()

	copied.StructureVersion = matchmakeSession.StructureVersion

	copied.Gathering = matchmakeSession.Gathering.Copy().(*Gathering)
	copied.GameMode = matchmakeSession.GameMode
	copied.Attributes = make(*types.List[*types.PrimitiveU32], len(matchmakeSession.Attributes))

	copy(copied.Attributes, matchmakeSession.Attributes)

	copied.OpenParticipation = matchmakeSession.OpenParticipation
	copied.MatchmakeSystemType = matchmakeSession.MatchmakeSystemType
	copied.ApplicationData = make([]byte, len(matchmakeSession.ApplicationData))

	copy(copied.ApplicationData, matchmakeSession.ApplicationData)

	copied.ParticipationCount = matchmakeSession.ParticipationCount
	copied.ProgressScore = matchmakeSession.ProgressScore
	copied.SessionKey = make([]byte, len(matchmakeSession.SessionKey))

	copy(copied.SessionKey, matchmakeSession.SessionKey)

	copied.Option = matchmakeSession.Option

	copied.MatchmakeParam = matchmakeSession.MatchmakeParam.Copy().(*MatchmakeParam)

	copied.StartedTime = matchmakeSession.StartedTime.Copy()

	copied.UserPassword = matchmakeSession.UserPassword
	copied.ReferGID = matchmakeSession.ReferGID
	copied.UserPasswordEnabled = matchmakeSession.UserPasswordEnabled
	copied.SystemPasswordEnabled = matchmakeSession.SystemPasswordEnabled
	copied.CodeWord = matchmakeSession.CodeWord

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeSession *MatchmakeSession) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeSession); !ok {
		return false
	}

	other := o.(*MatchmakeSession)

	if matchmakeSession.StructureVersion != other.StructureVersion {
		return false
	}

	if !matchmakeSession.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !matchmakeSession.GameMode.Equals(other.GameMode) {
		return false
	}

	if len(matchmakeSession.Attributes) != len(other.Attributes) {
		return false
	}

	for i := 0; i < len(matchmakeSession.Attributes); i++ {
		if matchmakeSession.Attributes[i] != other.Attributes[i] {
			return false
		}
	}

	if !matchmakeSession.OpenParticipation.Equals(other.OpenParticipation) {
		return false
	}

	if !matchmakeSession.MatchmakeSystemType.Equals(other.MatchmakeSystemType) {
		return false
	}

	if !matchmakeSession.ApplicationData.Equals(other.ApplicationData) {
		return false
	}

	if !matchmakeSession.ParticipationCount.Equals(other.ParticipationCount) {
		return false
	}

	if !matchmakeSession.ProgressScore.Equals(other.ProgressScore) {
		return false
	}

	if !matchmakeSession.SessionKey.Equals(other.SessionKey) {
		return false
	}

	if !matchmakeSession.Option.Equals(other.Option) {
		return false
	}

	if !matchmakeSession.MatchmakeParam.Equals(other.MatchmakeParam) {
		return false
	}

	if !matchmakeSession.StartedTime.Equals(other.StartedTime) {
		return false
	}

	if !matchmakeSession.UserPassword.Equals(other.UserPassword) {
		return false
	}

	if !matchmakeSession.ReferGID.Equals(other.ReferGID) {
		return false
	}

	if !matchmakeSession.UserPasswordEnabled.Equals(other.UserPasswordEnabled) {
		return false
	}

	if !matchmakeSession.SystemPasswordEnabled.Equals(other.SystemPasswordEnabled) {
		return false
	}

	if !matchmakeSession.CodeWord.Equals(other.CodeWord) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (matchmakeSession *MatchmakeSession) String() string {
	return matchmakeSession.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (matchmakeSession *MatchmakeSession) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeSession{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, matchmakeSession.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, matchmakeSession.StructureVersion))
	b.WriteString(fmt.Sprintf("%sGameMode: %d,\n", indentationValues, matchmakeSession.GameMode))
	b.WriteString(fmt.Sprintf("%sAttributes: %v,\n", indentationValues, matchmakeSession.Attributes))
	b.WriteString(fmt.Sprintf("%sOpenParticipation: %t,\n", indentationValues, matchmakeSession.OpenParticipation))
	b.WriteString(fmt.Sprintf("%sMatchmakeSystemType: %d,\n", indentationValues, matchmakeSession.MatchmakeSystemType))
	b.WriteString(fmt.Sprintf("%sApplicationData: %x,\n", indentationValues, matchmakeSession.ApplicationData))
	b.WriteString(fmt.Sprintf("%sParticipationCount: %d,\n", indentationValues, matchmakeSession.ParticipationCount))
	b.WriteString(fmt.Sprintf("%sProgressScore: %d,\n", indentationValues, matchmakeSession.ProgressScore))
	b.WriteString(fmt.Sprintf("%sSessionKey: %x,\n", indentationValues, matchmakeSession.SessionKey))
	b.WriteString(fmt.Sprintf("%sOption: %d,\n", indentationValues, matchmakeSession.Option))

	if matchmakeSession.MatchmakeParam != nil {
		b.WriteString(fmt.Sprintf("%sMatchmakeParam: %s,\n", indentationValues, matchmakeSession.MatchmakeParam.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sMatchmakeParam: nil,\n", indentationValues))
	}

	if matchmakeSession.StartedTime != nil {
		b.WriteString(fmt.Sprintf("%sStartedTime: %s,\n", indentationValues, matchmakeSession.StartedTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sStartedTime: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sUserPassword: %q,\n", indentationValues, matchmakeSession.UserPassword))
	b.WriteString(fmt.Sprintf("%sReferGID: %d,\n", indentationValues, matchmakeSession.ReferGID))
	b.WriteString(fmt.Sprintf("%sUserPasswordEnabled: %t,\n", indentationValues, matchmakeSession.UserPasswordEnabled))
	b.WriteString(fmt.Sprintf("%sSystemPasswordEnabled: %t,\n", indentationValues, matchmakeSession.SystemPasswordEnabled))
	b.WriteString(fmt.Sprintf("%sCodeWord: %q\n", indentationValues, matchmakeSession.CodeWord))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeSession returns a new MatchmakeSession
func NewMatchmakeSession() *MatchmakeSession {
	matchmakeSession := &MatchmakeSession{}
	matchmakeSession.Gathering = NewGathering()
	matchmakeSession.SetParentType(matchmakeSession.Gathering)

	return matchmakeSession
}
