// Package protocol implements the Ranking (Legacy) protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUploadScoreWithLimit(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UploadScoreWithLimit == nil {
		globals.Logger.Warning("RankingLegacy::UploadScoreWithLimit not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	uniqueID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.UploadScoreWithLimit(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, 0, 0, nil, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.UploadScoreWithLimit(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, 0, 0, nil, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	scores, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		_, errorCode = protocol.UploadScoreWithLimit(fmt.Errorf("Failed to read scores from parameters. %s", err.Error()), packet, callID, 0, 0, nil, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown1, err := parametersStream.ReadUInt8()
	if err != nil {
		_, errorCode = protocol.UploadScoreWithLimit(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), packet, callID, 0, 0, nil, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown2, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.UploadScoreWithLimit(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, 0, 0, nil, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown3, err := parametersStream.ReadUInt16LE()
	if err != nil {
		_, errorCode = protocol.UploadScoreWithLimit(fmt.Errorf("Failed to read unknown3 from parameters. %s", err.Error()), packet, callID, 0, 0, nil, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UploadScoreWithLimit(nil, packet, callID, uniqueID, category, scores, unknown1, unknown2, unknown3)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
