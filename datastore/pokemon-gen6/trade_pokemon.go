// Package protocol implements the DataStorePokemonGen6 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_gen6_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-gen6/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// TradePokemon sets the TradePokemon handler function
func (protocol *Protocol) TradePokemon(handler func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationTradePokemonParam) uint32) {
	protocol.tradePokemonHandler = handler
}

func (protocol *Protocol) handleTradePokemon(packet nex.PacketInterface) {
	if protocol.tradePokemonHandler == nil {
		globals.Logger.Warning("DataStorePokemonGen6::TradePokemon not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_pokemon_gen6_types.NewGlobalTradeStationTradePokemonParam())
	if err != nil {
		go protocol.tradePokemonHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.tradePokemonHandler(nil, client, callID, param.(*datastore_pokemon_gen6_types.GlobalTradeStationTradePokemonParam))
}
