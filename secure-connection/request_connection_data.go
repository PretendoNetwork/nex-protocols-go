// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRequestConnectionData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.RequestConnectionData == nil {
		globals.Logger.Warning("SecureConnection::RequestConnectionData not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	cidTarget, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.RequestConnectionData(fmt.Errorf("Failed to read cidTarget from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	pidTarget, err := parametersStream.ReadPID()
	if err != nil {
		_, errorCode = protocol.RequestConnectionData(fmt.Errorf("Failed to read pidTarget from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RequestConnectionData(nil, packet, callID, cidTarget, pidTarget)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
