package friends_wiiu

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
)

// BlacklistedPrincipal contains information about a blocked user
type BlacklistedPrincipal struct {
	PrincipalBasicInfo *PrincipalBasicInfo
	GameKey            *GameKey
	BlackListedSince   *nex.DateTime

	nex.Structure
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
	principalBasicInfoStructureInterface, err := stream.ReadStructure(NewPrincipalBasicInfo())
	if err != nil {
		return err
	}

	gameKeyStructureInterface, err := stream.ReadStructure(NewGameKey())
	if err != nil {
		return err
	}

	if len(stream.Bytes()[stream.ByteOffset():]) < 8 {
		return errors.New("[DataStorePersistenceTarget::ExtractFromStream] Data size too small")
	}

	principalBasicInfo := principalBasicInfoStructureInterface.(*PrincipalBasicInfo)
	gameKey := gameKeyStructureInterface.(*GameKey)
	blackListedSince := nex.NewDateTime(stream.ReadUInt64LE())

	blacklistedPrincipal.PrincipalBasicInfo = principalBasicInfo
	blacklistedPrincipal.GameKey = gameKey
	blacklistedPrincipal.BlackListedSince = blackListedSince

	return nil
}

// NewBlacklistedPrincipal returns a new BlacklistedPrincipal
func NewBlacklistedPrincipal() *BlacklistedPrincipal {
	return &BlacklistedPrincipal{}
}

// Comment contains data about a text comment
type Comment struct {
	Unknown     uint8
	Contents    string
	LastChanged *nex.DateTime

	nex.Structure
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
	if len(stream.Bytes()[stream.ByteOffset():]) < 9 { // unknown byte + datetime uint64
		return errors.New("[Comment::ExtractFromStream] Data size too small")
	}

	unknown := stream.ReadUInt8()
	contents, err := stream.ReadString()

	if err != nil {
		return err
	}

	lastChanged := nex.NewDateTime(stream.ReadUInt64LE())

	comment.Unknown = unknown
	comment.Contents = contents
	comment.LastChanged = lastChanged

	return nil
}

// NewComment returns a new Comment
func NewComment() *Comment {
	return &Comment{}
}

// FriendInfo contains information about a friend
type FriendInfo struct {
	NNAInfo      *NNAInfo
	Presence     *NintendoPresenceV2
	Status       *Comment
	BecameFriend *nex.DateTime
	LastOnline   *nex.DateTime
	Unknown      uint64

	nex.Structure
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

// NewFriendInfo returns a new FriendInfo
func NewFriendInfo() *FriendInfo {
	return &FriendInfo{}
}

// FriendRequest contains information about a friend request
type FriendRequest struct {
	PrincipalInfo *PrincipalBasicInfo
	Message       *FriendRequestMessage
	SentOn        *nex.DateTime

	nex.Structure
}

// Bytes encodes the FriendRequest and returns a byte array
func (friendRequest *FriendRequest) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(friendRequest.PrincipalInfo)
	stream.WriteStructure(friendRequest.Message)
	stream.WriteDateTime(friendRequest.SentOn)

	return stream.Bytes()
}

// NewFriendRequest returns a new FriendRequest
func NewFriendRequest() *FriendRequest {
	return &FriendRequest{}
}

