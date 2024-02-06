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

	if protocol.GetIntegerSettings == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Utility::GetIntegerSettings not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	integerSettingIndex := types.NewPrimitiveU32(0)
	err = integerSettingIndex.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetIntegerSettings(fmt.Errorf("Failed to read integerSettingIndex from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetIntegerSettings(nil, packet, callID, integerSettingIndex)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
