// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Follow sets the Follow handler function
func (protocol *Protocol) Follow(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.followHandler = handler
}

func (protocol *Protocol) handleFollow(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.followHandler == nil {
		globals.Logger.Warning("Subscriber::Follow not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Subscriber::Follow STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.followHandler(nil, client, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
