// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// TestCapability sets the TestCapability handler function
func (protocol *Protocol) TestCapability(handler func(err error, client *nex.Client, callID uint32, uiCapability uint32)) {
	protocol.testCapabilityHandler = handler
}

func (protocol *Protocol) handleTestCapability(packet nex.PacketInterface) {
	if protocol.testCapabilityHandler == nil {
		globals.Logger.Warning("AccountManagement::TestCapability not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiCapability, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.testCapabilityHandler(fmt.Errorf("Failed to read uiCapability from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.testCapabilityHandler(nil, client, callID, uiCapability)
}
