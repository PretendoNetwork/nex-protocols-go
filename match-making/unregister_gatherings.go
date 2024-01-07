// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUnregisterGatherings(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UnregisterGatherings == nil {
		globals.Logger.Warning("MatchMaking::UnregisterGatherings not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
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
		_, errorCode = protocol.UnregisterGatherings(fmt.Errorf("Failed to read lstGatherings from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UnregisterGatherings(nil, packet, callID, lstGatherings)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
