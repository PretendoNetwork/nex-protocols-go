// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostBankObject sets the CompletePostBankObject handler function
func (protocol *Protocol) CompletePostBankObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompletePostParam) uint32) {
	protocol.completePostBankObjectHandler = handler
}

func (protocol *Protocol) handleCompletePostBankObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.completePostBankObjectHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::CompletePostBankObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompletePostParam())
	if err != nil {
		errorCode = protocol.completePostBankObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.completePostBankObjectHandler(nil, client, callID, param.(*datastore_types.DataStoreCompletePostParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
