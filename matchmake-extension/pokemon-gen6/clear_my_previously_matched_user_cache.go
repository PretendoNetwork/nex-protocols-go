// Package protocol implements the Pokemon GEN 6 Matchmake Extension protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ClearMyPreviouslyMatchedUserCache sets the ClearMyPreviouslyMatchedUserCache handler function
func (protocol *Protocol) ClearMyPreviouslyMatchedUserCache(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.clearMyPreviouslyMatchedUserCacheHandler = handler
}

func (protocol *Protocol) handleClearMyPreviouslyMatchedUserCache(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.clearMyPreviouslyMatchedUserCacheHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::ClearMyPreviouslyMatchedUserCache not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.clearMyPreviouslyMatchedUserCacheHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
