// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetCommonData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetCommonData == nil {
		globals.Logger.Warning("Ranking2::GetCommonData not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	optionFlags, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetCommonData(fmt.Errorf("Failed to read optionFlags from parameters. %s", err.Error()), packet, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	principalID, err := parametersStream.ReadPID()
	if err != nil {
		_, errorCode = protocol.GetCommonData(fmt.Errorf("Failed to read principalID from parameters. %s", err.Error()), packet, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	nexUniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		_, errorCode = protocol.GetCommonData(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), packet, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetCommonData(nil, packet, callID, optionFlags, principalID, nexUniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
