// Package protocol implements the Debug protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetAPICalls(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetAPICalls == nil {
		globals.Logger.Warning("Debug::GetAPICalls not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	pids, err := parametersStream.ReadListPID()
	if err != nil {
		_, errorCode = protocol.GetAPICalls(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown, err := parametersStream.ReadDateTime()
	if err != nil {
		_, errorCode = protocol.GetAPICalls(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown2, err := parametersStream.ReadDateTime()
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
