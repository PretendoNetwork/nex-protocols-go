// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PreparePostObjectV1 sets the PreparePostObjectV1 handler function
func (protocol *Protocol) PreparePostObjectV1(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePreparePostParamV1) uint32) {
	protocol.preparePostObjectV1Handler = handler
}

func (protocol *Protocol) handlePreparePostObjectV1(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.preparePostObjectV1Handler == nil {
		globals.Logger.Warning("DataStore::PreparePostObjectV1 not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStorePreparePostParamV1())
	if err != nil {
		errorCode = protocol.preparePostObjectV1Handler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.preparePostObjectV1Handler(nil, client, callID, param.(*datastore_types.DataStorePreparePostParamV1))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
