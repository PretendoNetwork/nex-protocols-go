package nexproto

import (
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
	UpdateAndGetAllInformationHandler   func(client *nex.Client, callID uint32, nnaInfo *NNAInfo, presence *NintendoPresenceV2, birthday *nex.DateTime)
	AddFriendHandler                    func(client *nex.Client, callID uint32, pid uint32)
	AddFriendByNameHandler              func(client *nex.Client, callID uint32, username string)
	RemoveFriendHandler                 func(client *nex.Client, callID uint32, pid uint32)
	AddFriendRequestHandler             func(client *nex.Client, callID uint32, unknown1 uint32, unknown2 uint8, unknown3 string, unknown4 uint8, unknown5 string, gameKey *GameKey, unknown6 *nex.DateTime)
	CancelFriendRequestHandler          func(client *nex.Client, callID uint32, id uint64)
	AcceptFriendRequestHandler          func(client *nex.Client, callID uint32, id uint64)
	DeleteFriendRequestHandler          func(client *nex.Client, callID uint32, id uint64)
	DenyFriendRequestHandler            func(client *nex.Client, callID uint32, id uint64)
	MarkFriendRequestsAsReceivedHandler func(client *nex.Client, callID uint32, ids []uint64)
	AddBlackListHandler                 func(client *nex.Client, callID uint32, blacklistedPrincipal *BlacklistedPrincipal)
	RemoveBlackListHandler              func(client *nex.Client, callID uint32, pid uint32)
	UpdatePresenceHandler               func(client *nex.Client, callID uint32, presence *NintendoPresenceV2)
	UpdateMiiHandler                    func(client *nex.Client, callID uint32, mii *MiiV2)
	UpdateCommentHandler                func(client *nex.Client, callID uint32, comment *Comment)
	UpdatePreferenceHandler             func(client *nex.Client, callID uint32, preference *PrincipalPreference)
	GetBasicInfoHandler                 func(client *nex.Client, callID uint32, pids []uint32)
	DeleteFriendFlagsHandler            func(client *nex.Client, callID uint32, notifications []*PersistentNotification)
	CheckSettingStatusHandler           func(client *nex.Client, callID uint32)
	GetRequestBlockSettingsHandler      func(client *nex.Client, callID uint32, unknowns []uint32)
}

type BlacklistedPrincipal struct {
	principalInfo    *PrincipalBasicInfo
	gameKey          *GameKey
	blackListedSince *nex.DateTime
}

func (blacklistedPrincipal *BlacklistedPrincipal) ExtractFromStream(stream *nex.StreamIn) {
	blacklistedPrincipal.principalInfo = &PrincipalBasicInfo{}
	blacklistedPrincipal.principalInfo.ExtractFromStream(stream)
	blacklistedPrincipal.gameKey = &GameKey{}
	blacklistedPrincipal.gameKey.ExtractFromStream(stream)
	blacklistedPrincipal.blackListedSince = nex.NewDateTime(stream.ReadU64LENext(1)[0])
}

type Comment struct {
	unknown     uint8
	contents    string
	lastChanged *nex.DateTime
}

func (comment *Comment) ExtractFromStream(stream *nex.StreamIn) {
	comment.unknown = stream.ReadByteNext()
	comment.contents = stream.ReadStringNext()
	comment.lastChanged = nex.NewDateTime(stream.ReadU64LENext(1)[0])
}

type FriendInfo struct {
	nnaInfo      *NNAInfo
	presence     *NintendoPresenceV2
	status       *Comment
	becameFriend *nex.DateTime
	lastOnline   *nex.DateTime
	unknown      uint64
}

type FriendRequest struct {
	principalInfo *PrincipalBasicInfo
	message       *FriendRequestMessage
	sentOn        *nex.DateTime
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
}

type GameKey struct {
	titleID      uint64
	titleVersion uint16
}

func (gameKey *GameKey) ExtractFromStream(stream *nex.StreamIn) {
	gameKey.titleID = stream.ReadU64LENext(1)[0]
	gameKey.titleVersion = stream.ReadU16LENext(1)[0]
}

type MiiV2 struct {
	name     string
	unknown1 uint8
	unknown2 uint8
	data     []byte
	datetime *nex.DateTime
}

