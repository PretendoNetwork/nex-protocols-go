// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Friends (3DS) protocol
	ProtocolID = 0x65

	// MethodUpdateProfile is the method ID for method UpdateProfile
	MethodUpdateProfile = 0x1

	// MethodUpdateMii is the method ID for method UpdateMii
	MethodUpdateMii = 0x2

	// MethodUpdateMiiList is the method ID for method UpdateMiiList
	MethodUpdateMiiList = 0x3

	// MethodUpdatePlayedGames is the method ID for method UpdatePlayedGames
	MethodUpdatePlayedGames = 0x4

	// MethodUpdatePreference is the method ID for method UpdatePreference
	MethodUpdatePreference = 0x5

	// MethodGetFriendMii is the method ID for method GetFriendMii
	MethodGetFriendMii = 0x6

	// MethodGetFriendMiiList is the method ID for method GetFriendMiiList
	MethodGetFriendMiiList = 0x7

	// MethodIsActiveGame is the method ID for method IsActiveGame
	MethodIsActiveGame = 0x8

	// MethodGetPrincipalIDByLocalFriendCode is the method ID for method GetPrincipalIDByLocalFriendCode
	MethodGetPrincipalIDByLocalFriendCode = 0x9

	// MethodGetFriendRelationships is the method ID for method GetFriendRelationships
	MethodGetFriendRelationships = 0xA

	// MethodAddFriendByPrincipalID is the method ID for method AddFriendByPrincipalID
	MethodAddFriendByPrincipalID = 0xB

	// MethodAddFriendBylstPrincipalID is the method ID for method AddFriendBylstPrincipalID
	MethodAddFriendBylstPrincipalID = 0xC

	// MethodRemoveFriendByLocalFriendCode is the method ID for method RemoveFriendByLocalFriendCode
	MethodRemoveFriendByLocalFriendCode = 0xD

	// MethodRemoveFriendByPrincipalID is the method ID for method RemoveFriendByPrincipalID
	MethodRemoveFriendByPrincipalID = 0xE

	// MethodGetAllFriends is the method ID for method GetAllFriends
	MethodGetAllFriends = 0xF

	// MethodUpdateBlackList is the method ID for method UpdateBlackList
	MethodUpdateBlackList = 0x10

	// MethodSyncFriend is the method ID for method SyncFriend
	MethodSyncFriend = 0x11

	// MethodUpdatePresence is the method ID for method UpdatePresence
	MethodUpdatePresence = 0x12

	// MethodUpdateFavoriteGameKey is the method ID for method UpdateFavoriteGameKey
	MethodUpdateFavoriteGameKey = 0x13

	// MethodUpdateComment is the method ID for method UpdateComment
	MethodUpdateComment = 0x14

	// MethodUpdatePicture is the method ID for method UpdatePicture
	MethodUpdatePicture = 0x15

	// MethodGetFriendPresence is the method ID for method GetFriendPresence
	MethodGetFriendPresence = 0x16

	// MethodGetFriendComment is the method ID for method GetFriendComment
	MethodGetFriendComment = 0x17

	// MethodGetFriendPicture is the method ID for method GetFriendPicture
	MethodGetFriendPicture = 0x18

	// MethodGetFriendPersistentInfo is the method ID for method GetFriendPersistentInfo
	MethodGetFriendPersistentInfo = 0x19

	// MethodSendInvitation is the method ID for method SendInvitation
	MethodSendInvitation = 0x1A
)

