// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendUserStatuses sets the GetFriendUserStatuses handler function
func (protocol *Protocol) GetFriendUserStatuses(handler func(err error, client *nex.Client, callID uint32, unknown []uint8) uint32) {
	protocol.getFriendUserStatusesHandler = handler
}

func (protocol *Protocol) handleGetFriendUserStatuses(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getFriendUserStatusesHandler == nil {
		globals.Logger.Warning("Subscriber::GetFriendUserStatuses not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown, err := parametersStream.ReadListUInt8()
	if err != nil {
		errorCode = protocol.getFriendUserStatusesHandler(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getFriendUserStatusesHandler(nil, client, callID, unknown)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
