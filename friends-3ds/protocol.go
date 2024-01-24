// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
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
	server                          nex.ServerInterface
	UpdateProfile                   func(err error, packet nex.PacketInterface, callID uint32, profileData *friends_3ds_types.MyProfile) (*nex.RMCMessage, uint32)
	UpdateMii                       func(err error, packet nex.PacketInterface, callID uint32, mii *friends_3ds_types.Mii) (*nex.RMCMessage, uint32)
	UpdateMiiList                   func(err error, packet nex.PacketInterface, callID uint32, miiList *friends_3ds_types.MiiList) (*nex.RMCMessage, uint32)
	UpdatePlayedGames               func(err error, packet nex.PacketInterface, callID uint32, playedGames *types.List[*friends_3ds_types.PlayedGame]) (*nex.RMCMessage, uint32)
	UpdatePreference                func(err error, packet nex.PacketInterface, callID uint32, publicMode *types.PrimitiveBool, showGame *types.PrimitiveBool, showPlayedGame *types.PrimitiveBool) (*nex.RMCMessage, uint32)
	GetFriendMii                    func(err error, packet nex.PacketInterface, callID uint32, friends *types.List[*friends_3ds_types.FriendInfo]) (*nex.RMCMessage, uint32)
	GetFriendMiiList                func(err error, packet nex.PacketInterface, callID uint32, friends *types.List[*friends_3ds_types.FriendInfo]) (*nex.RMCMessage, uint32)
	IsActiveGame                    func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID], gameKey *friends_3ds_types.GameKey) (*nex.RMCMessage, uint32)
	GetPrincipalIDByLocalFriendCode func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64, lfcList *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, uint32)
	GetFriendRelationships          func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID]) (*nex.RMCMessage, uint32)
	AddFriendByPrincipalID          func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64, pid *types.PID) (*nex.RMCMessage, uint32)
	AddFriendBylstPrincipalID       func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64, pids *types.List[*types.PID]) (*nex.RMCMessage, uint32)
	RemoveFriendByLocalFriendCode   func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	RemoveFriendByPrincipalID       func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, uint32)
	GetAllFriends                   func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	UpdateBlackList                 func(err error, packet nex.PacketInterface, callID uint32, unknown *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)
	SyncFriend                      func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64, pids *types.List[*types.PID], lfcList *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, uint32)
	UpdatePresence                  func(err error, packet nex.PacketInterface, callID uint32, presence *friends_3ds_types.NintendoPresence, showGame *types.PrimitiveBool) (*nex.RMCMessage, uint32)
	UpdateFavoriteGameKey           func(err error, packet nex.PacketInterface, callID uint32, gameKey *friends_3ds_types.GameKey) (*nex.RMCMessage, uint32)
	UpdateComment                   func(err error, packet nex.PacketInterface, callID uint32, comment *types.String) (*nex.RMCMessage, uint32)
	UpdatePicture                   func(err error, packet nex.PacketInterface, callID uint32, unknown *types.PrimitiveU32, picture *types.Buffer) (*nex.RMCMessage, uint32)
	GetFriendPresence               func(err error, packet nex.PacketInterface, callID uint32, pidList *types.List[*types.PID]) (*nex.RMCMessage, uint32)
	GetFriendComment                func(err error, packet nex.PacketInterface, callID uint32, friends *types.List[*friends_3ds_types.FriendInfo]) (*nex.RMCMessage, uint32)
	GetFriendPicture                func(err error, packet nex.PacketInterface, callID uint32, unknown *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)
	GetFriendPersistentInfo         func(err error, packet nex.PacketInterface, callID uint32, pidList *types.List[*types.PID]) (*nex.RMCMessage, uint32)
	SendInvitation                  func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID]) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the Friends (3DS) protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerUpdateProfile(handler func(err error, packet nex.PacketInterface, callID uint32, profileData *friends_3ds_types.MyProfile) (*nex.RMCMessage, uint32))
	SetHandlerUpdateMii(handler func(err error, packet nex.PacketInterface, callID uint32, mii *friends_3ds_types.Mii) (*nex.RMCMessage, uint32))
	SetHandlerUpdateMiiList(handler func(err error, packet nex.PacketInterface, callID uint32, miiList *friends_3ds_types.MiiList) (*nex.RMCMessage, uint32))
	SetHandlerUpdatePlayedGames(handler func(err error, packet nex.PacketInterface, callID uint32, playedGames *types.List[*friends_3ds_types.PlayedGame]) (*nex.RMCMessage, uint32))
	SetHandlerUpdatePreference(handler func(err error, packet nex.PacketInterface, callID uint32, publicMode *types.PrimitiveBool, showGame *types.PrimitiveBool, showPlayedGame *types.PrimitiveBool) (*nex.RMCMessage, uint32))
	SetHandlerGetFriendMii(handler func(err error, packet nex.PacketInterface, callID uint32, friends *types.List[*friends_3ds_types.FriendInfo]) (*nex.RMCMessage, uint32))
	SetHandlerGetFriendMiiList(handler func(err error, packet nex.PacketInterface, callID uint32, friends *types.List[*friends_3ds_types.FriendInfo]) (*nex.RMCMessage, uint32))
	SetHandlerIsActiveGame(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID], gameKey *friends_3ds_types.GameKey) (*nex.RMCMessage, uint32))
	SetHandlerGetPrincipalIDByLocalFriendCode(handler func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64, lfcList *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, uint32))
	SetHandlerGetFriendRelationships(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID]) (*nex.RMCMessage, uint32))
	SetHandlerAddFriendByPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64, pid *types.PID) (*nex.RMCMessage, uint32))
	SetHandlerAddFriendBylstPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64, pids *types.List[*types.PID]) (*nex.RMCMessage, uint32))
	SetHandlerRemoveFriendByLocalFriendCode(handler func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerRemoveFriendByPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, uint32))
	SetHandlerGetAllFriends(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerUpdateBlackList(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32))
	SetHandlerSyncFriend(handler func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64, pids *types.List[*types.PID], lfcList *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, uint32))
	SetHandlerUpdatePresence(handler func(err error, packet nex.PacketInterface, callID uint32, presence *friends_3ds_types.NintendoPresence, showGame *types.PrimitiveBool) (*nex.RMCMessage, uint32))
	SetHandlerUpdateFavoriteGameKey(handler func(err error, packet nex.PacketInterface, callID uint32, gameKey *friends_3ds_types.GameKey) (*nex.RMCMessage, uint32))
	SetHandlerUpdateComment(handler func(err error, packet nex.PacketInterface, callID uint32, comment *types.String) (*nex.RMCMessage, uint32))
	SetHandlerUpdatePicture(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.PrimitiveU32, picture *types.Buffer) (*nex.RMCMessage, uint32))
	SetHandlerGetFriendPresence(handler func(err error, packet nex.PacketInterface, callID uint32, pidList *types.List[*types.PID]) (*nex.RMCMessage, uint32))
	SetHandlerGetFriendComment(handler func(err error, packet nex.PacketInterface, callID uint32, friends *types.List[*friends_3ds_types.FriendInfo]) (*nex.RMCMessage, uint32))
	SetHandlerGetFriendPicture(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32))
	SetHandlerGetFriendPersistentInfo(handler func(err error, packet nex.PacketInterface, callID uint32, pidList *types.List[*types.PID]) (*nex.RMCMessage, uint32))
	SetHandlerSendInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID]) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerUpdateProfile sets the handler for the UpdateProfile method
