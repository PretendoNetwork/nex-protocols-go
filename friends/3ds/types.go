package friends_3ds

import (
	"bytes"
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
)

type Mii struct {
	nex.Structure
	Name     string
	Unknown2 bool
	Unknown3 uint8
	MiiData  []byte
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

// Copy returns a new copied instance of Mii
func (mii *Mii) Copy() nex.StructureInterface {
	copied := NewMii()

	copied.Name = mii.Name
	copied.Unknown2 = mii.Unknown2
	copied.Unknown3 = mii.Unknown3
	copied.MiiData = make([]byte, len(mii.MiiData))

	copy(copied.MiiData, mii.MiiData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (mii *Mii) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Mii)

	if mii.Name != other.Name {
		return false
	}

	if mii.Unknown2 != other.Unknown2 {
		return false
	}

	if mii.Unknown3 != other.Unknown3 {
		return false
	}

	if !bytes.Equal(mii.MiiData, other.MiiData) {
		return false
	}

	return true
}

// NewMii returns a new Mii
func NewMii() *Mii {
	return &Mii{}
}

type FriendMii struct {
	nex.Structure
	PID        uint32
	Mii        *Mii
	ModifiedAt *nex.DateTime
}

// Bytes encodes the Mii and returns a byte array
func (friendMii *FriendMii) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(friendMii.PID)
	stream.WriteStructure(friendMii.Mii)
	stream.WriteDateTime(friendMii.ModifiedAt)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendMii
func (friendMii *FriendMii) Copy() nex.StructureInterface {
	copied := NewFriendMii()

	copied.PID = friendMii.PID
	copied.Mii = friendMii.Mii.Copy().(*Mii)
	copied.ModifiedAt = friendMii.ModifiedAt.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendMii *FriendMii) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendMii)

	if friendMii.PID != other.PID {
		return false
	}

	if !friendMii.Mii.Equals(other.Mii) {
		return false
	}

	if !friendMii.ModifiedAt.Equals(other.ModifiedAt) {
		return false
	}

	return true
}

// NewMii returns a new Mii
func NewFriendMii() *FriendMii {
	return &FriendMii{}
}

type MyProfile struct {
	nex.Structure
	Region   uint8
	Country  uint8
	Area     uint8
	Language uint8
	Platform uint8
	Unknown1 uint64
	Unknown2 string
	Unknown3 string
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

// Copy returns a new copied instance of MyProfile
func (myProfile *MyProfile) Copy() nex.StructureInterface {
	copied := NewMyProfile()

	copied.Region = myProfile.Region
	copied.Country = myProfile.Country
	copied.Area = myProfile.Area
	copied.Language = myProfile.Language
	copied.Platform = myProfile.Platform
	copied.Unknown1 = myProfile.Unknown1
	copied.Unknown2 = myProfile.Unknown2
	copied.Unknown3 = myProfile.Unknown3

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (myProfile *MyProfile) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MyProfile)

	if myProfile.Region != other.Region {
		return false
	}

	if myProfile.Country != other.Country {
		return false
	}

	if myProfile.Area != other.Area {
		return false
	}

	if myProfile.Language != other.Language {
		return false
	}

	if myProfile.Platform != other.Platform {
		return false
	}

	if myProfile.Unknown1 != other.Unknown1 {
		return false
	}

	if myProfile.Unknown2 != other.Unknown2 {
		return false
	}

	if myProfile.Unknown3 != other.Unknown3 {
		return false
	}

	return true
}

// NewMyProfile returns a new MyProfile
func NewMyProfile() *MyProfile {
	return &MyProfile{}
}

// NintendoPresence contains information about a users online presence
type NintendoPresence struct {
	nex.Structure
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

// Copy returns a new copied instance of NintendoPresence
func (presence *NintendoPresence) Copy() nex.StructureInterface {
	copied := NewNintendoPresence()

	copied.ChangedFlags = presence.ChangedFlags
	copied.GameKey = presence.GameKey.Copy().(*GameKey)
	copied.Message = presence.Message
	copied.JoinAvailableFlag = presence.JoinAvailableFlag
	copied.MatchmakeType = presence.MatchmakeType
	copied.JoinGameID = presence.JoinGameID
	copied.JoinGameMode = presence.JoinGameMode
	copied.OwnerPID = presence.OwnerPID
	copied.JoinGroupID = presence.JoinGroupID
	copied.ApplicationArg = make([]byte, len(presence.ApplicationArg))

	copy(copied.ApplicationArg, presence.ApplicationArg)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (presence *NintendoPresence) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NintendoPresence)

	if presence.ChangedFlags != other.ChangedFlags {
		return false
	}

	if !presence.GameKey.Equals(other.GameKey) {
		return false
	}

	if presence.Message != other.Message {
		return false
	}

	if presence.JoinAvailableFlag != other.JoinAvailableFlag {
		return false
	}

	if presence.MatchmakeType != other.MatchmakeType {
		return false
	}

	if presence.JoinGameID != other.JoinGameID {
		return false
	}

	if presence.JoinGameMode != other.JoinGameMode {
		return false
	}

	if presence.OwnerPID != other.OwnerPID {
		return false
	}

	if presence.JoinGroupID != other.JoinGroupID {
		return false
	}

	if !bytes.Equal(presence.ApplicationArg, other.ApplicationArg) {
		return false
	}

	return true
}

