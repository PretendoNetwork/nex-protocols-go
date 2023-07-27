// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendNotificationData sets the GetFriendNotificationData handler function
func (protocol *Protocol) GetFriendNotificationData(handler func(err error, client *nex.Client, callID uint32, uiType int32)) {
	protocol.getFriendNotificationDataHandler = handler
}

func (protocol *Protocol) handleGetFriendNotificationData(packet nex.PacketInterface) {
	if protocol.getFriendNotificationDataHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetFriendNotificationData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiType, err := parametersStream.ReadInt32LE()
	if err != nil {
		go protocol.getFriendNotificationDataHandler(fmt.Errorf("Failed to read uiType from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getFriendNotificationDataHandler(nil, client, callID, uiType)
}
