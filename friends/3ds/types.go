package friends_3ds

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
)

type Mii struct {
	Name     string
	Unknown2 bool
	Unknown3 uint8
	MiiData  []byte

	nex.Structure
}

// Bytes encodes the Mii and returns a byte array
func (mii *Mii) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(mii.Name)
	stream.WriteBool(mii.Unknown2)
	stream.WriteUInt8(mii.Unknown3)
	stream.WriteBuffer(mii.MiiData)

	return stream.Bytes()
}

// ExtractFromStream extracts a Mii from a stream
func (mii *Mii) ExtractFromStream(stream *nex.StreamIn) error {
	mii.Name, _ = stream.ReadString()
	mii.Unknown2 = (stream.ReadUInt8() == 1)
	mii.Unknown3 = stream.ReadUInt8()
	mii.MiiData, _ = stream.ReadBuffer()

	return nil
}

// NewMii returns a new Mii
func NewMii() *Mii {
	return &Mii{}
}

type FriendMii struct {
	PID        uint32
	Mii        *Mii
	ModifiedAt *nex.DateTime

	nex.Structure
}

// Bytes encodes the Mii and returns a byte array
func (friendMii *FriendMii) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(friendMii.PID)
	stream.WriteStructure(friendMii.Mii)
	stream.WriteDateTime(friendMii.ModifiedAt)

	return stream.Bytes()
}

// NewMii returns a new Mii
func NewFriendMii() *FriendMii {
	return &FriendMii{}
}

type MyProfile struct {
	Region   uint8
	Country  uint8
	Area     uint8
	Language uint8
	Platform uint8
	Unknown1 uint64
	Unknown2 string
	Unknown3 string

	nex.Structure
}

// ExtractFromStream extracts a MyProfile from a stream
func (myProfile *MyProfile) ExtractFromStream(stream *nex.StreamIn) error {
	myProfile.Region = stream.ReadUInt8()
	myProfile.Country = stream.ReadUInt8()
	myProfile.Area = stream.ReadUInt8()
	myProfile.Language = stream.ReadUInt8()
	myProfile.Platform = stream.ReadUInt8()
	myProfile.Unknown1 = stream.ReadUInt64LE()
	myProfile.Unknown2, _ = stream.ReadString()
	myProfile.Unknown3, _ = stream.ReadString()

	return nil
}

// NewMyProfile returns a new MyProfile
func NewMyProfile() *MyProfile {
	return &MyProfile{}
}

// NintendoPresence contains information about a users online presence
type NintendoPresence struct {
	ChangedFlags      uint32
	GameKey           *GameKey
	Message           string
	JoinAvailableFlag uint32
	MatchmakeType     uint8
	JoinGameID        uint32
	JoinGameMode      uint32
	OwnerPID          uint32
	JoinGroupID       uint32
	ApplicationArg    []byte

	nex.Structure
}

// Bytes encodes the NintendoPresence and returns a byte array
func (presence *NintendoPresence) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(presence.ChangedFlags)
	stream.WriteStructure(presence.GameKey)
	stream.WriteString(presence.Message)
	stream.WriteUInt32LE(presence.JoinAvailableFlag)
	stream.WriteUInt8(presence.MatchmakeType)
	stream.WriteUInt32LE(presence.JoinGameID)
	stream.WriteUInt32LE(presence.JoinGameMode)
	stream.WriteUInt32LE(presence.OwnerPID)
	stream.WriteUInt32LE(presence.JoinGroupID)
	stream.WriteBuffer(presence.ApplicationArg)

	return stream.Bytes()
}

