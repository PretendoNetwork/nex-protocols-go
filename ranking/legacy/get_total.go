// Package protocol implements the Ranking (Legacy) protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetTotal(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetTotal == nil {
		globals.Logger.Warning("RankingLegacy::GetTotal not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetTotal(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown1, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.GetTotal(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown2, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.GetTotal(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown3, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.GetTotal(fmt.Errorf("Failed to read unknown3 from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown4, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetTotal(fmt.Errorf("Failed to read unknown6 from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetTotal(nil, packet, callID, category, unknown1, unknown2, unknown3, unknown4)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
