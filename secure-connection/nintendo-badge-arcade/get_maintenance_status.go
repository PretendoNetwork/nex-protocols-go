// Package protocol implements the Nintendo Badge Arcade Secure Connection protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMaintenanceStatus sets the GetMaintenanceStatus function
func (protocol *Protocol) GetMaintenanceStatus(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.getMaintenanceStatusHandler = handler
}

func (protocol *Protocol) handleGetMaintenanceStatus(packet nex.PacketInterface) {
	if protocol.getMaintenanceStatusHandler == nil {
		globals.Logger.Warning("SecureConnectionNintendoBadgeArcade::GetMaintenanceStatus not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getMaintenanceStatusHandler(nil, client, callID)
}
