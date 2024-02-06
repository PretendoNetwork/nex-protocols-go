// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetCategorySetting(packet nex.PacketInterface) {
	if protocol.GetCategorySetting == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking2::GetCategorySetting not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	category := types.NewPrimitiveU32(0)

	err := category.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCategorySetting(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetCategorySetting(nil, packet, callID, category)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
