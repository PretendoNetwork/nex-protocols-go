// Package protocol implements the DataStorePokemonGen6 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore "github.com/PretendoNetwork/nex-protocols-go/datastore"
	datastore_pokemon_gen6_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-gen6/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the DataStore (Pokemon Gen6) protocol
	ProtocolID = 0x73

	// MethodUploadPokemon is the method ID for the UploadPokemon method
	MethodUploadPokemon = 0x29

	// MethodSearchPokemon is the method ID for the SearchPokemon method
	MethodSearchPokemon = 0x2A

	// MethodPrepareTradePokemon is the method ID for the PrepareTradePokemon method
	MethodPrepareTradePokemon = 0x2B

	// MethodTradePokemon is the method ID for the TradePokemon method
	MethodTradePokemon = 0x2C

	// MethodDownloadOtherPokemon is the method ID for the DownloadOtherPokemon method
	MethodDownloadOtherPokemon = 0x2D

	// MethodDownloadMyPokemon is the method ID for the DownloadMyPokemon method
	MethodDownloadMyPokemon = 0x2E

	// MethodDeletePokemon is the method ID for the DeletePokemon method
	MethodDeletePokemon = 0x2F
)

var patchedMethods = []uint32{
	MethodUploadPokemon,
	MethodSearchPokemon,
	MethodPrepareTradePokemon,
	MethodTradePokemon,
	MethodDownloadOtherPokemon,
	MethodDownloadMyPokemon,
	MethodDeletePokemon,
}

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Pokemon Gen6) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	server nex.ServerInterface
	dataStoreProtocol
	UploadPokemon        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationUploadPokemonParam) (*nex.RMCMessage, uint32)
	SearchPokemon        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationSearchPokemonParam) (*nex.RMCMessage, uint32)
	PrepareTradePokemon  func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationPrepareTradePokemonParam) (*nex.RMCMessage, uint32)
	TradePokemon         func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationTradePokemonParam) (*nex.RMCMessage, uint32)
	DownloadOtherPokemon func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationDownloadOtherPokemonParam) (*nex.RMCMessage, uint32)
	DownloadMyPokemon    func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationDownloadMyPokemonParam) (*nex.RMCMessage, uint32)
	DeletePokemon        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationDeletePokemonParam) (*nex.RMCMessage, uint32)
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if !slices.Contains(patchedMethods, message.MethodID) {
		protocol.dataStoreProtocol.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodUploadPokemon:
		protocol.handleUploadPokemon(packet)
	case MethodSearchPokemon:
		protocol.handleSearchPokemon(packet)
	case MethodPrepareTradePokemon:
		protocol.handlePrepareTradePokemon(packet)
	case MethodTradePokemon:
		protocol.handleTradePokemon(packet)
	case MethodDownloadOtherPokemon:
		protocol.handleDownloadOtherPokemon(packet)
	case MethodDownloadMyPokemon:
		protocol.handleDownloadMyPokemon(packet)
	case MethodDeletePokemon:
		protocol.handleDeletePokemon(packet)
	default:
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		fmt.Printf("Unsupported DataStore (Pokemon Gen6) method ID: %#v\n", message.MethodID)
	}
}

// NewProtocol returns a new DataStorePokemonGen6 protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}
	protocol.dataStoreProtocol.SetServer(server)

	return protocol
}