func (protocol *Protocol) SetHandlerUpdateProfile(handler func(err error, packet nex.PacketInterface, callID uint32, profileData *friends_3ds_types.MyProfile) (*nex.RMCMessage, uint32)) {
	protocol.UpdateProfile = handler
}

// SetHandlerUpdateMii sets the handler for the UpdateMii method
func (protocol *Protocol) SetHandlerUpdateMii(handler func(err error, packet nex.PacketInterface, callID uint32, mii *friends_3ds_types.Mii) (*nex.RMCMessage, uint32)) {
	protocol.UpdateMii = handler
}

// SetHandlerUpdateMiiList sets the handler for the UpdateMiiList method
func (protocol *Protocol) SetHandlerUpdateMiiList(handler func(err error, packet nex.PacketInterface, callID uint32, miiList *friends_3ds_types.MiiList) (*nex.RMCMessage, uint32)) {
	protocol.UpdateMiiList = handler
}

// SetHandlerUpdatePlayedGames sets the handler for the UpdatePlayedGames method
func (protocol *Protocol) SetHandlerUpdatePlayedGames(handler func(err error, packet nex.PacketInterface, callID uint32, playedGames *types.List[*friends_3ds_types.PlayedGame]) (*nex.RMCMessage, uint32)) {
	protocol.UpdatePlayedGames = handler
}

