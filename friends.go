package nexproto

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	FriendsProtocolID = 0x66

	FriendsMethodUpdateAndGetAllInformation   = 0x1
	FriendsMethodAddFriend                    = 0x2
	FriendsMethodAddFriendByName              = 0x3
	FriendsMethodRemoveFriend                 = 0x4
	FriendsMethodAddFriendRequest             = 0x5
	FriendsMethodCancelFriendRequest          = 0x6
	FriendsMethodAcceptFriendRequest          = 0x7
	FriendsMethodDeleteFriendRequest          = 0x8
	FriendsMethodDenyFriendRequest            = 0x9
	FriendsMethodMarkFriendRequestsAsReceived = 0xA
	FriendsMethodAddBlackList                 = 0xB
	FriendsMethodRemoveBlackList              = 0xC
	FriendsMethodUpdatePresence               = 0xD
	FriendsMethodUpdateMii                    = 0xE
	FriendsMethodUpdateComment                = 0xF
	FriendsMethodUpdatePreference             = 0x10
	FriendsMethodGetBasicInfo                 = 0x11
	FriendsMethodDeleteFriendFlags            = 0x12
	FriendsMethodCheckSettingStatus           = 0x13
	FriendsMethodGetRequestBlockSettings      = 0x14
)

type FriendsProtocol struct {
	server                              *nex.Server
	UpdateAndGetAllInformationHandler   func(err error, client *nex.Client, callID uint32, nnaInfo *NNAInfo, presence *NintendoPresenceV2, birthday *nex.DateTime)
	AddFriendHandler                    func(err error, client *nex.Client, callID uint32, pid uint32)
	AddFriendByNameHandler              func(err error, client *nex.Client, callID uint32, username string)
	RemoveFriendHandler                 func(err error, client *nex.Client, callID uint32, pid uint32)
	AddFriendRequestHandler             func(err error, client *nex.Client, callID uint32, unknown1 uint32, unknown2 uint8, unknown3 string, unknown4 uint8, unknown5 string, gameKey *GameKey, unknown6 *nex.DateTime)
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
	DeleteFriendFlagsHandler            func(err error, client *nex.Client, callID uint32, notifications []*PersistentNotification)
	CheckSettingStatusHandler           func(err error, client *nex.Client, callID uint32)
	GetRequestBlockSettingsHandler      func(err error, client *nex.Client, callID uint32, unknowns []uint32)
}

type BlacklistedPrincipal struct {
	principalBasicInfo *PrincipalBasicInfo
	gameKey            *GameKey
	blackListedSince   *nex.DateTime

	nex.Structure
}

func NewBlacklistedPrincipal() *BlacklistedPrincipal {
	return &BlacklistedPrincipal{}
}

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

	blacklistedPrincipal.principalBasicInfo = principalBasicInfo
	blacklistedPrincipal.gameKey = gameKey
	blacklistedPrincipal.blackListedSince = blackListedSince

	return nil
}

type Comment struct {
	unknown     uint8
	contents    string
	lastChanged *nex.DateTime

	nex.Structure
}

func NewComment() *Comment {
	return &Comment{}
}

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

	comment.unknown = unknown
	comment.contents = contents
	comment.lastChanged = lastChanged

	return nil
}

type FriendInfo struct {
	nnaInfo      *NNAInfo
	presence     *NintendoPresenceV2
	status       *Comment
	becameFriend *nex.DateTime
	lastOnline   *nex.DateTime
	unknown      uint64

	nex.Structure
}

func NewFriendInfo() *FriendInfo {
	return &FriendInfo{}
}

type FriendRequest struct {
	principalInfo *PrincipalBasicInfo
	message       *FriendRequestMessage
	sentOn        *nex.DateTime

	nex.Structure
}

func NewFriendRequest() *FriendRequest {
	return &FriendRequest{}
}

type FriendRequestMessage struct {
	unknown1  uint64
	unknown2  uint8
	unknown3  uint8
	message   string
	unknown4  uint8
	unknown5  string
	gameKey   *GameKey
	unknown6  *nex.DateTime
	expiresOn *nex.DateTime

	nex.Structure
}

func NewFriendRequestMessage() *FriendRequestMessage {
	return &FriendRequestMessage{}
}

type GameKey struct {
	titleID      uint64
	titleVersion uint16

	nex.Structure
}

func NewGameKey() *GameKey {
	return &GameKey{}
}

