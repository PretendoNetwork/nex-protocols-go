package nexproto

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// FriendsWiiUProtocolID is the protocol ID for the Friends (WiiU) protocol
	FriendsWiiUProtocolID = 0x66

	// FriendsWiiUMethodUpdateAndGetAllInformation is the method ID for method UpdateAndGetAllInformation
	FriendsWiiUMethodUpdateAndGetAllInformation = 0x1

	// FriendsWiiUMethodAddFriend is the method ID for method AddFriend
	FriendsWiiUMethodAddFriend = 0x2

	// FriendsWiiUMethodAddFriendByName is the method ID for method AddFriendByName
	FriendsWiiUMethodAddFriendByName = 0x3

	// FriendsWiiUMethodRemoveFriend is the method ID for method RemoveFriend
	FriendsWiiUMethodRemoveFriend = 0x4

	// FriendsWiiUMethodAddFriendRequest is the method ID for method AddFriendRequest
	FriendsWiiUMethodAddFriendRequest = 0x5

	// FriendsWiiUMethodCancelFriendRequest is the method ID for method CancelFriendRequest
	FriendsWiiUMethodCancelFriendRequest = 0x6

	// FriendsWiiUMethodAcceptFriendRequest is the method ID for method AcceptFriendRequest
	FriendsWiiUMethodAcceptFriendRequest = 0x7

	// FriendsWiiUMethodDeleteFriendRequest is the method ID for method DeleteFriendRequest
	FriendsWiiUMethodDeleteFriendRequest = 0x8

	// FriendsWiiUMethodDenyFriendRequest is the method ID for method DenyFriendRequest
	FriendsWiiUMethodDenyFriendRequest = 0x9

	// FriendsWiiUMethodMarkFriendRequestsAsReceived is the method ID for method MarkFriendRequestsAsReceived
	FriendsWiiUMethodMarkFriendRequestsAsReceived = 0xA

	// FriendsWiiUMethodAddBlackList is the method ID for method AddBlackList
	FriendsWiiUMethodAddBlackList = 0xB

	// FriendsWiiUMethodRemoveBlackList is the method ID for method RemoveBlackList
	FriendsWiiUMethodRemoveBlackList = 0xC

	// FriendsWiiUMethodUpdatePresence is the method ID for method UpdatePresence
	FriendsWiiUMethodUpdatePresence = 0xD

	// FriendsWiiUMethodUpdateMii is the method ID for method UpdateMii
	FriendsWiiUMethodUpdateMii = 0xE

	// FriendsWiiUMethodUpdateComment is the method ID for method UpdateComment
	FriendsWiiUMethodUpdateComment = 0xF

	// FriendsWiiUMethodUpdatePreference is the method ID for method UpdatePreference
	FriendsWiiUMethodUpdatePreference = 0x10

	// FriendsWiiUMethodGetBasicInfo is the method ID for method GetBasicInfo
	FriendsWiiUMethodGetBasicInfo = 0x11

	// FriendsWiiUMethodDeletePersistentNotification is the method ID for method DeletePersistentNotification
	FriendsWiiUMethodDeletePersistentNotification = 0x12

	// FriendsWiiUMethodCheckSettingStatus is the method ID for method CheckSettingStatus
	FriendsWiiUMethodCheckSettingStatus = 0x13

	// FriendsWiiUMethodGetRequestBlockSettings is the method ID for method GetRequestBlockSettings
	FriendsWiiUMethodGetRequestBlockSettings = 0x14
)

