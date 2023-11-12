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
	Server                                 nex.ServerInterface
	updateProfileHandler                   func(err error, packet nex.PacketInterface, callID uint32, profileData *friends_3ds_types.MyProfile) uint32
	updateMiiHandler                       func(err error, packet nex.PacketInterface, callID uint32, mii *friends_3ds_types.Mii) uint32
	updateMiiListHandler                   func(err error, packet nex.PacketInterface, callID uint32, miiList *friends_3ds_types.MiiList) uint32
	updatePlayedGamesHandler               func(err error, packet nex.PacketInterface, callID uint32, playedGames []*friends_3ds_types.PlayedGame) uint32
	updatePreferenceHandler                func(err error, packet nex.PacketInterface, callID uint32, publicMode bool, showGame bool, showPlayedGame bool) uint32
	getFriendMiiHandler                    func(err error, packet nex.PacketInterface, callID uint32, pidList []uint32) uint32
	getFriendMiiListHandler                func(err error, packet nex.PacketInterface, callID uint32, friends []*friends_3ds_types.FriendInfo) uint32
	isActiveGameHandler                    func(err error, packet nex.PacketInterface, callID uint32, pids []uint32, gameKey *friends_3ds_types.GameKey) uint32
	getPrincipalIDByLocalFriendCodeHandler func(err error, packet nex.PacketInterface, callID uint32, lfc uint64, lfcList []uint64) uint32
	getFriendRelationshipsHandler          func(err error, packet nex.PacketInterface, callID uint32, pids []uint32) uint32
	addFriendByPrincipalIDHandler          func(err error, packet nex.PacketInterface, callID uint32, lfc uint64, pid uint32) uint32
	addFriendBylstPrincipalIDHandler       func(err error, packet nex.PacketInterface, callID uint32, lfc uint64, pids []uint32) uint32
	removeFriendByLocalFriendCodeHandler   func(err error, packet nex.PacketInterface, callID uint32, lfc uint64) uint32
	removeFriendByPrincipalIDHandler       func(err error, packet nex.PacketInterface, callID uint32, pid uint32) uint32
	getAllFriendsHandler                   func(err error, packet nex.PacketInterface, callID uint32) uint32
	updateBlackListHandler                 func(err error, packet nex.PacketInterface, callID uint32, unknown []uint32) uint32
	syncFriendHandler                      func(err error, packet nex.PacketInterface, callID uint32, lfc uint64, pids []uint32, lfcList []uint64) uint32
	updatePresenceHandler                  func(err error, packet nex.PacketInterface, callID uint32, presence *friends_3ds_types.NintendoPresence, showGame bool) uint32
	updateFavoriteGameKeyHandler           func(err error, packet nex.PacketInterface, callID uint32, gameKey *friends_3ds_types.GameKey) uint32
	updateCommentHandler                   func(err error, packet nex.PacketInterface, callID uint32, comment string) uint32
	updatePictureHandler                   func(err error, packet nex.PacketInterface, callID uint32, unknown uint32, picture []byte) uint32
	getFriendPresenceHandler               func(err error, packet nex.PacketInterface, callID uint32, pidList []uint32) uint32
	getFriendCommentHandler                func(err error, packet nex.PacketInterface, callID uint32, friends []*friends_3ds_types.FriendInfo) uint32
	getFriendPictureHandler                func(err error, packet nex.PacketInterface, callID uint32, unknown []uint32) uint32
	getFriendPersistentInfoHandler         func(err error, packet nex.PacketInterface, callID uint32, pidList []uint32) uint32
	sendInvitationHandler                  func(err error, packet nex.PacketInterface, callID uint32, pids []uint32) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodUpdateProfile:
		go protocol.handleUpdateProfile(packet)
	case MethodUpdateMii:
		go protocol.handleUpdateMii(packet)
	case MethodUpdateMiiList:
		go protocol.handleUpdateMiiList(packet)
	case MethodUpdatePlayedGames:
		go protocol.handleUpdatePlayedGames(packet)
	case MethodUpdatePreference:
		go protocol.handleUpdatePreference(packet)
	case MethodGetFriendMii:
		go protocol.handleGetFriendMii(packet)
	case MethodGetFriendMiiList:
		go protocol.handleGetFriendMiiList(packet)
	case MethodIsActiveGame:
		go protocol.handleIsActiveGame(packet)
	case MethodGetPrincipalIDByLocalFriendCode:
		go protocol.handleGetPrincipalIDByLocalFriendCode(packet)
	case MethodGetFriendRelationships:
		go protocol.handleGetFriendRelationships(packet)
	case MethodAddFriendByPrincipalID:
		go protocol.handleAddFriendByPrincipalID(packet)
	case MethodAddFriendBylstPrincipalID:
		go protocol.handleAddFriendBylstPrincipalID(packet)
	case MethodRemoveFriendByLocalFriendCode:
		go protocol.handleRemoveFriendByLocalFriendCode(packet)
	case MethodRemoveFriendByPrincipalID:
		go protocol.handleRemoveFriendByPrincipalID(packet)
	case MethodGetAllFriends:
		go protocol.handleGetAllFriends(packet)
	case MethodUpdateBlackList:
		go protocol.handleUpdateBlackList(packet)
	case MethodSyncFriend:
		go protocol.handleSyncFriend(packet)
	case MethodUpdatePresence:
		go protocol.handleUpdatePresence(packet)
	case MethodUpdateFavoriteGameKey:
		go protocol.handleUpdateFavoriteGameKey(packet)
	case MethodUpdateComment:
		go protocol.handleUpdateComment(packet)
	case MethodUpdatePicture:
		go protocol.handleUpdatePicture(packet)
	case MethodGetFriendPresence:
		go protocol.handleGetFriendPresence(packet)
	case MethodGetFriendComment:
		go protocol.handleGetFriendComment(packet)
	case MethodGetFriendPicture:
		go protocol.handleGetFriendPicture(packet)
	case MethodGetFriendPersistentInfo:
		go protocol.handleGetFriendPersistentInfo(packet)
	case MethodSendInvitation:
		go protocol.handleSendInvitation(packet)
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