func (gameKey *GameKey) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 10 {
		return errors.New("[GameKey::ExtractFromStream] Data size too small")
	}

	gameKey.titleID = stream.ReadUInt64LE()
	gameKey.titleVersion = stream.ReadUInt16LE()

	return nil
}

type MiiV2 struct {
	name     string
	unknown1 uint8
	unknown2 uint8
	data     []byte
	datetime *nex.DateTime

	nex.Structure
}

func NewMiiV2() *MiiV2 {
	return &MiiV2{}
}

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

	mii.name = name
	mii.unknown1 = unknown1
	mii.unknown2 = unknown2
	mii.data = data
	mii.datetime = datetime

	return nil
}

type NintendoPresenceV2 struct {
	changedFlags    uint32
	isOnline        bool
	gameKey         *GameKey
	unknown1        uint8
	message         string
	unknown2        uint32
	unknown3        uint8
	gameServerID    uint32
	unknown4        uint32
	pid             uint32
	gatheringID     uint32
	applicationData []byte
	unknown5        uint8
	unknown6        uint8
	unknown7        uint8

	nex.Structure
}

func NewNintendoPresenceV2() *NintendoPresenceV2 {
	return &NintendoPresenceV2{}
}

func (presence *NintendoPresenceV2) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 40 {
		// length check for the following fixed-size data
		// changedFlags + isOnline + gameKey + gameKey + unknown1 + unknown2 + unknown3 + gameServerID + unknown4 + pid + gatheringID + unknown5 + unknown6 + unknown7
		return errors.New("[NintendoPresenceV2::ExtractFromStream] Data size too small")
	}

	changedFlags := stream.ReadUInt32LE()
	isOnline := (stream.ReadUInt8() == 1)
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

	presence.changedFlags = changedFlags
	presence.isOnline = isOnline
	presence.gameKey = gameKey
	presence.unknown1 = unknown1
	presence.message = message
	presence.unknown2 = unknown2
	presence.unknown3 = unknown3
	presence.gameServerID = gameServerID
	presence.unknown4 = unknown4
	presence.pid = pid
	presence.gatheringID = gatheringID
	presence.applicationData = applicationData
	presence.unknown5 = unknown5
	presence.unknown6 = unknown6
	presence.unknown7 = unknown7

	return nil
}

type NNAInfo struct {
	principalBasicInfo *PrincipalBasicInfo
	unknown1           uint8
	unknown2           uint8

	nex.Structure
}

func NewNNAInfo() *NNAInfo {
	return &NNAInfo{}
}

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

	nnaInfo.principalBasicInfo = principalBasicInfo
	nnaInfo.unknown1 = unknown1
	nnaInfo.unknown2 = unknown2

	return nil
}

type PersistentNotification struct {
	unknown1 uint64
	unknown2 uint32
	unknown3 uint32
	unknown4 uint32
	unknown5 string

	nex.Structure
}

func NewPersistentNotification() *PersistentNotification {
	return &PersistentNotification{}
}

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

	notification.unknown1 = unknown1
	notification.unknown2 = unknown2
	notification.unknown3 = unknown3
	notification.unknown4 = unknown4
	notification.unknown5 = unknown5

	return nil
}

type PrincipalBasicInfo struct {
	pid     uint32
	nnid    string
	mii     *MiiV2
	unknown uint8

	nex.Structure
}

func NewPrincipalBasicInfo() *PrincipalBasicInfo {
	return &PrincipalBasicInfo{}
}

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

	principalInfo.pid = pid
	principalInfo.nnid = nnid
	principalInfo.mii = miiV2
	principalInfo.unknown = unknown

	return nil
}

type PrincipalPreference struct {
	unknown1 bool
	unknown2 bool
	unknown3 bool

	nex.Structure
}

func NewPrincipalPreference() *PrincipalPreference {
	return &PrincipalPreference{}
}

func (preference *PrincipalPreference) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 1 {
		// length check for the following fixed-size data
		// unknown1 + unknown2 + unknown3
		return errors.New("[PrincipalPreference::ExtractFromStream] Data size too small")
	}

	preference.unknown1 = (stream.ReadUInt8() == 1)
	preference.unknown2 = (stream.ReadUInt8() == 1)
	preference.unknown3 = (stream.ReadUInt8() == 1)

	return nil
}

type PrincipalRequestBlockSetting struct {
	unknown1 uint32
	unknown2 bool
}

func NewPrincipalRequestBlockSetting() *PrincipalRequestBlockSetting {
	return &PrincipalRequestBlockSetting{}
}

