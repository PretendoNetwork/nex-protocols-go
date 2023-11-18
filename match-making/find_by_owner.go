// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindByOwner(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.FindByOwner == nil {
		globals.Logger.Warning("MatchMaking::FindByOwner not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id, err := parametersStream.ReadPID()
	if err != nil {
		_, errorCode = protocol.FindByOwner(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := nex.StreamReadStructure(parametersStream, nex.NewResultRange())
	if err != nil {
		_, errorCode = protocol.FindByOwner(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.FindByOwner(nil, packet, callID, id, resultRange)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
