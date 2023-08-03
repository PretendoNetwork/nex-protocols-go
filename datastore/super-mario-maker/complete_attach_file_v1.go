// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompleteAttachFileV1 sets the CompleteAttachFileV1 handler function
func (protocol *Protocol) CompleteAttachFileV1(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompletePostParamV1)) {
	protocol.completeAttachFileV1Handler = handler
}

func (protocol *Protocol) handleCompleteAttachFileV1(packet nex.PacketInterface) {
	if protocol.completeAttachFileV1Handler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::CompleteAttachFileV1 not implemented")
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
		go protocol.completeAttachFileV1Handler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.completeAttachFileV1Handler(nil, client, callID, param.(*datastore_types.DataStoreCompletePostParamV1))
}
