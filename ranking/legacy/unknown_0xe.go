// Package protocol implements the Ranking (Legacy) protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUnknown0xE(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.Unknown0xE == nil {
		globals.Logger.Warning("RankingLegacy::Unknown0xE not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	rankingMode, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.Unknown0xE(fmt.Errorf("Failed to read rankingMode from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.Unknown0xE(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	scoreIndex, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.Unknown0xE(fmt.Errorf("Failed to read scoreIndex from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown1, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.Unknown0xE(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown2, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.Unknown0xE(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown3, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.Unknown0xE(fmt.Errorf("Failed to read unknown3 from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown4, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.Unknown0xE(fmt.Errorf("Failed to read unknown4 from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown5, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.Unknown0xE(fmt.Errorf("Failed to read unknown5 from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown6, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.Unknown0xE(fmt.Errorf("Failed to read unknown6 from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	offset, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.Unknown0xE(fmt.Errorf("Failed to read offset from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	length, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.Unknown0xE(fmt.Errorf("Failed to read length from parameters. %s", err.Error()), packet, callID, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.Unknown0xE(nil, packet, callID, rankingMode, category, scoreIndex, unknown1, unknown2, unknown3, unknown4, unknown5, unknown6, offset, length)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
