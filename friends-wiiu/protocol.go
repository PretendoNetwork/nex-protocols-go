// Package friends_wiiu implements the Friends WiiU NEX protocol
package friends_wiiu

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Friends (WiiU) protocol
	ProtocolID = 0x66

	// MethodUpdateAndGetAllInformation is the method ID for method UpdateAndGetAllInformation
	MethodUpdateAndGetAllInformation = 0x1

	// MethodAddFriend is the method ID for method AddFriend
	MethodAddFriend = 0x2

	// MethodAddFriendByName is the method ID for method AddFriendByName
	MethodAddFriendByName = 0x3

	// MethodRemoveFriend is the method ID for method RemoveFriend
	MethodRemoveFriend = 0x4

	// MethodAddFriendRequest is the method ID for method AddFriendRequest
	MethodAddFriendRequest = 0x5

	// MethodCancelFriendRequest is the method ID for method CancelFriendRequest
	MethodCancelFriendRequest = 0x6

	// MethodAcceptFriendRequest is the method ID for method AcceptFriendRequest
	MethodAcceptFriendRequest = 0x7

	// MethodDeleteFriendRequest is the method ID for method DeleteFriendRequest
	MethodDeleteFriendRequest = 0x8

	// MethodDenyFriendRequest is the method ID for method DenyFriendRequest
	MethodDenyFriendRequest = 0x9

	// MethodMarkFriendRequestsAsReceived is the method ID for method MarkFriendRequestsAsReceived
	MethodMarkFriendRequestsAsReceived = 0xA

	// MethodAddBlackList is the method ID for method AddBlackList
	MethodAddBlackList = 0xB

	// MethodRemoveBlackList is the method ID for method RemoveBlackList
	MethodRemoveBlackList = 0xC

	// MethodUpdatePresence is the method ID for method UpdatePresence
	MethodUpdatePresence = 0xD

	// MethodUpdateMii is the method ID for method UpdateMii
	MethodUpdateMii = 0xE

	// MethodUpdateComment is the method ID for method UpdateComment
	MethodUpdateComment = 0xF

	// MethodUpdatePreference is the method ID for method UpdatePreference
	MethodUpdatePreference = 0x10

	// MethodGetBasicInfo is the method ID for method GetBasicInfo
	MethodGetBasicInfo = 0x11

	// MethodDeletePersistentNotification is the method ID for method DeletePersistentNotification
	MethodDeletePersistentNotification = 0x12

	// MethodCheckSettingStatus is the method ID for method CheckSettingStatus
	MethodCheckSettingStatus = 0x13

	// MethodGetRequestBlockSettings is the method ID for method GetRequestBlockSettings
	MethodGetRequestBlockSettings = 0x14
)

// FriendsWiiUProtocol handles the Friends (WiiU) NEX protocol
type FriendsWiiUProtocol struct {
	Server                              *nex.Server
	UpdateAndGetAllInformationHandler   func(err error, client *nex.Client, callID uint32, nnaInfo *friends_wiiu_types.NNAInfo, presence *friends_wiiu_types.NintendoPresenceV2, birthday *nex.DateTime)
	AddFriendHandler                    func(err error, client *nex.Client, callID uint32, pid uint32)
	AddFriendByNameHandler              func(err error, client *nex.Client, callID uint32, username string)
	RemoveFriendHandler                 func(err error, client *nex.Client, callID uint32, pid uint32)
	AddFriendRequestHandler             func(err error, client *nex.Client, callID uint32, pid uint32, unknown2 uint8, message string, unknown4 uint8, unknown5 string, gameKey *friends_wiiu_types.GameKey, unknown6 *nex.DateTime)
	CancelFriendRequestHandler          func(err error, client *nex.Client, callID uint32, id uint64)
	AcceptFriendRequestHandler          func(err error, client *nex.Client, callID uint32, id uint64)
	DeleteFriendRequestHandler          func(err error, client *nex.Client, callID uint32, id uint64)
	DenyFriendRequestHandler            func(err error, client *nex.Client, callID uint32, id uint64)
	MarkFriendRequestsAsReceivedHandler func(err error, client *nex.Client, callID uint32, ids []uint64)
	AddBlackListHandler                 func(err error, client *nex.Client, callID uint32, blacklistedPrincipal *friends_wiiu_types.BlacklistedPrincipal)
	RemoveBlackListHandler              func(err error, client *nex.Client, callID uint32, pid uint32)
	UpdatePresenceHandler               func(err error, client *nex.Client, callID uint32, presence *friends_wiiu_types.NintendoPresenceV2)
	UpdateMiiHandler                    func(err error, client *nex.Client, callID uint32, mii *friends_wiiu_types.MiiV2)
	UpdateCommentHandler                func(err error, client *nex.Client, callID uint32, comment *friends_wiiu_types.Comment)
	UpdatePreferenceHandler             func(err error, client *nex.Client, callID uint32, preference *friends_wiiu_types.PrincipalPreference)
	GetBasicInfoHandler                 func(err error, client *nex.Client, callID uint32, pids []uint32)
	DeletePersistentNotificationHandler func(err error, client *nex.Client, callID uint32, notifications []*friends_wiiu_types.PersistentNotification)
	CheckSettingStatusHandler           func(err error, client *nex.Client, callID uint32)
	GetRequestBlockSettingsHandler      func(err error, client *nex.Client, callID uint32, pids []uint32)
}

// Setup initializes the protocol
func (protocol *FriendsWiiUProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *FriendsWiiUProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodUpdateAndGetAllInformation:
		go protocol.handleUpdateAndGetAllInformation(packet)
	case MethodAddFriend:
		go protocol.handleAddFriend(packet)
	case MethodAddFriendByName:
		go protocol.handleAddFriendByName(packet)
	case MethodRemoveFriend:
		go protocol.handleRemoveFriend(packet)
	case MethodAddFriendRequest:
		go protocol.handleAddFriendRequest(packet)
	case MethodCancelFriendRequest:
		go protocol.handleCancelFriendRequest(packet)
	case MethodAcceptFriendRequest:
		go protocol.handleAcceptFriendRequest(packet)
	case MethodDeleteFriendRequest:
		go protocol.handleDeleteFriendRequest(packet)
	case MethodDenyFriendRequest:
		go protocol.handleDenyFriendRequest(packet)
	case MethodMarkFriendRequestsAsReceived:
		go protocol.handleMarkFriendRequestsAsReceived(packet)
	case MethodAddBlackList:
		go protocol.handleAddBlackList(packet)
	case MethodRemoveBlackList:
		go protocol.handleRemoveBlackList(packet)
	case MethodUpdatePresence:
		go protocol.handleUpdatePresence(packet)
	case MethodUpdateMii:
		go protocol.handleUpdateMii(packet)
	case MethodUpdateComment:
		go protocol.handleUpdateComment(packet)
	case MethodUpdatePreference:
		go protocol.handleUpdatePreference(packet)
	case MethodGetBasicInfo:
		go protocol.handleGetBasicInfo(packet)
	case MethodDeletePersistentNotification:
		go protocol.handleDeletePersistentNotification(packet)
	case MethodCheckSettingStatus:
		go protocol.handleCheckSettingStatus(packet)
	case MethodGetRequestBlockSettings:
		go protocol.handleGetRequestBlockSettings(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported Friends (WiiU) method ID: %#v\n", request.MethodID())
	}
}

// NewFriendsWiiUProtocol returns a new FriendsWiiUProtocol
func NewFriendsWiiUProtocol(server *nex.Server) *FriendsWiiUProtocol {
	protocol := &FriendsWiiUProtocol{Server: server}

	protocol.Setup()

	return protocol
}
