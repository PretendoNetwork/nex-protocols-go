// Package protocol implements the Simple Authentication protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLoginWithSubAccount(packet nex.PacketInterface) {
	if protocol.LoginWithSubAccount == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SimpleAuthentication::LoginWithSubAccount not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var loginData types.DataHolder

	err := loginData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.LoginWithSubAccount(fmt.Errorf("failed to read loginData from parameters. %s", err.Error()), packet, callID, loginData)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LoginWithSubAccount(nil, packet, callID, loginData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
