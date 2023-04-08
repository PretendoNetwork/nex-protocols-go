package match_making

import nex "github.com/PretendoNetwork/nex-go"

/*
	NEX and Rendez-Vous have multiple protocols for match making
	These protocols all share the same types
	In an effort to keep this library organized, each type used in all match making protocols is defined here
*/

// Gathering holds information about a matchmake gathering
type Gathering struct {
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

	*nex.Structure
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

// NewGathering returns a new Gathering
func NewGathering() *Gathering {
	return &Gathering{}
}

// MatchmakeSessionSearchCriteria holds information about a matchmaking search
type MatchmakeSessionSearchCriteria struct {
	Attribs             []string
	GameMode            string
	MinParticipants     string
	MaxParticipants     string
	MatchmakeSystemType string
	VacantOnly          bool
	ExcludeLocked       bool
	ExcludeNonHostPid   bool
	SelectionMethod     uint32
	VacantParticipants  uint16 // NEX v3.5.0+

	*nex.Structure
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

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 5 {
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

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 5 {
		stream.WriteUInt16LE(matchmakeSessionSearchCriteria.VacantParticipants)
	}

	return stream.Bytes()
}

// NewGathering returns a new Gathering
func NewMatchmakeSessionSearchCriteria() *MatchmakeSessionSearchCriteria {
	return &MatchmakeSessionSearchCriteria{}
}

// MatchmakeSession holds information about a matchmake session
type MatchmakeSession struct {
	GameMode              uint32
	Attributes            []uint32
	OpenParticipation     bool
	MatchmakeSystemType   uint32
	ApplicationData       []byte
	ParticipationCount    uint32
	ProgressScore         uint8           // NEX v3.5.0+
	SessionKey            []byte          // NEX v3.0.0+
	Option                uint32          // NEX v3.5.0+
	MatchmakeParam        *MatchmakeParam // NEX v4.0.0+
	StartedTime           *nex.DateTime   // NEX v4.0.0+
	UserPassword          string          // NEX v4.0.0+
	ReferGID              uint32          // NEX v4.0.0+
	UserPasswordEnabled   bool            // NEX v4.0.0+
	SystemPasswordEnabled bool            // NEX v4.0.0+
	CodeWord              string          // NEX v4.0.0+

	hierarchy []nex.StructureInterface
	*Gathering
}

// GetHierarchy returns the Structure hierarchy
func (matchmakeSession *MatchmakeSession) GetHierarchy() []nex.StructureInterface {
	return matchmakeSession.hierarchy
}

// ExtractFromStream extracts a MatchmakeSession structure from a stream
func (matchmakeSession *MatchmakeSession) ExtractFromStream(stream *nex.StreamIn) error {
	matchmakingVersion := stream.Server.MatchMakingProtocolVersion()

	var err error

	//matchmakeSession.Gathering = gathering
	matchmakeSession.GameMode = stream.ReadUInt32LE()
	matchmakeSession.Attributes = stream.ReadListUInt32LE()
	matchmakeSession.OpenParticipation = stream.ReadUInt8() == 1
	matchmakeSession.MatchmakeSystemType = stream.ReadUInt32LE()
	matchmakeSession.ApplicationData, err = stream.ReadBuffer()

	if err != nil {
		return err
	}

	matchmakeSession.ParticipationCount = stream.ReadUInt32LE()

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 5 {
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

	if matchmakingVersion.Major >= 4 {
		matchmakeParam, err := stream.ReadStructure(NewMatchmakeParam())

		if err != nil {
			return err
		}

		matchmakeSession.MatchmakeParam = matchmakeParam.(*MatchmakeParam)
		matchmakeSession.StartedTime = nex.NewDateTime(stream.ReadUInt64LE())
		matchmakeSession.UserPassword, err = stream.ReadString()

		if err != nil {
			return err
		}

		matchmakeSession.ReferGID = stream.ReadUInt32LE()
		matchmakeSession.UserPasswordEnabled = stream.ReadUInt8() == 1
		matchmakeSession.SystemPasswordEnabled = stream.ReadUInt8() == 1
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
	//stream.WriteStructure(matchmakeSession.Gathering)

	stream.WriteUInt32LE(matchmakeSession.GameMode)
	stream.WriteListUInt32LE(matchmakeSession.Attributes)
	stream.WriteBool(matchmakeSession.OpenParticipation)
	stream.WriteUInt32LE(matchmakeSession.MatchmakeSystemType)
	stream.WriteBuffer(matchmakeSession.ApplicationData)

	stream.WriteUInt32LE(matchmakeSession.ParticipationCount)

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 5 {
		stream.WriteUInt8(matchmakeSession.ProgressScore)
	}

	if matchmakingVersion.Major >= 3 {
		stream.WriteBuffer(matchmakeSession.SessionKey)
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 5 {
		stream.WriteUInt32LE(matchmakeSession.Option)
	}

	//unimplemented for now since MK7 didn't need it
	/*if server.NexVersion() >= 40000 {
		matchmakeParam, err := stream.ReadStructure(NewMatchmakeParam())

		if err != nil {
			return err
		}

		matchmakeSession.MatchmakeParam = matchmakeParam.(*MatchmakeParam)
		matchmakeSession.StartedTime = nex.NewDateTime(stream.ReadUInt64LE())
		matchmakeSession.UserPassword, err = stream.ReadString()

		if err != nil {
			return err
		}

		matchmakeSession.ReferGID = stream.ReadUInt32LE()
		matchmakeSession.UserPasswordEnabled = stream.ReadUInt8() == 1
		matchmakeSession.SystemPasswordEnabled = stream.ReadUInt8() == 1
		matchmakeSession.CodeWord, err = stream.ReadString()

		if err != nil {
			return err
		}
	}*/

	return stream.Bytes()
}

// NewMatchmakeSession returns a new MatchmakeSession
func NewMatchmakeSession() *MatchmakeSession {
	matchmakeSession := &MatchmakeSession{}

	gathering := NewGathering()

	matchmakeSession.Gathering = gathering

	matchmakeSession.hierarchy = []nex.StructureInterface{
		gathering,
	}

	return matchmakeSession
}

// MatchmakeParam holds parameters for a matchmake session
type MatchmakeParam struct {
	parameters map[interface{}]interface{}

	*nex.Structure
}

// ExtractFromStream extracts a MatchmakeParam structure from a stream
func (matchmakeParam *MatchmakeParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	matchmakeParam.parameters, err = stream.ReadMap(stream.ReadString, stream.ReadVariant)

	if err != nil {
		return err
	}

	return nil
}

// NewMatchmakeParam returns a new MatchmakeParam
func NewMatchmakeParam() *MatchmakeParam {
	return &MatchmakeParam{}
}

// CreateMatchmakeSessionParam holds parameters for a matchmake session
type CreateMatchmakeSessionParam struct {
	SourceMatchmakeSession       *MatchmakeSession
	AdditionalParticipants       []uint32
	GIDForParticipationCheck     uint32
	CreateMatchmakeSessionOption uint32
	JoinMessage                  string
	ParticipationCount           uint16

	*nex.Structure
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

// NewCreateMatchmakeSessionParam returns a new CreateMatchmakeSessionParam
func NewCreateMatchmakeSessionParam() *CreateMatchmakeSessionParam {
	return &CreateMatchmakeSessionParam{}
}

// JoinMatchmakeSessionParam holds parameters for a matchmake session
type JoinMatchmakeSessionParam struct {
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

	*nex.Structure
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

// NewJoinMatchmakeSessionParam returns a new JoinMatchmakeSessionParam
func NewJoinMatchmakeSessionParam() *JoinMatchmakeSessionParam {
	return &JoinMatchmakeSessionParam{}
}

// MatchmakeBlockListParam holds parameters for a matchmake session
type MatchmakeBlockListParam struct {
	OptionFlag uint32

	*nex.Structure
}

// ExtractFromStream extracts a MatchmakeBlockListParam structure from a stream
func (matchmakeBlockListParam *MatchmakeBlockListParam) ExtractFromStream(stream *nex.StreamIn) error {
	matchmakeBlockListParam.OptionFlag = stream.ReadUInt32LE()

	return nil
}

// NewMatchmakeBlockListParam returns a new MatchmakeBlockListParam
func NewMatchmakeBlockListParam() *MatchmakeBlockListParam {
	return &MatchmakeBlockListParam{}
}
