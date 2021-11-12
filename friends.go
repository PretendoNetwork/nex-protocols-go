package nexproto

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// FriendsProtocolID is the protocol ID for the Friends (WiiU) protocol
	FriendsProtocolID = 0x66

	// FriendsMethodUpdateAndGetAllInformation is the method ID for method UpdateAndGetAllInformation
	FriendsMethodUpdateAndGetAllInformation = 0x1

	// FriendsMethodAddFriend is the method ID for method AddFriend
	FriendsMethodAddFriend = 0x2

	// FriendsMethodAddFriendByName is the method ID for method AddFriendByName
	FriendsMethodAddFriendByName = 0x3

	// FriendsMethodRemoveFriend is the method ID for method RemoveFriend
	FriendsMethodRemoveFriend = 0x4

	// FriendsMethodAddFriendRequest is the method ID for method AddFriendRequest
	FriendsMethodAddFriendRequest = 0x5

	// FriendsMethodCancelFriendRequest is the method ID for method CancelFriendRequest
	FriendsMethodCancelFriendRequest = 0x6

	// FriendsMethodAcceptFriendRequest is the method ID for method AcceptFriendRequest
	FriendsMethodAcceptFriendRequest = 0x7

	// FriendsMethodDeleteFriendRequest is the method ID for method DeleteFriendRequest
	FriendsMethodDeleteFriendRequest = 0x8

	// FriendsMethodDenyFriendRequest is the method ID for method DenyFriendRequest
	FriendsMethodDenyFriendRequest = 0x9

	// FriendsMethodMarkFriendRequestsAsReceived is the method ID for method MarkFriendRequestsAsReceived
	FriendsMethodMarkFriendRequestsAsReceived = 0xA

	// FriendsMethodAddBlackList is the method ID for method AddBlackList
	FriendsMethodAddBlackList = 0xB

	// FriendsMethodRemoveBlackList is the method ID for method RemoveBlackList
	FriendsMethodRemoveBlackList = 0xC

	// FriendsMethodUpdatePresence is the method ID for method UpdatePresence
	FriendsMethodUpdatePresence = 0xD

	// FriendsMethodUpdateMii is the method ID for method UpdateMii
	FriendsMethodUpdateMii = 0xE

	// FriendsMethodUpdateComment is the method ID for method UpdateComment
	FriendsMethodUpdateComment = 0xF

	// FriendsMethodUpdatePreference is the method ID for method UpdatePreference
	FriendsMethodUpdatePreference = 0x10

	// FriendsMethodGetBasicInfo is the method ID for method GetBasicInfo
	FriendsMethodGetBasicInfo = 0x11

	// FriendsMethodDeletePersistentNotification is the method ID for method DeletePersistentNotification
	FriendsMethodDeletePersistentNotification = 0x12

	// FriendsMethodCheckSettingStatus is the method ID for method CheckSettingStatus
	FriendsMethodCheckSettingStatus = 0x13

	// FriendsMethodGetRequestBlockSettings is the method ID for method GetRequestBlockSettings
	FriendsMethodGetRequestBlockSettings = 0x14
)

