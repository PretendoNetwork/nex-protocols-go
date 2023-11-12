// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RollbackBankObject sets the RollbackBankObject handler function
func (protocol *Protocol) RollbackBankObject(handler func(err error, packet nex.PacketInterface, callID uint32, slotID uint16, transactionParam *datastore_pokemon_bank_types.BankTransactionParam, isForce bool) uint32) {
	protocol.rollbackBankObjectHandler = handler
}

func (protocol *Protocol) handleRollbackBankObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.rollbackBankObjectHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::RollbackBankObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	slotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		errorCode = protocol.rollbackBankObjectHandler(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, 0, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	transactionParam, err := parametersStream.ReadStructure(datastore_pokemon_bank_types.NewBankTransactionParam())
	if err != nil {
		errorCode = protocol.rollbackBankObjectHandler(fmt.Errorf("Failed to read transactionParam from parameters. %s", err.Error()), packet, callID, 0, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	isForce, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.rollbackBankObjectHandler(fmt.Errorf("Failed to read isForce from parameters. %s", err.Error()), packet, callID, 0, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.rollbackBankObjectHandler(nil, packet, callID, slotID, transactionParam.(*datastore_pokemon_bank_types.BankTransactionParam), isForce)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
