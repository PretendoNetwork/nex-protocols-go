// Package health implements the Health NEX protocol
package health

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PingDatabase sets the PingDatabase handler function
func (protocol *HealthProtocol) PingDatabase(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.PingDatabaseHandler = handler
}

func (protocol *HealthProtocol) handlePingDatabase(packet nex.PacketInterface) {
	if protocol.PingDatabaseHandler == nil {
		globals.Logger.Warning("Health::PingDatabase not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.PingDatabaseHandler(nil, client, callID)
}
