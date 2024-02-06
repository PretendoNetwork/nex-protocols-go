// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRollbackBankObject(packet nex.PacketInterface) {
	var err error

	if protocol.RollbackBankObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStorePokemonBank::RollbackBankObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	slotID := types.NewPrimitiveU16(0)
	err = slotID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RollbackBankObject(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	transactionParam := datastore_pokemon_bank_types.NewBankTransactionParam()
	err = transactionParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RollbackBankObject(fmt.Errorf("Failed to read transactionParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	isForce := types.NewPrimitiveBool(false)
	err = isForce.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RollbackBankObject(fmt.Errorf("Failed to read isForce from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RollbackBankObject(nil, packet, callID, slotID, transactionParam, isForce)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
