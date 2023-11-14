// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestURLs sets the RequestURLs handler function
func (protocol *Protocol) RequestURLs(handler func(err error, packet nex.PacketInterface, callID uint32, cidTarget uint32, pidTarget *nex.PID) uint32) {
	protocol.requestURLsHandler = handler
}

func (protocol *Protocol) handleRequestURLs(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.requestURLsHandler == nil {
		globals.Logger.Warning("SecureConnection::RequestURLs not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	cidTarget, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.requestURLsHandler(fmt.Errorf("Failed to read cidTarget from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	pidTarget, err := parametersStream.ReadPID()
	if err != nil {
		errorCode = protocol.requestURLsHandler(fmt.Errorf("Failed to read pidTarget from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.requestURLsHandler(nil, packet, callID, cidTarget, pidTarget)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
