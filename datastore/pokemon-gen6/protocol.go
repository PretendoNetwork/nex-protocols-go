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
	Server *nex.Server
	dataStoreProtocol
	uploadPokemonHandler        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationUploadPokemonParam) uint32
	searchPokemonHandler        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationSearchPokemonParam) uint32
	prepareTradePokemonHandler  func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationPrepareTradePokemonParam) uint32
	tradePokemonHandler         func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationTradePokemonParam) uint32
	downloadOtherPokemonHandler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationDownloadOtherPokemonParam) uint32
	downloadMyPokemonHandler    func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationDownloadMyPokemonParam) uint32
	deletePokemonHandler        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationDeletePokemonParam) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.dataStoreProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodUploadPokemon:
		go protocol.handleUploadPokemon(packet)
	case MethodSearchPokemon:
		go protocol.handleSearchPokemon(packet)
	case MethodPrepareTradePokemon:
		go protocol.handlePrepareTradePokemon(packet)
	case MethodTradePokemon:
		go protocol.handleTradePokemon(packet)
	case MethodDownloadOtherPokemon:
		go protocol.handleDownloadOtherPokemon(packet)
	case MethodDownloadMyPokemon:
		go protocol.handleDownloadMyPokemon(packet)
	case MethodDeletePokemon:
		go protocol.handleDeletePokemon(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported DataStore (Pokemon Gen6) method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new DataStorePokemonGen6 protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.dataStoreProtocol.Server = server

	protocol.Setup()

	return protocol
}
