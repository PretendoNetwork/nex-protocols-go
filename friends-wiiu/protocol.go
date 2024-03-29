// Package protocol implements the Friends WiiU protocol
package protocol

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

// Protocol stores all the RMC method handlers for the Friends (WiiU) protocol and listens for requests
type Protocol struct {
	Server                              *nex.Server
	updateAndGetAllInformationHandler   func(err error, packet nex.PacketInterface, callID uint32, nnaInfo *friends_wiiu_types.NNAInfo, presence *friends_wiiu_types.NintendoPresenceV2, birthday *nex.DateTime) uint32
	addFriendHandler                    func(err error, packet nex.PacketInterface, callID uint32, pid uint32) uint32
	addFriendByNameHandler              func(err error, packet nex.PacketInterface, callID uint32, username string) uint32
	removeFriendHandler                 func(err error, packet nex.PacketInterface, callID uint32, pid uint32) uint32
	addFriendRequestHandler             func(err error, packet nex.PacketInterface, callID uint32, pid uint32, unknown2 uint8, message string, unknown4 uint8, unknown5 string, gameKey *friends_wiiu_types.GameKey, unknown6 *nex.DateTime) uint32
	cancelFriendRequestHandler          func(err error, packet nex.PacketInterface, callID uint32, id uint64) uint32
	acceptFriendRequestHandler          func(err error, packet nex.PacketInterface, callID uint32, id uint64) uint32
	deleteFriendRequestHandler          func(err error, packet nex.PacketInterface, callID uint32, id uint64) uint32
	denyFriendRequestHandler            func(err error, packet nex.PacketInterface, callID uint32, id uint64) uint32
	markFriendRequestsAsReceivedHandler func(err error, packet nex.PacketInterface, callID uint32, ids []uint64) uint32
	addBlackListHandler                 func(err error, packet nex.PacketInterface, callID uint32, blacklistedPrincipal *friends_wiiu_types.BlacklistedPrincipal) uint32
	removeBlackListHandler              func(err error, packet nex.PacketInterface, callID uint32, pid uint32) uint32
	updatePresenceHandler               func(err error, packet nex.PacketInterface, callID uint32, presence *friends_wiiu_types.NintendoPresenceV2) uint32
	updateMiiHandler                    func(err error, packet nex.PacketInterface, callID uint32, mii *friends_wiiu_types.MiiV2) uint32
	updateCommentHandler                func(err error, packet nex.PacketInterface, callID uint32, comment *friends_wiiu_types.Comment) uint32
	updatePreferenceHandler             func(err error, packet nex.PacketInterface, callID uint32, preference *friends_wiiu_types.PrincipalPreference) uint32
	getBasicInfoHandler                 func(err error, packet nex.PacketInterface, callID uint32, pids []uint32) uint32
	deletePersistentNotificationHandler func(err error, packet nex.PacketInterface, callID uint32, notifications []*friends_wiiu_types.PersistentNotification) uint32
	checkSettingStatusHandler           func(err error, packet nex.PacketInterface, callID uint32) uint32
	getRequestBlockSettingsHandler      func(err error, packet nex.PacketInterface, callID uint32, pids []uint32) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
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
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Friends (WiiU) method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new Friends (WiiU) protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
