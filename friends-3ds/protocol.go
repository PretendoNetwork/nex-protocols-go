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

	// MethodUpdatePreference is the method ID for method UpdatePreference
	MethodUpdatePreference = 0x5

	// MethodGetFriendMii is the method ID for method GetFriendMii
	MethodGetFriendMii = 0x6

	// MethodAddFriendByPrincipalID is the method ID for method AddFriendByPrincipalID
	MethodAddFriendByPrincipalID = 0xb

	// MethodGetPrincipalIDByLocalFriendCode is the method ID for method GetPrincipalIDByLocalFriendCode
	MethodGetPrincipalIDByLocalFriendCode = 0x9

	// MethodRemoveFriendByLocalFriendCode is the method ID for method RemoveFriendByLocalFriendCode
	MethodRemoveFriendByLocalFriendCode = 0xd

	// MethodRemoveFriendByPrincipalID is the method ID for method RemoveFriendByPrincipalID
	MethodRemoveFriendByPrincipalID = 0xe

	// MethodGetAllFriends is the method ID for method MethodGetAllFriends
	MethodGetAllFriends = 0xf

	// MethodSyncFriend is the method ID for method SyncFriend
	MethodSyncFriend = 0x11

	// MethodUpdatePresence is the method ID for method UpdatePresence
	MethodUpdatePresence = 0x12

	// MethodUpdateFavoriteGameKey is the method ID for method UpdateFavoriteGameKey
	MethodUpdateFavoriteGameKey = 0x13

	// MethodUpdateComment is the method ID for method UpdateComment
	MethodUpdateComment = 0x14

	// MethodGetFriendPresence is the method ID for method GetFriendPresence
	MethodGetFriendPresence = 0x16

	// MethodGetFriendPersistentInfo is the method ID for method GetFriendPersistentInfo
	MethodGetFriendPersistentInfo = 0x19
)

// Friends3DSProtocol handles the Friends (3DS) NEX protocol
type Friends3DSProtocol struct {
	Server                                 *nex.Server
	UpdateProfileHandler                   func(err error, client *nex.Client, callID uint32, profileData *friends_3ds_types.MyProfile)
	UpdateMiiHandler                       func(err error, client *nex.Client, callID uint32, mii *friends_3ds_types.Mii)
	UpdatePreferenceHandler                func(err error, client *nex.Client, callID uint32, publicMode bool, showGame bool, showPlayedGame bool)
	UpdatePresenceHandler                  func(err error, client *nex.Client, callID uint32, presence *friends_3ds_types.NintendoPresence, showGame bool)
	UpdateFavoriteGameKeyHandler           func(err error, client *nex.Client, callID uint32, gameKey *friends_3ds_types.GameKey)
	UpdateCommentHandler                   func(err error, client *nex.Client, callID uint32, comment string)
	SyncFriendHandler                      func(err error, client *nex.Client, callID uint32, lfc uint64, pids []uint32, lfcList []uint64)
	GetPrincipalIDByLocalFriendCodeHandler func(err error, client *nex.Client, callID uint32, lfc uint64, lfcList []uint64)
	AddFriendByPrincipalIDHandler          func(err error, client *nex.Client, callID uint32, lfc uint64, pid uint32)
	RemoveFriendByLocalFriendCodeHandler   func(err error, client *nex.Client, callID uint32, lfc uint64)
	RemoveFriendByPrincipalIDHandler       func(err error, client *nex.Client, callID uint32, pid uint32)
	GetAllFriendsHandler                   func(err error, client *nex.Client, callID uint32)
	GetFriendPersistentInfoHandler         func(err error, client *nex.Client, callID uint32, pidList []uint32)
	GetFriendMiiHandler                    func(err error, client *nex.Client, callID uint32, pidList []uint32)
	GetFriendPresenceHandler               func(err error, client *nex.Client, callID uint32, pidList []uint32)
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
	case MethodUpdatePreference:
		go protocol.handleUpdatePreference(packet)
	case MethodSyncFriend:
		go protocol.handleSyncFriend(packet)
	case MethodUpdatePresence:
		go protocol.handleUpdatePresence(packet)
	case MethodUpdateFavoriteGameKey:
		go protocol.handleUpdateFavoriteGameKey(packet)
	case MethodUpdateComment:
		go protocol.handleUpdateComment(packet)
	case MethodGetPrincipalIDByLocalFriendCode:
		go protocol.handleGetPrincipalIDByLocalFriendCode(packet)
	case MethodAddFriendByPrincipalID:
		go protocol.handleAddFriendByPrincipalID(packet)
	case MethodRemoveFriendByLocalFriendCode:
		go protocol.handleRemoveFriendByLocalFriendCode(packet)
	case MethodRemoveFriendByPrincipalID:
		go protocol.handleRemoveFriendByPrincipalID(packet)
	case MethodGetAllFriends:
		go protocol.handleGetAllFriends(packet)
	case MethodGetFriendPersistentInfo:
		go protocol.handleGetFriendPersistentInfo(packet)
	case MethodGetFriendMii:
		go protocol.handleGetFriendMii(packet)
	case MethodGetFriendPresence:
		go protocol.handleGetFriendPresence(packet)
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
