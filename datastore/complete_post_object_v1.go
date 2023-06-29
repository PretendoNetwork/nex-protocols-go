package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostObjectV1 sets the CompletePostObjectV1 handler function
func (protocol *DataStoreProtocol) CompletePostObjectV1(handler func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParamV1 *datastore_types.DataStoreCompletePostParamV1)) {
	protocol.CompletePostObjectV1Handler = handler
}

func (protocol *DataStoreProtocol) handleCompletePostObjectV1(packet nex.PacketInterface) {
	if protocol.CompletePostObjectV1Handler == nil {
		globals.Logger.Warning("DataStore::CompletePostObjectV1 not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStoreCompletePostParamV1, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompletePostParamV1())
	if err != nil {
		go protocol.CompletePostObjectV1Handler(fmt.Errorf("Failed to read dataStoreCompletePostParamV1 from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.CompletePostObjectV1Handler(nil, client, callID, dataStoreCompletePostParamV1.(*datastore_types.DataStoreCompletePostParamV1))
}