// SetHandlerUpdatePreference sets the handler for the UpdatePreference method
func (protocol *Protocol) SetHandlerUpdatePreference(handler func(err error, packet nex.PacketInterface, callID uint32, publicMode *types.PrimitiveBool, showGame *types.PrimitiveBool, showPlayedGame *types.PrimitiveBool) (*nex.RMCMessage, uint32)) {
	protocol.UpdatePreference = handler
}

// SetHandlerGetFriendMii sets the handler for the GetFriendMii method
func (protocol *Protocol) SetHandlerGetFriendMii(handler func(err error, packet nex.PacketInterface, callID uint32, friends *types.List[*friends_3ds_types.FriendInfo]) (*nex.RMCMessage, uint32)) {
	protocol.GetFriendMii = handler
}

// SetHandlerGetFriendMiiList sets the handler for the GetFriendMiiList method
func (protocol *Protocol) SetHandlerGetFriendMiiList(handler func(err error, packet nex.PacketInterface, callID uint32, friends *types.List[*friends_3ds_types.FriendInfo]) (*nex.RMCMessage, uint32)) {
	protocol.GetFriendMiiList = handler
}

// SetHandlerIsActiveGame sets the handler for the IsActiveGame method
func (protocol *Protocol) SetHandlerIsActiveGame(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID], gameKey *friends_3ds_types.GameKey) (*nex.RMCMessage, uint32)) {
	protocol.IsActiveGame = handler
}

// SetHandlerGetPrincipalIDByLocalFriendCode sets the handler for the GetPrincipalIDByLocalFriendCode method
func (protocol *Protocol) SetHandlerGetPrincipalIDByLocalFriendCode(handler func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64, lfcList *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, uint32)) {
	protocol.GetPrincipalIDByLocalFriendCode = handler
}

// SetHandlerGetFriendRelationships sets the handler for the GetFriendRelationships method
func (protocol *Protocol) SetHandlerGetFriendRelationships(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID]) (*nex.RMCMessage, uint32)) {
	protocol.GetFriendRelationships = handler
}

// SetHandlerAddFriendByPrincipalID sets the handler for the AddFriendByPrincipalID method
func (protocol *Protocol) SetHandlerAddFriendByPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64, pid *types.PID) (*nex.RMCMessage, uint32)) {
	protocol.AddFriendByPrincipalID = handler
}

// SetHandlerAddFriendBylstPrincipalID sets the handler for the AddFriendBylstPrincipalID method
func (protocol *Protocol) SetHandlerAddFriendBylstPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64, pids *types.List[*types.PID]) (*nex.RMCMessage, uint32)) {
	protocol.AddFriendBylstPrincipalID = handler
}

// SetHandlerRemoveFriendByLocalFriendCode sets the handler for the RemoveFriendByLocalFriendCode method
func (protocol *Protocol) SetHandlerRemoveFriendByLocalFriendCode(handler func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.RemoveFriendByLocalFriendCode = handler
}

// SetHandlerRemoveFriendByPrincipalID sets the handler for the RemoveFriendByPrincipalID method
func (protocol *Protocol) SetHandlerRemoveFriendByPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, uint32)) {
	protocol.RemoveFriendByPrincipalID = handler
}

