// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMyBlockList sets the GetMyBlockList handler function
func (protocol *Protocol) GetMyBlockList(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.getMyBlockListHandler = handler
}

func (protocol *Protocol) handleGetMyBlockList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getMyBlockListHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetMyBlockList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.getMyBlockListHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
