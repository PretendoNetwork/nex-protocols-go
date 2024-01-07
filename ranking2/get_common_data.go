// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetCommonData(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetCommonData == nil {
		globals.Logger.Warning("Ranking2::GetCommonData not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	optionFlags := types.NewPrimitiveU32(0)
	err = optionFlags.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetCommonData(fmt.Errorf("Failed to read optionFlags from parameters. %s", err.Error()), packet, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	principalID := types.NewPID(0)
	err = principalID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetCommonData(fmt.Errorf("Failed to read principalID from parameters. %s", err.Error()), packet, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	nexUniqueID := types.NewPrimitiveU64(0)
	err = nexUniqueID.ExtractFrom(parametersStream)
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
