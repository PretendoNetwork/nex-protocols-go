package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SyncFriend sets the SyncFriend handler function
func (protocol *Friends3DSProtocol) SyncFriend(handler func(err error, client *nex.Client, callID uint32, lfc uint64, pids []uint32, lfcList []uint64)) {
	protocol.SyncFriendHandler = handler
}

func (protocol *Friends3DSProtocol) handleSyncFriend(packet nex.PacketInterface) {
	if protocol.SyncFriendHandler == nil {
		globals.Logger.Warning("Friends3DS::SyncFriend not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.SyncFriendHandler(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), client, callID, 0, nil, nil)
		return
	}

	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.SyncFriendHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), client, callID, 0, nil, nil)
		return
	}

	lfcList, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.SyncFriendHandler(fmt.Errorf("Failed to read lfcList from parameters. %s", err.Error()), client, callID, 0, nil, nil)
		return
	}

	go protocol.SyncFriendHandler(nil, client, callID, lfc, pids, lfcList)
}
