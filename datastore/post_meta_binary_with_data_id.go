// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostMetaBinaryWithDataID sets the PostMetaBinaryWithDataID handler function
func (protocol *Protocol) PostMetaBinaryWithDataID(handler func(err error, client *nex.Client, callID uint32, dataID uint64, param *datastore_types.DataStorePreparePostParam)) {
	protocol.postMetaBinaryWithDataIDHandler = handler
}

func (protocol *Protocol) handlePostMetaBinaryWithDataID(packet nex.PacketInterface) {
	if protocol.postMetaBinaryWithDataIDHandler == nil {
		globals.Logger.Warning("DataStore::PostMetaBinaryWithDataID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.postMetaBinaryWithDataIDHandler(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		go protocol.postMetaBinaryWithDataIDHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.postMetaBinaryWithDataIDHandler(nil, client, callID, dataID, param.(*datastore_types.DataStorePreparePostParam))
}