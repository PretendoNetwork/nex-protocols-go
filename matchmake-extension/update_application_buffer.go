// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateApplicationBuffer(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdateApplicationBuffer == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateApplicationBuffer not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	gid := types.NewPrimitiveU32(0)
	err = gid.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateApplicationBuffer(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	applicationBuffer := types.NewBuffer(nil)
	err = applicationBuffer.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateApplicationBuffer(fmt.Errorf("Failed to read applicationBuffer from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateApplicationBuffer(nil, packet, callID, gid, applicationBuffer)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
