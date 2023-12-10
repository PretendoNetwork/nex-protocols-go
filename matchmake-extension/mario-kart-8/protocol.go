// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension"
	matchmake_extension_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension/mario-kart-8/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Matchmake Extension (Mario Kart 8) protocol
	ProtocolID = 0x6D

	// MethodCreateSimpleSearchObject is the method ID for the CreateSimpleSearchObject method
	MethodCreateSimpleSearchObject = 0x24

	// MethodUpdateSimpleSearchObject is the method ID for the UpdateSimpleSearchObject method
	MethodUpdateSimpleSearchObject = 0x25

	// MethodDeleteSimpleSearchObject is the method ID for the DeleteSimpleSearchObject method
	MethodDeleteSimpleSearchObject = 0x26

	// MethodSearchSimpleSearchObject is the method ID for the SearchSimpleSearchObject method
	MethodSearchSimpleSearchObject = 0x27

	// MethodJoinMatchmakeSessionWithExtraParticipants is the method ID for the JoinMatchmakeSessionWithExtraParticipants method
	MethodJoinMatchmakeSessionWithExtraParticipants = 0x28

	// MethodSearchSimpleSearchObjectByObjectIDs is the method ID for the SearchSimpleSearchObjectByObjectIDs method
	MethodSearchSimpleSearchObjectByObjectIDs = 0x29
)

var patchedMethods = []uint32{
	MethodCreateSimpleSearchObject,
	MethodUpdateSimpleSearchObject,
	MethodDeleteSimpleSearchObject,
	MethodSearchSimpleSearchObject,
	MethodJoinMatchmakeSessionWithExtraParticipants,
	MethodSearchSimpleSearchObjectByObjectIDs,
}

type matchmakeExtensionProtocol = matchmake_extension.Protocol

// Protocol stores all the RMC method handlers for the Matchmake Extension (Mario Kart 8) protocol and listens for requests
// Embeds the Matchmake Extension protocol
type Protocol struct {
	server nex.ServerInterface
	matchmakeExtensionProtocol
	CreateSimpleSearchObject                  func(err error, packet nex.PacketInterface, callID uint32, object *matchmake_extension_mario_kart8_types.SimpleSearchObject) (*nex.RMCMessage, uint32)
	UpdateSimpleSearchObject                  func(err error, packet nex.PacketInterface, callID uint32, objectID uint32, newObject *matchmake_extension_mario_kart8_types.SimpleSearchObject) (*nex.RMCMessage, uint32)
	DeleteSimpleSearchObject                  func(err error, packet nex.PacketInterface, callID uint32, objectID uint32) (*nex.RMCMessage, uint32)
	SearchSimpleSearchObject                  func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_extension_mario_kart8_types.SimpleSearchParam) (*nex.RMCMessage, uint32)
	JoinMatchmakeSessionWithExtraParticipants func(err error, packet nex.PacketInterface, callID uint32, gid uint32, joinMessage string, ignoreBlacklist bool, participationCount uint16, extraParticipants uint32) (*nex.RMCMessage, uint32)
	SearchSimpleSearchObjectByObjectIDs       func(err error, packet nex.PacketInterface, callID uint32, objectIDs []uint32) (*nex.RMCMessage, uint32)
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
	case MethodCreateSimpleSearchObject:
		protocol.handleCreateSimpleSearchObject(packet)
	case MethodUpdateSimpleSearchObject:
		protocol.handleUpdateSimpleSearchObject(packet)
	case MethodDeleteSimpleSearchObject:
		protocol.handleDeleteSimpleSearchObject(packet)
	case MethodSearchSimpleSearchObject:
		protocol.handleSearchSimpleSearchObject(packet)
	case MethodJoinMatchmakeSessionWithExtraParticipants:
		protocol.handleJoinMatchmakeSessionWithExtraParticipants(packet)
	case MethodSearchSimpleSearchObjectByObjectIDs:
		protocol.handleSearchSimpleSearchObjectByObjectIDs(packet)
	default:
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Matchmake Extension (Mario Kart 8) method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new MatchmakeExtensionMarioKart8 protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}
	protocol.matchmakeExtensionProtocol.SetServer(server)

	protocol.Setup()

	return protocol
}
