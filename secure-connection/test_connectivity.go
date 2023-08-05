// Package protocol implements the Secure Connection protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// TestConnectivity sets the TestConnectivity handler function
func (protocol *Protocol) TestConnectivity(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.testConnectivityHandler = handler
}

func (protocol *Protocol) handleTestConnectivity(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.testConnectivityHandler == nil {
		globals.Logger.Warning("SecureConnection::TestConnectivity not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.testConnectivityHandler(nil, client, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
