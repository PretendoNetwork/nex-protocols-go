// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNumFollowers sets the GetNumFollowers handler function
func (protocol *Protocol) GetNumFollowers(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.getNumFollowersHandler = handler
}

func (protocol *Protocol) handleGetNumFollowers(packet nex.PacketInterface) {
	if protocol.getNumFollowersHandler == nil {
		globals.Logger.Warning("Subscriber::GetNumFollowers not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Subscriber::GetNumFollowers STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getNumFollowersHandler(nil, client, callID, packet.Payload())
}
