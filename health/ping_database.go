// Package protocol implements the Health protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PingDatabase sets the PingDatabase handler function
func (protocol *Protocol) PingDatabase(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.pingDatabaseHandler = handler
}

func (protocol *Protocol) handlePingDatabase(packet nex.PacketInterface) {
	if protocol.pingDatabaseHandler == nil {
		globals.Logger.Warning("Health::PingDatabase not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.pingDatabaseHandler(nil, client, callID)
}
