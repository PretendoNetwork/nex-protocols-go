// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteFromDeletions(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.DeleteFromDeletions == nil {
		globals.Logger.Warning("MatchMaking::DeleteFromDeletions not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstDeletions, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		_, errorCode = protocol.DeleteFromDeletions(fmt.Errorf("Failed to read lstDeletions from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DeleteFromDeletions(nil, packet, callID, lstDeletions)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
