// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRequestURLs(packet nex.PacketInterface) {
	var err error

	if protocol.RequestURLs == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SecureConnection::RequestURLs not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	cidTarget := types.NewPrimitiveU32(0)
	err = cidTarget.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RequestURLs(fmt.Errorf("Failed to read cidTarget from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	pidTarget := types.NewPID(0)
	err = pidTarget.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RequestURLs(fmt.Errorf("Failed to read pidTarget from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RequestURLs(nil, packet, callID, cidTarget, pidTarget)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
