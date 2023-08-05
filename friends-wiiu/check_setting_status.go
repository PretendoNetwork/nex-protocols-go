// Package protocol implements the Friends WiiU protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CheckSettingStatus sets the CheckSettingStatus handler function
func (protocol *Protocol) CheckSettingStatus(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.checkSettingStatusHandler = handler
}

func (protocol *Protocol) handleCheckSettingStatus(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.checkSettingStatusHandler == nil {
		globals.Logger.Warning("FriendsWiiU::CheckSettingStatus not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.checkSettingStatusHandler(nil, client, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