func (friendsProtocol *FriendsProtocol) Setup() {
	nexServer := friendsProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.GetRMCRequest()

		if FriendsProtocolID == request.GetProtocolID() {
			switch request.GetMethodID() {
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
			case FriendsMethodDeleteFriendFlags:
				go friendsProtocol.handleDeleteFriendFlags(packet)
			case FriendsMethodCheckSettingStatus:
				go friendsProtocol.handleCheckSettingStatus(packet)
			case FriendsMethodGetRequestBlockSettings:
				go friendsProtocol.handleGetRequestBlockSettings(packet)
			default:
				fmt.Printf("Unsupported Friends (WiiU) method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

func (friendsProtocol *FriendsProtocol) respondNotImplemented(packet nex.PacketInterface) {
	client := packet.GetSender()
	request := packet.GetRMCRequest()

	rmcResponse := nex.NewRMCResponse(FriendsProtocolID, request.GetCallID())
	rmcResponse.SetError(0x80010002)

	rmcResponseBytes := rmcResponse.Bytes()

	var responsePacket nex.PacketInterface
	if packet.GetVersion() == 1 {
		responsePacket, _ = nex.NewPacketV1(client, nil)
	} else {
		responsePacket, _ = nex.NewPacketV0(client, nil)
	}

	responsePacket.SetVersion(packet.GetVersion())
	responsePacket.SetSource(packet.GetDestination())
	responsePacket.SetDestination(packet.GetSource())
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	friendsProtocol.server.Send(responsePacket)
}

func (friendsProtocol *FriendsProtocol) UpdateAndGetAllInformation(handler func(err error, client *nex.Client, callID uint32, nnaInfo *NNAInfo, presence *NintendoPresenceV2, birthday *nex.DateTime)) {
	friendsProtocol.UpdateAndGetAllInformationHandler = handler
}

func (friendsProtocol *FriendsProtocol) AddFriend(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	friendsProtocol.AddFriendHandler = handler
}

func (friendsProtocol *FriendsProtocol) AddFriendByName(handler func(err error, client *nex.Client, callID uint32, username string)) {
	friendsProtocol.AddFriendByNameHandler = handler
}

func (friendsProtocol *FriendsProtocol) RemoveFriend(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	friendsProtocol.RemoveFriendHandler = handler
}

func (friendsProtocol *FriendsProtocol) AddFriendRequest(handler func(err error, client *nex.Client, callID uint32, unknown1 uint32, unknown2 uint8, unknown3 string, unknown4 uint8, unknown5 string, gameKey *GameKey, unknown6 *nex.DateTime)) {
	friendsProtocol.AddFriendRequestHandler = handler
}

func (friendsProtocol *FriendsProtocol) CancelFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	friendsProtocol.CancelFriendRequestHandler = handler
}

func (friendsProtocol *FriendsProtocol) AcceptFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	friendsProtocol.AcceptFriendRequestHandler = handler
}

func (friendsProtocol *FriendsProtocol) DeleteFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	friendsProtocol.DeleteFriendRequestHandler = handler
}

func (friendsProtocol *FriendsProtocol) DenyFriendRequest(handler func(err error, client *nex.Client, callID uint32, id uint64)) {
	friendsProtocol.DenyFriendRequestHandler = handler
}

func (friendsProtocol *FriendsProtocol) MarkFriendRequestsAsReceived(handler func(err error, client *nex.Client, callID uint32, ids []uint64)) {
	friendsProtocol.MarkFriendRequestsAsReceivedHandler = handler
}

func (friendsProtocol *FriendsProtocol) AddBlackList(handler func(err error, client *nex.Client, callID uint32, blacklistedPrincipal *BlacklistedPrincipal)) {
	friendsProtocol.AddBlackListHandler = handler
}

func (friendsProtocol *FriendsProtocol) RemoveBlackList(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	friendsProtocol.RemoveBlackListHandler = handler
}

func (friendsProtocol *FriendsProtocol) UpdatePresence(handler func(err error, client *nex.Client, callID uint32, presence *NintendoPresenceV2)) {
	friendsProtocol.UpdatePresenceHandler = handler
}

func (friendsProtocol *FriendsProtocol) UpdateMii(handler func(err error, client *nex.Client, callID uint32, mii *MiiV2)) {
	friendsProtocol.UpdateMiiHandler = handler
}

func (friendsProtocol *FriendsProtocol) UpdateComment(handler func(err error, client *nex.Client, callID uint32, comment *Comment)) {
	friendsProtocol.UpdateCommentHandler = handler
}

func (friendsProtocol *FriendsProtocol) UpdatePreference(handler func(err error, client *nex.Client, callID uint32, preference *PrincipalPreference)) {
	friendsProtocol.UpdatePreferenceHandler = handler
}

func (friendsProtocol *FriendsProtocol) GetBasicInfo(handler func(err error, client *nex.Client, callID uint32, pids []uint32)) {
	friendsProtocol.GetBasicInfoHandler = handler
}

func (friendsProtocol *FriendsProtocol) DeleteFriendFlags(handler func(err error, client *nex.Client, callID uint32, notifications []*PersistentNotification)) {
	friendsProtocol.DeleteFriendFlagsHandler = handler
}

func (friendsProtocol *FriendsProtocol) CheckSettingStatus(handler func(err error, client *nex.Client, callID uint32)) {
	friendsProtocol.CheckSettingStatusHandler = handler
}

func (friendsProtocol *FriendsProtocol) GetRequestBlockSettings(handler func(err error, client *nex.Client, callID uint32, unknowns []uint32)) {
	friendsProtocol.GetRequestBlockSettingsHandler = handler
}

func (friendsProtocol *FriendsProtocol) handleUpdateAndGetAllInformation(packet nex.PacketInterface) {
	if friendsProtocol.UpdateAndGetAllInformationHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::UpdateAndGetAllInformation not implemented")
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4+1+1+8 {
		// length check for the following fixed-size data
		// unknown1 + unknown2 + unknown4 + gameKey + unknown6
		err := errors.New("[FriendsProtocol::AddFriendRequest] Data holder not long enough for PID")
		go friendsProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", nil, nil)
		return
	}

	unknown1 := parametersStream.ReadUInt32LE()
	unknown2 := parametersStream.ReadUInt8()
	unknown3, err := parametersStream.ReadString()

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

	go friendsProtocol.AddFriendRequestHandler(nil, client, callID, unknown1, unknown2, unknown3, unknown4, unknown5, gameKey, unknown6)
}

func (friendsProtocol *FriendsProtocol) handleCancelFriendRequest(packet nex.PacketInterface) {
	if friendsProtocol.CancelFriendRequestHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::CancelFriendRequest not implemented")
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := NewStreamIn(parameters, friendsProtocol.server)

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsProtocol::GetBasicInfo] Data missing list length")
		go friendsProtocol.GetBasicInfoHandler(err, client, callID, make([]uint32, 0))
		return
	}

	pids := parametersStream.ReadListUInt32LE()

	go friendsProtocol.GetBasicInfoHandler(nil, client, callID, pids)
}

