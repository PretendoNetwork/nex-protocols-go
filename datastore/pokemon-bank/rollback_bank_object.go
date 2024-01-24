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
	var errorCode uint32

	if protocol.RollbackBankObject == nil {
		globals.Logger.Warning("DataStorePokemonBank::RollbackBankObject not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	slotID := types.NewPrimitiveU16(0)
	err = slotID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RollbackBankObject(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	transactionParam := datastore_pokemon_bank_types.NewBankTransactionParam()
	err = transactionParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RollbackBankObject(fmt.Errorf("Failed to read transactionParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	isForce := types.NewPrimitiveBool(false)
	err = isForce.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RollbackBankObject(fmt.Errorf("Failed to read isForce from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RollbackBankObject(nil, packet, callID, slotID, transactionParam, isForce)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
