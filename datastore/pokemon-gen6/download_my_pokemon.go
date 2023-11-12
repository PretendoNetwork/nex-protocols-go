// Package protocol implements the DataStorePokemonGen6 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_gen6_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-gen6/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DownloadMyPokemon sets the DownloadMyPokemon handler function
func (protocol *Protocol) DownloadMyPokemon(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationDownloadMyPokemonParam) uint32) {
	protocol.downloadMyPokemonHandler = handler
}

func (protocol *Protocol) handleDownloadMyPokemon(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.downloadMyPokemonHandler == nil {
		globals.Logger.Warning("DataStorePokemonGen6::DownloadMyPokemon not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_pokemon_gen6_types.NewGlobalTradeStationDownloadMyPokemonParam())
	if err != nil {
		errorCode = protocol.downloadMyPokemonHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.downloadMyPokemonHandler(nil, packet, callID, param.(*datastore_pokemon_gen6_types.GlobalTradeStationDownloadMyPokemonParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
