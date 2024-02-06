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

	if protocol.DebugNotifyEvent == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::DebugNotifyEvent not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	pid := types.NewPID(0)
	err = pid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DebugNotifyEvent(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	mainType := types.NewPrimitiveU32(0)
	err = mainType.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DebugNotifyEvent(fmt.Errorf("Failed to read mainType from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	subType := types.NewPrimitiveU32(0)
	err = subType.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DebugNotifyEvent(fmt.Errorf("Failed to read subType from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	param1 := types.NewPrimitiveU64(0)
	err = param1.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DebugNotifyEvent(fmt.Errorf("Failed to read param1 from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	param2 := types.NewPrimitiveU64(0)
	err = param2.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DebugNotifyEvent(fmt.Errorf("Failed to read param2 from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	stringParam := types.NewString("")
	err = stringParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DebugNotifyEvent(fmt.Errorf("Failed to read stringParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DebugNotifyEvent(nil, packet, callID, pid, mainType, subType, param1, param2, stringParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
