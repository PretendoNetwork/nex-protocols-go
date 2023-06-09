package friends_wiiu

import (
	"bytes"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

// BlacklistedPrincipal contains information about a blocked user
type BlacklistedPrincipal struct {
	nex.Structure
	PrincipalBasicInfo *PrincipalBasicInfo
	GameKey            *GameKey
	BlackListedSince   *nex.DateTime
}

// Bytes encodes the BlacklistedPrincipal and returns a byte array
func (blacklistedPrincipal *BlacklistedPrincipal) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(blacklistedPrincipal.PrincipalBasicInfo)
	stream.WriteStructure(blacklistedPrincipal.GameKey)
	stream.WriteDateTime(blacklistedPrincipal.BlackListedSince)

	return stream.Bytes()
}

// ExtractFromStream extracts a BlacklistedPrincipal structure from a stream
func (blacklistedPrincipal *BlacklistedPrincipal) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	principalBasicInfo, err := stream.ReadStructure(NewPrincipalBasicInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal.PrincipalBasicInfo. %s", err.Error())
	}

	blacklistedPrincipal.PrincipalBasicInfo = principalBasicInfo.(*PrincipalBasicInfo)
	gameKey, err := stream.ReadStructure(NewGameKey())
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal.GameKey. %s", err.Error())
	}

	blacklistedPrincipal.GameKey = gameKey.(*GameKey)
	blacklistedPrincipal.BlackListedSince, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract BlacklistedPrincipal.BlackListedSince. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BlacklistedPrincipal
func (blacklistedPrincipal *BlacklistedPrincipal) Copy() nex.StructureInterface {
	copied := NewBlacklistedPrincipal()

	copied.PrincipalBasicInfo = blacklistedPrincipal.PrincipalBasicInfo.Copy().(*PrincipalBasicInfo)
	copied.GameKey = blacklistedPrincipal.GameKey.Copy().(*GameKey)
	copied.BlackListedSince = blacklistedPrincipal.BlackListedSince.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (blacklistedPrincipal *BlacklistedPrincipal) Equals(structure nex.StructureInterface) bool {
	other := structure.(*BlacklistedPrincipal)

	if !blacklistedPrincipal.PrincipalBasicInfo.Equals(other.PrincipalBasicInfo) {
		return false
	}

	if !blacklistedPrincipal.GameKey.Equals(other.GameKey) {
		return false
	}

	if !blacklistedPrincipal.BlackListedSince.Equals(other.BlackListedSince) {
		return false
	}

	return true
}

// NewBlacklistedPrincipal returns a new BlacklistedPrincipal
func NewBlacklistedPrincipal() *BlacklistedPrincipal {
	return &BlacklistedPrincipal{}
}

// Comment contains data about a text comment
type Comment struct {
	nex.Structure
	Unknown     uint8
	Contents    string
	LastChanged *nex.DateTime
}

// Bytes encodes the Comment and returns a byte array
func (comment *Comment) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(comment.Unknown)
	stream.WriteString(comment.Contents)
	stream.WriteDateTime(comment.LastChanged)

	return stream.Bytes()
}

// ExtractFromStream extracts a Comment structure from a stream
func (comment *Comment) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	comment.Unknown, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Comment.Unknown. %s", err.Error())
	}

	comment.Contents, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract Comment.Contents. %s", err.Error())
	}

	comment.LastChanged, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract Comment.LastChanged. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Comment
func (comment *Comment) Copy() nex.StructureInterface {
	copied := NewComment()

	copied.Unknown = comment.Unknown
	copied.Contents = comment.Contents
	copied.LastChanged = comment.LastChanged.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (comment *Comment) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Comment)

	if comment.Unknown != other.Unknown {
		return false
	}

	if comment.Contents != other.Contents {
		return false
	}

	if !comment.LastChanged.Equals(other.LastChanged) {
		return false
	}

	return true
}

// NewComment returns a new Comment
func NewComment() *Comment {
	return &Comment{}
}

// FriendInfo contains information about a friend
type FriendInfo struct {
	nex.Structure
	NNAInfo      *NNAInfo
	Presence     *NintendoPresenceV2
	Status       *Comment
	BecameFriend *nex.DateTime
	LastOnline   *nex.DateTime
	Unknown      uint64
}

