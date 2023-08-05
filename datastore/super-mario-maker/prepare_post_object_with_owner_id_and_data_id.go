// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PreparePostObjectWithOwnerIDAndDataID sets the PreparePostObjectWithOwnerIDAndDataID handler function
func (protocol *Protocol) PreparePostObjectWithOwnerIDAndDataID(handler func(err error, client *nex.Client, callID uint32, ownerID uint32, dataID uint64, param *datastore_types.DataStorePreparePostParam) uint32) {
	protocol.preparePostObjectWithOwnerIDAndDataIDHandler = handler
}

func (protocol *Protocol) handlePreparePostObjectWithOwnerIDAndDataID(packet nex.PacketInterface) {
	if protocol.preparePostObjectWithOwnerIDAndDataIDHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::PreparePostObjectWithOwnerIDAndDataID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	ownerID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.preparePostObjectWithOwnerIDAndDataIDHandler(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), client, callID, 0, 0, nil)
		return
	}

	dataID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.preparePostObjectWithOwnerIDAndDataIDHandler(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), client, callID, 0, 0, nil)
		return
	}

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		go protocol.preparePostObjectWithOwnerIDAndDataIDHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, 0, 0, nil)
		return
	}

	go protocol.preparePostObjectWithOwnerIDAndDataIDHandler(nil, client, callID, ownerID, dataID, param.(*datastore_types.DataStorePreparePostParam))
}
