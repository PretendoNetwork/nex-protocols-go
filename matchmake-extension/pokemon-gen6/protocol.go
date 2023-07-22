// Package matchmake_extension_pokemon_gen6 implements the Pokemon GEN 6 Matchmake Extension NEX protocol
package matchmake_extension_pokemon_gen6

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the protocol ID for the Matchmake Extension (Pokemon GEN 6) protocol. ID is the same as the Matchmake Extension protocol
	ProtocolID = 0x6D

	// MethodClearMyPreviouslyMatchedUserCache is the method ID for the method ClearMyPreviouslyMatchedUserCache
	MethodClearMyPreviouslyMatchedUserCache = 0x22
)

var patchedMethods = []uint32{
	MethodClearMyPreviouslyMatchedUserCache,
}

// MatchmakeExtensionPokemonGen6Protocol handles the Matchmake Extension (Pokemon GEN 6) NEX protocol. Embeds MatchmakeExtensionProtocol
type MatchmakeExtensionPokemonGen6Protocol struct {
	Server *nex.Server
	matchmake_extension.MatchmakeExtensionProtocol
	clearMyPreviouslyMatchedUserCacheHandler func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *MatchmakeExtensionPokemonGen6Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.MatchmakeExtensionProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *MatchmakeExtensionPokemonGen6Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodClearMyPreviouslyMatchedUserCache:
		go protocol.handleClearMyPreviouslyMatchedUserCache(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported MatchmakeExtension (Pokemon GEN 6) method ID: %#v\n", request.MethodID())
	}
}

// NewMatchmakeExtensionPokemonGen6Protocol returns a new MatchmakeExtensionPokemonGen6Protocol
func NewMatchmakeExtensionPokemonGen6Protocol(server *nex.Server) *MatchmakeExtensionPokemonGen6Protocol {
	protocol := &MatchmakeExtensionPokemonGen6Protocol{Server: server}
	protocol.MatchmakeExtensionProtocol.Server = server

	protocol.Setup()

	return protocol
}
