// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Hello sets the Hello handler function
func (protocol *Protocol) Hello(handler func(err error, client *nex.Client, callID uint32, name string)) {
	protocol.helloHandler = handler
}

func (protocol *Protocol) handleHello(packet nex.PacketInterface) {
	if protocol.helloHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::Hello not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	name, err := parametersStream.ReadString()
	if err != nil {
		go protocol.helloHandler(fmt.Errorf("Failed to read name from parameters. %s", err.Error()), client, callID, "")
		return
	}

	go protocol.helloHandler(nil, client, callID, name)
}