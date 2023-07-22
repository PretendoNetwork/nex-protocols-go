// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByType sets the FindByType handler function
func (protocol *MatchMakingProtocol) FindByType(handler func(err error, client *nex.Client, callID uint32, strType string, resultRange *nex.ResultRange)) {
	protocol.findByTypeHandler = handler
}

func (protocol *MatchMakingProtocol) handleFindByType(packet nex.PacketInterface) {
	if protocol.findByTypeHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByType not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strType, err := parametersStream.ReadString()
	if err != nil {
		go protocol.findByTypeHandler(fmt.Errorf("Failed to read strType from parameters. %s", err.Error()), client, callID, "", nil)
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findByTypeHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, "", nil)
	}

	go protocol.findByTypeHandler(nil, client, callID, strType, resultRange.(*nex.ResultRange))
}