// FriendsWiiUProtocol handles the Friends (WiiU) nex protocol
type FriendsWiiUProtocol struct {
	server                              *nex.Server
	UpdateAndGetAllInformationHandler   func(err error, client *nex.Client, callID uint32, nnaInfo *NNAInfo, presence *NintendoPresenceV2, birthday *nex.DateTime)
	AddFriendHandler                    func(err error, client *nex.Client, callID uint32, pid uint32)
	AddFriendByNameHandler              func(err error, client *nex.Client, callID uint32, username string)
	RemoveFriendHandler                 func(err error, client *nex.Client, callID uint32, pid uint32)
	AddFriendRequestHandler             func(err error, client *nex.Client, callID uint32, pid uint32, unknown2 uint8, message string, unknown4 uint8, unknown5 string, gameKey *GameKey, unknown6 *nex.DateTime)
	CancelFriendRequestHandler          func(err error, client *nex.Client, callID uint32, id uint64)
	AcceptFriendRequestHandler          func(err error, client *nex.Client, callID uint32, id uint64)
	DeleteFriendRequestHandler          func(err error, client *nex.Client, callID uint32, id uint64)
	DenyFriendRequestHandler            func(err error, client *nex.Client, callID uint32, id uint64)
	MarkFriendRequestsAsReceivedHandler func(err error, client *nex.Client, callID uint32, ids []uint64)
	AddBlackListHandler                 func(err error, client *nex.Client, callID uint32, blacklistedPrincipal *BlacklistedPrincipal)
	RemoveBlackListHandler              func(err error, client *nex.Client, callID uint32, pid uint32)
	UpdatePresenceHandler               func(err error, client *nex.Client, callID uint32, presence *NintendoPresenceV2)
	UpdateMiiHandler                    func(err error, client *nex.Client, callID uint32, mii *MiiV2)
	UpdateCommentHandler                func(err error, client *nex.Client, callID uint32, comment *Comment)
	UpdatePreferenceHandler             func(err error, client *nex.Client, callID uint32, preference *PrincipalPreference)
	GetBasicInfoHandler                 func(err error, client *nex.Client, callID uint32, pids []uint32)
	DeletePersistentNotificationHandler func(err error, client *nex.Client, callID uint32, notifications []*PersistentNotification)
	CheckSettingStatusHandler           func(err error, client *nex.Client, callID uint32)
	GetRequestBlockSettingsHandler      func(err error, client *nex.Client, callID uint32, pids []uint32)
}

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

// Setup initializes the protocol
func (friendsWiiUProtocol *FriendsWiiUProtocol) Setup() {
	nexServer := friendsWiiUProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if FriendsWiiUProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case FriendsWiiUMethodUpdateAndGetAllInformation:
				go friendsWiiUProtocol.handleUpdateAndGetAllInformation(packet)
			case FriendsWiiUMethodAddFriend:
				go friendsWiiUProtocol.handleAddFriend(packet)
			case FriendsWiiUMethodAddFriendByName:
				go friendsWiiUProtocol.handleAddFriendByName(packet)
			case FriendsWiiUMethodRemoveFriend:
				go friendsWiiUProtocol.handleRemoveFriend(packet)
			case FriendsWiiUMethodAddFriendRequest:
				go friendsWiiUProtocol.handleAddFriendRequest(packet)
			case FriendsWiiUMethodCancelFriendRequest:
				go friendsWiiUProtocol.handleCancelFriendRequest(packet)
			case FriendsWiiUMethodAcceptFriendRequest:
				go friendsWiiUProtocol.handleAcceptFriendRequest(packet)
			case FriendsWiiUMethodDeleteFriendRequest:
				go friendsWiiUProtocol.handleDeleteFriendRequest(packet)
			case FriendsWiiUMethodDenyFriendRequest:
				go friendsWiiUProtocol.handleDenyFriendRequest(packet)
			case FriendsWiiUMethodMarkFriendRequestsAsReceived:
				go friendsWiiUProtocol.handleMarkFriendRequestsAsReceived(packet)
			case FriendsWiiUMethodAddBlackList:
				go friendsWiiUProtocol.handleAddBlackList(packet)
			case FriendsWiiUMethodRemoveBlackList:
				go friendsWiiUProtocol.handleRemoveBlackList(packet)
			case FriendsWiiUMethodUpdatePresence:
				go friendsWiiUProtocol.handleUpdatePresence(packet)
			case FriendsWiiUMethodUpdateMii:
				go friendsWiiUProtocol.handleUpdateMii(packet)
			case FriendsWiiUMethodUpdateComment:
				go friendsWiiUProtocol.handleUpdateComment(packet)
			case FriendsWiiUMethodUpdatePreference:
				go friendsWiiUProtocol.handleUpdatePreference(packet)
			case FriendsWiiUMethodGetBasicInfo:
				go friendsWiiUProtocol.handleGetBasicInfo(packet)
			case FriendsWiiUMethodDeletePersistentNotification:
				go friendsWiiUProtocol.handleDeletePersistentNotification(packet)
			case FriendsWiiUMethodCheckSettingStatus:
				go friendsWiiUProtocol.handleCheckSettingStatus(packet)
			case FriendsWiiUMethodGetRequestBlockSettings:
				go friendsWiiUProtocol.handleGetRequestBlockSettings(packet)
			default:
				go respondNotImplemented(packet, FriendsWiiUProtocolID)
				fmt.Printf("Unsupported Friends (WiiU) method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// UpdateAndGetAllInformation sets the UpdateAndGetAllInformation handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) UpdateAndGetAllInformation(handler func(err error, client *nex.Client, callID uint32, nnaInfo *NNAInfo, presence *NintendoPresenceV2, birthday *nex.DateTime)) {
	friendsWiiUProtocol.UpdateAndGetAllInformationHandler = handler
}

// AddFriend sets the AddFriend handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) AddFriend(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	friendsWiiUProtocol.AddFriendHandler = handler
}

