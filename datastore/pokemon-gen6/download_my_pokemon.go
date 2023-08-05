// Package protocol implements the DataStorePokemonGen6 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_gen6_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-gen6/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DownloadMyPokemon sets the DownloadMyPokemon handler function
func (protocol *Protocol) DownloadMyPokemon(handler func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationDownloadMyPokemonParam) uint32) {
	protocol.downloadMyPokemonHandler = handler
}

func (protocol *Protocol) handleDownloadMyPokemon(packet nex.PacketInterface) {
	if protocol.downloadMyPokemonHandler == nil {
		globals.Logger.Warning("DataStorePokemonGen6::DownloadMyPokemon not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_pokemon_gen6_types.NewGlobalTradeStationDownloadMyPokemonParam())
	if err != nil {
		go protocol.downloadMyPokemonHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.downloadMyPokemonHandler(nil, client, callID, param.(*datastore_pokemon_gen6_types.GlobalTradeStationDownloadMyPokemonParam))
}
