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
	var errorCode uint32

	if protocol.RequestURLs == nil {
		globals.Logger.Warning("SecureConnection::RequestURLs not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	cidTarget := types.NewPrimitiveU32(0)
	err = cidTarget.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RequestURLs(fmt.Errorf("Failed to read cidTarget from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	pidTarget := types.NewPID(0)
	err = pidTarget.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RequestURLs(fmt.Errorf("Failed to read pidTarget from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RequestURLs(nil, packet, callID, cidTarget, pidTarget)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
