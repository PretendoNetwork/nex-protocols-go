// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostObjectWithOwnerID sets the CompletePostObjectWithOwnerID handler function
func (protocol *Protocol) CompletePostObjectWithOwnerID(handler func(err error, packet nex.PacketInterface, callID uint32, ownerID uint32, param *datastore_types.DataStoreCompletePostParam) uint32) {
	protocol.completePostObjectWithOwnerIDHandler = handler
}

func (protocol *Protocol) handleCompletePostObjectWithOwnerID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.completePostObjectWithOwnerIDHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::CompletePostObjectWithOwnerID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	ownerID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.completePostObjectWithOwnerIDHandler(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompletePostParam())
	if err != nil {
		errorCode = protocol.completePostObjectWithOwnerIDHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.completePostObjectWithOwnerIDHandler(nil, packet, callID, ownerID, param.(*datastore_types.DataStoreCompletePostParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
