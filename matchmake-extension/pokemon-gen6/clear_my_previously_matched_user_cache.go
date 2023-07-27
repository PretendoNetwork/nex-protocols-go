// Package protocol implements the Pokemon GEN 6 Matchmake Extension protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ClearMyPreviouslyMatchedUserCache sets the ClearMyPreviouslyMatchedUserCache handler function
func (protocol *Protocol) ClearMyPreviouslyMatchedUserCache(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.clearMyPreviouslyMatchedUserCacheHandler = handler
}

func (protocol *Protocol) handleClearMyPreviouslyMatchedUserCache(packet nex.PacketInterface) {
	if protocol.clearMyPreviouslyMatchedUserCacheHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::ClearMyPreviouslyMatchedUserCache not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.clearMyPreviouslyMatchedUserCacheHandler(nil, client, callID)
}
