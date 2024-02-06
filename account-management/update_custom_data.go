// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateCustomData(packet nex.PacketInterface) {
	var err error

	if protocol.UpdateCustomData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::UpdateCustomData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	oPublicData := types.NewAnyDataHolder()
	err = oPublicData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateCustomData(fmt.Errorf("Failed to read oPublicData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	oPrivateData := types.NewAnyDataHolder()
	err = oPrivateData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateCustomData(fmt.Errorf("Failed to read oPrivateData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateCustomData(nil, packet, callID, oPublicData, oPrivateData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
