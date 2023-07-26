// Package protocol implements the Nintendo Badge Arcade Secure Connection protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMaintenanceStatus sets the GetMaintenanceStatus function
func (protocol *Protocol) GetMaintenanceStatus(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetMaintenanceStatusHandler = handler
}

func (protocol *Protocol) handleGetMaintenanceStatus(packet nex.PacketInterface) {
	if protocol.GetMaintenanceStatusHandler == nil {
		globals.Logger.Warning("SecureConnectionBadgeArcade::GetMaintenanceStatus not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.GetMaintenanceStatusHandler(nil, client, callID)
}
