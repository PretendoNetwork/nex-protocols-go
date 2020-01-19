package nexproto

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

const (
	FriendsProtocolID = 0x66

	FriendsMethodUpdateAndGetAllInformation = 0x1
	FriendsMethodAddFriend = 0x2
	FriendsMethodAddFriendByName = 0x3
	FriendsMethodRemoveFriend = 0x4
	FriendsMethodAddFriendRequest = 0x5
	FriendsMethodCancelFriendRequest = 0x6
	FriendsMethodAcceptFriendRequest = 0x7
	FriendsMethodDeleteFriendRequest = 0x8
	FriendsMethodDenyFriendRequest = 0x9
	FriendsMethodMarkFriendRequestsAsReceived = 0xA
	FriendsMethodAddBlackList = 0xB
	FriendsMethodRemoveBlackList = 0xC
	FriendsMethodUpdatePresence = 0xD
	FriendsMethodUpdateMii = 0xE
	FriendsMethodUpdateComment = 0xF
	FriendsMethodUpdatePreference = 0x10
	FriendsMethodGetBasicInfo = 0x11
	FriendsMethodDeleteFriendFlags = 0x12
	FriendsMethodCheckSettingStatus = 0x13
	FriendsMethodGetRequestBlockSettings = 0x14
)

type MiiV2 struct {
	name string
	unknown1 uint8
	unknown2 uint8
	data []byte
	datetime *nex.DateTime
}

func (mii *MiiV2) ExtractFromStreamNext(stream *nex.Stream) {
	mii.name = stream.ReadNEXStringNext()
	mii.unknown1 = stream.ReadByteNext()
	mii.unknown2 = stream.ReadByteNext()
	mii.data = stream.ReadNEXBufferNext()
	mii.datetime = nex.NewDateTime(stream.ReadU64LENext(1)[0])
}

type PrincipalBasicInfo struct {
	pid uint32
	nnid string
	mii *MiiV2
	unknown uint8
}

func (principalInfo *PrincipalBasicInfo) ExtractFromStreamNext(stream *nex.Stream) {
	principalInfo.pid = stream.ReadU32LENext(1)[0]
	principalInfo.nnid = stream.ReadNEXStringNext()
	principalInfo.mii = &MiiV2{}
	principalInfo.mii.ExtractFromStreamNext(stream)
	principalInfo.unknown = stream.ReadByteNext()
}

type NNAInfo struct {
	principalInfo *PrincipalBasicInfo
	unknown1 uint8 
	unknown2 uint8 
}

func (nnaInfo *NNAInfo) ExtractFromStreamNext(stream *nex.Stream) {
	nnaInfo.principalInfo = &PrincipalBasicInfo{}
	nnaInfo.principalInfo.ExtractFromStreamNext(stream)
	nnaInfo.unknown1 = stream.ReadByteNext()
	nnaInfo.unknown2 = stream.ReadByteNext()
}

type GameKey struct {
	titleID uint64
	titleVersion uint16
}

func (gameKey *GameKey) ExtractFromStreamNext(stream *nex.Stream) {
	gameKey.titleID = stream.ReadU64LENext(1)[0]
	gameKey.titleVersion = stream.ReadU16LENext(1)[0]
}

type NintendoPresenceV2 struct {
	changedFlags uint32
	isOnline bool
	gameKey *GameKey
	unknown1 uint8
	message string
	unknown2 uint32
	unknown3 uint8
	gameServerID uint32
	unknown4 uint32
	pid uint32
	gatheringID uint32
	applicationData []byte
	unknown5 uint8
	unknown6 uint8
	unknown7 uint8
}

func (presence *NintendoPresenceV2) ExtractFromStreamNext(stream *nex.Stream) {
	presence.changedFlags = stream.ReadU32LENext(1)[0]
	presence.isOnline = (stream.ReadByteNext() == 1)
	presence.gameKey = &GameKey{}
	presence.gameKey.ExtractFromStreamNext(stream)
	presence.unknown1 = stream.ReadByteNext()
	presence.message = stream.ReadNEXStringNext()
	presence.unknown2 = stream.ReadU32LENext(1)[0]
	presence.unknown3 = stream.ReadByteNext()
	presence.gameServerID = stream.ReadU32LENext(1)[0]
	presence.unknown4 = stream.ReadU32LENext(1)[0]
	presence.pid = stream.ReadU32LENext(1)[0]
	presence.gatheringID = stream.ReadU32LENext(1)[0]
	presence.applicationData = stream.ReadNEXBufferNext()
	presence.unknown5 = stream.ReadByteNext()
	presence.unknown6 = stream.ReadByteNext()
	presence.unknown7 = stream.ReadByteNext()
}

type FriendsProtocol struct {
	server *nex.Server
	UpdateAndGetAllInformationHandler func(client *nex.Client, callID uint32, nnaInfo *NNAInfo, presence *NintendoPresenceV2, birthday *nex.DateTime)
	CheckSettingStatusHandler func(client *nex.Client, callID uint32)
}

func (friendsProtocol *FriendsProtocol) Setup() {
	nexServer := friendsProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.GetRMCRequest()

		if FriendsProtocolID == request.GetProtocolID() {
			switch request.GetMethodID() {
			case FriendsMethodUpdateAndGetAllInformation:
				go friendsProtocol.handleUpdateAndGetAllInformation(packet)
			case FriendsMethodCheckSettingStatus:
				go friendsProtocol.handleCheckSettingStatus(packet)
			default:
				fmt.Printf("Unsupported Friends (WiiU) method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

func (friendsProtocol *FriendsProtocol) UpdateAndGetAllInformation(handler func(client *nex.Client, callID uint32, nnaInfo *NNAInfo, presence *NintendoPresenceV2, birthday *nex.DateTime)) {
	friendsProtocol.UpdateAndGetAllInformationHandler = handler
}

func (friendsProtocol *FriendsProtocol) CheckSettingStatus(handler func(client *nex.Client, callID uint32)) {
	friendsProtocol.CheckSettingStatusHandler = handler
}

func (friendsProtocol *FriendsProtocol) handleUpdateAndGetAllInformation(packet nex.PacketInterface) {
	if friendsProtocol.UpdateAndGetAllInformationHandler == nil {
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	nnaInfo := &NNAInfo{}
	presence := &NintendoPresenceV2{}
	dateTime := nex.NewDateTime(0)

	nnaInfo.ExtractFromStreamNext(parametersStream)
	presence.ExtractFromStreamNext(parametersStream)

	go friendsProtocol.UpdateAndGetAllInformationHandler(client, callID, nnaInfo, presence, dateTime)
}

func (friendsProtocol *FriendsProtocol) handleCheckSettingStatus(packet nex.PacketInterface) {
	if friendsProtocol.CheckSettingStatusHandler == nil {
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()

	go friendsProtocol.CheckSettingStatusHandler(client, callID)
}

func NewFriendsProtocol(server *nex.Server) *FriendsProtocol {
	friendsProtocol := &FriendsProtocol{server: server}

	friendsProtocol.Setup()

	return friendsProtocol
}