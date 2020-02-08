package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	DataStoreSMMProtocolID = 0x73
)

type DataStoreSMMProtocol struct {
	server *nex.Server
	DataStoreProtocol
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) Setup() {
	nexServer := dataStoreSMMProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.GetRMCRequest()

		if DataStoreSMMProtocolID == request.GetProtocolID() {
			switch request.GetMethodID() {
			case DataStoreMethodGetMeta:
				go dataStoreSMMProtocol.handleGetMeta(packet)
			default:
				fmt.Printf("Unsupported DataStoreSMM method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) respondNotImplemented(packet nex.PacketInterface) {
	client := packet.GetSender()
	request := packet.GetRMCRequest()

	rmcResponse := nex.NewRMCResponse(DataStoreSMMProtocolID, request.GetCallID())
	rmcResponse.SetError(0x80010002)

	rmcResponseBytes := rmcResponse.Bytes()

	var responsePacket nex.PacketInterface
	if packet.GetVersion() == 1 {
		responsePacket, _ = nex.NewPacketV0(client, nil)
	} else {
		responsePacket, _ = nex.NewPacketV1(client, nil)
	}

	responsePacket.SetVersion(packet.GetVersion())
	responsePacket.SetSource(packet.GetDestination())
	responsePacket.SetDestination(packet.GetSource())
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	dataStoreSMMProtocol.server.Send(responsePacket)
}

func NewDataStoreSMMProtocol(server *nex.Server) *DataStoreSMMProtocol {
	dataStoreSMMProtocol := &DataStoreSMMProtocol{server: server}
	dataStoreSMMProtocol.DataStoreProtocol.server = server

	dataStoreSMMProtocol.Setup()

	return dataStoreSMMProtocol
}
