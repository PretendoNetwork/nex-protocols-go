// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ChangeMetaV1 sets the ChangeMetaV1 handler function
func (protocol *Protocol) ChangeMetaV1(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreChangeMetaParamV1)) {
	protocol.changeMetaV1Handler = handler
}

func (protocol *Protocol) handleChangeMetaV1(packet nex.PacketInterface) {
	if protocol.changeMetaV1Handler == nil {
		globals.Logger.Warning("DataStore::ChangeMetaV1 not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreChangeMetaParamV1())
	if err != nil {
		go protocol.changeMetaV1Handler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.changeMetaV1Handler(nil, client, callID, param.(*datastore_types.DataStoreChangeMetaParamV1))
}