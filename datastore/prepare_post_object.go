// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PreparePostObject sets the PreparePostObject handler function
func (protocol *Protocol) PreparePostObject(handler func(err error, client *nex.Client, callID uint32, dataStorePreparePostParam *datastore_types.DataStorePreparePostParam)) {
	protocol.PreparePostObjectHandler = handler
}

func (protocol *Protocol) handlePreparePostObject(packet nex.PacketInterface) {
	if protocol.PreparePostObjectHandler == nil {
		globals.Logger.Warning("DataStore::PreparePostObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStorePreparePostParam, err := parametersStream.ReadStructure(datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		go protocol.PreparePostObjectHandler(fmt.Errorf("Failed to read dataStorePreparePostParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.PreparePostObjectHandler(nil, client, callID, dataStorePreparePostParam.(*datastore_types.DataStorePreparePostParam))
}
