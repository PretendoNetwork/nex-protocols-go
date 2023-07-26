// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SetExcludeJoinedMatchmakeSession sets the SetExcludeJoinedMatchmakeSession handler function
func (protocol *Protocol) SetExcludeJoinedMatchmakeSession(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.SetExcludeJoinedMatchmakeSessionHandler = handler
}

func (protocol *Protocol) handleSetExcludeJoinedMatchmakeSession(packet nex.PacketInterface) {
	globals.Logger.Warning("Debug::SetExcludeJoinedMatchmakeSession STUBBED")

	if protocol.SetExcludeJoinedMatchmakeSessionHandler == nil {
		globals.Logger.Warning("Debug::SetExcludeJoinedMatchmakeSession not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
}