// FriendsProtocol handles the Friends (WiiU) nex protocol
type FriendsProtocol struct {
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
	stream.WriteUInt64LE(comment.LastChanged.Value())

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
	stream.WriteUInt64LE(friendInfo.BecameFriend.Value())
	stream.WriteUInt64LE(friendInfo.LastOnline.Value())
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
	stream.WriteUInt64LE(friendRequest.SentOn.Value())

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
	stream.WriteUInt64LE(friendRequestMessage.Unknown5.Value())
	stream.WriteUInt64LE(friendRequestMessage.ExpiresOn.Value())

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
	stream.WriteUInt64LE(mii.Datetime.Value())

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
func (friendsProtocol *FriendsProtocol) Setup() {
	nexServer := friendsProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if FriendsProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case FriendsMethodUpdateAndGetAllInformation:
				go friendsProtocol.handleUpdateAndGetAllInformation(packet)
			case FriendsMethodAddFriend:
				go friendsProtocol.handleAddFriend(packet)
			case FriendsMethodAddFriendByName:
				go friendsProtocol.handleAddFriendByName(packet)
			case FriendsMethodRemoveFriend:
				go friendsProtocol.handleRemoveFriend(packet)
			case FriendsMethodAddFriendRequest:
				go friendsProtocol.handleAddFriendRequest(packet)
			case FriendsMethodCancelFriendRequest:
				go friendsProtocol.handleCancelFriendRequest(packet)
			case FriendsMethodAcceptFriendRequest:
				go friendsProtocol.handleAcceptFriendRequest(packet)
			case FriendsMethodDeleteFriendRequest:
				go friendsProtocol.handleDeleteFriendRequest(packet)
			case FriendsMethodDenyFriendRequest:
				go friendsProtocol.handleDenyFriendRequest(packet)
			case FriendsMethodMarkFriendRequestsAsReceived:
				go friendsProtocol.handleMarkFriendRequestsAsReceived(packet)
			case FriendsMethodAddBlackList:
				go friendsProtocol.handleAddBlackList(packet)
			case FriendsMethodRemoveBlackList:
				go friendsProtocol.handleRemoveBlackList(packet)
			case FriendsMethodUpdatePresence:
				go friendsProtocol.handleUpdatePresence(packet)
			case FriendsMethodUpdateMii:
				go friendsProtocol.handleUpdateMii(packet)
			case FriendsMethodUpdateComment:
				go friendsProtocol.handleUpdateComment(packet)
			case FriendsMethodUpdatePreference:
				go friendsProtocol.handleUpdatePreference(packet)
			case FriendsMethodGetBasicInfo:
				go friendsProtocol.handleGetBasicInfo(packet)
			case FriendsMethodDeletePersistentNotification:
				go friendsProtocol.handleDeletePersistentNotification(packet)
			case FriendsMethodCheckSettingStatus:
				go friendsProtocol.handleCheckSettingStatus(packet)
			case FriendsMethodGetRequestBlockSettings:
				go friendsProtocol.handleGetRequestBlockSettings(packet)
			default:
				go respondNotImplemented(packet, FriendsProtocolID)
				fmt.Printf("Unsupported Friends (WiiU) method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// UpdateAndGetAllInformation sets the UpdateAndGetAllInformation handler function
func (friendsProtocol *FriendsProtocol) UpdateAndGetAllInformation(handler func(err error, client *nex.Client, callID uint32, nnaInfo *NNAInfo, presence *NintendoPresenceV2, birthday *nex.DateTime)) {
	friendsProtocol.UpdateAndGetAllInformationHandler = handler
}

// AddFriend sets the AddFriend handler function
func (friendsProtocol *FriendsProtocol) AddFriend(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	friendsProtocol.AddFriendHandler = handler
}

// AddFriendByName sets the AddFriendByName handler function
func (friendsProtocol *FriendsProtocol) AddFriendByName(handler func(err error, client *nex.Client, callID uint32, username string)) {
	friendsProtocol.AddFriendByNameHandler = handler
}

// RemoveFriend sets the RemoveFriend handler function
func (friendsProtocol *FriendsProtocol) RemoveFriend(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	friendsProtocol.RemoveFriendHandler = handler
}

// AddFriendRequest sets the AddFriendRequest handler function
func (friendsProtocol *FriendsProtocol) AddFriendRequest(handler func(err error, client *nex.Client, callID uint32, pid uint32, unknown2 uint8, message string, unknown4 uint8, unknown5 string, gameKey *GameKey, unknown6 *nex.DateTime)) {
	friendsProtocol.AddFriendRequestHandler = handler
}

// CancelFriendRequest sets the CancelFriendRequest handler function
func (friendsProtocol *FriendsProtocol) CancelFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	friendsProtocol.CancelFriendRequestHandler = handler
}

// AcceptFriendRequest sets the AcceptFriendRequest handler function
func (friendsProtocol *FriendsProtocol) AcceptFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	friendsProtocol.AcceptFriendRequestHandler = handler
}

// DeleteFriendRequest sets the DeleteFriendRequest handler function
func (friendsProtocol *FriendsProtocol) DeleteFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	friendsProtocol.DeleteFriendRequestHandler = handler
}

