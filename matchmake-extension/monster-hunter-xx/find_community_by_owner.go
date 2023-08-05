// Package protocol implements the MatchmakeExtensionMonsterHunterXX protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindCommunityByOwner sets the FindCommunityByOwner handler function
func (protocol *Protocol) FindCommunityByOwner(handler func(err error, client *nex.Client, callID uint32, id uint64, resultRange *nex.ResultRange) uint32) {
	protocol.findCommunityByOwnerHandler = handler
}

func (protocol *Protocol) handleFindCommunityByOwner(packet nex.PacketInterface) {
	if protocol.findCommunityByOwnerHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMonsterHunterXX::FindCommunityByOwner not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.findCommunityByOwnerHandler(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findCommunityByOwnerHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.findCommunityByOwnerHandler(nil, client, callID, id, resultRange.(*nex.ResultRange))
}
