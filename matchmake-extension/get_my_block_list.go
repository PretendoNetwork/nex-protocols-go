// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMyBlockList sets the GetMyBlockList handler function
func (protocol *MatchmakeExtensionProtocol) GetMyBlockList(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getMyBlockListHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleGetMyBlockList(packet nex.PacketInterface) {
	if protocol.getMyBlockListHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetMyBlockList not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getMyBlockListHandler(nil, client, callID)
}
