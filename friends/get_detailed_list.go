// Package friends implements the Friends QRV protocol
package friends

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetDetailedList sets the GetDetailedList handler function
func (protocol *FriendsProtocol) GetDetailedList(handler func(err error, client *nex.Client, callID uint32, byRelationship uint8, bReversed bool)) {
	protocol.getDetailedListHandler = handler
}

func (protocol *FriendsProtocol) handleGetDetailedList(packet nex.PacketInterface) {
	if protocol.getDetailedListHandler == nil {
		globals.Logger.Warning("Friends::GetDetailedList not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	byRelationship, err := parametersStream.ReadUInt8()
	if err != nil {
		go protocol.getDetailedListHandler(fmt.Errorf("Failed to read byRelationship from parameters. %s", err.Error()), client, callID, 0, false)
		return
	}

	bReversed, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.getDetailedListHandler(fmt.Errorf("Failed to read bReversed from parameters. %s", err.Error()), client, callID, 0, false)
		return
	}

	go protocol.getDetailedListHandler(nil, client, callID, byRelationship, bReversed)
}
