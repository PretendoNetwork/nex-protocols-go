// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// TestCapability sets the TestCapability handler function
func (protocol *Protocol) TestCapability(handler func(err error, packet nex.PacketInterface, callID uint32, uiCapability uint32) uint32) {
	protocol.testCapabilityHandler = handler
}

func (protocol *Protocol) handleTestCapability(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.testCapabilityHandler == nil {
		globals.Logger.Warning("AccountManagement::TestCapability not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiCapability, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.testCapabilityHandler(fmt.Errorf("Failed to read uiCapability from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.testCapabilityHandler(nil, packet, callID, uiCapability)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
