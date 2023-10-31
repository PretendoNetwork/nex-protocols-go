// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Hello sets the Hello handler function
func (protocol *Protocol) Hello(handler func(err error, packet nex.PacketInterface, callID uint32, unknown string) uint32) {
	protocol.helloHandler = handler
}

func (protocol *Protocol) handleHello(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.helloHandler == nil {
		globals.Logger.Warning("Subscriber::Hello not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.helloHandler(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.helloHandler(nil, packet, callID, unknown)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