func (friendsProtocol *FriendsProtocol) handleDeleteFriendFlags(packet nex.PacketInterface) {
	if friendsProtocol.DeleteFriendFlagsHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::DeleteFriendFlags not implemented")
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := NewStreamIn(parameters, friendsProtocol.server)

	persistentNotifications, err := parametersStream.ReadListPersistentNotification()

	if err != nil {
		go friendsProtocol.DeleteFriendFlagsHandler(err, client, callID, nil)
		return
	}

	go friendsProtocol.DeleteFriendFlagsHandler(nil, client, callID, persistentNotifications)
}

func (friendsProtocol *FriendsProtocol) handleCheckSettingStatus(packet nex.PacketInterface) {
	if friendsProtocol.CheckSettingStatusHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::CheckSettingStatus not implemented")
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()

	go friendsProtocol.CheckSettingStatusHandler(nil, client, callID)
}

func (friendsProtocol *FriendsProtocol) handleGetRequestBlockSettings(packet nex.PacketInterface) {
	if friendsProtocol.GetRequestBlockSettingsHandler == nil {
		fmt.Println("[Warning] FriendsProtocol::GetRequestBlockSettings not implemented")
		go friendsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsProtocol::GetRequestBlockSettings] Data missing list length")
		go friendsProtocol.GetRequestBlockSettingsHandler(err, client, callID, make([]uint32, 0))
		return
	}

	unknowns := parametersStream.ReadListUInt32LE()

	go friendsProtocol.GetRequestBlockSettingsHandler(nil, client, callID, unknowns)
}

func NewFriendsProtocol(server *nex.Server) *FriendsProtocol {
	friendsProtocol := &FriendsProtocol{server: server}

	friendsProtocol.Setup()

	return friendsProtocol
}
