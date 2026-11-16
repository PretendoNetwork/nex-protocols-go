// Package protocol implements the Uplay Win protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetVirtualCurrencyUserBalance(packet nex.PacketInterface) {
	if protocol.GetVirtualCurrencyUserBalance == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UplayWin::GetVirtualCurrencyUserBalance not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var gameCode types.String

	if err := gameCode.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetVirtualCurrencyUserBalance(fmt.Errorf("failed to read gameCode from parameters. %s", err.Error()), packet, callID, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetVirtualCurrencyUserBalance(nil, packet, callID, gameCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