// ExtractFromStream extracts a NintendoPresence structure from a stream
func (presence *NintendoPresence) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 25 {
		// length check for the following fixed-size data
		// changedFlags + JoinAvailableFlag + MatchmakeType + JoinGameID + JoinGameMode + OwnerPID + JoinGroupID
		return errors.New("[NintendoPresence::ExtractFromStream] Data size too small")
	}

	changedFlags := stream.ReadUInt32LE()
	gameKeyStructureInterface, err := stream.ReadStructure(NewGameKey())
	if err != nil {
		return err
	}
	gameKey := gameKeyStructureInterface.(*GameKey)
	message, err := stream.ReadString()
	if err != nil {
		return err
	}
	JoinAvailableFlag := stream.ReadUInt32LE()
	MatchmakeType := stream.ReadUInt8()
	JoinGameID := stream.ReadUInt32LE()
	JoinGameMode := stream.ReadUInt32LE()
	OwnerPID := stream.ReadUInt32LE()
	JoinGroupID := stream.ReadUInt32LE()
	ApplicationArg, err := stream.ReadBuffer()
	if err != nil {
		return err
	}

	presence.ChangedFlags = changedFlags
	presence.GameKey = gameKey
	presence.Message = message
	presence.JoinAvailableFlag = JoinAvailableFlag
	presence.MatchmakeType = MatchmakeType
	presence.JoinGameID = JoinGameID
	presence.JoinGameMode = JoinGameMode
	presence.OwnerPID = OwnerPID
	presence.JoinGroupID = JoinGroupID
	presence.ApplicationArg = ApplicationArg

	return nil
}

// NewNintendoPresence returns a new NintendoPresence
func NewNintendoPresence() *NintendoPresence {
	return &NintendoPresence{}
}

// FriendPresence contains information about a users online presence
type FriendPresence struct {
	PID      uint32
	Presence *NintendoPresence

	nex.Structure
}

// Bytes encodes the FriendPresence and returns a byte array
func (presence *FriendPresence) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(presence.PID)
	stream.WriteStructure(presence.Presence)

	return stream.Bytes()
}

// NewFriendPresence returns a new FriendPresence
func NewFriendPresence() *FriendPresence {
	return &FriendPresence{}
}

// FriendRelationship contains information about a users relationship with another PID
type FriendRelationship struct {
	PID              uint32
	LFC              uint64
	RelationshipType uint8

	nex.Structure
}

// Bytes encodes the FriendRelationship and returns a byte array
func (relationship *FriendRelationship) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(relationship.PID)
	stream.WriteUInt64LE(relationship.LFC)
	stream.WriteUInt8(relationship.RelationshipType)

	return stream.Bytes()
}

// NewFriendRelationship returns a new FriendRelationship
func NewFriendRelationship() *FriendRelationship {
	return &FriendRelationship{}
}

// FriendPersistentInfo contains user settings
type FriendPersistentInfo struct {
	PID              uint32
	Region           uint8
	Country          uint8
	Area             uint8
	Language         uint8
	Platform         uint8
	GameKey          *GameKey
	Message          string
	MessageUpdatedAt *nex.DateTime //appears to be correct, but not 100% sure.
	FriendedAt       *nex.DateTime
	LastOnline       *nex.DateTime

	nex.Structure
}

// Bytes encodes the FriendPersistentInfo and returns a byte array
func (friendPersistentInfo *FriendPersistentInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(friendPersistentInfo.PID)
	stream.WriteUInt8(friendPersistentInfo.Region)
	stream.WriteUInt8(friendPersistentInfo.Country)
	stream.WriteUInt8(friendPersistentInfo.Area)
	stream.WriteUInt8(friendPersistentInfo.Language)
	stream.WriteUInt8(friendPersistentInfo.Platform)
	stream.WriteStructure(friendPersistentInfo.GameKey)
	stream.WriteString(friendPersistentInfo.Message)
	stream.WriteDateTime(friendPersistentInfo.MessageUpdatedAt)
	stream.WriteDateTime(friendPersistentInfo.FriendedAt)
	stream.WriteDateTime(friendPersistentInfo.LastOnline)

	return stream.Bytes()
}

// NewFriendPersistentInfo returns a new FriendPersistentInfo
func NewFriendPersistentInfo() *FriendPersistentInfo {
	return &FriendPersistentInfo{}
}

// GameKey contains the title ID and version for a title
type GameKey struct {
	TitleID      uint64
	TitleVersion uint16

	nex.Structure
}

// Bytes encodes the GameKey and returns a byte array
func (gameKey *GameKey) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(gameKey.TitleID)
	stream.WriteUInt16LE(gameKey.TitleVersion)

	return stream.Bytes()
}

// ExtractFromStream extracts a GameKey structure from a stream
func (gameKey *GameKey) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 10 {
		return errors.New("[GameKey::ExtractFromStream] Data size too small")
	}

	gameKey.TitleID = stream.ReadUInt64LE()
	gameKey.TitleVersion = stream.ReadUInt16LE()

	return nil
}

// NewGameKey returns a new GameKey
func NewGameKey() *GameKey {
	return &GameKey{}
}
