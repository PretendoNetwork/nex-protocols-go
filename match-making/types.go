package match_making

import (
	"bytes"

	nex "github.com/PretendoNetwork/nex-go"
)

/*
	NEX and Rendez-Vous have multiple protocols for match making
	These protocols all share the same types
	In an effort to keep this library organized, each type used in all match making protocols is defined here
*/

// Gathering holds information about a matchmake gathering
type Gathering struct {
	nex.Structure
	ID                  uint32
	OwnerPID            uint32
	HostPID             uint32
	MinimumParticipants uint16
	MaximumParticipants uint16
	ParticipationPolicy uint32
	PolicyArgument      uint32
	Flags               uint32
	State               uint32
	Description         string
}

// ExtractFromStream extracts a Gathering structure from a stream
func (gathering *Gathering) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	gathering.ID = stream.ReadUInt32LE()
	gathering.OwnerPID = stream.ReadUInt32LE()
	gathering.HostPID = stream.ReadUInt32LE()
	gathering.MinimumParticipants = stream.ReadUInt16LE()
	gathering.MaximumParticipants = stream.ReadUInt16LE()
	gathering.ParticipationPolicy = stream.ReadUInt32LE()
	gathering.PolicyArgument = stream.ReadUInt32LE()
	gathering.Flags = stream.ReadUInt32LE()
	gathering.State = stream.ReadUInt32LE()
	gathering.Description, err = stream.ReadString()

	if err != nil {
		return err
	}

	return nil
}

// Bytes encodes the Gathering and returns a byte array
func (gathering *Gathering) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(gathering.ID)
	stream.WriteUInt32LE(gathering.OwnerPID)
	stream.WriteUInt32LE(gathering.HostPID)
	stream.WriteUInt16LE(gathering.MinimumParticipants)
	stream.WriteUInt16LE(gathering.MaximumParticipants)
	stream.WriteUInt32LE(gathering.ParticipationPolicy)
	stream.WriteUInt32LE(gathering.PolicyArgument)
	stream.WriteUInt32LE(gathering.Flags)
	stream.WriteUInt32LE(gathering.State)
	stream.WriteString(gathering.Description)

	return stream.Bytes()
}

// Copy returns a new copied instance of Gathering
func (gathering *Gathering) Copy() nex.StructureInterface {
	copied := NewGathering()

	copied.ID = gathering.ID
	copied.OwnerPID = gathering.OwnerPID
	copied.HostPID = gathering.HostPID
	copied.MinimumParticipants = gathering.MinimumParticipants
	copied.MaximumParticipants = gathering.MaximumParticipants
	copied.ParticipationPolicy = gathering.ParticipationPolicy
	copied.PolicyArgument = gathering.PolicyArgument
	copied.Flags = gathering.Flags
	copied.State = gathering.State
	copied.Description = gathering.Description

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (gathering *Gathering) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Gathering)

	if gathering.ID != other.ID {
		return false
	}

	if gathering.OwnerPID != other.OwnerPID {
		return false
	}

	if gathering.HostPID != other.HostPID {
		return false
	}

	if gathering.MinimumParticipants != other.MinimumParticipants {
		return false
	}

	if gathering.MaximumParticipants != other.MaximumParticipants {
		return false
	}

	if gathering.ParticipationPolicy != other.ParticipationPolicy {
		return false
	}

	if gathering.PolicyArgument != other.PolicyArgument {
		return false
	}

	if gathering.Flags != other.Flags {
		return false
	}

	if gathering.State != other.State {
		return false
	}

	if gathering.Description != other.Description {
		return false
	}

	return true
}

// NewGathering returns a new Gathering
func NewGathering() *Gathering {
	return &Gathering{}
}

// MatchmakeSessionSearchCriteria holds information about a matchmaking search
type MatchmakeSessionSearchCriteria struct {
	nex.Structure
	Attribs             []string
	GameMode            string
	MinParticipants     string
	MaxParticipants     string
	MatchmakeSystemType string
	VacantOnly          bool
	ExcludeLocked       bool
	ExcludeNonHostPid   bool
	SelectionMethod     uint32
	VacantParticipants  uint16 // NEX v3.4.0+
}

