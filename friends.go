package nexproto

import (
	"fmt"
	"errors"
	nex "../nex-go"
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
	principalInfo    *PrincipalBasicInfo
	gameKey          *GameKey
	blackListedSince *nex.DateTime
}

func (blacklistedPrincipal *BlacklistedPrincipal) ExtractFromStream(stream *nex.StreamIn) error {
	principalInfo := &PrincipalBasicInfo{}
	err := principalInfo.ExtractFromStream(stream)

	if err != nil {
		return err
	}

	gameKey := &GameKey{}
	err = principalInfo.ExtractFromStream(stream)

	if err != nil {
		return err
	}

	if len(stream.Bytes()[stream.ByteOffset():]) < 8 {
		return errors.New("[DataStorePersistenceTarget::ExtractFromStream] Data size too small")
	}

	blacklistedPrincipal.principalInfo = principalInfo
	blacklistedPrincipal.gameKey = gameKey
	blacklistedPrincipal.blackListedSince = nex.NewDateTime(stream.ReadU64LENext(1)[0])

	return nil
}

type Comment struct {
	unknown     uint8
	contents    string
	lastChanged *nex.DateTime
}

func (comment *Comment) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 9 { // unknown byte + datetime uint64
		return errors.New("[Comment::ExtractFromStream] Data size too small")
	}

	unknown := stream.ReadByteNext()

	contents, err := stream.ReadStringNext()

	if err != nil {
		return err
	}

	lastChanged := nex.NewDateTime(stream.ReadU64LENext(1)[0])

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

func (gameKey *GameKey) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 10 {
		return errors.New("[GameKey::ExtractFromStream] Data size too small")
	}

	gameKey.titleID = stream.ReadU64LENext(1)[0]
	gameKey.titleVersion = stream.ReadU16LENext(1)[0]

	return nil
}

type MiiV2 struct {
	name     string
	unknown1 uint8
	unknown2 uint8
	data     []byte
	datetime *nex.DateTime
}

func (mii *MiiV2) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 10 { // 2 unknown bytes + datetime uint64
		return errors.New("[MiiV2::ExtractFromStream] Data size too small")
	}

	name, err := stream.ReadStringNext()

	if err != nil {
		return err
	}

	unknown1 := stream.ReadByteNext()
	unknown2 := stream.ReadByteNext()
	data, err := stream.ReadBufferNext()

	if err != nil {
		return err
	}

	datetime := nex.NewDateTime(stream.ReadU64LENext(1)[0])

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
}

func (presence *NintendoPresenceV2) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 40 {
		// length check for the following fixed-size data
		// changedFlags + isOnline + gameKey + gameKey + unknown1 + unknown2 + unknown3 + gameServerID + unknown4 + pid + gatheringID + unknown5 + unknown6 + unknown7
		return errors.New("[NintendoPresenceV2::ExtractFromStream] Data size too small")
	}

	changedFlags := stream.ReadU32LENext(1)[0]

	isOnline := (stream.ReadByteNext() == 1)

	gameKey := &GameKey{}
	err := gameKey.ExtractFromStream(stream)

	if err != nil {
		return err
	}

	unknown1 := stream.ReadByteNext()

	message, err := stream.ReadStringNext()

	if err != nil {
		return err
	}

	unknown2 := stream.ReadU32LENext(1)[0]

	unknown3 := stream.ReadByteNext()

	gameServerID := stream.ReadU32LENext(1)[0]

	unknown4 := stream.ReadU32LENext(1)[0]

	pid := stream.ReadU32LENext(1)[0]

	gatheringID := stream.ReadU32LENext(1)[0]

	applicationData, err := stream.ReadBufferNext()

	if err != nil {
		return err
	}

	unknown5 := stream.ReadByteNext()

	unknown6 := stream.ReadByteNext()

	unknown7 := stream.ReadByteNext()


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
	principalInfo *PrincipalBasicInfo
	unknown1      uint8
	unknown2      uint8
}

func (nnaInfo *NNAInfo) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 2 {
		// length check for the following fixed-size data
		// unknown1 + unknown2
		return errors.New("[NNAInfo::ExtractFromStream] Data size too small")
	}

	principalInfo := &PrincipalBasicInfo{}
	err := principalInfo.ExtractFromStream(stream)

	if err != nil {
		return err
	}

	unknown1 := stream.ReadByteNext()
	unknown2 := stream.ReadByteNext()

	nnaInfo.principalInfo = principalInfo
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
}

