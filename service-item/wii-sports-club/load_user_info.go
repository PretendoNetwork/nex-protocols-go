// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// LoadUserInfo sets the LoadUserInfo handler function
func (protocol *Protocol) LoadUserInfo(handler func(err error, packet nex.PacketInterface, callID uint32) uint32) {
	protocol.loadUserInfoHandler = handler
}

func (protocol *Protocol) handleLoadUserInfo(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.loadUserInfoHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::LoadUserInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.loadUserInfoHandler(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
