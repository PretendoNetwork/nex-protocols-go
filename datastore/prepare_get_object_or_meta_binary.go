// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareGetObjectOrMetaBinary sets the PrepareGetObjectOrMetaBinary handler function
func (protocol *Protocol) PrepareGetObjectOrMetaBinary(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePrepareGetParam) uint32) {
	protocol.prepareGetObjectOrMetaBinaryHandler = handler
}

func (protocol *Protocol) handlePrepareGetObjectOrMetaBinary(packet nex.PacketInterface) {
	if protocol.prepareGetObjectOrMetaBinaryHandler == nil {
		globals.Logger.Warning("DataStore::PrepareGetObjectOrMetaBinary not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStorePrepareGetParam())
	if err != nil {
		go protocol.prepareGetObjectOrMetaBinaryHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.prepareGetObjectOrMetaBinaryHandler(nil, client, callID, param.(*datastore_types.DataStorePrepareGetParam))
}