// Bytes encodes the FriendInfo and returns a byte array
func (friendInfo *FriendInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(friendInfo.NNAInfo)
	stream.WriteStructure(friendInfo.Presence)
	stream.WriteStructure(friendInfo.Status)
	stream.WriteDateTime(friendInfo.BecameFriend)
	stream.WriteDateTime(friendInfo.LastOnline)
	stream.WriteUInt64LE(friendInfo.Unknown)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendInfo
func (friendInfo *FriendInfo) Copy() nex.StructureInterface {
	copied := NewFriendInfo()

	copied.NNAInfo = friendInfo.NNAInfo.Copy().(*NNAInfo)
	copied.Presence = friendInfo.Presence.Copy().(*NintendoPresenceV2)
	copied.Status = friendInfo.Status.Copy().(*Comment)
	copied.BecameFriend = friendInfo.BecameFriend.Copy()
	copied.LastOnline = friendInfo.LastOnline.Copy()
	copied.Unknown = friendInfo.Unknown

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendInfo *FriendInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendInfo)

	if !friendInfo.NNAInfo.Equals(other.NNAInfo) {
		return false
	}

	if !friendInfo.Presence.Equals(other.Presence) {
		return false
	}

	if !friendInfo.Status.Equals(other.Status) {
		return false
	}

	if !friendInfo.BecameFriend.Equals(other.BecameFriend) {
		return false
	}

	if !friendInfo.LastOnline.Equals(other.LastOnline) {
		return false
	}

	if friendInfo.Unknown != other.Unknown {
		return false
	}

	return true
}

// NewFriendInfo returns a new FriendInfo
func NewFriendInfo() *FriendInfo {
	return &FriendInfo{}
}

// FriendRequest contains information about a friend request
type FriendRequest struct {
	nex.Structure
	PrincipalInfo *PrincipalBasicInfo
	Message       *FriendRequestMessage
	SentOn        *nex.DateTime
}

// Bytes encodes the FriendRequest and returns a byte array
func (friendRequest *FriendRequest) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(friendRequest.PrincipalInfo)
	stream.WriteStructure(friendRequest.Message)
	stream.WriteDateTime(friendRequest.SentOn)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendRequest
func (friendRequest *FriendRequest) Copy() nex.StructureInterface {
	copied := NewFriendRequest()

	copied.PrincipalInfo = friendRequest.PrincipalInfo.Copy().(*PrincipalBasicInfo)
	copied.Message = friendRequest.Message.Copy().(*FriendRequestMessage)
	copied.SentOn = friendRequest.SentOn.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendRequest *FriendRequest) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendRequest)

	if !friendRequest.PrincipalInfo.Equals(other.PrincipalInfo) {
		return false
	}

	if !friendRequest.Message.Equals(other.Message) {
		return false
	}

	if !friendRequest.SentOn.Equals(other.SentOn) {
		return false
	}

	return true
}

// NewFriendRequest returns a new FriendRequest
func NewFriendRequest() *FriendRequest {
	return &FriendRequest{}
}

// FriendRequestMessage contains message data for a FriendRequest
type FriendRequestMessage struct {
	nex.Structure
	FriendRequestID uint64
	Received        bool
	Unknown2        uint8
	Message         string
	Unknown3        uint8
	Unknown4        string
	GameKey         *GameKey
	Unknown5        *nex.DateTime
	ExpiresOn       *nex.DateTime
}

// Bytes encodes the FriendRequestMessage and returns a byte array
func (friendRequestMessage *FriendRequestMessage) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(friendRequestMessage.FriendRequestID)
	stream.WriteBool(friendRequestMessage.Received)
	stream.WriteUInt8(friendRequestMessage.Unknown2)
	stream.WriteString(friendRequestMessage.Message)
	stream.WriteUInt8(friendRequestMessage.Unknown3)
	stream.WriteString(friendRequestMessage.Unknown4)
	stream.WriteStructure(friendRequestMessage.GameKey)
	stream.WriteDateTime(friendRequestMessage.Unknown5)
	stream.WriteDateTime(friendRequestMessage.ExpiresOn)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendRequestMessage
