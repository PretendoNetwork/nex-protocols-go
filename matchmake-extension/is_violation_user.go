// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// IsViolationUser sets the IsViolationUser handler function
func (protocol *MatchmakeExtensionProtocol) IsViolationUser(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.isViolationUserHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleIsViolationUser(packet nex.PacketInterface) {
	if protocol.isViolationUserHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::IsViolationUser not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.isViolationUserHandler(nil, client, callID)
}
