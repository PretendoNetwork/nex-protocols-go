// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteSimpleSearchObject(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.DeleteSimpleSearchObject == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::DeleteSimpleSearchObject not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	objectID := types.NewPrimitiveU32(0)
	err = objectID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DeleteSimpleSearchObject(fmt.Errorf("Failed to read objectID from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DeleteSimpleSearchObject(nil, packet, callID, objectID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