func (mii *MiiV2) ExtractFromStream(stream *nex.StreamIn) {
	mii.name = stream.ReadStringNext()
	mii.unknown1 = stream.ReadByteNext()
	mii.unknown2 = stream.ReadByteNext()
	mii.data = stream.ReadBufferNext()
	mii.datetime = nex.NewDateTime(stream.ReadU64LENext(1)[0])
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
}

func (presence *NintendoPresenceV2) ExtractFromStream(stream *nex.StreamIn) {
	presence.changedFlags = stream.ReadU32LENext(1)[0]
	presence.isOnline = (stream.ReadByteNext() == 1)
	presence.gameKey = &GameKey{}
	presence.gameKey.ExtractFromStream(stream)
	presence.unknown1 = stream.ReadByteNext()
	presence.message = stream.ReadStringNext()
	presence.unknown2 = stream.ReadU32LENext(1)[0]
	presence.unknown3 = stream.ReadByteNext()
	presence.gameServerID = stream.ReadU32LENext(1)[0]
	presence.unknown4 = stream.ReadU32LENext(1)[0]
	presence.pid = stream.ReadU32LENext(1)[0]
	presence.gatheringID = stream.ReadU32LENext(1)[0]
	presence.applicationData = stream.ReadBufferNext()
	presence.unknown5 = stream.ReadByteNext()
	presence.unknown6 = stream.ReadByteNext()
	presence.unknown7 = stream.ReadByteNext()
}

type NNAInfo struct {
	principalInfo *PrincipalBasicInfo
	unknown1      uint8
	unknown2      uint8
}

func (nnaInfo *NNAInfo) ExtractFromStream(stream *nex.StreamIn) {
	nnaInfo.principalInfo = &PrincipalBasicInfo{}
	nnaInfo.principalInfo.ExtractFromStream(stream)
	nnaInfo.unknown1 = stream.ReadByteNext()
	nnaInfo.unknown2 = stream.ReadByteNext()
}

type PersistentNotification struct {
	unknown1 uint64
	unknown2 uint32
	unknown3 uint32
	unknown4 uint32
	unknown5 string
}

func (notification *PersistentNotification) ExtractFromStream(stream *nex.StreamIn) {
	notification.unknown1 = stream.ReadU64LENext(1)[0]
	notification.unknown2 = stream.ReadU32LENext(1)[0]
	notification.unknown3 = stream.ReadU32LENext(1)[0]
	notification.unknown4 = stream.ReadU32LENext(1)[0]
	notification.unknown5 = stream.ReadStringNext()
}

type PrincipalBasicInfo struct {
	pid     uint32
	nnid    string
	mii     *MiiV2
	unknown uint8
}

func (principalInfo *PrincipalBasicInfo) ExtractFromStream(stream *nex.StreamIn) {
	principalInfo.pid = stream.ReadU32LENext(1)[0]
	principalInfo.nnid = stream.ReadStringNext()
	principalInfo.mii = &MiiV2{}
	principalInfo.mii.ExtractFromStream(stream)
	principalInfo.unknown = stream.ReadByteNext()
}

type PrincipalPreference struct {
	unknown1 bool
	unknown2 bool
	unknown3 bool
}

func (preference *PrincipalPreference) ExtractFromStream(stream *nex.StreamIn) {
	preference.unknown1 = (stream.ReadByteNext() == 1)
	preference.unknown2 = (stream.ReadByteNext() == 1)
	preference.unknown3 = (stream.ReadByteNext() == 1)
}

