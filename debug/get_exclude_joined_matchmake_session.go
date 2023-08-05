// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetExcludeJoinedMatchmakeSession sets the GetExcludeJoinedMatchmakeSession handler function
func (protocol *Protocol) GetExcludeJoinedMatchmakeSession(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.getExcludeJoinedMatchmakeSessionHandler = handler
}

func (protocol *Protocol) handleGetExcludeJoinedMatchmakeSession(packet nex.PacketInterface) {
	if protocol.getExcludeJoinedMatchmakeSessionHandler == nil {
		globals.Logger.Warning("Debug::GetExcludeJoinedMatchmakeSession not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Debug::GetExcludeJoinedMatchmakeSession STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	go protocol.getExcludeJoinedMatchmakeSessionHandler(nil, client, callID, packet.Payload())
}