func (friendRequestMessage *FriendRequestMessage) Copy() nex.StructureInterface {
	copied := NewFriendRequestMessage()

	copied.FriendRequestID = friendRequestMessage.FriendRequestID
	copied.Received = friendRequestMessage.Received
	copied.Unknown2 = friendRequestMessage.Unknown2
	copied.Message = friendRequestMessage.Message
	copied.Unknown3 = friendRequestMessage.Unknown3
	copied.Unknown4 = friendRequestMessage.Unknown4
	copied.GameKey = friendRequestMessage.GameKey.Copy().(*GameKey)
	copied.Unknown5 = friendRequestMessage.Unknown5.Copy()
	copied.ExpiresOn = friendRequestMessage.ExpiresOn.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendRequestMessage *FriendRequestMessage) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendRequestMessage)

	if friendRequestMessage.FriendRequestID != other.FriendRequestID {
		return false
	}

	if friendRequestMessage.Received != other.Received {
		return false
	}

	if friendRequestMessage.Unknown2 != other.Unknown2 {
		return false
	}

	if friendRequestMessage.Message != other.Message {
		return false
	}

	if friendRequestMessage.Unknown3 != other.Unknown3 {
		return false
	}

	if friendRequestMessage.Unknown4 != other.Unknown4 {
		return false
	}

	if !friendRequestMessage.GameKey.Equals(other.GameKey) {
		return false
	}

	if !friendRequestMessage.Unknown5.Equals(other.Unknown5) {
		return false
	}

	if !friendRequestMessage.ExpiresOn.Equals(other.ExpiresOn) {
		return false
	}

	return true
}

// NewFriendRequestMessage returns a new FriendRequestMessage
func NewFriendRequestMessage() *FriendRequestMessage {
	return &FriendRequestMessage{}
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
	var err error

	gameKey.TitleID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GameKey.TitleID. %s", err.Error())
	}

	gameKey.TitleVersion, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GameKey.TitleVersion. %s", err.Error())
	}

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

// MiiV2 contains data about a Mii
type MiiV2 struct {
	nex.Structure
	Name     string
	Unknown1 uint8
	Unknown2 uint8
	Data     []byte
	Datetime *nex.DateTime
}

// Bytes encodes the MiiV2 and returns a byte array
func (mii *MiiV2) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(mii.Name)
	stream.WriteUInt8(mii.Unknown1)
	stream.WriteUInt8(mii.Unknown2)
	stream.WriteBuffer(mii.Data)
	stream.WriteDateTime(mii.Datetime)

	return stream.Bytes()
}

// ExtractFromStream extracts a MiiV2 structure from a stream
func (mii *MiiV2) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	mii.Name, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Name. %s", err.Error())
	}

	mii.Unknown1, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Unknown1. %s", err.Error())
	}

	mii.Unknown2, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Unknown2. %s", err.Error())
	}

	mii.Data, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Data. %s", err.Error())
	}

	mii.Datetime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Datetime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MiiV2
func (mii *MiiV2) Copy() nex.StructureInterface {
	copied := NewMiiV2()

	copied.Name = mii.Name
	copied.Unknown1 = mii.Unknown1
	copied.Unknown2 = mii.Unknown2
	copied.Data = make([]byte, len(mii.Data))

	copy(copied.Data, mii.Data)

	copied.Datetime = mii.Datetime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (mii *MiiV2) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MiiV2)

	if mii.Name != other.Name {
		return false
	}

	if mii.Unknown1 != other.Unknown1 {
		return false
	}

	if mii.Unknown2 != other.Unknown2 {
		return false
	}

	if !bytes.Equal(mii.Data, other.Data) {
		return false
	}

	if !mii.Datetime.Equals(other.Datetime) {
		return false
	}

	return true
}

// NewMiiV2 returns a new MiiV2
func NewMiiV2() *MiiV2 {
	return &MiiV2{}
}

// NintendoPresenceV2 contains information about a users online presence
type NintendoPresenceV2 struct {
	nex.Structure
	ChangedFlags    uint32
	Online          bool
	GameKey         *GameKey
	Unknown1        uint8
	Message         string
	Unknown2        uint32
	Unknown3        uint8
	GameServerID    uint32
	Unknown4        uint32
	PID             uint32
	GatheringID     uint32
	ApplicationData []byte
	Unknown5        uint8
	Unknown6        uint8
	Unknown7        uint8
}

