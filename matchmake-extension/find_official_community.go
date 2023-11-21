// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindOfficialCommunity(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.FindOfficialCommunity == nil {
		globals.Logger.Warning("MatchmakeExtension::FindOfficialCommunity not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	isAvailableOnly, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.FindOfficialCommunity(fmt.Errorf("Failed to read isAvailableOnly from parameters. %s", err.Error()), packet, callID, false, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := nex.StreamReadStructure(parametersStream, nex.NewResultRange())
	if err != nil {
		_, errorCode = protocol.FindOfficialCommunity(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, false, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.FindOfficialCommunity(nil, packet, callID, isAvailableOnly, resultRange)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
