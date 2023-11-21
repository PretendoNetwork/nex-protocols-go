// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindByDescriptionLike(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.FindByDescriptionLike == nil {
		globals.Logger.Warning("MatchMaking::FindByDescriptionLike not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strDescriptionLike, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.FindByDescriptionLike(fmt.Errorf("Failed to read strDescriptionLike from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := nex.StreamReadStructure(parametersStream, nex.NewResultRange())
	if err != nil {
		_, errorCode = protocol.FindByDescriptionLike(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.FindByDescriptionLike(nil, packet, callID, strDescriptionLike, resultRange)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
