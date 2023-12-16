// Package protocol implements the Ranking (Legacy) protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUploadCommonData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UploadCommonData == nil {
		globals.Logger.Warning("RankingLegacy::UploadCommonData not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	uniqueID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.UploadCommonData(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	commonData, err := parametersStream.ReadBuffer()
	if err != nil {
		_, errorCode = protocol.UploadCommonData(fmt.Errorf("Failed to read commonData from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UploadCommonData(nil, packet, callID, uniqueID, commonData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
