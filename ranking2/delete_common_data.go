// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteCommonData(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.DeleteCommonData == nil {
		globals.Logger.Warning("Ranking2::DeleteCommonData not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	nexUniqueID := types.NewPrimitiveU64(0)
	err = nexUniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DeleteCommonData(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DeleteCommonData(nil, packet, callID, nexUniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
