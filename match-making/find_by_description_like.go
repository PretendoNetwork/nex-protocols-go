// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByDescriptionLike sets the FindByDescriptionLike handler function
func (protocol *Protocol) FindByDescriptionLike(handler func(err error, client *nex.Client, callID uint32, strDescriptionLike string, resultRange *nex.ResultRange)) {
	protocol.findByDescriptionLikeHandler = handler
}

func (protocol *Protocol) handleFindByDescriptionLike(packet nex.PacketInterface) {
	if protocol.findByDescriptionLikeHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByDescriptionLike not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strDescriptionLike, err := parametersStream.ReadString()
	if err != nil {
		go protocol.findByDescriptionLikeHandler(fmt.Errorf("Failed to read strDescriptionLike from parameters. %s", err.Error()), client, callID, "", nil)
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findByDescriptionLikeHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, "", nil)
	}

	go protocol.findByDescriptionLikeHandler(nil, client, callID, strDescriptionLike, resultRange.(*nex.ResultRange))
}
