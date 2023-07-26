// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNextReplay sets the GetNextReplay handler function
func (protocol *Protocol) GetNextReplay(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetNextReplayHandler = handler
}

func (protocol *Protocol) handleGetNextReplay(packet nex.PacketInterface) {
	if protocol.GetNextReplayHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::GetNextReplay not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.GetNextReplayHandler(nil, client, callID)
}
