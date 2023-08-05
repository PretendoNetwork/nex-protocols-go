// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareGetObjectV1 sets the PrepareGetObjectV1 handler function
func (protocol *Protocol) PrepareGetObjectV1(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePrepareGetParamV1) uint32) {
	protocol.prepareGetObjectV1Handler = handler
}

func (protocol *Protocol) handlePrepareGetObjectV1(packet nex.PacketInterface) {
	if protocol.prepareGetObjectV1Handler == nil {
		globals.Logger.Warning("DataStore::PrepareGetObjectV1 not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStorePrepareGetParamV1())
	if err != nil {
		go protocol.prepareGetObjectV1Handler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.prepareGetObjectV1Handler(nil, client, callID, param.(*datastore_types.DataStorePrepareGetParamV1))
}