// NewNintendoPresence returns a new NintendoPresence
func NewNintendoPresence() *NintendoPresence {
	return &NintendoPresence{}
}

// FriendPresence contains information about a users online presence
type FriendPresence struct {
	nex.Structure
	PID      uint32
	Presence *NintendoPresence
}

// Bytes encodes the FriendPresence and returns a byte array
func (presence *FriendPresence) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(presence.PID)
	stream.WriteStructure(presence.Presence)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendPresence
func (presence *FriendPresence) Copy() nex.StructureInterface {
	copied := NewFriendPresence()

	copied.PID = presence.PID
	copied.Presence = presence.Presence.Copy().(*NintendoPresence)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (presence *FriendPresence) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendPresence)

	if presence.PID != other.PID {
		return false
	}

	if !presence.Presence.Equals(other.Presence) {
		return false
	}

	return true
}

// NewFriendPresence returns a new FriendPresence
func NewFriendPresence() *FriendPresence {
	return &FriendPresence{}
}

// FriendRelationship contains information about a users relationship with another PID
type FriendRelationship struct {
	nex.Structure
	PID              uint32
	LFC              uint64
	RelationshipType uint8
}

// Bytes encodes the FriendRelationship and returns a byte array
func (relationship *FriendRelationship) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(relationship.PID)
	stream.WriteUInt64LE(relationship.LFC)
	stream.WriteUInt8(relationship.RelationshipType)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendRelationship
func (relationship *FriendRelationship) Copy() nex.StructureInterface {
	copied := NewFriendRelationship()

	copied.PID = relationship.PID
	copied.LFC = relationship.LFC
	copied.RelationshipType = relationship.RelationshipType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (relationship *FriendRelationship) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendRelationship)

	if relationship.PID != other.PID {
		return false
	}

	if relationship.LFC != other.LFC {
		return false
	}

	if relationship.RelationshipType != other.RelationshipType {
		return false
	}

	return true
}

// NewFriendRelationship returns a new FriendRelationship
func NewFriendRelationship() *FriendRelationship {
	return &FriendRelationship{}
}

// FriendPersistentInfo contains user settings
type FriendPersistentInfo struct {
	nex.Structure
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

// Copy returns a new copied instance of FriendPersistentInfo
func (friendPersistentInfo *FriendPersistentInfo) Copy() nex.StructureInterface {
	copied := NewFriendPersistentInfo()

	copied.PID = friendPersistentInfo.PID
	copied.Region = friendPersistentInfo.Region
	copied.Country = friendPersistentInfo.Country
	copied.Area = friendPersistentInfo.Area
	copied.Language = friendPersistentInfo.Language
	copied.Platform = friendPersistentInfo.Platform
	copied.GameKey = friendPersistentInfo.GameKey.Copy().(*GameKey)
	copied.Message = friendPersistentInfo.Message
	copied.MessageUpdatedAt = friendPersistentInfo.MessageUpdatedAt.Copy()
	copied.FriendedAt = friendPersistentInfo.FriendedAt.Copy()
	copied.LastOnline = friendPersistentInfo.LastOnline.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendPersistentInfo *FriendPersistentInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendPersistentInfo)

	if friendPersistentInfo.PID != other.PID {
		return false
	}

	if friendPersistentInfo.Region != other.Region {
		return false
	}

	if friendPersistentInfo.Country != other.Country {
		return false
	}

	if friendPersistentInfo.Area != other.Area {
		return false
	}

	if friendPersistentInfo.Language != other.Language {
		return false
	}

	if friendPersistentInfo.Platform != other.Platform {
		return false
	}

	if !friendPersistentInfo.GameKey.Equals(other.GameKey) {
		return false
	}

	if friendPersistentInfo.Message != other.Message {
		return false
	}

	if !friendPersistentInfo.MessageUpdatedAt.Equals(other.MessageUpdatedAt) {
		return false
	}

	if !friendPersistentInfo.FriendedAt.Equals(other.FriendedAt) {
		return false
	}

	if !friendPersistentInfo.LastOnline.Equals(other.LastOnline) {
		return false
	}

	return true
}

// NewFriendPersistentInfo returns a new FriendPersistentInfo
func NewFriendPersistentInfo() *FriendPersistentInfo {
	return &FriendPersistentInfo{}
}

// GameKey contains the title ID and version for a title
type GameKey struct {
	nex.Structure
	TitleID      uint64
	TitleVersion uint16
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

// Copy returns a new copied instance of GameKey
func (gameKey *GameKey) Copy() nex.StructureInterface {
	copied := NewGameKey()

	copied.TitleID = gameKey.TitleID
	copied.TitleVersion = gameKey.TitleVersion

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (gameKey *GameKey) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GameKey)

	if gameKey.TitleID != other.TitleID {
		return false
	}

	if gameKey.TitleVersion != other.TitleVersion {
		return false
	}

	return true
}

// NewGameKey returns a new GameKey
func NewGameKey() *GameKey {
	return &GameKey{}
}
