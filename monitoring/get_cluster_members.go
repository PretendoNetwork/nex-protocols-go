// Package protocol implements the Monitoring protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetClusterMembers sets the GetClusterMembers handler function
func (protocol *Protocol) GetClusterMembers(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.getClusterMembersHandler = handler
}

func (protocol *Protocol) handleGetClusterMembers(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getClusterMembersHandler == nil {
		globals.Logger.Warning("Monitoring::GetClusterMembers not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getClusterMembersHandler(nil, client, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
