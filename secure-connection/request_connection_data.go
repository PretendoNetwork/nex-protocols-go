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
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	cidTarget, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.RequestConnectionData(fmt.Errorf("Failed to read cidTarget from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	pidTarget, err := parametersStream.ReadPID()
	if err != nil {
		errorCode = protocol.RequestConnectionData(fmt.Errorf("Failed to read pidTarget from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.RequestConnectionData(nil, packet, callID, cidTarget, pidTarget)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
