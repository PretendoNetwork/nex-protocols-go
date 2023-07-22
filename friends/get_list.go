// Package friends implements the Friends QRV protocol
package friends

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetList sets the GetList handler function
func (protocol *FriendsProtocol) GetList(handler func(err error, client *nex.Client, callID uint32, byRelationship uint8, bReversed bool)) {
	protocol.getListHandler = handler
}

func (protocol *FriendsProtocol) handleGetList(packet nex.PacketInterface) {
	if protocol.getListHandler == nil {
		globals.Logger.Warning("Friends::GetList not implemented")
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
		go protocol.getListHandler(fmt.Errorf("Failed to read byRelationship from parameters. %s", err.Error()), client, callID, 0, false)
		return
	}

	bReversed, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.getListHandler(fmt.Errorf("Failed to read bReversed from parameters. %s", err.Error()), client, callID, 0, false)
		return
	}

	go protocol.getListHandler(nil, client, callID, byRelationship, bReversed)
}