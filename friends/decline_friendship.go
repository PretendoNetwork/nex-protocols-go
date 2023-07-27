// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeclineFriendship sets the DeclineFriendship handler function
func (protocol *Protocol) DeclineFriendship(handler func(err error, client *nex.Client, callID uint32, uiPlayer uint32)) {
	protocol.declineFriendshipHandler = handler
}

func (protocol *Protocol) handleDeclineFriendship(packet nex.PacketInterface) {
	if protocol.declineFriendshipHandler == nil {
		globals.Logger.Warning("Friends::DeclineFriendship not implemented")
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
		go protocol.declineFriendshipHandler(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.declineFriendshipHandler(nil, client, callID, uiPlayer)
}
