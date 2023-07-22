// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetlstFriendNotificationData sets the GetlstFriendNotificationData handler function
func (protocol *MatchmakeExtensionProtocol) GetlstFriendNotificationData(handler func(err error, client *nex.Client, callID uint32, lstTypes []uint32)) {
	protocol.getlstFriendNotificationDataHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleGetlstFriendNotificationData(packet nex.PacketInterface) {
	if protocol.getlstFriendNotificationDataHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetlstFriendNotificationData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstTypes, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.getlstFriendNotificationDataHandler(fmt.Errorf("Failed to read lstTypes from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getlstFriendNotificationDataHandler(nil, client, callID, lstTypes)
}