type PrincipalRequestBlockSetting struct {
	unknown1 uint32
	unknown2 bool
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
		responsePacket = nex.NewPacketV0(client, nil)
	} else {
		responsePacket = nex.NewPacketV1(client, nil)
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

func (friendsProtocol *FriendsProtocol) UpdateAndGetAllInformation(handler func(client *nex.Client, callID uint32, nnaInfo *NNAInfo, presence *NintendoPresenceV2, birthday *nex.DateTime)) {
	friendsProtocol.UpdateAndGetAllInformationHandler = handler
}

func (friendsProtocol *FriendsProtocol) AddFriend(handler func(client *nex.Client, callID uint32, pid uint32)) {
	friendsProtocol.AddFriendHandler = handler
}

func (friendsProtocol *FriendsProtocol) AddFriendByName(handler func(client *nex.Client, callID uint32, username string)) {
	friendsProtocol.AddFriendByNameHandler = handler
}

func (friendsProtocol *FriendsProtocol) RemoveFriend(handler func(client *nex.Client, callID uint32, pid uint32)) {
	friendsProtocol.RemoveFriendHandler = handler
}

func (friendsProtocol *FriendsProtocol) AddFriendRequest(handler func(client *nex.Client, callID uint32, unknown1 uint32, unknown2 uint8, unknown3 string, unknown4 uint8, unknown5 string, gameKey *GameKey, unknown6 *nex.DateTime)) {
	friendsProtocol.AddFriendRequestHandler = handler
}

func (friendsProtocol *FriendsProtocol) CancelFriendRequest(handler func(client *nex.Client, callID uint32, id uint64)) {
	friendsProtocol.CancelFriendRequestHandler = handler
}

func (friendsProtocol *FriendsProtocol) AcceptFriendRequest(handler func(client *nex.Client, callID uint32, id uint64)) {
	friendsProtocol.AcceptFriendRequestHandler = handler
}

func (friendsProtocol *FriendsProtocol) DeleteFriendRequest(handler func(client *nex.Client, callID uint32, id uint64)) {
	friendsProtocol.DeleteFriendRequestHandler = handler
}

func (friendsProtocol *FriendsProtocol) DenyFriendRequest(handler func(client *nex.Client, callID uint32, id uint64)) {
	friendsProtocol.DenyFriendRequestHandler = handler
}

func (friendsProtocol *FriendsProtocol) MarkFriendRequestsAsReceived(handler func(client *nex.Client, callID uint32, ids []uint64)) {
	friendsProtocol.MarkFriendRequestsAsReceivedHandler = handler
}

func (friendsProtocol *FriendsProtocol) AddBlackList(handler func(client *nex.Client, callID uint32, blacklistedPrincipal *BlacklistedPrincipal)) {
	friendsProtocol.AddBlackListHandler = handler
}

func (friendsProtocol *FriendsProtocol) RemoveBlackList(handler func(client *nex.Client, callID uint32, pid uint32)) {
	friendsProtocol.RemoveBlackListHandler = handler
}

func (friendsProtocol *FriendsProtocol) UpdatePresence(handler func(client *nex.Client, callID uint32, presence *NintendoPresenceV2)) {
	friendsProtocol.UpdatePresenceHandler = handler
}

func (friendsProtocol *FriendsProtocol) UpdateMii(handler func(client *nex.Client, callID uint32, mii *MiiV2)) {
	friendsProtocol.UpdateMiiHandler = handler
}

func (friendsProtocol *FriendsProtocol) UpdateComment(handler func(client *nex.Client, callID uint32, comment *Comment)) {
	friendsProtocol.UpdateCommentHandler = handler
}

func (friendsProtocol *FriendsProtocol) UpdatePreference(handler func(client *nex.Client, callID uint32, preference *PrincipalPreference)) {
	friendsProtocol.UpdatePreferenceHandler = handler
}

func (friendsProtocol *FriendsProtocol) GetBasicInfo(handler func(client *nex.Client, callID uint32, pids []uint32)) {
	friendsProtocol.GetBasicInfoHandler = handler
}

func (friendsProtocol *FriendsProtocol) DeleteFriendFlags(handler func(client *nex.Client, callID uint32, notifications []*PersistentNotification)) {
	friendsProtocol.DeleteFriendFlagsHandler = handler
}

func (friendsProtocol *FriendsProtocol) CheckSettingStatus(handler func(client *nex.Client, callID uint32)) {
	friendsProtocol.CheckSettingStatusHandler = handler
}

func (friendsProtocol *FriendsProtocol) GetRequestBlockSettings(handler func(client *nex.Client, callID uint32, unknowns []uint32)) {
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

	nnaInfo := &NNAInfo{}
	presence := &NintendoPresenceV2{}
	dateTime := nex.NewDateTime(0)

	nnaInfo.ExtractFromStream(parametersStream)
	presence.ExtractFromStream(parametersStream)

	go friendsProtocol.UpdateAndGetAllInformationHandler(client, callID, nnaInfo, presence, dateTime)
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

	pid := parametersStream.ReadU32LENext(1)[0]

	go friendsProtocol.AddFriendHandler(client, callID, pid)
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

	username := parametersStream.ReadStringNext()

	go friendsProtocol.AddFriendByNameHandler(client, callID, username)
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

	pid := parametersStream.ReadU32LENext(1)[0]

	go friendsProtocol.RemoveFriendHandler(client, callID, pid)
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

	unknown1 := parametersStream.ReadU32LENext(1)[0]
	unknown2 := parametersStream.ReadByteNext()
	unknown3 := parametersStream.ReadStringNext()
	unknown4 := parametersStream.ReadByteNext()
	unknown5 := parametersStream.ReadStringNext()
	gameKey := &GameKey{}
	gameKey.ExtractFromStream(parametersStream)
	unknown6 := nex.NewDateTime(parametersStream.ReadU64LENext(1)[0])

	go friendsProtocol.AddFriendRequestHandler(client, callID, unknown1, unknown2, unknown3, unknown4, unknown5, gameKey, unknown6)
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

	id := parametersStream.ReadU64LENext(1)[0]

	go friendsProtocol.CancelFriendRequestHandler(client, callID, id)
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

	id := parametersStream.ReadU64LENext(1)[0]

	go friendsProtocol.AcceptFriendRequestHandler(client, callID, id)
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

	id := parametersStream.ReadU64LENext(1)[0]

	go friendsProtocol.DeleteFriendRequestHandler(client, callID, id)
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

	id := parametersStream.ReadU64LENext(1)[0]

	go friendsProtocol.DenyFriendRequestHandler(client, callID, id)
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

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	idCount := parametersStream.ReadU32LENext(1)[0]
	ids := make([]uint64, 0)

	for i := 0; i < int(idCount); i++ {
		id := parametersStream.ReadU64LENext(1)[0]
		ids = append(ids, id)
	}

	go friendsProtocol.MarkFriendRequestsAsReceivedHandler(client, callID, ids)
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

	blacklistedPrincipal := &BlacklistedPrincipal{}
	blacklistedPrincipal.ExtractFromStream(parametersStream)

	go friendsProtocol.AddBlackListHandler(client, callID, blacklistedPrincipal)
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

	pid := parametersStream.ReadU32LENext(1)[0]

	go friendsProtocol.RemoveBlackListHandler(client, callID, pid)
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

	presence := &NintendoPresenceV2{}
	presence.ExtractFromStream(parametersStream)

	go friendsProtocol.UpdatePresenceHandler(client, callID, presence)
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

	mii := &MiiV2{}
	mii.ExtractFromStream(parametersStream)

	go friendsProtocol.UpdateMiiHandler(client, callID, mii)
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

	comment := &Comment{}
	comment.ExtractFromStream(parametersStream)

	go friendsProtocol.UpdateCommentHandler(client, callID, comment)
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

	preference := &PrincipalPreference{}
	preference.ExtractFromStream(parametersStream)

	go friendsProtocol.UpdatePreferenceHandler(client, callID, preference)
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

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	pidCount := parametersStream.ReadU32LENext(1)[0]
	pids := make([]uint32, 0)

	for i := 0; i < int(pidCount); i++ {
		pid := parametersStream.ReadU32LENext(1)[0]
		pids = append(pids, pid)
	}

	go friendsProtocol.GetBasicInfoHandler(client, callID, pids)
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

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	notificationCount := parametersStream.ReadU32LENext(1)[0]
	notifications := make([]*PersistentNotification, 0)

	for i := 0; i < int(notificationCount); i++ {
		notification := &PersistentNotification{}
		notification.ExtractFromStream(parametersStream)
		notifications = append(notifications, notification)
	}

	go friendsProtocol.DeleteFriendFlagsHandler(client, callID, notifications)
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

	go friendsProtocol.CheckSettingStatusHandler(client, callID)
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

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	unknownCount := parametersStream.ReadU32LENext(1)[0]
	unknowns := make([]uint32, 0)

	for i := 0; i < int(unknownCount); i++ {
		unknown := parametersStream.ReadU32LENext(1)[0]
		unknowns = append(unknowns, unknown)
	}

	go friendsProtocol.GetRequestBlockSettingsHandler(client, callID, unknowns)
}

func NewFriendsProtocol(server *nex.Server) *FriendsProtocol {
	friendsProtocol := &FriendsProtocol{server: server}

	friendsProtocol.Setup()

	return friendsProtocol
}
