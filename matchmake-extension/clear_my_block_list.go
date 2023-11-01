// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ClearMyBlockList sets the ClearMyBlockList handler function
func (protocol *Protocol) ClearMyBlockList(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.clearMyBlockListHandler = handler
}

func (protocol *Protocol) handleClearMyBlockList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.clearMyBlockListHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::ClearMyBlockList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.clearMyBlockListHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