// Bytes encodes the NintendoPresenceV2 and returns a byte array
func (presence *NintendoPresenceV2) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(presence.ChangedFlags)
	if presence.Online {
		stream.WriteUInt8(1)
	} else {
		stream.WriteUInt8(0)
	}
	stream.WriteStructure(presence.GameKey)
	stream.WriteUInt8(presence.Unknown1)
	stream.WriteString(presence.Message)
	stream.WriteUInt32LE(presence.Unknown2)
	stream.WriteUInt8(presence.Unknown3)
	stream.WriteUInt32LE(presence.GameServerID)
	stream.WriteUInt32LE(presence.Unknown4)
	stream.WriteUInt32LE(presence.PID)
	stream.WriteUInt32LE(presence.GatheringID)
	stream.WriteBuffer(presence.ApplicationData)
	stream.WriteUInt8(presence.Unknown5)
	stream.WriteUInt8(presence.Unknown6)
	stream.WriteUInt8(presence.Unknown7)

	return stream.Bytes()
}

// ExtractFromStream extracts a NintendoPresenceV2 structure from a stream
func (presence *NintendoPresenceV2) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	presence.ChangedFlags, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.ChangedFlags. %s", err.Error())
	}

	presence.Online, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Online. %s", err.Error())
	}

	gameKey, err := stream.ReadStructure(NewGameKey())
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.GameKey. %s", err.Error())
	}

	presence.GameKey = gameKey.(*GameKey)
	presence.Unknown1, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown1. %s", err.Error())
	}

	presence.Message, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Message. %s", err.Error())
	}

	presence.Unknown2, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown2. %s", err.Error())
	}

	presence.Unknown3, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown3. %s", err.Error())
	}

	presence.GameServerID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.GameServerID. %s", err.Error())
	}

	presence.Unknown4, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown4. %s", err.Error())
	}

	presence.PID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.PID. %s", err.Error())
	}

	presence.GatheringID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.GatheringID. %s", err.Error())
	}

	presence.ApplicationData, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.ApplicationData. %s", err.Error())
	}

	presence.Unknown5, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown5. %s", err.Error())
	}

	presence.Unknown6, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown6. %s", err.Error())
	}

	presence.Unknown7, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown7. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoPresenceV2
func (presence *NintendoPresenceV2) Copy() nex.StructureInterface {
	copied := NewNintendoPresenceV2()

	copied.ChangedFlags = presence.ChangedFlags
	copied.Online = presence.Online
	copied.GameKey = presence.GameKey.Copy().(*GameKey)
	copied.Unknown1 = presence.Unknown1
	copied.Message = presence.Message
	copied.Unknown2 = presence.Unknown2
	copied.Unknown3 = presence.Unknown3
	copied.GameServerID = presence.GameServerID
	copied.Unknown4 = presence.Unknown4
	copied.PID = presence.PID
	copied.GatheringID = presence.GatheringID
	copied.ApplicationData = make([]byte, len(presence.ApplicationData))

	copy(copied.ApplicationData, presence.ApplicationData)

	copied.Unknown5 = presence.Unknown5
	copied.Unknown6 = presence.Unknown6
	copied.Unknown7 = presence.Unknown7

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (presence *NintendoPresenceV2) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NintendoPresenceV2)

	if presence.ChangedFlags != other.ChangedFlags {
		return false
	}

	if presence.Online != other.Online {
		return false
	}

	if !presence.GameKey.Equals(other.GameKey) {
		return false
	}

	if presence.Unknown1 != other.Unknown1 {
		return false
	}

	if presence.Message != other.Message {
		return false
	}

	if presence.Unknown2 != other.Unknown2 {
		return false
	}

	if presence.Unknown3 != other.Unknown3 {
		return false
	}

	if presence.GameServerID != other.GameServerID {
		return false
	}

	if presence.Unknown4 != other.Unknown4 {
		return false
	}

	if presence.PID != other.PID {
		return false
	}

	if presence.GatheringID != other.GatheringID {
		return false
	}

	if !bytes.Equal(presence.ApplicationData, other.ApplicationData) {
		return false
	}

	if presence.Unknown5 != other.Unknown5 {
		return false
	}

	if presence.Unknown6 != other.Unknown6 {
		return false
	}

	if presence.Unknown7 != other.Unknown7 {
		return false
	}

	return true
}

// NewNintendoPresenceV2 returns a new NintendoPresenceV2
func NewNintendoPresenceV2() *NintendoPresenceV2 {
	return &NintendoPresenceV2{}
}

