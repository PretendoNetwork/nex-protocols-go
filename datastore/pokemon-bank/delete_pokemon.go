// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeletePokemon sets the DeletePokemon handler function
func (protocol *Protocol) DeletePokemon(handler func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationDeletePokemonParam) uint32) {
	protocol.deletePokemonHandler = handler
}

func (protocol *Protocol) handleDeletePokemon(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deletePokemonHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::DeletePokemon not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_pokemon_bank_types.NewGlobalTradeStationDeletePokemonParam())
	if err != nil {
		errorCode = protocol.deletePokemonHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.deletePokemonHandler(nil, client, callID, param.(*datastore_pokemon_bank_types.GlobalTradeStationDeletePokemonParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
