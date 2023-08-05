// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompleteUpdateBankObject sets the CompleteUpdateBankObject handler function
func (protocol *Protocol) CompleteUpdateBankObject(handler func(err error, client *nex.Client, callID uint32, slotID uint16, transactionParam *datastore_pokemon_bank_types.BankTransactionParam, isForce bool) uint32) {
	protocol.completeUpdateBankObjectHandler = handler
}

func (protocol *Protocol) handleCompleteUpdateBankObject(packet nex.PacketInterface) {
	if protocol.completeUpdateBankObjectHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::CompleteUpdateBankObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	slotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		go protocol.completeUpdateBankObjectHandler(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), client, callID, 0, nil, false)
		return
	}

	transactionParam, err := parametersStream.ReadStructure(datastore_pokemon_bank_types.NewBankTransactionParam())
	if err != nil {
		go protocol.completeUpdateBankObjectHandler(fmt.Errorf("Failed to read transactionParam from parameters. %s", err.Error()), client, callID, 0, nil, false)
		return
	}

	isForce, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.completeUpdateBankObjectHandler(fmt.Errorf("Failed to read isForce from parameters. %s", err.Error()), client, callID, 0, nil, false)
		return
	}

	go protocol.completeUpdateBankObjectHandler(nil, client, callID, slotID, transactionParam.(*datastore_pokemon_bank_types.BankTransactionParam), isForce)
}
