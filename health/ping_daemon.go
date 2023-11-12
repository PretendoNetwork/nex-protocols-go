// Package protocol implements the Health protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PingDaemon sets the PingDaemon handler function
func (protocol *Protocol) PingDaemon(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.pingDaemonHandler = handler
}

func (protocol *Protocol) handlePingDaemon(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.pingDaemonHandler == nil {
		globals.Logger.Warning("Health::PingDaemon not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.pingDaemonHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
