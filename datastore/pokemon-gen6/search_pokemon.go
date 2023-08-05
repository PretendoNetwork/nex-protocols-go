// Package protocol implements the DataStorePokemonGen6 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_gen6_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-gen6/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchPokemon sets the SearchPokemon handler function
func (protocol *Protocol) SearchPokemon(handler func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationSearchPokemonParam) uint32) {
	protocol.searchPokemonHandler = handler
}

func (protocol *Protocol) handleSearchPokemon(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.searchPokemonHandler == nil {
		globals.Logger.Warning("DataStorePokemonGen6::SearchPokemon not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_pokemon_gen6_types.NewGlobalTradeStationSearchPokemonParam())
	if err != nil {
		errorCode = protocol.searchPokemonHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.searchPokemonHandler(nil, client, callID, param.(*datastore_pokemon_gen6_types.GlobalTradeStationSearchPokemonParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
