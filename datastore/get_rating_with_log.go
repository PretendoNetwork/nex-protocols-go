// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetRatingWithLog(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetRatingWithLog == nil {
		globals.Logger.Warning("DataStore::GetRatingWithLog not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	target := datastore_types.NewDataStoreRatingTarget()
	err = target.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetRatingWithLog(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	accessPassword := types.NewPrimitiveU64(0)
	err = accessPassword.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetRatingWithLog(fmt.Errorf("Failed to read accessPassword from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetRatingWithLog(nil, packet, callID, target, accessPassword)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
