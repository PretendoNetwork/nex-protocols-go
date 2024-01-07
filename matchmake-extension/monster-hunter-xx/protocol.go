// Package protocol implements the MatchmakeExtensionMonsterHunterXX protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
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
	server nex.ServerInterface
	matchmakeExtensionProtocol
	UpdateFriendUserProfile func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_extension_monster_hunter_x_x_types.FriendUserParam) (*nex.RMCMessage, uint32)
	GetFriendUserProfiles   func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID]) (*nex.RMCMessage, uint32)
	AddFriends              func(err error, packet nex.PacketInterface, callID uint32, pids *types.List[*types.PID]) (*nex.RMCMessage, uint32)
	RemoveFriend            func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID) (*nex.RMCMessage, uint32)
	FindCommunityByOwner    func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU64, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			if slices.Contains(patchedMethods, message.MethodID) {
				protocol.HandlePacket(packet)
			} else {
				protocol.matchmakeExtensionProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodUpdateFriendUserProfile:
		protocol.handleUpdateFriendUserProfile(packet)
	case MethodGetFriendUserProfiles:
		protocol.handleGetFriendUserProfiles(packet)
	case MethodAddFriends:
		protocol.handleAddFriends(packet)
	case MethodRemoveFriend:
		protocol.handleRemoveFriend(packet)
	case MethodFindCommunityByOwner:
		protocol.handleFindCommunityByOwner(packet)
	default:
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Matchmake Extension (Monster Hunter XX) method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new MatchmakeExtensionMonsterHunterXX protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}
	protocol.matchmakeExtensionProtocol.SetServer(server)

	protocol.Setup()

	return protocol
}
