// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Register sets the Register handler function
func (protocol *Protocol) Register(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs []*nex.StationURL) uint32) {
	protocol.registerHandler = handler
}

func (protocol *Protocol) handleRegister(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.registerHandler == nil {
		globals.Logger.Warning("SecureConnection::Register not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	vecMyURLs, err := parametersStream.ReadListStationURL()
	if err != nil {
		errorCode = protocol.registerHandler(fmt.Errorf("Failed to read hCustomData from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.registerHandler(nil, packet, callID, vecMyURLs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
