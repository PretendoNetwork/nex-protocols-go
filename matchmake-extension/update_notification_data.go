package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateNotificationData sets the UpdateNotificationData handler function
func (protocol *MatchmakeExtensionProtocol) UpdateNotificationData(handler func(err error, client *nex.Client, callID uint32, uiType uint32, uiParam1 uint32, uiParam2 uint32, strParam string)) {
	protocol.UpdateNotificationDataHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleUpdateNotificationData(packet nex.PacketInterface) {
	if protocol.UpdateNotificationDataHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateNotificationData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiType, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.UpdateNotificationDataHandler(fmt.Errorf("Failed to read uiType from parameters. %s", err.Error()), client, callID, 0, 0, 0, "")
		return
	}

	uiParam1, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.UpdateNotificationDataHandler(fmt.Errorf("Failed to read uiParam1 from parameters. %s", err.Error()), client, callID, 0, 0, 0, "")
		return
	}

	uiParam2, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.UpdateNotificationDataHandler(fmt.Errorf("Failed to read uiParam2 from parameters. %s", err.Error()), client, callID, 0, 0, 0, "")
		return
	}

	strParam, err := parametersStream.ReadString()
	if err != nil {
		go protocol.UpdateNotificationDataHandler(fmt.Errorf("Failed to read strParam from parameters. %s", err.Error()), client, callID, 0, 0, 0, "")
		return
	}

	go protocol.UpdateNotificationDataHandler(nil, client, callID, uiType, uiParam1, uiParam2, strParam)
}
