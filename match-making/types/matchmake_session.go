// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/match-making/constants"
)

// MatchmakeSession is a type within the Matchmaking protocol
type MatchmakeSession struct {
	types.Structure
	Gathering
	GameMode              types.UInt32
	Attributes            types.List[types.UInt32]
	OpenParticipation     types.Bool
	MatchmakeSystemType   constants.MatchmakeSystemType
	ApplicationBuffer     types.Buffer
	ParticipationCount    types.UInt32
	ProgressScore         types.UInt8                       // * NEX v3.4.0
	SessionKey            types.Buffer                      // * NEX v3.0.0
	Option0               constants.MatchmakeSessionOption0 // * NEX v3.5.0
	MatchmakeParam        MatchmakeParam                    // * NEX v3.6.0
	StartedTime           types.DateTime                    // * NEX v3.6.0
	UserPassword          types.String                      // * NEX v3.7.0
	ReferGID              types.UInt32                      // * NEX v3.8.0
	UserPasswordEnabled   types.Bool                        // * NEX v3.8.0
	SystemPasswordEnabled types.Bool                        // * NEX v3.8.0
	CodeWord              types.String                      // * NEX v4.0.0
}

// ObjectID returns the object identifier of the type
func (ms MatchmakeSession) ObjectID() types.RVType {
	return ms.GatheringObjectID()
}

// DataObjectID returns the object identifier of the type embedding Gathering
func (ms MatchmakeSession) GatheringObjectID() types.RVType {
	return types.NewString("MatchmakeSession")
}

