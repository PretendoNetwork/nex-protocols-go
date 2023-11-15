// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetParticipantsURLs(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetParticipantsURLs == nil {
		globals.Logger.Warning("MatchMakingExt::GetParticipantsURLs not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstGatherings, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetParticipantsURLs(fmt.Errorf("Failed to read lstGatherings from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetParticipantsURLs(nil, packet, callID, lstGatherings)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
