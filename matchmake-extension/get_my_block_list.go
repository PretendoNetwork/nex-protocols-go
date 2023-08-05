// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMyBlockList sets the GetMyBlockList handler function
func (protocol *Protocol) GetMyBlockList(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.getMyBlockListHandler = handler
}

func (protocol *Protocol) handleGetMyBlockList(packet nex.PacketInterface) {
	if protocol.getMyBlockListHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetMyBlockList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getMyBlockListHandler(nil, client, callID)
}
