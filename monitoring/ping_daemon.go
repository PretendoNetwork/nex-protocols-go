// Package protocol implements the Monitoring protocol
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
	if protocol.pingDaemonHandler == nil {
		globals.Logger.Warning("Monitoring::PingDaemon not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.pingDaemonHandler(nil, client, callID)
}
