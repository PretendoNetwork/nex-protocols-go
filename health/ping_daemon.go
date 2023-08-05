// Package protocol implements the Health protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PingDaemon sets the PingDaemon handler function
func (protocol *Protocol) PingDaemon(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.pingDaemonHandler = handler
}

func (protocol *Protocol) handlePingDaemon(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.pingDaemonHandler == nil {
		globals.Logger.Warning("Health::PingDaemon not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.pingDaemonHandler(nil, client, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
