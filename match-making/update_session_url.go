// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateSessionURL sets the UpdateSessionURL handler function
func (protocol *MatchMakingProtocol) UpdateSessionURL(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, strURL string)) {
	protocol.updateSessionURLHandler = handler
}

func (protocol *MatchMakingProtocol) handleUpdateSessionURL(packet nex.PacketInterface) {
	if protocol.updateSessionURLHandler == nil {
		globals.Logger.Warning("MatchMaking::UpdateSessionURL not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.updateSessionURLHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, "")
	}

	strURL, err := parametersStream.ReadString()
	if err != nil {
		go protocol.updateSessionURLHandler(fmt.Errorf("Failed to read strURL from parameters. %s", err.Error()), client, callID, 0, "")
	}

	go protocol.updateSessionURLHandler(nil, client, callID, idGathering, strURL)
}
