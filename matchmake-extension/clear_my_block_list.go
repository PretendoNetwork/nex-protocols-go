// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ClearMyBlockList sets the ClearMyBlockList handler function
func (protocol *Protocol) ClearMyBlockList(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.clearMyBlockListHandler = handler
}

func (protocol *Protocol) handleClearMyBlockList(packet nex.PacketInterface) {
	if protocol.clearMyBlockListHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::ClearMyBlockList not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.clearMyBlockListHandler(nil, client, callID)
}
