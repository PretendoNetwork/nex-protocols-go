// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ClearMyBlockList sets the ClearMyBlockList handler function
func (protocol *MatchmakeExtensionProtocol) ClearMyBlockList(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.clearMyBlockListHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleClearMyBlockList(packet nex.PacketInterface) {
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
