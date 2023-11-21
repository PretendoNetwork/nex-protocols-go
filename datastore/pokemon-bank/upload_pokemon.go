// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUploadPokemon(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UploadPokemon == nil {
		globals.Logger.Warning("DataStorePokemonBank::UploadPokemon not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := nex.StreamReadStructure(parametersStream, datastore_pokemon_bank_types.NewGlobalTradeStationUploadPokemonParam())
	if err != nil {
		_, errorCode = protocol.UploadPokemon(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UploadPokemon(nil, packet, callID, param)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
