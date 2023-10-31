// Package protocol implements the Health protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PingDatabase sets the PingDatabase handler function
func (protocol *Protocol) PingDatabase(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.pingDatabaseHandler = handler
}

func (protocol *Protocol) handlePingDatabase(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.pingDatabaseHandler == nil {
		globals.Logger.Warning("Health::PingDatabase not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.pingDatabaseHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
