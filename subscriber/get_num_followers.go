// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNumFollowers sets the GetNumFollowers handler function
func (protocol *Protocol) GetNumFollowers(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.getNumFollowersHandler = handler
}

func (protocol *Protocol) handleGetNumFollowers(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getNumFollowersHandler == nil {
		globals.Logger.Warning("Subscriber::GetNumFollowers not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Subscriber::GetNumFollowers STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getNumFollowersHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