// WriteTo writes the MatchmakeSession to the given writable
func (ms MatchmakeSession) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.MatchMaking

	ms.Gathering.WriteTo(writable)

	contentWritable := writable.CopyNew()

	ms.GameMode.WriteTo(contentWritable)
	ms.Attributes.WriteTo(contentWritable)
	ms.OpenParticipation.WriteTo(contentWritable)
	ms.MatchmakeSystemType.WriteTo(contentWritable)
	ms.ApplicationBuffer.WriteTo(contentWritable)
	ms.ParticipationCount.WriteTo(contentWritable)

	if libraryVersion.GreaterOrEqual("3.4.0") {
		ms.ProgressScore.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.0.0") {
		ms.SessionKey.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.5.0") {
		ms.Option0.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.6.0") {
		ms.MatchmakeParam.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.6.0") {
		ms.StartedTime.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.7.0") {
		ms.UserPassword.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.8.0") {
		ms.ReferGID.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.8.0") {
		ms.UserPasswordEnabled.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("3.8.0") {
		ms.SystemPasswordEnabled.WriteTo(contentWritable)
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		ms.CodeWord.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	ms.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeSession from the given readable
func (ms *MatchmakeSession) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.MatchMaking

	var err error

	err = ms.Gathering.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.Gathering. %s", err.Error())
	}

	err = ms.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession header. %s", err.Error())
	}

	err = ms.GameMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.GameMode. %s", err.Error())
	}

	err = ms.Attributes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.Attributes. %s", err.Error())
	}

	err = ms.OpenParticipation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.OpenParticipation. %s", err.Error())
	}

	err = ms.MatchmakeSystemType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.MatchmakeSystemType. %s", err.Error())
	}

	err = ms.ApplicationBuffer.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.ApplicationBuffer. %s", err.Error())
	}

	err = ms.ParticipationCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.ParticipationCount. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("3.4.0") {
		err = ms.ProgressScore.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.ProgressScore. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.0.0") {
		err = ms.SessionKey.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.SessionKey. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.5.0") {
		err = ms.Option0.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.Option0. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.6.0") {
		err = ms.MatchmakeParam.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.MatchmakeParam. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.6.0") {
		err = ms.StartedTime.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.StartedTime. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.7.0") {
		err = ms.UserPassword.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.UserPassword. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.8.0") {
		err = ms.ReferGID.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.ReferGID. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.8.0") {
		err = ms.UserPasswordEnabled.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.UserPasswordEnabled. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("3.8.0") {
		err = ms.SystemPasswordEnabled.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.SystemPasswordEnabled. %s", err.Error())
		}
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		err = ms.CodeWord.ExtractFrom(readable)
		if err = ms.CodeWord.ExtractFrom(readable); err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.CodeWord. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeSession
func (ms MatchmakeSession) Copy() types.RVType {
	copied := NewMatchmakeSession()

	copied.StructureVersion = ms.StructureVersion
	copied.Gathering = ms.Gathering.Copy().(Gathering)
	copied.GameMode = ms.GameMode.Copy().(types.UInt32)
	copied.Attributes = ms.Attributes.Copy().(types.List[types.UInt32])
	copied.OpenParticipation = ms.OpenParticipation.Copy().(types.Bool)
	copied.MatchmakeSystemType = ms.MatchmakeSystemType
	copied.ApplicationBuffer = ms.ApplicationBuffer.Copy().(types.Buffer)
	copied.ParticipationCount = ms.ParticipationCount.Copy().(types.UInt32)
	copied.ProgressScore = ms.ProgressScore.Copy().(types.UInt8)
	copied.SessionKey = ms.SessionKey.Copy().(types.Buffer)
	copied.Option0 = ms.Option0
	copied.MatchmakeParam = ms.MatchmakeParam.Copy().(MatchmakeParam)
	copied.StartedTime = ms.StartedTime.Copy().(types.DateTime)
	copied.UserPassword = ms.UserPassword.Copy().(types.String)
	copied.ReferGID = ms.ReferGID.Copy().(types.UInt32)
	copied.UserPasswordEnabled = ms.UserPasswordEnabled.Copy().(types.Bool)
	copied.SystemPasswordEnabled = ms.SystemPasswordEnabled.Copy().(types.Bool)
	copied.CodeWord = ms.CodeWord.Copy().(types.String)

	return copied
}

// Equals checks if the given MatchmakeSession contains the same data as the current MatchmakeSession
func (ms MatchmakeSession) Equals(o types.RVType) bool {
	if _, ok := o.(MatchmakeSession); !ok {
		return false
	}

	other := o.(MatchmakeSession)

	if ms.StructureVersion != other.StructureVersion {
		return false
	}

	if !ms.Gathering.Equals(other.Gathering) {
		return false
	}

	if !ms.GameMode.Equals(other.GameMode) {
		return false
	}

	if !ms.Attributes.Equals(other.Attributes) {
		return false
	}

	if !ms.OpenParticipation.Equals(other.OpenParticipation) {
		return false
	}

	if ms.MatchmakeSystemType != other.MatchmakeSystemType {
		return false
	}

	if !ms.ApplicationBuffer.Equals(other.ApplicationBuffer) {
		return false
	}

	if !ms.ParticipationCount.Equals(other.ParticipationCount) {
		return false
	}

	if !ms.ProgressScore.Equals(other.ProgressScore) {
		return false
	}

	if !ms.SessionKey.Equals(other.SessionKey) {
		return false
	}

	if ms.Option0 != other.Option0 {
		return false
	}

	if !ms.MatchmakeParam.Equals(other.MatchmakeParam) {
		return false
	}

	if !ms.StartedTime.Equals(other.StartedTime) {
		return false
	}

	if !ms.UserPassword.Equals(other.UserPassword) {
		return false
	}

	if !ms.ReferGID.Equals(other.ReferGID) {
		return false
	}

	if !ms.UserPasswordEnabled.Equals(other.UserPasswordEnabled) {
		return false
	}

	if !ms.SystemPasswordEnabled.Equals(other.SystemPasswordEnabled) {
		return false
	}

	return ms.CodeWord.Equals(other.CodeWord)
}

// CopyRef copies the current value of the MatchmakeSession
// and returns a pointer to the new copy
func (ms MatchmakeSession) CopyRef() types.RVTypePtr {
	copied := ms.Copy().(MatchmakeSession)
	return &copied
}

// Deref takes a pointer to the MatchmakeSession
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (ms *MatchmakeSession) Deref() types.RVType {
	return *ms
}

// String returns the string representation of the MatchmakeSession
func (ms MatchmakeSession) String() string {
	return ms.FormatToString(0)
}

// FormatToString pretty-prints the MatchmakeSession using the provided indentation level
func (ms MatchmakeSession) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeSession{\n")
	b.WriteString(fmt.Sprintf("%sGathering (parent): %s,\n", indentationValues, ms.Gathering.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sGameMode: %s,\n", indentationValues, ms.GameMode))
	b.WriteString(fmt.Sprintf("%sAttributes: %s,\n", indentationValues, ms.Attributes))
	b.WriteString(fmt.Sprintf("%sOpenParticipation: %s,\n", indentationValues, ms.OpenParticipation))
	b.WriteString(fmt.Sprintf("%sMatchmakeSystemType: %s,\n", indentationValues, ms.MatchmakeSystemType))
	b.WriteString(fmt.Sprintf("%sApplicationBuffer: %s,\n", indentationValues, ms.ApplicationBuffer))
	b.WriteString(fmt.Sprintf("%sParticipationCount: %s,\n", indentationValues, ms.ParticipationCount))
	b.WriteString(fmt.Sprintf("%sProgressScore: %s,\n", indentationValues, ms.ProgressScore))
	b.WriteString(fmt.Sprintf("%sSessionKey: %s,\n", indentationValues, ms.SessionKey))
	b.WriteString(fmt.Sprintf("%sOption: %s,\n", indentationValues, ms.Option0))
	b.WriteString(fmt.Sprintf("%sMatchmakeParam: %s,\n", indentationValues, ms.MatchmakeParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStartedTime: %s,\n", indentationValues, ms.StartedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUserPassword: %s,\n", indentationValues, ms.UserPassword))
	b.WriteString(fmt.Sprintf("%sReferGID: %s,\n", indentationValues, ms.ReferGID))
	b.WriteString(fmt.Sprintf("%sUserPasswordEnabled: %s,\n", indentationValues, ms.UserPasswordEnabled))
	b.WriteString(fmt.Sprintf("%sSystemPasswordEnabled: %s,\n", indentationValues, ms.SystemPasswordEnabled))
	b.WriteString(fmt.Sprintf("%sCodeWord: %s,\n", indentationValues, ms.CodeWord))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeSession returns a new MatchmakeSession
func NewMatchmakeSession() MatchmakeSession {
	return MatchmakeSession{
		Gathering:             NewGathering(),
		GameMode:              types.NewUInt32(0),
		Attributes:            types.NewList[types.UInt32](),
		OpenParticipation:     types.NewBool(false),
		MatchmakeSystemType:   constants.MatchmakeSystemTypeInvalid,
		ApplicationBuffer:     types.NewBuffer(nil),
		ParticipationCount:    types.NewUInt32(0),
		ProgressScore:         types.NewUInt8(0),
		SessionKey:            types.NewBuffer(nil),
		Option0:               constants.MatchmakeSessionOption0None,
		MatchmakeParam:        NewMatchmakeParam(),
		StartedTime:           types.NewDateTime(0),
		UserPassword:          types.NewString(""),
		ReferGID:              types.NewUInt32(0),
		UserPasswordEnabled:   types.NewBool(false),
		SystemPasswordEnabled: types.NewBool(false),
		CodeWord:              types.NewString(""),
	}

}
