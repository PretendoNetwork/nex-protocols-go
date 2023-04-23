package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendNotificationData sets the GetFriendNotificationData handler function
func (protocol *MatchmakeExtensionProtocol) GetFriendNotificationData(handler func(err error, client *nex.Client, callID uint32, uiType int32)) {
	protocol.GetFriendNotificationDataHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleGetFriendNotificationData(packet nex.PacketInterface) {
	if protocol.GetFriendNotificationDataHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::GetFriendNotificationData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiType := int32(parametersStream.ReadUInt32LE())

	go protocol.GetFriendNotificationDataHandler(nil, client, callID, uiType)
}