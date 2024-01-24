// Package protocol implements the Utility protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetStringSettings(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetStringSettings == nil {
		globals.Logger.Warning("Utility::GetStringSettings not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	stringSettingIndex := types.NewPrimitiveU32(0)
	err = stringSettingIndex.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetStringSettings(fmt.Errorf("Failed to read stringSettingIndex from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetStringSettings(nil, packet, callID, stringSettingIndex)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
