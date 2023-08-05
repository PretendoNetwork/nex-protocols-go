// Package protocol implements the MatchmakeExtensionMonsterHunterXX protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension"
	matchmake_extension_monster_hunter_x_x_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension/monster-hunter-xx/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Matchmake Extension (Monster Hunter XX) protocol
	ProtocolID = 0x6D

	// MethodUpdateFriendUserProfile is the method ID for the UpdateFriendUserProfile method
	MethodUpdateFriendUserProfile = 0x36

	// MethodGetFriendUserProfiles is the method ID for the GetFriendUserProfiles method
	MethodGetFriendUserProfiles = 0x37

	// MethodAddFriends is the method ID for the AddFriends method
	MethodAddFriends = 0x39

	// MethodRemoveFriend is the method ID for the RemoveFriend method
	MethodRemoveFriend = 0x3A

	// MethodFindCommunityByOwner is the method ID for the FindCommunityByOwner method
	MethodFindCommunityByOwner = 0x3B
)

var patchedMethods = []uint32{
	MethodUpdateFriendUserProfile,
	MethodGetFriendUserProfiles,
	MethodAddFriends,
	MethodRemoveFriend,
	MethodFindCommunityByOwner,
}

type matchmakeExtensionProtocol = matchmake_extension.Protocol

// Protocol stores all the RMC method handlers for the Matchmake Extension (Monster Hunter XX) protocol and listens for requests
// Embeds the Matchmake Extension protocol
type Protocol struct {
	Server *nex.Server
	matchmakeExtensionProtocol
	updateFriendUserProfileHandler func(err error, client *nex.Client, callID uint32, param *matchmake_extension_monster_hunter_x_x_types.FriendUserParam) uint32
	getFriendUserProfilesHandler   func(err error, client *nex.Client, callID uint32, pids []uint64) uint32
	addFriendsHandler              func(err error, client *nex.Client, callID uint32, pids []uint64) uint32
	removeFriendHandler            func(err error, client *nex.Client, callID uint32, pid uint64) uint32
	findCommunityByOwnerHandler    func(err error, client *nex.Client, callID uint32, id uint64, resultRange *nex.ResultRange) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.matchmakeExtensionProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodUpdateFriendUserProfile:
		go protocol.handleUpdateFriendUserProfile(packet)
	case MethodGetFriendUserProfiles:
		go protocol.handleGetFriendUserProfiles(packet)
	case MethodAddFriends:
		go protocol.handleAddFriends(packet)
	case MethodRemoveFriend:
		go protocol.handleRemoveFriend(packet)
	case MethodFindCommunityByOwner:
		go protocol.handleFindCommunityByOwner(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Matchmake Extension (Monster Hunter XX) method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new MatchmakeExtensionMonsterHunterXX protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.matchmakeExtensionProtocol.Server = server

	protocol.Setup()

	return protocol
}
