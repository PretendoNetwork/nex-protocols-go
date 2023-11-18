// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleCompleteUpdateBankObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.CompleteUpdateBankObject == nil {
		globals.Logger.Warning("DataStorePokemonBank::CompleteUpdateBankObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	slotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		_, errorCode = protocol.CompleteUpdateBankObject(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, 0, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	transactionParam, err := nex.StreamReadStructure(parametersStream, datastore_pokemon_bank_types.NewBankTransactionParam())
	if err != nil {
		_, errorCode = protocol.CompleteUpdateBankObject(fmt.Errorf("Failed to read transactionParam from parameters. %s", err.Error()), packet, callID, 0, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	isForce, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.CompleteUpdateBankObject(fmt.Errorf("Failed to read isForce from parameters. %s", err.Error()), packet, callID, 0, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.CompleteUpdateBankObject(nil, packet, callID, slotID, transactionParam, isForce)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
