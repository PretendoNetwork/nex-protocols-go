// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostMetaBinary sets the PostMetaBinary handler function
func (protocol *Protocol) PostMetaBinary(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePreparePostParam) uint32) {
	protocol.postMetaBinaryHandler = handler
}

func (protocol *Protocol) handlePostMetaBinary(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.postMetaBinaryHandler == nil {
		globals.Logger.Warning("DataStore::PostMetaBinary not implemented")
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
		errorCode = protocol.postMetaBinaryHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.postMetaBinaryHandler(nil, client, callID, param.(*datastore_types.DataStorePreparePostParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