// FriendRequestMessage contains message data for a FriendRequest
type FriendRequestMessage struct {
	FriendRequestID uint64
	Received        bool
	Unknown2        uint8
	Message         string
	Unknown3        uint8
	Unknown4        string
	GameKey         *GameKey
	Unknown5        *nex.DateTime
	ExpiresOn       *nex.DateTime

	nex.Structure
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

// NewFriendRequestMessage returns a new FriendRequestMessage
func NewFriendRequestMessage() *FriendRequestMessage {
	return &FriendRequestMessage{}
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

// MiiV2 contains data about a Mii
type MiiV2 struct {
	Name     string
	Unknown1 uint8
	Unknown2 uint8
	Data     []byte
	Datetime *nex.DateTime

	nex.Structure
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
	if len(stream.Bytes()[stream.ByteOffset():]) < 10 { // 2 unknown bytes + datetime uint64
		return errors.New("[MiiV2::ExtractFromStream] Data size too small")
	}

	name, err := stream.ReadString()

	if err != nil {
		return err
	}

	unknown1 := stream.ReadUInt8()
	unknown2 := stream.ReadUInt8()
	data, err := stream.ReadBuffer()

	if err != nil {
		return err
	}

	datetime := nex.NewDateTime(stream.ReadUInt64LE())

	mii.Name = name
	mii.Unknown1 = unknown1
	mii.Unknown2 = unknown2
	mii.Data = data
	mii.Datetime = datetime

	return nil
}

// NewMiiV2 returns a new MiiV2
func NewMiiV2() *MiiV2 {
	return &MiiV2{}
}

// NintendoPresenceV2 contains information about a users online presence
type NintendoPresenceV2 struct {
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

	nex.Structure
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
	if len(stream.Bytes()[stream.ByteOffset():]) < 40 {
		// length check for the following fixed-size data
		// changedFlags + isOnline + gameKey + gameKey + unknown1 + unknown2 + unknown3 + gameServerID + unknown4 + pid + gatheringID + unknown5 + unknown6 + unknown7
		return errors.New("[NintendoPresenceV2::ExtractFromStream] Data size too small")
	}

	changedFlags := stream.ReadUInt32LE()
	Online := (stream.ReadUInt8() == 1)
	gameKeyStructureInterface, err := stream.ReadStructure(NewGameKey())
	if err != nil {
		return err
	}
	gameKey := gameKeyStructureInterface.(*GameKey)
	unknown1 := stream.ReadUInt8()
	message, err := stream.ReadString()
	if err != nil {
		return err
	}
	unknown2 := stream.ReadUInt32LE()
	unknown3 := stream.ReadUInt8()
	gameServerID := stream.ReadUInt32LE()
	unknown4 := stream.ReadUInt32LE()
	pid := stream.ReadUInt32LE()
	gatheringID := stream.ReadUInt32LE()
	applicationData, err := stream.ReadBuffer()
	if err != nil {
		return err
	}
	unknown5 := stream.ReadUInt8()
	unknown6 := stream.ReadUInt8()
	unknown7 := stream.ReadUInt8()

	presence.ChangedFlags = changedFlags
	presence.Online = Online
	presence.GameKey = gameKey
	presence.Unknown1 = unknown1
	presence.Message = message
	presence.Unknown2 = unknown2
	presence.Unknown3 = unknown3
	presence.GameServerID = gameServerID
	presence.Unknown4 = unknown4
	presence.PID = pid
	presence.GatheringID = gatheringID
	presence.ApplicationData = applicationData
	presence.Unknown5 = unknown5
	presence.Unknown6 = unknown6
	presence.Unknown7 = unknown7

	return nil
}

// NewNintendoPresenceV2 returns a new NintendoPresenceV2
func NewNintendoPresenceV2() *NintendoPresenceV2 {
	return &NintendoPresenceV2{}
}

// NNAInfo contains information about a Nintendo Network Account
type NNAInfo struct {
	PrincipalBasicInfo *PrincipalBasicInfo
	Unknown1           uint8
	Unknown2           uint8

	nex.Structure
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
	if len(stream.Bytes()[stream.ByteOffset():]) < 2 {
		// length check for the following fixed-size data
		// unknown1 + unknown2
		return errors.New("[NNAInfo::ExtractFromStream] Data size too small")
	}

	principalBasicInfoStructureInterface, err := stream.ReadStructure(NewPrincipalBasicInfo())
	if err != nil {
		return err
	}

	principalBasicInfo := principalBasicInfoStructureInterface.(*PrincipalBasicInfo)
	unknown1 := stream.ReadUInt8()
	unknown2 := stream.ReadUInt8()

	nnaInfo.PrincipalBasicInfo = principalBasicInfo
	nnaInfo.Unknown1 = unknown1
	nnaInfo.Unknown2 = unknown2

	return nil
}

// NewNNAInfo returns a new NNAInfo
func NewNNAInfo() *NNAInfo {
	return &NNAInfo{}
}

// PersistentNotification contains unknown data
type PersistentNotification struct {
	Unknown1 uint64
	Unknown2 uint32
	Unknown3 uint32
	Unknown4 uint32
	Unknown5 string

	nex.Structure
}

// ExtractFromStream extracts a PersistentNotification structure from a stream
func (notification *PersistentNotification) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 20 {
		// length check for the following fixed-size data
		// unknown1 + unknown2 + unknown3 + unknown4
		return errors.New("[PersistentNotification::ExtractFromStream] Data size too small")
	}

	unknown1 := stream.ReadUInt64LE()
	unknown2 := stream.ReadUInt32LE()
	unknown3 := stream.ReadUInt32LE()
	unknown4 := stream.ReadUInt32LE()
	unknown5, err := stream.ReadString()
	if err != nil {
		return err
	}

	notification.Unknown1 = unknown1
	notification.Unknown2 = unknown2
	notification.Unknown3 = unknown3
	notification.Unknown4 = unknown4
	notification.Unknown5 = unknown5

	return nil
}

// NewPersistentNotification returns a new PersistentNotification
func NewPersistentNotification() *PersistentNotification {
	return &PersistentNotification{}
}

// PrincipalBasicInfo contains user account and Mii data
type PrincipalBasicInfo struct {
	PID     uint32
	NNID    string
	Mii     *MiiV2
	Unknown uint8

	nex.Structure
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

	if len(stream.Bytes()[stream.ByteOffset():]) < 4 {
		return errors.New("[PrincipalBasicInfo::ExtractFromStream] Data size too small")
	}

	pid := stream.ReadUInt32LE()
	nnid, err := stream.ReadString()

	if err != nil {
		return err
	}

	miiV2StructureInterface, err := stream.ReadStructure(NewMiiV2())
	if err != nil {
		return err
	}
	miiV2 := miiV2StructureInterface.(*MiiV2)

	if len(stream.Bytes()[stream.ByteOffset():]) < 1 {
		return errors.New("[PrincipalBasicInfo::ExtractFromStream] Data size too small")
	}

	unknown := stream.ReadUInt8()

	principalInfo.PID = pid
	principalInfo.NNID = nnid
	principalInfo.Mii = miiV2
	principalInfo.Unknown = unknown

	return nil
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
	if len(stream.Bytes()[stream.ByteOffset():]) < 1 {
		// length check for the following fixed-size data
		// unknown1 + unknown2 + unknown3
		return errors.New("[PrincipalPreference::ExtractFromStream] Data size too small")
	}

	principalPreference.ShowOnlinePresence = (stream.ReadUInt8() == 1)
	principalPreference.ShowCurrentTitle = (stream.ReadUInt8() == 1)
	principalPreference.BlockFriendRequests = (stream.ReadUInt8() == 1)

	return nil
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

// NewPrincipalRequestBlockSetting returns a new PrincipalRequestBlockSetting
func NewPrincipalRequestBlockSetting() *PrincipalRequestBlockSetting {
	return &PrincipalRequestBlockSetting{}
}
