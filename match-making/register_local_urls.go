// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterLocalURLs sets the RegisterLocalURLs handler function
func (protocol *MatchMakingProtocol) RegisterLocalURLs(handler func(err error, client *nex.Client, callID uint32, gid uint32, lstURLs []*nex.StationURL)) {
	protocol.registerLocalURLsHandler = handler
}

func (protocol *MatchMakingProtocol) handleRegisterLocalURLs(packet nex.PacketInterface) {
	if protocol.registerLocalURLsHandler == nil {
		globals.Logger.Warning("MatchMaking::RegisterLocalURLs not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.registerLocalURLsHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, nil)
	}

	lstURLs, err := parametersStream.ReadListStationURL()
	if err != nil {
		go protocol.registerLocalURLsHandler(fmt.Errorf("Failed to read lstURLs from parameters. %s", err.Error()), client, callID, 0, nil)
	}

	go protocol.registerLocalURLsHandler(nil, client, callID, gid, lstURLs)
}
