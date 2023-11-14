// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestConnectionData sets the RequestConnectionData handler function
func (protocol *Protocol) RequestConnectionData(handler func(err error, packet nex.PacketInterface, callID uint32, cidTarget uint32, pidTarget *nex.PID) uint32) {
	protocol.requestConnectionDataHandler = handler
}

func (protocol *Protocol) handleRequestConnectionData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.requestConnectionDataHandler == nil {
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
		errorCode = protocol.requestConnectionDataHandler(fmt.Errorf("Failed to read cidTarget from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	pidTarget, err := parametersStream.ReadPID()
	if err != nil {
		errorCode = protocol.requestConnectionDataHandler(fmt.Errorf("Failed to read pidTarget from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.requestConnectionDataHandler(nil, packet, callID, cidTarget, pidTarget)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
