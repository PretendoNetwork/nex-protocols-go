// Package friends implements the Friends QRV protocol
package friends

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRelationships sets the GetRelationships handler function
func (protocol *FriendsProtocol) GetRelationships(handler func(err error, client *nex.Client, callID uint32, resultRange *nex.ResultRange)) {
	protocol.getRelationshipsHandler = handler
}

func (protocol *FriendsProtocol) handleGetRelationships(packet nex.PacketInterface) {
	if protocol.getRelationshipsHandler == nil {
		globals.Logger.Warning("Friends::GetRelationships not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.getRelationshipsHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getRelationshipsHandler(nil, client, callID, resultRange.(*nex.ResultRange))
}
