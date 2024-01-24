// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDebugNotifyEvent(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.DebugNotifyEvent == nil {
		globals.Logger.Warning("MatchmakeExtension::DebugNotifyEvent not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	pid := types.NewPID(0)
	err = pid.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DebugNotifyEvent(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	mainType := types.NewPrimitiveU32(0)
	err = mainType.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DebugNotifyEvent(fmt.Errorf("Failed to read mainType from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	subType := types.NewPrimitiveU32(0)
	err = subType.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DebugNotifyEvent(fmt.Errorf("Failed to read subType from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param1 := types.NewPrimitiveU64(0)
	err = param1.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DebugNotifyEvent(fmt.Errorf("Failed to read param1 from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param2 := types.NewPrimitiveU64(0)
	err = param2.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DebugNotifyEvent(fmt.Errorf("Failed to read param2 from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	stringParam := types.NewString("")
	err = stringParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DebugNotifyEvent(fmt.Errorf("Failed to read stringParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DebugNotifyEvent(nil, packet, callID, pid, mainType, subType, param1, param2, stringParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
