// Package protocol implements the Debug protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetAPICalls(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetAPICalls == nil {
		globals.Logger.Warning("Debug::GetAPICalls not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	pids := types.NewList[*types.PID]()
	pids.Type = types.NewPID(0)
	err = pids.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetAPICalls(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown := types.NewDateTime(0)
	err = unknown.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetAPICalls(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown2 := types.NewDateTime(0)
	err = unknown2.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetAPICalls(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetAPICalls(nil, packet, callID, pids, unknown, unknown2)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
