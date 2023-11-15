// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handlePrepareUpdateBankObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.PrepareUpdateBankObject == nil {
		globals.Logger.Warning("DataStorePokemonBank::PrepareUpdateBankObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	transactionParam, err := parametersStream.ReadStructure(datastore_pokemon_bank_types.NewBankTransactionParam())
	if err != nil {
		_, errorCode = protocol.PrepareUpdateBankObject(fmt.Errorf("Failed to read transactionParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.PrepareUpdateBankObject(nil, packet, callID, transactionParam.(*datastore_pokemon_bank_types.BankTransactionParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
