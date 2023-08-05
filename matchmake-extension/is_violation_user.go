// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// IsViolationUser sets the IsViolationUser handler function
func (protocol *Protocol) IsViolationUser(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.isViolationUserHandler = handler
}

func (protocol *Protocol) handleIsViolationUser(packet nex.PacketInterface) {
	if protocol.isViolationUserHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::IsViolationUser not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.isViolationUserHandler(nil, client, callID)
}