// NNAInfo contains information about a Nintendo Network Account
type NNAInfo struct {
	nex.Structure
	PrincipalBasicInfo *PrincipalBasicInfo
	Unknown1           uint8
	Unknown2           uint8
}

// Bytes encodes the NNAInfo and returns a byte array
func (nnaInfo *NNAInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(nnaInfo.PrincipalBasicInfo)
	stream.WriteUInt8(nnaInfo.Unknown1)
	stream.WriteUInt8(nnaInfo.Unknown2)

	return stream.Bytes()
}

// ExtractFromStream extracts a NNAInfo structure from a stream
func (nnaInfo *NNAInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	principalBasicInfo, err := stream.ReadStructure(NewPrincipalBasicInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.PrincipalBasicInfo. %s", err.Error())
	}

	nnaInfo.PrincipalBasicInfo = principalBasicInfo.(*PrincipalBasicInfo)
	nnaInfo.Unknown1, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.Unknown1. %s", err.Error())
	}

	nnaInfo.Unknown2, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.Unknown2. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NNAInfo
func (nnaInfo *NNAInfo) Copy() nex.StructureInterface {
	copied := NewNNAInfo()

	copied.PrincipalBasicInfo = nnaInfo.PrincipalBasicInfo.Copy().(*PrincipalBasicInfo)
	copied.Unknown1 = nnaInfo.Unknown1
	copied.Unknown2 = nnaInfo.Unknown2

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nnaInfo *NNAInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NNAInfo)

	if !nnaInfo.PrincipalBasicInfo.Equals(other.PrincipalBasicInfo) {
		return false
	}

	if nnaInfo.Unknown1 != other.Unknown1 {
		return false
	}

	if nnaInfo.Unknown2 != other.Unknown2 {
		return false
	}

	return true
}

// NewNNAInfo returns a new NNAInfo
func NewNNAInfo() *NNAInfo {
	return &NNAInfo{}
}

// PersistentNotification contains unknown data
type PersistentNotification struct {
	nex.Structure
	Unknown1 uint64
	Unknown2 uint32
	Unknown3 uint32
	Unknown4 uint32
	Unknown5 string
}

// ExtractFromStream extracts a PersistentNotification structure from a stream
func (notification *PersistentNotification) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	notification.Unknown1, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown1. %s", err.Error())
	}

	notification.Unknown2, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown2. %s", err.Error())
	}

	notification.Unknown3, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown3. %s", err.Error())
	}

	notification.Unknown4, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown4. %s", err.Error())
	}

	notification.Unknown5, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract PersistentNotification.Unknown5. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PersistentNotification
func (notification *PersistentNotification) Copy() nex.StructureInterface {
	copied := NewPersistentNotification()

	copied.Unknown1 = notification.Unknown1
	copied.Unknown2 = notification.Unknown2
	copied.Unknown3 = notification.Unknown3
	copied.Unknown4 = notification.Unknown4
	copied.Unknown5 = notification.Unknown5

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (notification *PersistentNotification) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PersistentNotification)

	if notification.Unknown1 != other.Unknown1 {
		return false
	}

	if notification.Unknown2 != other.Unknown2 {
		return false
	}

	if notification.Unknown3 != other.Unknown3 {
		return false
	}

	if notification.Unknown4 != other.Unknown4 {
		return false
	}

	if notification.Unknown5 != other.Unknown5 {
		return false
	}

	return true
}

// NewPersistentNotification returns a new PersistentNotification
func NewPersistentNotification() *PersistentNotification {
	return &PersistentNotification{}
}

// PrincipalBasicInfo contains user account and Mii data
type PrincipalBasicInfo struct {
	nex.Structure
	PID     uint32
	NNID    string
	Mii     *MiiV2
	Unknown uint8
}

// Bytes encodes the PrincipalBasicInfo and returns a byte array
func (principalInfo *PrincipalBasicInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(principalInfo.PID)
	stream.WriteString(principalInfo.NNID)
	stream.WriteStructure(principalInfo.Mii)
	stream.WriteUInt8(principalInfo.Unknown)

	return stream.Bytes()
}

