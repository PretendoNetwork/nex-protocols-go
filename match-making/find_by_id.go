// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByID sets the FindByID handler function
func (protocol *MatchMakingProtocol) FindByID(handler func(err error, client *nex.Client, callID uint32, lstID []uint32)) {
	protocol.findByIDHandler = handler
}

func (protocol *MatchMakingProtocol) handleFindByID(packet nex.PacketInterface) {
	if protocol.findByIDHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.findByIDHandler(fmt.Errorf("Failed to read lstID from parameters. %s", err.Error()), client, callID, nil)
	}

	go protocol.findByIDHandler(nil, client, callID, lstID)
}
