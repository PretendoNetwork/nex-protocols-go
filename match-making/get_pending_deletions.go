// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPendingDeletions sets the GetPendingDeletions handler function
func (protocol *MatchMakingProtocol) GetPendingDeletions(handler func(err error, client *nex.Client, callID uint32, uiReason uint32, resultRange *nex.ResultRange)) {
	protocol.getPendingDeletionsHandler = handler
}

func (protocol *MatchMakingProtocol) handleGetPendingDeletions(packet nex.PacketInterface) {
	if protocol.getPendingDeletionsHandler == nil {
		globals.Logger.Warning("MatchMaking::GetPendingDeletions not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiReason, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getPendingDeletionsHandler(fmt.Errorf("Failed to read uiReason from parameters. %s", err.Error()), client, callID, 0, nil)
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.getPendingDeletionsHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, 0, nil)
	}

	go protocol.getPendingDeletionsHandler(nil, client, callID, uiReason, resultRange.(*nex.ResultRange))
}
