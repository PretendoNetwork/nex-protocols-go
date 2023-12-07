// Package protocol implements the Ranking (Legacy) protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetSpecificPeriodTotal(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetSpecificPeriodTotal == nil {
		globals.Logger.Warning("RankingLegacy::GetSpecificPeriodTotal not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetSpecificPeriodTotal(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetSpecificPeriodTotal(nil, packet, callID, category)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
