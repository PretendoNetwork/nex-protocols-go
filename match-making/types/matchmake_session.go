package match_making_types

import (
	"bytes"
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

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

	matchmakeSession.GameMode, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.GameMode. %s", err.Error())
	}

	matchmakeSession.Attributes, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.Attributes. %s", err.Error())
	}

	matchmakeSession.OpenParticipation, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.OpenParticipation. %s", err.Error())
	}

	matchmakeSession.MatchmakeSystemType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.MatchmakeSystemType. %s", err.Error())
	}

	matchmakeSession.ApplicationData, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.ApplicationData. %s", err.Error())
	}

	matchmakeSession.ParticipationCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSession.ParticipationCount. %s", err.Error())
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 4 {
		matchmakeSession.ProgressScore, err = stream.ReadUInt8()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.ProgressScore. %s", err.Error())
		}
	}

	if matchmakingVersion.Major >= 3 {
		matchmakeSession.SessionKey, err = stream.ReadBuffer()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.SessionKey. %s", err.Error())
		}
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 5 {
		matchmakeSession.Option, err = stream.ReadUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.Option. %s", err.Error())
		}
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 6 {
		matchmakeParam, err := stream.ReadStructure(NewMatchmakeParam())
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.MatchmakeParam. %s", err.Error())
		}

		matchmakeSession.MatchmakeParam = matchmakeParam.(*MatchmakeParam)
		matchmakeSession.StartedTime, err = stream.ReadDateTime()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.StartedTime. %s", err.Error())
		}
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 7 {
		matchmakeSession.UserPassword, err = stream.ReadString()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.UserPassword. %s", err.Error())
		}
	}

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 8 {
		matchmakeSession.ReferGID, err = stream.ReadUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.ReferGID. %s", err.Error())
		}

		matchmakeSession.UserPasswordEnabled, err = stream.ReadBool()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.UserPasswordEnabled. %s", err.Error())
		}

		matchmakeSession.SystemPasswordEnabled, err = stream.ReadBool()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.SystemPasswordEnabled. %s", err.Error())
		}
	}

	if matchmakingVersion.Major >= 4 {
		matchmakeSession.CodeWord, err = stream.ReadString()
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSession.CodeWord. %s", err.Error())
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

	copied.Gathering = matchmakeSession.Gathering.Copy().(*Gathering)
	copied.SetParentType(copied.Gathering)
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
