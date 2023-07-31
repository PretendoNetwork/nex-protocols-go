// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SetExcludeJoinedMatchmakeSession sets the SetExcludeJoinedMatchmakeSession handler function
func (protocol *Protocol) SetExcludeJoinedMatchmakeSession(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.setExcludeJoinedMatchmakeSessionHandler = handler
}

func (protocol *Protocol) handleSetExcludeJoinedMatchmakeSession(packet nex.PacketInterface) {
	if protocol.setExcludeJoinedMatchmakeSessionHandler == nil {
		globals.Logger.Warning("Debug::SetExcludeJoinedMatchmakeSession not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("Debug::SetExcludeJoinedMatchmakeSession STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	go protocol.setExcludeJoinedMatchmakeSessionHandler(nil, client, callID, packet.Payload())
}
