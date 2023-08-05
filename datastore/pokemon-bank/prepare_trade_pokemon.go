// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareTradePokemon sets the PrepareTradePokemon handler function
func (protocol *Protocol) PrepareTradePokemon(handler func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationPrepareTradePokemonParam) uint32) {
	protocol.prepareTradePokemonHandler = handler
}

func (protocol *Protocol) handlePrepareTradePokemon(packet nex.PacketInterface) {
	if protocol.prepareTradePokemonHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::PrepareTradePokemon not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_pokemon_bank_types.NewGlobalTradeStationPrepareTradePokemonParam())
	if err != nil {
		go protocol.prepareTradePokemonHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.prepareTradePokemonHandler(nil, client, callID, param.(*datastore_pokemon_bank_types.GlobalTradeStationPrepareTradePokemonParam))
}
