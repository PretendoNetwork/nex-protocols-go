// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// TradePokemon sets the TradePokemon handler function
func (protocol *Protocol) TradePokemon(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationTradePokemonParam) uint32) {
	protocol.tradePokemonHandler = handler
}

func (protocol *Protocol) handleTradePokemon(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.tradePokemonHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::TradePokemon not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_pokemon_bank_types.NewGlobalTradeStationTradePokemonParam())
	if err != nil {
		errorCode = protocol.tradePokemonHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.tradePokemonHandler(nil, packet, callID, param.(*datastore_pokemon_bank_types.GlobalTradeStationTradePokemonParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
