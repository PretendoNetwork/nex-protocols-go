// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetUserStatuses(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetUserStatuses == nil {
		globals.Logger.Warning("Subscriber::GetUserStatuses not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
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
		_, errorCode = protocol.GetUserStatuses(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown := types.NewList[*types.PrimitiveU8]()
	unknown.Type = types.NewPrimitiveU8(0)
	err = unknown.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetUserStatuses(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetUserStatuses(nil, packet, callID, pids, unknown)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