// DenyFriendRequest sets the DenyFriendRequest handler function
func (friendsProtocol *FriendsProtocol) DenyFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	friendsProtocol.DenyFriendRequestHandler = handler
}

// MarkFriendRequestsAsReceived sets the MarkFriendRequestsAsReceived handler function
func (friendsProtocol *FriendsProtocol) MarkFriendRequestsAsReceived(handler func(err error, client *nex.Client, callID uint32, ids []uint64)) {
	friendsProtocol.MarkFriendRequestsAsReceivedHandler = handler
}

// AddBlackList sets the AddBlackList handler function
func (friendsProtocol *FriendsProtocol) AddBlackList(handler func(err error, client *nex.Client, callID uint32, blacklistedPrincipal *BlacklistedPrincipal)) {
	friendsProtocol.AddBlackListHandler = handler
}

// RemoveBlackList sets the RemoveBlackList handler function
func (friendsProtocol *FriendsProtocol) RemoveBlackList(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	friendsProtocol.RemoveBlackListHandler = handler
}

// UpdatePresence sets the UpdatePresence handler function
func (friendsProtocol *FriendsProtocol) UpdatePresence(handler func(err error, client *nex.Client, callID uint32, presence *NintendoPresenceV2)) {
	friendsProtocol.UpdatePresenceHandler = handler
}

// UpdateMii sets the UpdateMii handler function
func (friendsProtocol *FriendsProtocol) UpdateMii(handler func(err error, client *nex.Client, callID uint32, mii *MiiV2)) {
	friendsProtocol.UpdateMiiHandler = handler
}

// UpdateComment sets the UpdateComment handler function
func (friendsProtocol *FriendsProtocol) UpdateComment(handler func(err error, client *nex.Client, callID uint32, comment *Comment)) {
	friendsProtocol.UpdateCommentHandler = handler
}

// UpdatePreference sets the UpdatePreference handler function
func (friendsProtocol *FriendsProtocol) UpdatePreference(handler func(err error, client *nex.Client, callID uint32, preference *PrincipalPreference)) {
	friendsProtocol.UpdatePreferenceHandler = handler
}

// GetBasicInfo sets the GetBasicInfo handler function
func (friendsProtocol *FriendsProtocol) GetBasicInfo(handler func(err error, client *nex.Client, callID uint32, pids []uint32)) {
	friendsProtocol.GetBasicInfoHandler = handler
}

// DeletePersistentNotification sets the DeletePersistentNotification handler function
func (friendsProtocol *FriendsProtocol) DeletePersistentNotification(handler func(err error, client *nex.Client, callID uint32, notifications []*PersistentNotification)) {
	friendsProtocol.DeletePersistentNotificationHandler = handler
}

// CheckSettingStatus sets the CheckSettingStatus handler function
func (friendsProtocol *FriendsProtocol) CheckSettingStatus(handler func(err error, client *nex.Client, callID uint32)) {
	friendsProtocol.CheckSettingStatusHandler = handler
}

// GetRequestBlockSettings sets the GetRequestBlockSettings handler function
func (friendsProtocol *FriendsProtocol) GetRequestBlockSettings(handler func(err error, client *nex.Client, callID uint32, unknowns []uint32)) {
	friendsProtocol.GetRequestBlockSettingsHandler = handler
}

