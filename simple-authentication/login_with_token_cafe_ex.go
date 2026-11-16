// Package protocol implements the Simple Authentication protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLoginWithTokenCafeEx(packet nex.PacketInterface) {
	if protocol.LoginWithTokenCafeEx == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SimpleAuthentication::LoginWithTokenCafeEx not implemented")

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
	var oAnyData types.DataHolder

	var err error

	err = strNintendoToken.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.LoginWithTokenCafeEx(fmt.Errorf("failed to read strNintendoToken from parameters. %s", err.Error()), packet, callID, strNintendoToken, oAnyData)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = oAnyData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.LoginWithTokenCafeEx(fmt.Errorf("failed to read oAnyData from parameters. %s", err.Error()), packet, callID, strNintendoToken, oAnyData)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LoginWithTokenEx(nil, packet, callID, strNintendoToken, oAnyData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
