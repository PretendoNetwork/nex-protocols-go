// Package friends implements the Friends QRV protocol
package friends

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// BlackList sets the BlackList handler function
func (protocol *FriendsProtocol) BlackList(handler func(err error, client *nex.Client, callID uint32, uiPlayer uint32, uiDetails uint32)) {
	protocol.blackListHandler = handler
}

func (protocol *FriendsProtocol) handleBlackList(packet nex.PacketInterface) {
	if protocol.blackListHandler == nil {
		globals.Logger.Warning("Friends::BlackList not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiPlayer, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.blackListHandler(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	uiDetails, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.blackListHandler(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.blackListHandler(nil, client, callID, uiPlayer, uiDetails)
}