func (friendsProtocol *FriendsProtocol) handleUpdateAndGetAllInformation(packet nex.PacketInterface) {
	if friendsProtocol.UpdateAndGetAllInformationHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::UpdateAndGetAllInformation not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	nnaInfoStructureInterface, err := parametersStream.ReadStructure(NewNNAInfo())
	if err != nil {
		go friendsProtocol.UpdateAndGetAllInformationHandler(err, client, callID, nil, nil, nil)
		return
	}

	presenceStructureInterface, err := parametersStream.ReadStructure(NewNintendoPresenceV2())
	if err != nil {
		go friendsProtocol.UpdateAndGetAllInformationHandler(err, client, callID, nil, nil, nil)
		return
	}

	nnaInfo := nnaInfoStructureInterface.(*NNAInfo)
	presence := presenceStructureInterface.(*NintendoPresenceV2)
	birthday := nex.NewDateTime(parametersStream.ReadUInt64LE())

	go friendsProtocol.UpdateAndGetAllInformationHandler(nil, client, callID, nnaInfo, presence, birthday)
}

func (friendsProtocol *FriendsProtocol) handleAddFriend(packet nex.PacketInterface) {
	if friendsProtocol.AddFriendHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::AddFriend not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsProtocol::AddFriend] Data holder not long enough for PID")
		go friendsProtocol.AddFriendHandler(err, client, callID, 0)
		return
	}

	pid := parametersStream.ReadUInt32LE()

	go friendsProtocol.AddFriendHandler(nil, client, callID, pid)
}

func (friendsProtocol *FriendsProtocol) handleAddFriendByName(packet nex.PacketInterface) {
	if friendsProtocol.AddFriendByNameHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::AddFriendByName not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	username, err := parametersStream.ReadString()

	if err != nil {
		go friendsProtocol.AddFriendByNameHandler(err, client, callID, "")
		return
	}

	go friendsProtocol.AddFriendByNameHandler(nil, client, callID, username)
}

func (friendsProtocol *FriendsProtocol) handleRemoveFriend(packet nex.PacketInterface) {
	if friendsProtocol.RemoveFriendHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::RemoveFriend not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsProtocol::RemoveFriend] Data holder not long enough for PID")
		go friendsProtocol.RemoveFriendHandler(err, client, callID, 0)
		return
	}

	pid := parametersStream.ReadUInt32LE()

	go friendsProtocol.RemoveFriendHandler(nil, client, callID, pid)
}

