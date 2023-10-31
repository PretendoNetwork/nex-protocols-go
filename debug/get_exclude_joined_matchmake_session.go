// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetExcludeJoinedMatchmakeSession sets the GetExcludeJoinedMatchmakeSession handler function
func (protocol *Protocol) GetExcludeJoinedMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.getExcludeJoinedMatchmakeSessionHandler = handler
}

func (protocol *Protocol) handleGetExcludeJoinedMatchmakeSession(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getExcludeJoinedMatchmakeSessionHandler == nil {
		globals.Logger.Warning("Debug::GetExcludeJoinedMatchmakeSession not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("Debug::GetExcludeJoinedMatchmakeSession STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	errorCode = protocol.getExcludeJoinedMatchmakeSessionHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