// ExtractFromStream extracts a PrincipalBasicInfo structure from a stream
func (principalInfo *PrincipalBasicInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	principalInfo.PID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.PID. %s", err.Error())
	}

	principalInfo.NNID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.NNID. %s", err.Error())
	}

	miiV2, err := stream.ReadStructure(NewMiiV2())
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.Mii. %s", err.Error())
	}

	principalInfo.Mii = miiV2.(*MiiV2)
	principalInfo.Unknown, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.Unknown. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PrincipalBasicInfo
func (principalInfo *PrincipalBasicInfo) Copy() nex.StructureInterface {
	copied := NewPrincipalBasicInfo()

	copied.PID = principalInfo.PID
	copied.NNID = principalInfo.NNID
	copied.Mii = principalInfo.Mii.Copy().(*MiiV2)
	copied.Unknown = principalInfo.Unknown

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (principalInfo *PrincipalBasicInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PrincipalBasicInfo)

	if principalInfo.PID != other.PID {
		return false
	}

	if principalInfo.NNID != other.NNID {
		return false
	}

	if !principalInfo.Mii.Equals(other.Mii) {
		return false
	}

	if principalInfo.Unknown != other.Unknown {
		return false
	}

	return true
}

// NewPrincipalBasicInfo returns a new PrincipalBasicInfo
func NewPrincipalBasicInfo() *PrincipalBasicInfo {
	return &PrincipalBasicInfo{}
}

// PrincipalPreference contains unknown data
type PrincipalPreference struct {
	nex.Structure
	ShowOnlinePresence  bool
	ShowCurrentTitle    bool
	BlockFriendRequests bool
}

// Bytes encodes the PrincipalPreference and returns a byte array
func (principalPreference *PrincipalPreference) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteBool(principalPreference.ShowOnlinePresence)
	stream.WriteBool(principalPreference.ShowCurrentTitle)
	stream.WriteBool(principalPreference.BlockFriendRequests)

	return stream.Bytes()
}

// ExtractFromStream extracts a PrincipalPreference structure from a stream
func (principalPreference *PrincipalPreference) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	principalPreference.ShowOnlinePresence, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference.ShowOnlinePresence. %s", err.Error())
	}

	principalPreference.ShowCurrentTitle, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference.ShowCurrentTitle. %s", err.Error())
	}

	principalPreference.BlockFriendRequests, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalPreference.BlockFriendRequests. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PrincipalPreference
func (principalPreference *PrincipalPreference) Copy() nex.StructureInterface {
	copied := NewPrincipalPreference()

	copied.ShowOnlinePresence = principalPreference.ShowOnlinePresence
	copied.ShowCurrentTitle = principalPreference.ShowCurrentTitle
	copied.BlockFriendRequests = principalPreference.BlockFriendRequests

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (principalPreference *PrincipalPreference) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PrincipalPreference)

	if principalPreference.ShowOnlinePresence != other.ShowOnlinePresence {
		return false
	}

	if principalPreference.ShowCurrentTitle != other.ShowCurrentTitle {
		return false
	}

	if principalPreference.BlockFriendRequests != other.BlockFriendRequests {
		return false
	}

	return true
}

// NewPrincipalPreference returns a new PrincipalPreference
func NewPrincipalPreference() *PrincipalPreference {
	return &PrincipalPreference{}
}

// PrincipalRequestBlockSetting contains unknow data
type PrincipalRequestBlockSetting struct {
	nex.Structure
	PID       uint32
	IsBlocked bool
}

// Bytes encodes the PrincipalRequestBlockSetting and returns a byte array
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(principalRequestBlockSetting.PID)
	stream.WriteBool(principalRequestBlockSetting.IsBlocked)

	return stream.Bytes()
}

// Copy returns a new copied instance of PrincipalRequestBlockSetting
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) Copy() nex.StructureInterface {
	copied := NewPrincipalRequestBlockSetting()

	copied.PID = principalRequestBlockSetting.PID
	copied.IsBlocked = principalRequestBlockSetting.IsBlocked

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PrincipalRequestBlockSetting)

	if principalRequestBlockSetting.PID != other.PID {
		return false
	}

	if principalRequestBlockSetting.IsBlocked != other.IsBlocked {
		return false
	}

	return true
}

// NewPrincipalRequestBlockSetting returns a new PrincipalRequestBlockSetting
func NewPrincipalRequestBlockSetting() *PrincipalRequestBlockSetting {
	return &PrincipalRequestBlockSetting{}
}