// ExtractFromStream extracts a Gathering structure from a stream
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) ExtractFromStream(stream *nex.StreamIn) error {
	matchmakingVersion := stream.Server.MatchMakingProtocolVersion()

	var err error

	matchmakeSessionSearchCriteria.Attribs = stream.ReadListString()
	matchmakeSessionSearchCriteria.GameMode, err = stream.ReadString()
	if err != nil {
		return err
	}
	matchmakeSessionSearchCriteria.MinParticipants, err = stream.ReadString()
	if err != nil {
		return err
	}
	matchmakeSessionSearchCriteria.MaxParticipants, err = stream.ReadString()
	if err != nil {
		return err
	}
	matchmakeSessionSearchCriteria.MatchmakeSystemType, err = stream.ReadString()
	if err != nil {
		return err
	}
	matchmakeSessionSearchCriteria.VacantOnly = stream.ReadBool()
	matchmakeSessionSearchCriteria.ExcludeLocked = stream.ReadBool()
	matchmakeSessionSearchCriteria.ExcludeNonHostPid = stream.ReadBool()
	matchmakeSessionSearchCriteria.SelectionMethod = stream.ReadUInt32LE()

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 4 {
		matchmakeSessionSearchCriteria.VacantParticipants = stream.ReadUInt16LE()
	}

	return nil
}

// Bytes encodes the Gathering and returns a byte array
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) Bytes(stream *nex.StreamOut) []byte {
	matchmakingVersion := stream.Server.MatchMakingProtocolVersion()

	stream.WriteListString(matchmakeSessionSearchCriteria.Attribs)
	stream.WriteString(matchmakeSessionSearchCriteria.GameMode)
	stream.WriteString(matchmakeSessionSearchCriteria.MinParticipants)
	stream.WriteString(matchmakeSessionSearchCriteria.MaxParticipants)
	stream.WriteString(matchmakeSessionSearchCriteria.MatchmakeSystemType)
	stream.WriteBool(matchmakeSessionSearchCriteria.VacantOnly)
	stream.WriteBool(matchmakeSessionSearchCriteria.ExcludeLocked)
	stream.WriteBool(matchmakeSessionSearchCriteria.ExcludeNonHostPid)
	stream.WriteUInt32LE(matchmakeSessionSearchCriteria.SelectionMethod)

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 4 {
		stream.WriteUInt16LE(matchmakeSessionSearchCriteria.VacantParticipants)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of Gathering
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) Copy() nex.StructureInterface {
	copied := NewMatchmakeSessionSearchCriteria()

	copied.Attribs = make([]string, len(matchmakeSessionSearchCriteria.Attribs))

	copy(copied.Attribs, matchmakeSessionSearchCriteria.Attribs)

	copied.GameMode = matchmakeSessionSearchCriteria.GameMode
	copied.MinParticipants = matchmakeSessionSearchCriteria.MinParticipants
	copied.MaxParticipants = matchmakeSessionSearchCriteria.MaxParticipants
	copied.MatchmakeSystemType = matchmakeSessionSearchCriteria.MatchmakeSystemType
	copied.VacantOnly = matchmakeSessionSearchCriteria.VacantOnly
	copied.ExcludeLocked = matchmakeSessionSearchCriteria.ExcludeLocked
	copied.ExcludeNonHostPid = matchmakeSessionSearchCriteria.ExcludeNonHostPid
	copied.SelectionMethod = matchmakeSessionSearchCriteria.SelectionMethod
	copied.VacantParticipants = matchmakeSessionSearchCriteria.VacantParticipants

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeSessionSearchCriteria)

	if len(matchmakeSessionSearchCriteria.Attribs) != len(other.Attribs) {
		return false
	}

	for i := 0; i < len(matchmakeSessionSearchCriteria.Attribs); i++ {
		if matchmakeSessionSearchCriteria.Attribs[i] != other.Attribs[i] {
			return false
		}
	}

	if matchmakeSessionSearchCriteria.GameMode != other.GameMode {
		return false
	}

	if matchmakeSessionSearchCriteria.MinParticipants != other.MinParticipants {
		return false
	}

	if matchmakeSessionSearchCriteria.MaxParticipants != other.MaxParticipants {
		return false
	}

	if matchmakeSessionSearchCriteria.MatchmakeSystemType != other.MatchmakeSystemType {
		return false
	}

	if matchmakeSessionSearchCriteria.VacantOnly != other.VacantOnly {
		return false
	}

	if matchmakeSessionSearchCriteria.ExcludeLocked != other.ExcludeLocked {
		return false
	}

	if matchmakeSessionSearchCriteria.ExcludeNonHostPid != other.ExcludeNonHostPid {
		return false
	}

	if matchmakeSessionSearchCriteria.SelectionMethod != other.SelectionMethod {
		return false
	}

	if matchmakeSessionSearchCriteria.VacantParticipants != other.VacantParticipants {
		return false
	}

	return true
}