func (friendsProtocol *FriendsProtocol) handleAddFriendRequest(packet nex.PacketInterface) {
	if friendsProtocol.AddFriendRequestHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::AddFriendRequest not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4+1+1+8 {
		// length check for the following fixed-size data
		// unknown1 + unknown2 + unknown4 + gameKey + unknown6
		err := errors.New("[FriendsProtocol::AddFriendRequest] Data holder not long enough for PID")
		go friendsProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	pid := parametersStream.ReadUInt32LE()
	unknown2 := parametersStream.ReadUInt8()
	message, err := parametersStream.ReadString()

	if err != nil {
		go friendsProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	unknown4 := parametersStream.ReadUInt8()
	unknown5, err := parametersStream.ReadString()

	if err != nil {
		go friendsProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	gameKeyStructureInterface, err := parametersStream.ReadStructure(NewGameKey())
	if err != nil {
		go friendsProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	gameKey := gameKeyStructureInterface.(*GameKey)

	if err != nil {
		go friendsProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	unknown6 := nex.NewDateTime(parametersStream.ReadUInt64LE())

	go friendsProtocol.AddFriendRequestHandler(nil, client, callID, pid, unknown2, message, unknown4, unknown5, gameKey, unknown6)
}

func (friendsProtocol *FriendsProtocol) handleCancelFriendRequest(packet nex.PacketInterface) {
	if friendsProtocol.CancelFriendRequestHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::CancelFriendRequest not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[FriendsProtocol::CancelFriendRequest] Data missing list length")
		go friendsProtocol.CancelFriendRequestHandler(err, client, callID, 0)
		return
	}

	id := parametersStream.ReadUInt64LE()

	go friendsProtocol.CancelFriendRequestHandler(nil, client, callID, id)
}

func (friendsProtocol *FriendsProtocol) handleAcceptFriendRequest(packet nex.PacketInterface) {
	if friendsProtocol.AcceptFriendRequestHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::AcceptFriendRequest not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[FriendsProtocol::AcceptFriendRequest] Data missing list length")
		go friendsProtocol.AcceptFriendRequestHandler(err, client, callID, 0)
		return
	}

	id := parametersStream.ReadUInt64LE()

	go friendsProtocol.AcceptFriendRequestHandler(nil, client, callID, id)
}

func (friendsProtocol *FriendsProtocol) handleDeleteFriendRequest(packet nex.PacketInterface) {
	if friendsProtocol.DeleteFriendRequestHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::DeleteFriendRequest not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[FriendsProtocol::DeleteFriendRequest] Data missing list length")
		go friendsProtocol.DeleteFriendRequestHandler(err, client, callID, 0)
		return
	}

	id := parametersStream.ReadUInt64LE()

	go friendsProtocol.DeleteFriendRequestHandler(nil, client, callID, id)
}

func (friendsProtocol *FriendsProtocol) handleDenyFriendRequest(packet nex.PacketInterface) {
	if friendsProtocol.DenyFriendRequestHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::DenyFriendRequest not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[FriendsProtocol::DenyFriendRequest] Data missing list length")
		go friendsProtocol.DenyFriendRequestHandler(err, client, callID, 0)
		return
	}

	id := parametersStream.ReadUInt64LE()

	go friendsProtocol.DenyFriendRequestHandler(nil, client, callID, id)
}

func (friendsProtocol *FriendsProtocol) handleMarkFriendRequestsAsReceived(packet nex.PacketInterface) {
	if friendsProtocol.MarkFriendRequestsAsReceivedHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::MarkFriendRequestsAsReceived not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsProtocol::MarkFriendRequestsAsReceived] Data missing list length")
		go friendsProtocol.MarkFriendRequestsAsReceivedHandler(err, client, callID, make([]uint64, 0))
		return
	}

	ids := parametersStream.ReadListUInt64LE()

	go friendsProtocol.MarkFriendRequestsAsReceivedHandler(nil, client, callID, ids)
}

func (friendsProtocol *FriendsProtocol) handleAddBlackList(packet nex.PacketInterface) {
	if friendsProtocol.AddBlackListHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::AddBlackList not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	blacklistedPrincipalStructureInterface, err := parametersStream.ReadStructure(NewBlacklistedPrincipal())
	if err != nil {
		go friendsProtocol.AddBlackListHandler(err, client, callID, nil)
		return
	}

	blacklistedPrincipal := blacklistedPrincipalStructureInterface.(*BlacklistedPrincipal)

	go friendsProtocol.AddBlackListHandler(nil, client, callID, blacklistedPrincipal)
}

func (friendsProtocol *FriendsProtocol) handleRemoveBlackList(packet nex.PacketInterface) {
	if friendsProtocol.RemoveBlackListHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::RemoveBlackList not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsProtocol::RemoveBlackList] Data missing list length")
		go friendsProtocol.RemoveBlackListHandler(err, client, callID, 0)
		return
	}

	pid := parametersStream.ReadUInt32LE()

	go friendsProtocol.RemoveBlackListHandler(nil, client, callID, pid)
}

func (friendsProtocol *FriendsProtocol) handleUpdatePresence(packet nex.PacketInterface) {
	if friendsProtocol.UpdatePresenceHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::UpdatePresence not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	nintendoPresenceV2StructureInterface, err := parametersStream.ReadStructure(NewNintendoPresenceV2())
	if err != nil {
		go friendsProtocol.UpdatePresenceHandler(err, client, callID, nil)
		return
	}

	nintendoPresenceV2 := nintendoPresenceV2StructureInterface.(*NintendoPresenceV2)

	go friendsProtocol.UpdatePresenceHandler(nil, client, callID, nintendoPresenceV2)
}