// AddFriendByName sets the AddFriendByName handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) AddFriendByName(handler func(err error, client *nex.Client, callID uint32, username string)) {
	friendsWiiUProtocol.AddFriendByNameHandler = handler
}

// RemoveFriend sets the RemoveFriend handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) RemoveFriend(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	friendsWiiUProtocol.RemoveFriendHandler = handler
}

// AddFriendRequest sets the AddFriendRequest handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) AddFriendRequest(handler func(err error, client *nex.Client, callID uint32, pid uint32, unknown2 uint8, message string, unknown4 uint8, unknown5 string, gameKey *GameKey, unknown6 *nex.DateTime)) {
	friendsWiiUProtocol.AddFriendRequestHandler = handler
}

// CancelFriendRequest sets the CancelFriendRequest handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) CancelFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	friendsWiiUProtocol.CancelFriendRequestHandler = handler
}

// AcceptFriendRequest sets the AcceptFriendRequest handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) AcceptFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	friendsWiiUProtocol.AcceptFriendRequestHandler = handler
}

// DeleteFriendRequest sets the DeleteFriendRequest handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) DeleteFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	friendsWiiUProtocol.DeleteFriendRequestHandler = handler
}

// DenyFriendRequest sets the DenyFriendRequest handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) DenyFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	friendsWiiUProtocol.DenyFriendRequestHandler = handler
}

// MarkFriendRequestsAsReceived sets the MarkFriendRequestsAsReceived handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) MarkFriendRequestsAsReceived(handler func(err error, client *nex.Client, callID uint32, ids []uint64)) {
	friendsWiiUProtocol.MarkFriendRequestsAsReceivedHandler = handler
}

// AddBlackList sets the AddBlackList handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) AddBlackList(handler func(err error, client *nex.Client, callID uint32, blacklistedPrincipal *BlacklistedPrincipal)) {
	friendsWiiUProtocol.AddBlackListHandler = handler
}

// RemoveBlackList sets the RemoveBlackList handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) RemoveBlackList(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	friendsWiiUProtocol.RemoveBlackListHandler = handler
}

// UpdatePresence sets the UpdatePresence handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) UpdatePresence(handler func(err error, client *nex.Client, callID uint32, presence *NintendoPresenceV2)) {
	friendsWiiUProtocol.UpdatePresenceHandler = handler
}

// UpdateMii sets the UpdateMii handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) UpdateMii(handler func(err error, client *nex.Client, callID uint32, mii *MiiV2)) {
	friendsWiiUProtocol.UpdateMiiHandler = handler
}

// UpdateComment sets the UpdateComment handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) UpdateComment(handler func(err error, client *nex.Client, callID uint32, comment *Comment)) {
	friendsWiiUProtocol.UpdateCommentHandler = handler
}

// UpdatePreference sets the UpdatePreference handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) UpdatePreference(handler func(err error, client *nex.Client, callID uint32, preference *PrincipalPreference)) {
	friendsWiiUProtocol.UpdatePreferenceHandler = handler
}

// GetBasicInfo sets the GetBasicInfo handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) GetBasicInfo(handler func(err error, client *nex.Client, callID uint32, pids []uint32)) {
	friendsWiiUProtocol.GetBasicInfoHandler = handler
}

// DeletePersistentNotification sets the DeletePersistentNotification handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) DeletePersistentNotification(handler func(err error, client *nex.Client, callID uint32, notifications []*PersistentNotification)) {
	friendsWiiUProtocol.DeletePersistentNotificationHandler = handler
}

// CheckSettingStatus sets the CheckSettingStatus handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) CheckSettingStatus(handler func(err error, client *nex.Client, callID uint32)) {
	friendsWiiUProtocol.CheckSettingStatusHandler = handler
}

