// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateSessionHost sets the UpdateSessionHost handler function
func (protocol *MatchMakingProtocol) UpdateSessionHost(handler func(err error, client *nex.Client, callID uint32, gid uint32, isMigrateOwner bool)) {
	protocol.updateSessionHostHandler = handler
}

func (protocol *MatchMakingProtocol) handleUpdateSessionHost(packet nex.PacketInterface) {
	if protocol.updateSessionHostHandler == nil {
		fmt.Println("[Warning] MatchMaking::UpdateSessionHost not implemented")
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
		go protocol.updateSessionHostHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, false)
	}

	isMigrateOwner, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.updateSessionHostHandler(fmt.Errorf("Failed to read isMigrateOwner from parameters. %s", err.Error()), client, callID, 0, false)
	}

	go protocol.updateSessionHostHandler(nil, client, callID, gid, isMigrateOwner)
}
