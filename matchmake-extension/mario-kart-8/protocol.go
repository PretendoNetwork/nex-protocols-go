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
	Server *nex.Server
	matchmakeExtensionProtocol
	createSimpleSearchObjectHandler                  func(err error, client *nex.Client, callID uint32, object *matchmake_extension_mario_kart8_types.SimpleSearchObject)
	updateSimpleSearchObjectHandler                  func(err error, client *nex.Client, callID uint32, objectID uint32, newObject *matchmake_extension_mario_kart8_types.SimpleSearchObject)
	deleteSimpleSearchObjectHandler                  func(err error, client *nex.Client, callID uint32, objectID uint32)
	searchSimpleSearchObjectHandler                  func(err error, client *nex.Client, callID uint32, param *matchmake_extension_mario_kart8_types.SimpleSearchParam)
	joinMatchmakeSessionWithExtraParticipantsHandler func(err error, client *nex.Client, callID uint32, packetPayload []byte)
	searchSimpleSearchObjectByObjectIDsHandler       func(err error, client *nex.Client, callID uint32, objectIDs []uint32)
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
	case MethodCreateSimpleSearchObject:
		go protocol.handleCreateSimpleSearchObject(packet)
	case MethodUpdateSimpleSearchObject:
		go protocol.handleUpdateSimpleSearchObject(packet)
	case MethodDeleteSimpleSearchObject:
		go protocol.handleDeleteSimpleSearchObject(packet)
	case MethodSearchSimpleSearchObject:
		go protocol.handleSearchSimpleSearchObject(packet)
	case MethodJoinMatchmakeSessionWithExtraParticipants:
		go protocol.handleJoinMatchmakeSessionWithExtraParticipants(packet)
	case MethodSearchSimpleSearchObjectByObjectIDs:
		go protocol.handleSearchSimpleSearchObjectByObjectIDs(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported Matchmake Extension (Mario Kart 8) method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new MatchmakeExtensionMarioKart8 protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.matchmakeExtensionProtocol.Server = server

	protocol.Setup()

	return protocol
}