// SetHandlerGetAllFriends sets the handler for the GetAllFriends method
func (protocol *Protocol) SetHandlerGetAllFriends(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.GetAllFriends = handler
}

// SetHandlerUpdateBlackList sets the handler for the UpdateBlackList method
func (protocol *Protocol) SetHandlerUpdateBlackList(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)) {
	protocol.UpdateBlackList = handler
}

// SetHandlerSyncFriend sets the handler for the SyncFriend method
func (protocol *Protocol) SetHandlerSyncFriend(handler func(err error, packet nex.PacketInterface, callID uint32, lfc *types.PrimitiveU64, pids *types.List[*types.PID], lfcList *types.List[*types.PrimitiveU64]) (*nex.RMCMessage, uint32)) {
	protocol.SyncFriend = handler
}

// SetHandlerUpdatePresence sets the handler for the UpdatePresence method
func (protocol *Protocol) SetHandlerUpdatePresence(handler func(err error, packet nex.PacketInterface, callID uint32, presence *friends_3ds_types.NintendoPresence, showGame *types.PrimitiveBool) (*nex.RMCMessage, uint32)) {
	protocol.UpdatePresence = handler
}

// SetHandlerUpdateFavoriteGameKey sets the handler for the UpdateFavoriteGameKey method
func (protocol *Protocol) SetHandlerUpdateFavoriteGameKey(handler func(err error, packet nex.PacketInterface, callID uint32, gameKey *friends_3ds_types.GameKey) (*nex.RMCMessage, uint32)) {
	protocol.UpdateFavoriteGameKey = handler
}

// SetHandlerUpdateComment sets the handler for the UpdateComment method
func (protocol *Protocol) SetHandlerUpdateComment(handler func(err error, packet nex.PacketInterface, callID uint32, comment *types.String) (*nex.RMCMessage, uint32)) {
	protocol.UpdateComment = handler
}

// SetHandlerUpdatePicture sets the handler for the UpdatePicture method
func (protocol *Protocol) SetHandlerUpdatePicture(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.PrimitiveU32, picture *types.Buffer) (*nex.RMCMessage, uint32)) {
	protocol.UpdatePicture = handler
}

// SetHandlerGetFriendPresence sets the handler for the GetFriendPresence method
func (protocol *Protocol) SetHandlerGetFriendPresence(handler func(err error, packet nex.PacketInterface, callID uint32, pidList *types.List[*types.PID]) (*nex.RMCMessage, uint32)) {
	protocol.GetFriendPresence = handler
}

// SetHandlerGetFriendComment sets the handler for the GetFriendComment method
func (protocol *Protocol) SetHandlerGetFriendComment(handler func(err error, packet nex.PacketInterface, callID uint32, friends *types.List[*friends_3ds_types.FriendInfo]) (*nex.RMCMessage, uint32)) {
	protocol.GetFriendComment = handler
}

// SetHandlerGetFriendPicture sets the handler for the GetFriendPicture method
func (protocol *Protocol) SetHandlerGetFriendPicture(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)) {
	protocol.GetFriendPicture = handler
}

// SetHandlerGetFriendPersistentInfo sets the handler for the GetFriendPersistentInfo method
func (protocol *Protocol) SetHandlerGetFriendPersistentInfo(handler func(err error, packet nex.PacketInterface, callID uint32, pidList *types.List[*types.PID]) (*nex.RMCMessage, uint32)) {
	protocol.GetFriendPersistentInfo = handler
}

// SetHandlerSendInvitation sets the handler for the SendInvitation method
func (protocol *Protocol) SetHandlerSendInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID]) (*nex.RMCMessage, uint32)) {
	protocol.SendInvitation = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	switch message.MethodID {
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
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		fmt.Printf("Unsupported Friends (3DS) method ID: %#v\n", message.MethodID)
	}
}

// NewProtocol returns a new Friends (3DS) protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	return &Protocol{server: server}
}
