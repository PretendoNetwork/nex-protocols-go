// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// LoadUserInfo sets the LoadUserInfo handler function
func (protocol *Protocol) LoadUserInfo(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.loadUserInfoHandler = handler
}

func (protocol *Protocol) handleLoadUserInfo(packet nex.PacketInterface) {
	if protocol.loadUserInfoHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::LoadUserInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.loadUserInfoHandler(nil, client, callID)
}
