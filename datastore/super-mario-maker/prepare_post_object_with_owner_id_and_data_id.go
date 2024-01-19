// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handlePreparePostObjectWithOwnerIDAndDataID(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.PreparePostObjectWithOwnerIDAndDataID == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::PreparePostObjectWithOwnerIDAndDataID not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	ownerID := types.NewPrimitiveU32(0)
	err = ownerID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.PreparePostObjectWithOwnerIDAndDataID(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dataID := types.NewPrimitiveU64(0)
	err = dataID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.PreparePostObjectWithOwnerIDAndDataID(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param := datastore_types.NewDataStorePreparePostParam()
	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.PreparePostObjectWithOwnerIDAndDataID(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.PreparePostObjectWithOwnerIDAndDataID(nil, packet, callID, ownerID, dataID, param)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
