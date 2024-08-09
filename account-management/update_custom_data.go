// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUpdateCustomData(packet nex.PacketInterface) {
	if protocol.UpdateCustomData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::UpdateCustomData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var oPublicData types.AnyDataHolder
	var oPrivateData types.AnyDataHolder

	var err error

	err = oPublicData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateCustomData(fmt.Errorf("Failed to read oPublicData from parameters. %s", err.Error()), packet, callID, oPublicData, oPrivateData)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = oPrivateData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateCustomData(fmt.Errorf("Failed to read oPrivateData from parameters. %s", err.Error()), packet, callID, oPublicData, oPrivateData)
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
