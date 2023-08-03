// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DownloadMyPokemon sets the DownloadMyPokemon handler function
func (protocol *Protocol) DownloadMyPokemon(handler func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationDownloadMyPokemonParam)) {
	protocol.downloadMyPokemonHandler = handler
}

func (protocol *Protocol) handleDownloadMyPokemon(packet nex.PacketInterface) {
	if protocol.downloadMyPokemonHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::DownloadMyPokemon not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_pokemon_bank_types.NewGlobalTradeStationDownloadMyPokemonParam())
	if err != nil {
		go protocol.downloadMyPokemonHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.downloadMyPokemonHandler(nil, client, callID, param.(*datastore_pokemon_bank_types.GlobalTradeStationDownloadMyPokemonParam))
}
