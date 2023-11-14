// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleLoadUserInfo(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.LoadUserInfo == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::LoadUserInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.LoadUserInfo(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