// NewGathering returns a new Gathering
func NewMatchmakeSessionSearchCriteria() *MatchmakeSessionSearchCriteria {
	return &MatchmakeSessionSearchCriteria{}
}

// MatchmakeSession holds information about a matchmake session
type MatchmakeSession struct {
	nex.Structure
	*Gathering
	GameMode              uint32
	Attributes            []uint32
	OpenParticipation     bool
	MatchmakeSystemType   uint32
	ApplicationData       []byte
	ParticipationCount    uint32
	ProgressScore         uint8           // NEX v3.4.0+
	SessionKey            []byte          // NEX v3.0.0+
	Option                uint32          // NEX v3.5.0+
	MatchmakeParam        *MatchmakeParam // NEX v3.6.0+
	StartedTime           *nex.DateTime   // NEX v3.6.0+
	UserPassword          string          // NEX v3.7.0+
	ReferGID              uint32          // NEX v3.8.0+
	UserPasswordEnabled   bool            // NEX v3.8.0+
	SystemPasswordEnabled bool            // NEX v3.8.0+
	CodeWord              string          // NEX v4.0.0+
}

// ExtractFromStream extracts a MatchmakeSession structure from a stream
func (matchmakeSession *MatchmakeSession) ExtractFromStream(stream *nex.StreamIn) error {
	matchmakingVersion := stream.Server.MatchMakingProtocolVersion()

	var err error

	matchmakeSession.GameMode = stream.ReadUInt32LE()
	matchmakeSession.Attributes = stream.ReadListUInt32LE()
	matchmakeSession.OpenParticipation = stream.ReadUInt8() == 1
	matchmakeSession.MatchmakeSystemType = stream.ReadUInt32LE()
	matchmakeSession.ApplicationData, err = stream.ReadBuffer()

	if err != nil {
		return err
	}

	matchmakeSession.ParticipationCount = stream.ReadUInt32LE()

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 4 {
		matchmakeSession.ProgressScore = stream.ReadUInt8()
	}

	if matchmakingVersion.Major >= 3 {
		matchmakeSession.SessionKey, err = stream.ReadBuffer()

		if err != nil {
			return err
		}
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 5 {
		matchmakeSession.Option = stream.ReadUInt32LE()
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 6 {
		matchmakeParam, err := stream.ReadStructure(NewMatchmakeParam())

		if err != nil {
			return err
		}

		matchmakeSession.MatchmakeParam = matchmakeParam.(*MatchmakeParam)
		matchmakeSession.StartedTime = stream.ReadDateTime()
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 7 {
		matchmakeSession.UserPassword, err = stream.ReadString()

		if err != nil {
			return err
		}
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 8 {
		matchmakeSession.ReferGID = stream.ReadUInt32LE()
		matchmakeSession.UserPasswordEnabled = stream.ReadBool()
		matchmakeSession.SystemPasswordEnabled = stream.ReadBool()
	}

	if matchmakingVersion.Major >= 4 {
		matchmakeSession.CodeWord, err = stream.ReadString()

		if err != nil {
			return err
		}
	}

	return nil
}

// Bytes extracts a MatchmakeSession structure from a stream
func (matchmakeSession *MatchmakeSession) Bytes(stream *nex.StreamOut) []byte {
	matchmakingVersion := stream.Server.MatchMakingProtocolVersion()

	stream.WriteUInt32LE(matchmakeSession.GameMode)
	stream.WriteListUInt32LE(matchmakeSession.Attributes)
	stream.WriteBool(matchmakeSession.OpenParticipation)
	stream.WriteUInt32LE(matchmakeSession.MatchmakeSystemType)
	stream.WriteBuffer(matchmakeSession.ApplicationData)

	stream.WriteUInt32LE(matchmakeSession.ParticipationCount)

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 4 {
		stream.WriteUInt8(matchmakeSession.ProgressScore)
	}

	if matchmakingVersion.Major >= 3 {
		stream.WriteBuffer(matchmakeSession.SessionKey)
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 5 {
		stream.WriteUInt32LE(matchmakeSession.Option)
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 6 {
		stream.WriteStructure(matchmakeSession.MatchmakeParam)
		stream.WriteDateTime(matchmakeSession.StartedTime)
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 7 {
		stream.WriteString(matchmakeSession.UserPassword)
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 8 {
		stream.WriteUInt32LE(matchmakeSession.ReferGID)
		stream.WriteBool(matchmakeSession.UserPasswordEnabled)
		stream.WriteBool(matchmakeSession.SystemPasswordEnabled)
	}

	if matchmakingVersion.Major >= 4 {
		stream.WriteString(matchmakeSession.CodeWord)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of MatchmakeSession
func (matchmakeSession *MatchmakeSession) Copy() nex.StructureInterface {
	copied := NewMatchmakeSession()

	copied.SetParentType(matchmakeSession.ParentType().Copy())
	copied.GameMode = matchmakeSession.GameMode
	copied.Attributes = make([]uint32, len(matchmakeSession.Attributes))

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

	if matchmakeSession.MatchmakeParam != nil {
		copied.MatchmakeParam = matchmakeSession.MatchmakeParam.Copy().(*MatchmakeParam)
	}

	if matchmakeSession.StartedTime != nil {
		copied.StartedTime = matchmakeSession.StartedTime.Copy()
	}

	copied.UserPassword = matchmakeSession.UserPassword
	copied.ReferGID = matchmakeSession.ReferGID
	copied.UserPasswordEnabled = matchmakeSession.UserPasswordEnabled
	copied.SystemPasswordEnabled = matchmakeSession.SystemPasswordEnabled
	copied.CodeWord = matchmakeSession.CodeWord

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeSession *MatchmakeSession) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeSession)

	if !matchmakeSession.ParentType().Equals(other.ParentType()) {
		return false
	}

	if matchmakeSession.GameMode != other.GameMode {
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

	if matchmakeSession.OpenParticipation != other.OpenParticipation {
		return false
	}

	if matchmakeSession.MatchmakeSystemType != other.MatchmakeSystemType {
		return false
	}

	if !bytes.Equal(matchmakeSession.ApplicationData, other.ApplicationData) {
		return false
	}

	if matchmakeSession.ParticipationCount != other.ParticipationCount {
		return false
	}

	if matchmakeSession.ProgressScore != other.ProgressScore {
		return false
	}

	if !bytes.Equal(matchmakeSession.SessionKey, other.SessionKey) {
		return false
	}

	if matchmakeSession.Option != other.Option {
		return false
	}

	if matchmakeSession.MatchmakeParam != nil && other.MatchmakeParam == nil {
		return false
	}

	if matchmakeSession.MatchmakeParam == nil && other.MatchmakeParam != nil {
		return false
	}

	if matchmakeSession.MatchmakeParam != nil && other.MatchmakeParam != nil {
		if !matchmakeSession.MatchmakeParam.Equals(other.MatchmakeParam) {
			return false
		}
	}

	if matchmakeSession.StartedTime != nil && other.StartedTime == nil {
		return false
	}

	if matchmakeSession.StartedTime == nil && other.StartedTime != nil {
		return false
	}

	if matchmakeSession.StartedTime != nil && other.StartedTime != nil {
		if !matchmakeSession.StartedTime.Equals(other.StartedTime) {
			return false
		}
	}

	if matchmakeSession.UserPassword != other.UserPassword {
		return false
	}

	if matchmakeSession.ReferGID != other.ReferGID {
		return false
	}

	if matchmakeSession.UserPasswordEnabled != other.UserPasswordEnabled {
		return false
	}

	if matchmakeSession.SystemPasswordEnabled != other.SystemPasswordEnabled {
		return false
	}

	if matchmakeSession.CodeWord != other.CodeWord {
		return false
	}

	return true
}

// NewMatchmakeSession returns a new MatchmakeSession
func NewMatchmakeSession() *MatchmakeSession {
	matchmakeSession := &MatchmakeSession{}
	matchmakeSession.Gathering = NewGathering()
	matchmakeSession.SetParentType(matchmakeSession.Gathering)

	return matchmakeSession
}

// MatchmakeParam holds parameters for a matchmake session
type MatchmakeParam struct {
	nex.Structure
	parameters map[string]*nex.Variant
}

// ExtractFromStream extracts a MatchmakeParam structure from a stream
func (matchmakeParam *MatchmakeParam) ExtractFromStream(stream *nex.StreamIn) error {
	parameters, err := stream.ReadMap(stream.ReadString, stream.ReadVariant)

	if err != nil {
		return err
	}

	matchmakeParam.parameters = make(map[string]*nex.Variant, len(parameters))

	for key, value := range parameters {
		matchmakeParam.parameters[key.(string)] = value.(*nex.Variant)
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeParam
func (matchmakeParam *MatchmakeParam) Copy() nex.StructureInterface {
	copied := NewMatchmakeParam()

	copied.parameters = make(map[string]*nex.Variant, len(matchmakeParam.parameters))

	for key, value := range matchmakeParam.parameters {
		copied.parameters[key] = value.Copy()
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeParam *MatchmakeParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeParam)

	if len(matchmakeParam.parameters) != len(other.parameters) {
		return false
	}

	for key, value := range matchmakeParam.parameters {
		if !value.Equals(other.parameters[key]) {
			return false
		}
	}

	return true
}

// NewMatchmakeParam returns a new MatchmakeParam
func NewMatchmakeParam() *MatchmakeParam {
	return &MatchmakeParam{}
}

// CreateMatchmakeSessionParam holds parameters for a matchmake session
type CreateMatchmakeSessionParam struct {
	nex.Structure
	SourceMatchmakeSession       *MatchmakeSession
	AdditionalParticipants       []uint32
	GIDForParticipationCheck     uint32
	CreateMatchmakeSessionOption uint32
	JoinMessage                  string
	ParticipationCount           uint16
}

// ExtractFromStream extracts a CreateMatchmakeSessionParam structure from a stream
func (createMatchmakeSessionParam *CreateMatchmakeSessionParam) ExtractFromStream(stream *nex.StreamIn) error {
	sourceMatchmakeSession, err := stream.ReadStructure(NewMatchmakeSession())
	if err != nil {
		return err
	}

	createMatchmakeSessionParam.SourceMatchmakeSession = sourceMatchmakeSession.(*MatchmakeSession)
	createMatchmakeSessionParam.AdditionalParticipants = stream.ReadListUInt32LE()
	createMatchmakeSessionParam.GIDForParticipationCheck = stream.ReadUInt32LE()
	createMatchmakeSessionParam.CreateMatchmakeSessionOption = stream.ReadUInt32LE()
	createMatchmakeSessionParam.JoinMessage, _ = stream.ReadString()
	createMatchmakeSessionParam.ParticipationCount = stream.ReadUInt16LE()

	return nil
}

// Copy returns a new copied instance of CreateMatchmakeSessionParam
func (createMatchmakeSessionParam *CreateMatchmakeSessionParam) Copy() nex.StructureInterface {
	copied := NewCreateMatchmakeSessionParam()

	if createMatchmakeSessionParam.SourceMatchmakeSession != nil {
		copied.SourceMatchmakeSession = createMatchmakeSessionParam.SourceMatchmakeSession.Copy().(*MatchmakeSession)
	}

	copied.AdditionalParticipants = make([]uint32, len(createMatchmakeSessionParam.AdditionalParticipants))

	copy(copied.AdditionalParticipants, createMatchmakeSessionParam.AdditionalParticipants)

	copied.GIDForParticipationCheck = createMatchmakeSessionParam.GIDForParticipationCheck
	copied.CreateMatchmakeSessionOption = createMatchmakeSessionParam.CreateMatchmakeSessionOption
	copied.JoinMessage = createMatchmakeSessionParam.JoinMessage
	copied.ParticipationCount = createMatchmakeSessionParam.ParticipationCount

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (createMatchmakeSessionParam *CreateMatchmakeSessionParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*CreateMatchmakeSessionParam)

	if createMatchmakeSessionParam.SourceMatchmakeSession != nil && other.SourceMatchmakeSession == nil {
		return false
	}

	if createMatchmakeSessionParam.SourceMatchmakeSession == nil && other.SourceMatchmakeSession != nil {
		return false
	}

	if createMatchmakeSessionParam.SourceMatchmakeSession != nil && other.SourceMatchmakeSession != nil {
		if !createMatchmakeSessionParam.SourceMatchmakeSession.Equals(other.SourceMatchmakeSession) {
			return false
		}
	}

	if len(createMatchmakeSessionParam.AdditionalParticipants) != len(other.AdditionalParticipants) {
		return false
	}

	for i := 0; i < len(createMatchmakeSessionParam.AdditionalParticipants); i++ {
		if createMatchmakeSessionParam.AdditionalParticipants[i] != other.AdditionalParticipants[i] {
			return false
		}
	}

	if createMatchmakeSessionParam.GIDForParticipationCheck != other.GIDForParticipationCheck {
		return false
	}

	if createMatchmakeSessionParam.CreateMatchmakeSessionOption != other.CreateMatchmakeSessionOption {
		return false
	}

	if createMatchmakeSessionParam.JoinMessage != other.JoinMessage {
		return false
	}

	if createMatchmakeSessionParam.ParticipationCount != other.ParticipationCount {
		return false
	}

	return true
}

// NewCreateMatchmakeSessionParam returns a new CreateMatchmakeSessionParam
func NewCreateMatchmakeSessionParam() *CreateMatchmakeSessionParam {
	return &CreateMatchmakeSessionParam{}
}

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
	// TODO - Check errors

	joinMatchmakeSessionParam.GID = stream.ReadUInt32LE()
	joinMatchmakeSessionParam.AdditionalParticipants = stream.ReadListUInt32LE()
	joinMatchmakeSessionParam.GIDForParticipationCheck = stream.ReadUInt32LE()
	joinMatchmakeSessionParam.JoinMatchmakeSessionOption = stream.ReadUInt32LE()
	joinMatchmakeSessionParam.JoinMatchmakeSessionBehavior = stream.ReadUInt8()
	joinMatchmakeSessionParam.StrUserPassword, _ = stream.ReadString()
	joinMatchmakeSessionParam.StrSystemPassword, _ = stream.ReadString()
	joinMatchmakeSessionParam.JoinMessage, _ = stream.ReadString()
	joinMatchmakeSessionParam.ParticipationCount = stream.ReadUInt16LE()
	joinMatchmakeSessionParam.ExtraParticipants = stream.ReadUInt16LE()

	blockListParam, err := stream.ReadStructure(NewMatchmakeBlockListParam())
	if err != nil {
		return err
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

// NewJoinMatchmakeSessionParam returns a new JoinMatchmakeSessionParam
func NewJoinMatchmakeSessionParam() *JoinMatchmakeSessionParam {
	return &JoinMatchmakeSessionParam{}
}

// MatchmakeBlockListParam holds parameters for a matchmake session
type MatchmakeBlockListParam struct {
	nex.Structure
	OptionFlag uint32
}

// ExtractFromStream extracts a MatchmakeBlockListParam structure from a stream
func (matchmakeBlockListParam *MatchmakeBlockListParam) ExtractFromStream(stream *nex.StreamIn) error {
	matchmakeBlockListParam.OptionFlag = stream.ReadUInt32LE()

	return nil
}

// Copy returns a new copied instance of MatchmakeBlockListParam
func (matchmakeBlockListParam *MatchmakeBlockListParam) Copy() nex.StructureInterface {
	copied := NewMatchmakeBlockListParam()

	copied.OptionFlag = matchmakeBlockListParam.OptionFlag

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeBlockListParam *MatchmakeBlockListParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeBlockListParam)

	return matchmakeBlockListParam.OptionFlag == other.OptionFlag
}

// NewMatchmakeBlockListParam returns a new MatchmakeBlockListParam
func NewMatchmakeBlockListParam() *MatchmakeBlockListParam {
	return &MatchmakeBlockListParam{}
}

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
	sourceMatchmakeSession, err := stream.ReadStructure(NewMatchmakeSession())
	if err != nil {
		return err
	}

	autoMatchmakeParam.SourceMatchmakeSession = sourceMatchmakeSession.(*MatchmakeSession)
	autoMatchmakeParam.AdditionalParticipants = stream.ReadListUInt32LE()
	autoMatchmakeParam.GIDForParticipationCheck = stream.ReadUInt32LE()
	autoMatchmakeParam.AutoMatchmakeOption = stream.ReadUInt32LE()
	autoMatchmakeParam.JoinMessage, _ = stream.ReadString()
	autoMatchmakeParam.ParticipationCount = stream.ReadUInt16LE()

	lstSearchCriteria, err := stream.ReadListStructure(NewMatchmakeSessionSearchCriteria())
	if err != nil {
		return err
	}

	autoMatchmakeParam.LstSearchCriteria = lstSearchCriteria.([]*MatchmakeSessionSearchCriteria)
	autoMatchmakeParam.TargetGIDs = stream.ReadListUInt32LE()

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

// PersistentGathering holds parameters for a matchmake session
type PersistentGathering struct {
	nex.Structure
	*Gathering
	M_CommunityType          uint32
	M_Password               string
	M_Attribs                []uint32
	M_ApplicationBuffer      []byte
	M_ParticipationStartDate *nex.DateTime
	M_ParticipationEndDate   *nex.DateTime
	M_MatchmakeSessionCount  uint32
	M_ParticipationCount     uint32
}

// ExtractFromStream extracts a PersistentGathering structure from a stream
func (persistentGathering *PersistentGathering) ExtractFromStream(stream *nex.StreamIn) error {
	persistentGathering.M_CommunityType = stream.ReadUInt32LE()
	persistentGathering.M_Password, _ = stream.ReadString()
	persistentGathering.M_Attribs = stream.ReadListUInt32LE()
	persistentGathering.M_ApplicationBuffer, _ = stream.ReadBuffer()
	persistentGathering.M_ParticipationStartDate = stream.ReadDateTime()
	persistentGathering.M_ParticipationEndDate = stream.ReadDateTime()
	persistentGathering.M_MatchmakeSessionCount = stream.ReadUInt32LE()
	persistentGathering.M_ParticipationCount = stream.ReadUInt32LE()

	return nil
}

// Copy returns a new copied instance of PersistentGathering
func (persistentGathering *PersistentGathering) Copy() nex.StructureInterface {
	copied := NewPersistentGathering()

	copied.SetParentType(persistentGathering.ParentType().Copy())
	copied.M_CommunityType = persistentGathering.M_CommunityType
	copied.M_Password = persistentGathering.M_Password
	copied.M_Attribs = make([]uint32, len(persistentGathering.M_Attribs))

	copy(copied.M_Attribs, persistentGathering.M_Attribs)

	copied.M_ApplicationBuffer = make([]byte, len(persistentGathering.M_ApplicationBuffer))

	copy(copied.M_ApplicationBuffer, persistentGathering.M_ApplicationBuffer)

	if persistentGathering.M_ParticipationStartDate != nil {
		copied.M_ParticipationStartDate = persistentGathering.M_ParticipationStartDate.Copy()
	}

	if persistentGathering.M_ParticipationEndDate != nil {
		copied.M_ParticipationEndDate = persistentGathering.M_ParticipationEndDate.Copy()
	}

	copied.M_MatchmakeSessionCount = persistentGathering.M_MatchmakeSessionCount
	copied.M_ParticipationCount = persistentGathering.M_ParticipationCount

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (persistentGathering *PersistentGathering) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PersistentGathering)

	if !persistentGathering.ParentType().Equals(other.ParentType()) {
		return false
	}

	if persistentGathering.M_CommunityType != other.M_CommunityType {
		return false
	}

	if persistentGathering.M_Password != other.M_Password {
		return false
	}

	if len(persistentGathering.M_Attribs) != len(other.M_Attribs) {
		return false
	}

	for i := 0; i < len(persistentGathering.M_Attribs); i++ {
		if persistentGathering.M_Attribs[i] != other.M_Attribs[i] {
			return false
		}
	}

	if !bytes.Equal(persistentGathering.M_ApplicationBuffer, other.M_ApplicationBuffer) {
		return false
	}

	if persistentGathering.M_ParticipationStartDate != nil && other.M_ParticipationStartDate == nil {
		return false
	}

	if persistentGathering.M_ParticipationStartDate == nil && other.M_ParticipationStartDate != nil {
		return false
	}

	if persistentGathering.M_ParticipationStartDate != nil && other.M_ParticipationStartDate != nil {
		if persistentGathering.M_ParticipationStartDate.Equals(other.M_ParticipationStartDate) {
			return false
		}
	}

	if persistentGathering.M_ParticipationEndDate != nil && other.M_ParticipationEndDate == nil {
		return false
	}

	if persistentGathering.M_ParticipationEndDate == nil && other.M_ParticipationEndDate != nil {
		return false
	}

	if persistentGathering.M_ParticipationEndDate != nil && other.M_ParticipationEndDate != nil {
		if persistentGathering.M_ParticipationEndDate.Equals(other.M_ParticipationEndDate) {
			return false
		}
	}

	if persistentGathering.M_MatchmakeSessionCount != other.M_MatchmakeSessionCount {
		return false
	}

	if persistentGathering.M_ParticipationCount != other.M_ParticipationCount {
		return false
	}

	return true
}

// NewPersistentGathering returns a new PersistentGathering
func NewPersistentGathering() *PersistentGathering {
	persistentGathering := &PersistentGathering{}
	persistentGathering.Gathering = NewGathering()
	persistentGathering.SetParentType(persistentGathering.Gathering)

	return persistentGathering
}

// AutoMatchmakeParam holds parameters for a matchmake session
type SimpleCommunity struct {
	nex.Structure
	M_GatheringID           uint32
	M_MatchmakeSessionCount uint32
}

// ExtractFromStream extracts a SimpleCommunity structure from a stream
func (simpleCommunity *SimpleCommunity) ExtractFromStream(stream *nex.StreamIn) error {
	simpleCommunity.M_GatheringID = stream.ReadUInt32LE()
	simpleCommunity.M_MatchmakeSessionCount = stream.ReadUInt32LE()

	return nil
}

// Copy returns a new copied instance of SimpleCommunity
func (simpleCommunity *SimpleCommunity) Copy() nex.StructureInterface {
	copied := NewSimpleCommunity()

	copied.M_GatheringID = simpleCommunity.M_GatheringID
	copied.M_MatchmakeSessionCount = simpleCommunity.M_MatchmakeSessionCount

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleCommunity *SimpleCommunity) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SimpleCommunity)

	if simpleCommunity.M_GatheringID != other.M_GatheringID {
		return false
	}

	if simpleCommunity.M_MatchmakeSessionCount != other.M_MatchmakeSessionCount {
		return false
	}

	return true
}

// NewSimpleCommunity returns a new SimpleCommunity
func NewSimpleCommunity() *SimpleCommunity {
	return &SimpleCommunity{}
}
