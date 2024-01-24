// Package protocol implements the Utility protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetIntegerSettings(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetIntegerSettings == nil {
		globals.Logger.Warning("Utility::GetIntegerSettings not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	integerSettingIndex := types.NewPrimitiveU32(0)
	err = integerSettingIndex.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetIntegerSettings(fmt.Errorf("Failed to read integerSettingIndex from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetIntegerSettings(nil, packet, callID, integerSettingIndex)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
