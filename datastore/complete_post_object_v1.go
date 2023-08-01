// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostObjectV1 sets the CompletePostObjectV1 handler function
func (protocol *Protocol) CompletePostObjectV1(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompletePostParamV1)) {
	protocol.completePostObjectV1Handler = handler
}

func (protocol *Protocol) handleCompletePostObjectV1(packet nex.PacketInterface) {
	if protocol.completePostObjectV1Handler == nil {
		globals.Logger.Warning("DataStore::CompletePostObjectV1 not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompletePostParamV1())
	if err != nil {
		go protocol.completePostObjectV1Handler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.completePostObjectV1Handler(nil, client, callID, param.(*datastore_types.DataStoreCompletePostParamV1))
}