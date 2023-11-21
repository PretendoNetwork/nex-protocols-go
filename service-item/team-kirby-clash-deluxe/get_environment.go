// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetEnvironment(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetEnvironment == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetEnvironment not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uniqueID, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.GetEnvironment(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, "", 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	platform, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.GetEnvironment(fmt.Errorf("Failed to read platform from parameters. %s", err.Error()), packet, callID, "", 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetEnvironment(nil, packet, callID, uniqueID, platform)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
