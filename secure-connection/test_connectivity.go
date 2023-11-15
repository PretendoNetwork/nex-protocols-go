// Package protocol implements the Secure Connection protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleTestConnectivity(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.TestConnectivity == nil {
		globals.Logger.Warning("SecureConnection::TestConnectivity not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.TestConnectivity(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