func (notification *PersistentNotification) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 20 {
		// length check for the following fixed-size data
		// unknown1 + unknown2 + unknown3 + unknown4
		return errors.New("[PersistentNotification::ExtractFromStream] Data size too small")
	}

	unknown1 := stream.ReadU64LENext(1)[0]
	unknown2 := stream.ReadU32LENext(1)[0]
	unknown3 := stream.ReadU32LENext(1)[0]
	unknown4 := stream.ReadU32LENext(1)[0]
	unknown5, err := stream.ReadStringNext()

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
}

func (principalInfo *PrincipalBasicInfo) ExtractFromStream(stream *nex.StreamIn) error {

	if len(stream.Bytes()[stream.ByteOffset():]) < 4 {
		return errors.New("[PrincipalBasicInfo::ExtractFromStream] Data size too small")
	}

	pid := stream.ReadU32LENext(1)[0]

	nnid, err := stream.ReadStringNext()

	if err != nil {
		return err
	}

	mii := &MiiV2{}
	err = mii.ExtractFromStream(stream)

	if err != nil {
		return err
	}

	if len(stream.Bytes()[stream.ByteOffset():]) < 1 {
		return errors.New("[PrincipalBasicInfo::ExtractFromStream] Data size too small")
	}

	unknown := stream.ReadByteNext()

	principalInfo.pid = pid
	principalInfo.nnid = nnid
	principalInfo.mii = mii
	principalInfo.unknown = unknown

	return nil
}

type PrincipalPreference struct {
	unknown1 bool
	unknown2 bool
	unknown3 bool
}