func (friendsProtocol *FriendsProtocol) handleUpdateMii(packet nex.PacketInterface) {
	if friendsProtocol.UpdateMiiHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::UpdateMii not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	miiV2StructureInterface, err := parametersStream.ReadStructure(NewMiiV2())
	if err != nil {
		go friendsProtocol.UpdateMiiHandler(err, client, callID, nil)
		return
	}

	miiV2 := miiV2StructureInterface.(*MiiV2)

	go friendsProtocol.UpdateMiiHandler(nil, client, callID, miiV2)
}

func (friendsProtocol *FriendsProtocol) handleUpdateComment(packet nex.PacketInterface) {
	if friendsProtocol.UpdateCommentHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::UpdateComment not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	commentStructureInterface, err := parametersStream.ReadStructure(NewComment())
	if err != nil {
		go friendsProtocol.UpdateCommentHandler(err, client, callID, nil)
		return
	}

	comment := commentStructureInterface.(*Comment)

	go friendsProtocol.UpdateCommentHandler(nil, client, callID, comment)
}

func (friendsProtocol *FriendsProtocol) handleUpdatePreference(packet nex.PacketInterface) {
	if friendsProtocol.UpdatePreferenceHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::UpdatePreference not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	principalPreferenceStructureInterface, err := parametersStream.ReadStructure(NewPrincipalPreference())
	if err != nil {
		go friendsProtocol.UpdatePreferenceHandler(err, client, callID, nil)
		return
	}

	principalPreference := principalPreferenceStructureInterface.(*PrincipalPreference)

	go friendsProtocol.UpdatePreferenceHandler(nil, client, callID, principalPreference)
}

func (friendsProtocol *FriendsProtocol) handleGetBasicInfo(packet nex.PacketInterface) {
	if friendsProtocol.GetBasicInfoHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::GetBasicInfo not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsProtocol::GetBasicInfo] Data missing list length")
		go friendsProtocol.GetBasicInfoHandler(err, client, callID, make([]uint32, 0))
		return
	}

	pids := parametersStream.ReadListUInt32LE()

	go friendsProtocol.GetBasicInfoHandler(nil, client, callID, pids)
}

func (friendsProtocol *FriendsProtocol) handleDeletePersistentNotification(packet nex.PacketInterface) {
	if friendsProtocol.DeletePersistentNotificationHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::DeletePersistentNotification not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	persistentNotifications, err := ReadListPersistentNotification(parametersStream)

	if err != nil {
		go friendsProtocol.DeletePersistentNotificationHandler(err, client, callID, nil)
		return
	}

	go friendsProtocol.DeletePersistentNotificationHandler(nil, client, callID, persistentNotifications)
}

func (friendsProtocol *FriendsProtocol) handleCheckSettingStatus(packet nex.PacketInterface) {
	if friendsProtocol.CheckSettingStatusHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::CheckSettingStatus not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go friendsProtocol.CheckSettingStatusHandler(nil, client, callID)
}

func (friendsProtocol *FriendsProtocol) handleGetRequestBlockSettings(packet nex.PacketInterface) {
	if friendsProtocol.GetRequestBlockSettingsHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::GetRequestBlockSettings not implemented")
		go respondNotImplemented(packet, FriendsProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsProtocol::GetRequestBlockSettings] Data missing list length")
		go friendsProtocol.GetRequestBlockSettingsHandler(err, client, callID, make([]uint32, 0))
		return
	}

	pids := parametersStream.ReadListUInt32LE()

	go friendsProtocol.GetRequestBlockSettingsHandler(nil, client, callID, pids)
}

// NewFriendsProtocol returns a new FriendsProtocol
func NewFriendsProtocol(server *nex.Server) *FriendsProtocol {
	friendsProtocol := &FriendsProtocol{server: server}

	friendsProtocol.Setup()

	return friendsProtocol
}
