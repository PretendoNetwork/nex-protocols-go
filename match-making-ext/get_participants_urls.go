// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetParticipantsURLs(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetParticipantsURLs == nil {
		globals.Logger.Warning("MatchMakingExt::GetParticipantsURLs not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lstGatherings := types.NewList[*types.PrimitiveU32]()
	lstGatherings.Type = types.NewPrimitiveU32(0)
	err = lstGatherings.ExtractFrom(parametersStream)
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
