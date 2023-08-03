// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SearchPokemon sets the SearchPokemon handler function
func (protocol *Protocol) SearchPokemon(handler func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationSearchPokemonParam)) {
	protocol.searchPokemonHandler = handler
}

func (protocol *Protocol) handleSearchPokemon(packet nex.PacketInterface) {
	if protocol.searchPokemonHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::SearchPokemon not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_pokemon_bank_types.NewGlobalTradeStationSearchPokemonParam())
	if err != nil {
		go protocol.searchPokemonHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.searchPokemonHandler(nil, client, callID, param.(*datastore_pokemon_bank_types.GlobalTradeStationSearchPokemonParam))
}