func (preference *PrincipalPreference) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 1 {
		// length check for the following fixed-size data
		// unknown1 + unknown2 + unknown3
		return errors.New("[PrincipalPreference::ExtractFromStream] Data size too small")
	}

	preference.unknown1 = (stream.ReadByteNext() == 1)
	preference.unknown2 = (stream.ReadByteNext() == 1)
	preference.unknown3 = (stream.ReadByteNext() == 1)

	return nil
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
		responsePacket, _ = nex.NewPacketV0(client, nil)
	} else {
		responsePacket, _ = nex.NewPacketV1(client, nil)
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

	nnaInfo := &NNAInfo{}
	presence := &NintendoPresenceV2{}

	err := nnaInfo.ExtractFromStream(parametersStream)

	if err != nil {
		go friendsProtocol.UpdateAndGetAllInformationHandler(err, client, callID, nnaInfo, presence, nex.NewDateTime(0))
		return
	}

	err = presence.ExtractFromStream(parametersStream)

	if err != nil {
		go friendsProtocol.UpdateAndGetAllInformationHandler(err, client, callID, nnaInfo, presence, nex.NewDateTime(0))
		return
	}


	birthday := nex.NewDateTime(parametersStream.ReadU64LENext(1)[0])

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
		go friendsProtocol.AddFriendHandler(errors.New("[FriendsProtocol::AddFriend] Data holder not long enough for PID"), client, callID, 0)
		return
	}

	pid := parametersStream.ReadU32LENext(1)[0]

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

	username, err := parametersStream.ReadStringNext()

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
		go friendsProtocol.RemoveFriendHandler(errors.New("[FriendsProtocol::RemoveFriend] Data holder not long enough for PID"), client, callID, 0)
		return
	}

	pid := parametersStream.ReadU32LENext(1)[0]

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

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 + 1 +1 + 8 {
		// length check for the following fixed-size data
		// unknown1 + unknown2 + unknown4 + gameKey + unknown6
		err := errors.New("[FriendsProtocol::AddFriendRequest] Data holder not long enough for PID")
		go friendsProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", &GameKey{}, nex.NewDateTime(0))
		return
	}

	unknown1 := parametersStream.ReadU32LENext(1)[0]
	unknown2 := parametersStream.ReadByteNext()
	unknown3, err := parametersStream.ReadStringNext()

	if err != nil {
		go friendsProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", &GameKey{}, nex.NewDateTime(0))
		return
	}

	unknown4 := parametersStream.ReadByteNext()
	unknown5, err := parametersStream.ReadStringNext()

	if err != nil {
		go friendsProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", &GameKey{}, nex.NewDateTime(0))
		return
	}

	gameKey := &GameKey{}
	err = gameKey.ExtractFromStream(parametersStream)

	if err != nil {
		go friendsProtocol.AddFriendRequestHandler(err, client, callID, 0, 0, "", 0, "", &GameKey{}, nex.NewDateTime(0))
		return
	}

	unknown6 := nex.NewDateTime(parametersStream.ReadU64LENext(1)[0])

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

	id := parametersStream.ReadU64LENext(1)[0]

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

	id := parametersStream.ReadU64LENext(1)[0]

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

	id := parametersStream.ReadU64LENext(1)[0]

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

	id := parametersStream.ReadU64LENext(1)[0]

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

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsProtocol::MarkFriendRequestsAsReceived] Data missing list length")
		go friendsProtocol.MarkFriendRequestsAsReceivedHandler(err, client, callID, make([]uint64, 0))
		return
	}

	idCount := parametersStream.ReadU32LENext(1)[0]
	ids := make([]uint64, 0)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < (8 * int(idCount)) {
		err := errors.New("[FriendsProtocol::MarkFriendRequestsAsReceived] Data length less than content length")
		go friendsProtocol.MarkFriendRequestsAsReceivedHandler(err, client, callID, ids)
		return
	}

	for i := 0; i < int(idCount); i++ {
		id := parametersStream.ReadU64LENext(1)[0]
		ids = append(ids, id)
	}

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

	blacklistedPrincipal := &BlacklistedPrincipal{}
	err := blacklistedPrincipal.ExtractFromStream(parametersStream)

	if err != nil {
		go friendsProtocol.AddBlackListHandler(err, client, callID, &BlacklistedPrincipal{})
		return
	}

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

	pid := parametersStream.ReadU32LENext(1)[0]

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

	presence := &NintendoPresenceV2{}
	err := presence.ExtractFromStream(parametersStream)

	if err != nil {
		go friendsProtocol.UpdatePresenceHandler(err, client, callID, &NintendoPresenceV2{})
		return
	}

	go friendsProtocol.UpdatePresenceHandler(nil, client, callID, presence)
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
	err := mii.ExtractFromStream(parametersStream)

	if err != nil {
		go friendsProtocol.UpdateMiiHandler(err, client, callID, &MiiV2{})
		return
	}

	go friendsProtocol.UpdateMiiHandler(nil, client, callID, mii)
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
	err := comment.ExtractFromStream(parametersStream)

	if err != nil {
		go friendsProtocol.UpdateCommentHandler(err, client, callID, &Comment{})
		return
	}

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

	preference := &PrincipalPreference{}
	err := preference.ExtractFromStream(parametersStream)

	if err != nil {
		go friendsProtocol.UpdatePreferenceHandler(err, client, callID, &PrincipalPreference{})
		return
	}

	go friendsProtocol.UpdatePreferenceHandler(nil, client, callID, preference)
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

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsProtocol::GetBasicInfo] Data missing list length")
		go friendsProtocol.GetBasicInfoHandler(err, client, callID, make([]uint32, 0))
		return
	}

	pidCount := parametersStream.ReadU32LENext(1)[0]
	pids := make([]uint32, 0)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < (4 * int(pidCount)) {
		err := errors.New("[FriendsProtocol::GetBasicInfo] Data length less than content length")
		go friendsProtocol.GetBasicInfoHandler(err, client, callID, pids)
		return
	}

	for i := 0; i < int(pidCount); i++ {
		pid := parametersStream.ReadU32LENext(1)[0]
		pids = append(pids, pid)
	}

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

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	notificationCount := parametersStream.ReadU32LENext(1)[0]
	notifications := make([]*PersistentNotification, 0)

	for i := 0; i < int(notificationCount); i++ {
		notification := &PersistentNotification{}
		notification.ExtractFromStream(parametersStream)
		notifications = append(notifications, notification)
	}

	go friendsProtocol.DeleteFriendFlagsHandler(nil, client, callID, notifications)
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

	parametersStream := nex.NewStreamIn(parameters, friendsProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsProtocol::GetRequestBlockSettings] Data missing list length")
		go friendsProtocol.GetRequestBlockSettingsHandler(err, client, callID, make([]uint32, 0))
		return
	}

	unknownCount := parametersStream.ReadU32LENext(1)[0]
	unknowns := make([]uint32, 0)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < (4 * int(unknownCount)) {
		err := errors.New("[FriendsProtocol::GetRequestBlockSettings] Data length less than content length")
		go friendsProtocol.GetRequestBlockSettingsHandler(err, client, callID, unknowns)
		return
	}

	for i := 0; i < int(unknownCount); i++ {
		unknown := parametersStream.ReadU32LENext(1)[0]
		unknowns = append(unknowns, unknown)
	}

	go friendsProtocol.GetRequestBlockSettingsHandler(nil, client, callID, unknowns)
}

func NewFriendsProtocol(server *nex.Server) *FriendsProtocol {
	friendsProtocol := &FriendsProtocol{server: server}

	friendsProtocol.Setup()

	return friendsProtocol
}
