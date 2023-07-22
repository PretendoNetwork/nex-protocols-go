// Package friends_3ds implements the Friends 3DS NEX protocol
package friends_3ds

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

// Friends3DSProtocol handles the Friends (3DS) NEX protocol
type Friends3DSProtocol struct {
	Server                                 *nex.Server
	updateProfileHandler                   func(err error, client *nex.Client, callID uint32, profileData *friends_3ds_types.MyProfile)
	updateMiiHandler                       func(err error, client *nex.Client, callID uint32, mii *friends_3ds_types.Mii)
	updateMiiListHandler                   func(err error, client *nex.Client, callID uint32, miiList *friends_3ds_types.MiiList)
	updatePlayedGamesHandler               func(err error, client *nex.Client, callID uint32, playedGames []*friends_3ds_types.PlayedGame)
	updatePreferenceHandler                func(err error, client *nex.Client, callID uint32, publicMode bool, showGame bool, showPlayedGame bool)
	getFriendMiiHandler                    func(err error, client *nex.Client, callID uint32, pidList []uint32)
	getFriendMiiListHandler                func(err error, client *nex.Client, callID uint32, friends []*friends_3ds_types.FriendInfo)
	isActiveGameHandler                    func(err error, client *nex.Client, callID uint32, pids []uint32, gameKey *friends_3ds_types.GameKey)
	getPrincipalIDByLocalFriendCodeHandler func(err error, client *nex.Client, callID uint32, lfc uint64, lfcList []uint64)
	getFriendRelationshipsHandler          func(err error, client *nex.Client, callID uint32, pids []uint32)
	addFriendByPrincipalIDHandler          func(err error, client *nex.Client, callID uint32, lfc uint64, pid uint32)
	addFriendBylstPrincipalIDHandler       func(err error, client *nex.Client, callID uint32, lfc uint64, pids []uint32)
	removeFriendByLocalFriendCodeHandler   func(err error, client *nex.Client, callID uint32, lfc uint64)
	removeFriendByPrincipalIDHandler       func(err error, client *nex.Client, callID uint32, pid uint32)
	getAllFriendsHandler                   func(err error, client *nex.Client, callID uint32)
	updateBlackListHandler                 func(err error, client *nex.Client, callID uint32, unknown []uint32)
	syncFriendHandler                      func(err error, client *nex.Client, callID uint32, lfc uint64, pids []uint32, lfcList []uint64)
	updatePresenceHandler                  func(err error, client *nex.Client, callID uint32, presence *friends_3ds_types.NintendoPresence, showGame bool)
	updateFavoriteGameKeyHandler           func(err error, client *nex.Client, callID uint32, gameKey *friends_3ds_types.GameKey)
	updateCommentHandler                   func(err error, client *nex.Client, callID uint32, comment string)
	updatePictureHandler                   func(err error, client *nex.Client, callID uint32, unknown uint32, picture []byte)
	getFriendPresenceHandler               func(err error, client *nex.Client, callID uint32, pidList []uint32)
	getFriendCommentHandler                func(err error, client *nex.Client, callID uint32, friends []*friends_3ds_types.FriendInfo)
	getFriendPictureHandler                func(err error, client *nex.Client, callID uint32, unknown []uint32)
	getFriendPersistentInfoHandler         func(err error, client *nex.Client, callID uint32, pidList []uint32)
	sendInvitationHandler                  func(err error, client *nex.Client, callID uint32, pids []uint32)
}

// Setup initializes the protocol
func (protocol *Friends3DSProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Friends3DSProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
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
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported Friends (3DS) method ID: %#v\n", request.MethodID())
	}
}

// NewFriends3DSProtocol returns a new Friends3DSProtocol
func NewFriends3DSProtocol(server *nex.Server) *Friends3DSProtocol {
	protocol := &Friends3DSProtocol{Server: server}

	protocol.Setup()

	return protocol
}
