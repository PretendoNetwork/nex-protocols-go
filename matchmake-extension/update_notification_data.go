// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateNotificationData(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdateNotificationData == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateNotificationData not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiType := types.NewPrimitiveU32(0)
	err = uiType.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateNotificationData(fmt.Errorf("Failed to read uiType from parameters. %s", err.Error()), packet, callID, 0, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiParam1 := types.NewPrimitiveU32(0)
	err = uiParam1.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateNotificationData(fmt.Errorf("Failed to read uiParam1 from parameters. %s", err.Error()), packet, callID, 0, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiParam2 := types.NewPrimitiveU32(0)
	err = uiParam2.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateNotificationData(fmt.Errorf("Failed to read uiParam2 from parameters. %s", err.Error()), packet, callID, 0, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strParam := types.NewString("")
	err = strParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateNotificationData(fmt.Errorf("Failed to read strParam from parameters. %s", err.Error()), packet, callID, 0, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateNotificationData(nil, packet, callID, uiType, uiParam1, uiParam2, strParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
