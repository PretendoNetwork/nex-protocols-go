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
	var errorCode uint32

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
		errorCode = protocol.preparePostObjectWithOwnerIDAndDataIDHandler(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), client, callID, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dataID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.preparePostObjectWithOwnerIDAndDataIDHandler(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), client, callID, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		errorCode = protocol.preparePostObjectWithOwnerIDAndDataIDHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.preparePostObjectWithOwnerIDAndDataIDHandler(nil, client, callID, ownerID, dataID, param.(*datastore_types.DataStorePreparePostParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