// GetRequestBlockSettings sets the GetRequestBlockSettings handler function
func (friendsWiiUProtocol *FriendsWiiUProtocol) GetRequestBlockSettings(handler func(err error, client *nex.Client, callID uint32, unknowns []uint32)) {
	friendsWiiUProtocol.GetRequestBlockSettingsHandler = handler
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleUpdateAndGetAllInformation(packet nex.PacketInterface) {
	if friendsWiiUProtocol.UpdateAndGetAllInformationHandler == nil {
		logger.Warning("FriendsWiiUProtocol::UpdateAndGetAllInformation not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	nnaInfoStructureInterface, err := parametersStream.ReadStructure(NewNNAInfo())
	if err != nil {
		go friendsWiiUProtocol.UpdateAndGetAllInformationHandler(err, client, callID, nil, nil, nil)
		return
	}

	presenceStructureInterface, err := parametersStream.ReadStructure(NewNintendoPresenceV2())
	if err != nil {
		go friendsWiiUProtocol.UpdateAndGetAllInformationHandler(err, client, callID, nil, nil, nil)
		return
	}

	nnaInfo := nnaInfoStructureInterface.(*NNAInfo)
	presence := presenceStructureInterface.(*NintendoPresenceV2)
	birthday := nex.NewDateTime(parametersStream.ReadUInt64LE())

	go friendsWiiUProtocol.UpdateAndGetAllInformationHandler(nil, client, callID, nnaInfo, presence, birthday)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleAddFriend(packet nex.PacketInterface) {
	if friendsWiiUProtocol.AddFriendHandler == nil {
		logger.Warning("FriendsWiiUProtocol::AddFriend not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsWiiUProtocol::AddFriend] Data holder not long enough for PID")
		go friendsWiiUProtocol.AddFriendHandler(err, client, callID, 0)
		return
	}

	pid := parametersStream.ReadUInt32LE()

	go friendsWiiUProtocol.AddFriendHandler(nil, client, callID, pid)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleAddFriendByName(packet nex.PacketInterface) {
	if friendsWiiUProtocol.AddFriendByNameHandler == nil {
		logger.Warning("FriendsWiiUProtocol::AddFriendByName not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	username, err := parametersStream.ReadString()

	if err != nil {
		go friendsWiiUProtocol.AddFriendByNameHandler(err, client, callID, "")
		return
	}

	go friendsWiiUProtocol.AddFriendByNameHandler(nil, client, callID, username)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleRemoveFriend(packet nex.PacketInterface) {
	if friendsWiiUProtocol.RemoveFriendHandler == nil {
		logger.Warning("FriendsWiiUProtocol::RemoveFriend not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsWiiUProtocol::RemoveFriend] Data holder not long enough for PID")
		go friendsWiiUProtocol.RemoveFriendHandler(err, client, callID, 0)
		return
	}

	pid := parametersStream.ReadUInt32LE()

	go friendsWiiUProtocol.RemoveFriendHandler(nil, client, callID, pid)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleAddFriendRequest(packet nex.PacketInterface) {
	if friendsWiiUProtocol.AddFriendRequestHandler == nil {
		logger.Warning("FriendsWiiUProtocol::AddFriendRequest not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4+1+1+8 {
		// length check for the following fixed-size data
		// unknown1 + unknown2 + unknown4 + gameKey + unknown6
		err := errors.New("[FriendsWiiUProtocol::AddFriendRequest] Data holder not long enough for PID")
		go friendsWiiUProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	pid := parametersStream.ReadUInt32LE()
	unknown2 := parametersStream.ReadUInt8()
	message, err := parametersStream.ReadString()

	if err != nil {
		go friendsWiiUProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	unknown4 := parametersStream.ReadUInt8()
	unknown5, err := parametersStream.ReadString()

	if err != nil {
		go friendsWiiUProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	gameKeyStructureInterface, err := parametersStream.ReadStructure(NewGameKey())
	if err != nil {
		go friendsWiiUProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	gameKey := gameKeyStructureInterface.(*GameKey)

	if err != nil {
		go friendsWiiUProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	unknown6 := nex.NewDateTime(parametersStream.ReadUInt64LE())

	go friendsWiiUProtocol.AddFriendRequestHandler(nil, client, callID, pid, unknown2, message, unknown4, unknown5, gameKey, unknown6)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleCancelFriendRequest(packet nex.PacketInterface) {
	if friendsWiiUProtocol.CancelFriendRequestHandler == nil {
		logger.Warning("FriendsWiiUProtocol::CancelFriendRequest not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[FriendsWiiUProtocol::CancelFriendRequest] Data missing list length")
		go friendsWiiUProtocol.CancelFriendRequestHandler(err, client, callID, 0)
		return
	}

	id := parametersStream.ReadUInt64LE()

	go friendsWiiUProtocol.CancelFriendRequestHandler(nil, client, callID, id)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleAcceptFriendRequest(packet nex.PacketInterface) {
	if friendsWiiUProtocol.AcceptFriendRequestHandler == nil {
		logger.Warning("FriendsWiiUProtocol::AcceptFriendRequest not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[FriendsWiiUProtocol::AcceptFriendRequest] Data missing list length")
		go friendsWiiUProtocol.AcceptFriendRequestHandler(err, client, callID, 0)
		return
	}

	id := parametersStream.ReadUInt64LE()

	go friendsWiiUProtocol.AcceptFriendRequestHandler(nil, client, callID, id)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleDeleteFriendRequest(packet nex.PacketInterface) {
	if friendsWiiUProtocol.DeleteFriendRequestHandler == nil {
		logger.Warning("FriendsWiiUProtocol::DeleteFriendRequest not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[FriendsWiiUProtocol::DeleteFriendRequest] Data missing list length")
		go friendsWiiUProtocol.DeleteFriendRequestHandler(err, client, callID, 0)
		return
	}

	id := parametersStream.ReadUInt64LE()

	go friendsWiiUProtocol.DeleteFriendRequestHandler(nil, client, callID, id)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleDenyFriendRequest(packet nex.PacketInterface) {
	if friendsWiiUProtocol.DenyFriendRequestHandler == nil {
		logger.Warning("FriendsWiiUProtocol::DenyFriendRequest not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[FriendsWiiUProtocol::DenyFriendRequest] Data missing list length")
		go friendsWiiUProtocol.DenyFriendRequestHandler(err, client, callID, 0)
		return
	}

	id := parametersStream.ReadUInt64LE()

	go friendsWiiUProtocol.DenyFriendRequestHandler(nil, client, callID, id)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleMarkFriendRequestsAsReceived(packet nex.PacketInterface) {
	if friendsWiiUProtocol.MarkFriendRequestsAsReceivedHandler == nil {
		logger.Warning("FriendsWiiUProtocol::MarkFriendRequestsAsReceived not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsWiiUProtocol::MarkFriendRequestsAsReceived] Data missing list length")
		go friendsWiiUProtocol.MarkFriendRequestsAsReceivedHandler(err, client, callID, make([]uint64, 0))
		return
	}

	ids := parametersStream.ReadListUInt64LE()

	go friendsWiiUProtocol.MarkFriendRequestsAsReceivedHandler(nil, client, callID, ids)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleAddBlackList(packet nex.PacketInterface) {
	if friendsWiiUProtocol.AddBlackListHandler == nil {
		logger.Warning("FriendsWiiUProtocol::AddBlackList not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	blacklistedPrincipalStructureInterface, err := parametersStream.ReadStructure(NewBlacklistedPrincipal())
	if err != nil {
		go friendsWiiUProtocol.AddBlackListHandler(err, client, callID, nil)
		return
	}

	blacklistedPrincipal := blacklistedPrincipalStructureInterface.(*BlacklistedPrincipal)

	go friendsWiiUProtocol.AddBlackListHandler(nil, client, callID, blacklistedPrincipal)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleRemoveBlackList(packet nex.PacketInterface) {
	if friendsWiiUProtocol.RemoveBlackListHandler == nil {
		logger.Warning("FriendsWiiUProtocol::RemoveBlackList not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsWiiUProtocol::RemoveBlackList] Data missing list length")
		go friendsWiiUProtocol.RemoveBlackListHandler(err, client, callID, 0)
		return
	}

	pid := parametersStream.ReadUInt32LE()

	go friendsWiiUProtocol.RemoveBlackListHandler(nil, client, callID, pid)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleUpdatePresence(packet nex.PacketInterface) {
	if friendsWiiUProtocol.UpdatePresenceHandler == nil {
		logger.Warning("FriendsWiiUProtocol::UpdatePresence not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	nintendoPresenceV2StructureInterface, err := parametersStream.ReadStructure(NewNintendoPresenceV2())
	if err != nil {
		go friendsWiiUProtocol.UpdatePresenceHandler(err, client, callID, nil)
		return
	}

	nintendoPresenceV2 := nintendoPresenceV2StructureInterface.(*NintendoPresenceV2)

	go friendsWiiUProtocol.UpdatePresenceHandler(nil, client, callID, nintendoPresenceV2)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleUpdateMii(packet nex.PacketInterface) {
	if friendsWiiUProtocol.UpdateMiiHandler == nil {
		logger.Warning("FriendsWiiUProtocol::UpdateMii not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	miiV2StructureInterface, err := parametersStream.ReadStructure(NewMiiV2())
	if err != nil {
		go friendsWiiUProtocol.UpdateMiiHandler(err, client, callID, nil)
		return
	}

	miiV2 := miiV2StructureInterface.(*MiiV2)

	go friendsWiiUProtocol.UpdateMiiHandler(nil, client, callID, miiV2)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleUpdateComment(packet nex.PacketInterface) {
	if friendsWiiUProtocol.UpdateCommentHandler == nil {
		logger.Warning("FriendsWiiUProtocol::UpdateComment not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	commentStructureInterface, err := parametersStream.ReadStructure(NewComment())
	if err != nil {
		go friendsWiiUProtocol.UpdateCommentHandler(err, client, callID, nil)
		return
	}

	comment := commentStructureInterface.(*Comment)

	go friendsWiiUProtocol.UpdateCommentHandler(nil, client, callID, comment)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleUpdatePreference(packet nex.PacketInterface) {
	if friendsWiiUProtocol.UpdatePreferenceHandler == nil {
		logger.Warning("FriendsWiiUProtocol::UpdatePreference not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	principalPreferenceStructureInterface, err := parametersStream.ReadStructure(NewPrincipalPreference())
	if err != nil {
		go friendsWiiUProtocol.UpdatePreferenceHandler(err, client, callID, nil)
		return
	}

	principalPreference := principalPreferenceStructureInterface.(*PrincipalPreference)

	go friendsWiiUProtocol.UpdatePreferenceHandler(nil, client, callID, principalPreference)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleGetBasicInfo(packet nex.PacketInterface) {
	if friendsWiiUProtocol.GetBasicInfoHandler == nil {
		logger.Warning("FriendsWiiUProtocol::GetBasicInfo not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsWiiUProtocol::GetBasicInfo] Data missing list length")
		go friendsWiiUProtocol.GetBasicInfoHandler(err, client, callID, make([]uint32, 0))
		return
	}

	pids := parametersStream.ReadListUInt32LE()

	go friendsWiiUProtocol.GetBasicInfoHandler(nil, client, callID, pids)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleDeletePersistentNotification(packet nex.PacketInterface) {
	if friendsWiiUProtocol.DeletePersistentNotificationHandler == nil {
		logger.Warning("FriendsWiiUProtocol::DeletePersistentNotification not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	persistentNotifications, err := ReadListPersistentNotification(parametersStream)

	if err != nil {
		go friendsWiiUProtocol.DeletePersistentNotificationHandler(err, client, callID, nil)
		return
	}

	go friendsWiiUProtocol.DeletePersistentNotificationHandler(nil, client, callID, persistentNotifications)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleCheckSettingStatus(packet nex.PacketInterface) {
	if friendsWiiUProtocol.CheckSettingStatusHandler == nil {
		logger.Warning("FriendsWiiUProtocol::CheckSettingStatus not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go friendsWiiUProtocol.CheckSettingStatusHandler(nil, client, callID)
}

func (friendsWiiUProtocol *FriendsWiiUProtocol) handleGetRequestBlockSettings(packet nex.PacketInterface) {
	if friendsWiiUProtocol.GetRequestBlockSettingsHandler == nil {
		logger.Warning("FriendsWiiUProtocol::GetRequestBlockSettings not implemented")
		go respondNotImplemented(packet, FriendsWiiUProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsWiiUProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsWiiUProtocol::GetRequestBlockSettings] Data missing list length")
		go friendsWiiUProtocol.GetRequestBlockSettingsHandler(err, client, callID, make([]uint32, 0))
		return
	}

	pids := parametersStream.ReadListUInt32LE()

	go friendsWiiUProtocol.GetRequestBlockSettingsHandler(nil, client, callID, pids)
}

// NewFriendsWiiUProtocol returns a new FriendsWiiUProtocol
func NewFriendsWiiUProtocol(server *nex.Server) *FriendsWiiUProtocol {
	friendsWiiUProtocol := &FriendsWiiUProtocol{server: server}

	friendsWiiUProtocol.Setup()

	return friendsWiiUProtocol
}
