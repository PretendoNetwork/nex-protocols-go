// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateDetails(packet nex.PacketInterface) {
	if protocol.UpdateDetails == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends::UpdateDetails not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiPlayer := types.NewPrimitiveU32(0)
	uiDetails := types.NewPrimitiveU32(0)

	var err error

	err = uiPlayer.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateDetails(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = uiDetails.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateDetails(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateDetails(nil, packet, callID, uiPlayer, uiDetails)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
