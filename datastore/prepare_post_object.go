// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PreparePostObject sets the PreparePostObject handler function
func (protocol *Protocol) PreparePostObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePreparePostParam) uint32) {
	protocol.preparePostObjectHandler = handler
}

func (protocol *Protocol) handlePreparePostObject(packet nex.PacketInterface) {
	if protocol.preparePostObjectHandler == nil {
		globals.Logger.Warning("DataStore::PreparePostObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		go protocol.preparePostObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.preparePostObjectHandler(nil, client, callID, param.(*datastore_types.DataStorePreparePostParam))
}
