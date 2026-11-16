// Package protocol implements the Simple Authentication protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLoginWithTokenCafe(packet nex.PacketInterface) {
	if protocol.LoginWithTokenCafe == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SimpleAuthentication::LoginWithTokenCafe not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var strNintendoToken types.String

	err := strNintendoToken.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.LoginWithTokenCafe(fmt.Errorf("failed to read strNintendoToken from parameters. %s", err.Error()), packet, callID, strNintendoToken)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LoginWithTokenCafe(nil, packet, callID, strNintendoToken)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
