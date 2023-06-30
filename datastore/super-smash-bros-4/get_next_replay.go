// Package datastore_super_smash_bros_4 implements the Super Smash Bros. 4 DataStore NEX protocol
package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNextReplay sets the GetNextReplay handler function
func (protocol *DataStoreSuperSmashBros4Protocol) GetNextReplay(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetNextReplayHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) handleGetNextReplay(packet nex.PacketInterface) {
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
