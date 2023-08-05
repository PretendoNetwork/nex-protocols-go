// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareUpdateBankObject sets the PrepareUpdateBankObject handler function
func (protocol *Protocol) PrepareUpdateBankObject(handler func(err error, client *nex.Client, callID uint32, transactionParam *datastore_pokemon_bank_types.BankTransactionParam) uint32) {
	protocol.prepareUpdateBankObjectHandler = handler
}

func (protocol *Protocol) handlePrepareUpdateBankObject(packet nex.PacketInterface) {
	if protocol.prepareUpdateBankObjectHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::PrepareUpdateBankObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	transactionParam, err := parametersStream.ReadStructure(datastore_pokemon_bank_types.NewBankTransactionParam())
	if err != nil {
		go protocol.prepareUpdateBankObjectHandler(fmt.Errorf("Failed to read transactionParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.prepareUpdateBankObjectHandler(nil, client, callID, transactionParam.(*datastore_pokemon_bank_types.BankTransactionParam))
}
