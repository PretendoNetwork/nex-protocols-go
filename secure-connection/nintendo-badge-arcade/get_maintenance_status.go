// Package protocol implements the Nintendo Badge Arcade Secure Connection protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMaintenanceStatus sets the GetMaintenanceStatus function
func (protocol *Protocol) GetMaintenanceStatus(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.getMaintenanceStatusHandler = handler
}

func (protocol *Protocol) handleGetMaintenanceStatus(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getMaintenanceStatusHandler == nil {
		globals.Logger.Warning("SecureConnectionNintendoBadgeArcade::GetMaintenanceStatus not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.getMaintenanceStatusHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
