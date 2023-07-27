// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

// UpdateCommunity sets the UpdateCommunity handler function
func (protocol *Protocol) UpdateCommunity(handler func(err error, client *nex.Client, callID uint32, community *match_making_types.PersistentGathering)) {
	protocol.updateCommunityHandler = handler
}

func (protocol *Protocol) handleUpdateCommunity(packet nex.PacketInterface) {
	if protocol.updateCommunityHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateCommunity not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	community, err := parametersStream.ReadStructure(match_making_types.NewPersistentGathering())
	if err != nil {
		go protocol.updateCommunityHandler(fmt.Errorf("Failed to read community from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updateCommunityHandler(nil, client, callID, community.(*match_making_types.PersistentGathering))
}
