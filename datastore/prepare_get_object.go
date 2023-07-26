// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareGetObject sets the PrepareGetObject handler function
func (protocol *Protocol) PrepareGetObject(handler func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParam *datastore_types.DataStorePrepareGetParam)) {
	protocol.PrepareGetObjectHandler = handler
}

func (protocol *Protocol) handlePrepareGetObject(packet nex.PacketInterface) {
	if protocol.PrepareGetObjectHandler == nil {
		globals.Logger.Warning("DataStore::PrepareGetObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStorePrepareGetParam, err := parametersStream.ReadStructure(datastore_types.NewDataStorePrepareGetParam())
	if err != nil {
		go protocol.PrepareGetObjectHandler(fmt.Errorf("Failed to read dataStorePrepareGetParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.PrepareGetObjectHandler(nil, client, callID, dataStorePrepareGetParam.(*datastore_types.DataStorePrepareGetParam))
}
