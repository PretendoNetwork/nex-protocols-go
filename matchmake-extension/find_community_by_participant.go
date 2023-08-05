// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindCommunityByParticipant sets the FindCommunityByParticipant handler function
func (protocol *Protocol) FindCommunityByParticipant(handler func(err error, client *nex.Client, callID uint32, pid uint32, resultRange *nex.ResultRange) uint32) {
	protocol.findCommunityByParticipantHandler = handler
}

func (protocol *Protocol) handleFindCommunityByParticipant(packet nex.PacketInterface) {
	if protocol.findCommunityByParticipantHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindCommunityByParticipant not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.findCommunityByParticipantHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findCommunityByParticipantHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.findCommunityByParticipantHandler(nil, client, callID, pid, resultRange.(*nex.ResultRange))
}
