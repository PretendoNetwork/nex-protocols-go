// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Hello sets the Hello handler function
func (protocol *Protocol) Hello(handler func(err error, packet nex.PacketInterface, callID uint32, name string) uint32) {
	protocol.helloHandler = handler
}

func (protocol *Protocol) handleHello(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.helloHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::Hello not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	name, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.helloHandler(fmt.Errorf("Failed to read name from parameters. %s", err.Error()), packet, callID, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.helloHandler(nil, packet, callID, name)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
