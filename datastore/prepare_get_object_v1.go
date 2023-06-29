package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareGetObjectV1 sets the PrepareGetObjectV1 handler function
func (protocol *DataStoreProtocol) PrepareGetObjectV1(handler func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParamV1 *datastore_types.DataStorePrepareGetParamV1)) {
	protocol.PrepareGetObjectV1Handler = handler
}

func (protocol *DataStoreProtocol) handlePrepareGetObjectV1(packet nex.PacketInterface) {
	if protocol.PrepareGetObjectV1Handler == nil {
		globals.Logger.Warning("DataStore::PrepareGetObjectV1 not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStorePrepareGetParamV1, err := parametersStream.ReadStructure(datastore_types.NewDataStorePrepareGetParamV1())
	if err != nil {
		go protocol.PrepareGetObjectV1Handler(fmt.Errorf("Failed to read dataStorePrepareGetParamV1 from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.PrepareGetObjectV1Handler(nil, client, callID, dataStorePrepareGetParamV1.(*datastore_types.DataStorePrepareGetParamV1))
}