// Protocol stores all the RMC method handlers for the Friends (3DS) protocol and listens for requests
type Protocol struct {
	Server                          nex.ServerInterface
	UpdateProfile                   func(err error, packet nex.PacketInterface, callID uint32, profileData *friends_3ds_types.MyProfile) (*nex.RMCMessage, uint32)
	UpdateMii                       func(err error, packet nex.PacketInterface, callID uint32, mii *friends_3ds_types.Mii) (*nex.RMCMessage, uint32)
	UpdateMiiList                   func(err error, packet nex.PacketInterface, callID uint32, miiList *friends_3ds_types.MiiList) (*nex.RMCMessage, uint32)
	UpdatePlayedGames               func(err error, packet nex.PacketInterface, callID uint32, playedGames []*friends_3ds_types.PlayedGame) (*nex.RMCMessage, uint32)
	UpdatePreference                func(err error, packet nex.PacketInterface, callID uint32, publicMode bool, showGame bool, showPlayedGame bool) (*nex.RMCMessage, uint32)
	GetFriendMii                    func(err error, packet nex.PacketInterface, callID uint32, friends []*friends_3ds_types.FriendInfo) (*nex.RMCMessage, uint32)
	GetFriendMiiList                func(err error, packet nex.PacketInterface, callID uint32, friends []*friends_3ds_types.FriendInfo) (*nex.RMCMessage, uint32)
	IsActiveGame                    func(err error, packet nex.PacketInterface, callID uint32, pids []*nex.PID, gameKey *friends_3ds_types.GameKey) (*nex.RMCMessage, uint32)
	GetPrincipalIDByLocalFriendCode func(err error, packet nex.PacketInterface, callID uint32, lfc uint64, lfcList []uint64) (*nex.RMCMessage, uint32)
	GetFriendRelationships          func(err error, packet nex.PacketInterface, callID uint32, pids []*nex.PID) (*nex.RMCMessage, uint32)
	AddFriendByPrincipalID          func(err error, packet nex.PacketInterface, callID uint32, lfc uint64, pid *nex.PID) (*nex.RMCMessage, uint32)
	AddFriendBylstPrincipalID       func(err error, packet nex.PacketInterface, callID uint32, lfc uint64, pids []*nex.PID) (*nex.RMCMessage, uint32)
	RemoveFriendByLocalFriendCode   func(err error, packet nex.PacketInterface, callID uint32, lfc uint64) (*nex.RMCMessage, uint32)
	RemoveFriendByPrincipalID       func(err error, packet nex.PacketInterface, callID uint32, pid *nex.PID) (*nex.RMCMessage, uint32)
	GetAllFriends                   func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	UpdateBlackList                 func(err error, packet nex.PacketInterface, callID uint32, unknown []uint32) (*nex.RMCMessage, uint32)
	SyncFriend                      func(err error, packet nex.PacketInterface, callID uint32, lfc uint64, pids []*nex.PID, lfcList []uint64) (*nex.RMCMessage, uint32)
	UpdatePresence                  func(err error, packet nex.PacketInterface, callID uint32, presence *friends_3ds_types.NintendoPresence, showGame bool) (*nex.RMCMessage, uint32)
	UpdateFavoriteGameKey           func(err error, packet nex.PacketInterface, callID uint32, gameKey *friends_3ds_types.GameKey) (*nex.RMCMessage, uint32)
	UpdateComment                   func(err error, packet nex.PacketInterface, callID uint32, comment string) (*nex.RMCMessage, uint32)
	UpdatePicture                   func(err error, packet nex.PacketInterface, callID uint32, unknown uint32, picture []byte) (*nex.RMCMessage, uint32)
	GetFriendPresence               func(err error, packet nex.PacketInterface, callID uint32, pidList []*nex.PID) (*nex.RMCMessage, uint32)
	GetFriendComment                func(err error, packet nex.PacketInterface, callID uint32, friends []*friends_3ds_types.FriendInfo) (*nex.RMCMessage, uint32)
	GetFriendPicture                func(err error, packet nex.PacketInterface, callID uint32, unknown []uint32) (*nex.RMCMessage, uint32)
	GetFriendPersistentInfo         func(err error, packet nex.PacketInterface, callID uint32, pidList []*nex.PID) (*nex.RMCMessage, uint32)
	SendInvitation                  func(err error, packet nex.PacketInterface, callID uint32, pids []*nex.PID) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodUpdateProfile:
		protocol.handleUpdateProfile(packet)
	case MethodUpdateMii:
		protocol.handleUpdateMii(packet)
	case MethodUpdateMiiList:
		protocol.handleUpdateMiiList(packet)
	case MethodUpdatePlayedGames:
		protocol.handleUpdatePlayedGames(packet)
	case MethodUpdatePreference:
		protocol.handleUpdatePreference(packet)
	case MethodGetFriendMii:
		protocol.handleGetFriendMii(packet)
	case MethodGetFriendMiiList:
		protocol.handleGetFriendMiiList(packet)
	case MethodIsActiveGame:
		protocol.handleIsActiveGame(packet)
	case MethodGetPrincipalIDByLocalFriendCode:
		protocol.handleGetPrincipalIDByLocalFriendCode(packet)
	case MethodGetFriendRelationships:
		protocol.handleGetFriendRelationships(packet)
	case MethodAddFriendByPrincipalID:
		protocol.handleAddFriendByPrincipalID(packet)
	case MethodAddFriendBylstPrincipalID:
		protocol.handleAddFriendBylstPrincipalID(packet)
	case MethodRemoveFriendByLocalFriendCode:
		protocol.handleRemoveFriendByLocalFriendCode(packet)
	case MethodRemoveFriendByPrincipalID:
		protocol.handleRemoveFriendByPrincipalID(packet)
	case MethodGetAllFriends:
		protocol.handleGetAllFriends(packet)
	case MethodUpdateBlackList:
		protocol.handleUpdateBlackList(packet)
	case MethodSyncFriend:
		protocol.handleSyncFriend(packet)
	case MethodUpdatePresence:
		protocol.handleUpdatePresence(packet)
	case MethodUpdateFavoriteGameKey:
		protocol.handleUpdateFavoriteGameKey(packet)
	case MethodUpdateComment:
		protocol.handleUpdateComment(packet)
	case MethodUpdatePicture:
		protocol.handleUpdatePicture(packet)
	case MethodGetFriendPresence:
		protocol.handleGetFriendPresence(packet)
	case MethodGetFriendComment:
		protocol.handleGetFriendComment(packet)
	case MethodGetFriendPicture:
		protocol.handleGetFriendPicture(packet)
	case MethodGetFriendPersistentInfo:
		protocol.handleGetFriendPersistentInfo(packet)
	case MethodSendInvitation:
		protocol.handleSendInvitation(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Friends (3DS) method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Friends (3DS) protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
