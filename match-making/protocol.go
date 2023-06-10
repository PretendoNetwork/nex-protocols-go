package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Match Making protocol
	ProtocolID = 0x15

	// MethodUnregisterGathering is the method ID for the method UnregisterGathering
	MethodUnregisterGathering = 0x2

	// MethodUnregisterGatherings is the method ID for the method UnregisterGatherings
	MethodUnregisterGatherings = 0x3

	// MethodFindBySingleID is the method ID for the method FindBySingleID
	MethodFindBySingleID = 0x15

	// MethodUpdateSessionHostV1 is the method ID for the method UpdateSessionHostV1
	MethodUpdateSessionHostV1 = 0x28

	// MethodGetSessionURLs is the method ID for the method GetSessionURLs
	MethodGetSessionURLs = 0x29

	// MethodUpdateSessionHost is the method ID for the method UpdateSessionHost
	MethodUpdateSessionHost = 0x2A
)

// AuthenticationProtocol handles the Authentication nex protocol
type MatchMakingProtocol struct {
	Server                      *nex.Server
	UnregisterGatheringHandler  func(err error, client *nex.Client, callID uint32, idGathering uint32)
	UnregisterGatheringsHandler func(err error, client *nex.Client, callID uint32, lstGatherings []uint32)
	FindBySingleIDHandler       func(err error, client *nex.Client, callID uint32, id uint32)
	UpdateSessionHostV1Handler  func(err error, client *nex.Client, callID uint32, gid uint32)
	GetSessionURLsHandler       func(err error, client *nex.Client, callID uint32, gid uint32)
	UpdateSessionHostHandler    func(err error, client *nex.Client, callID uint32, gid uint32, isMigrateOwner bool)
}

// Setup initializes the protocol
func (protocol *MatchMakingProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

func (protocol *MatchMakingProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodUnregisterGathering:
		go protocol.HandleUnregisterGathering(packet)
	case MethodUnregisterGatherings:
		go protocol.HandleUnregisterGatherings(packet)
	case MethodFindBySingleID:
		go protocol.HandleFindBySingleID(packet)
	case MethodUpdateSessionHostV1:
		go protocol.HandleUpdateSessionHostV1(packet)
	case MethodGetSessionURLs:
		go protocol.HandleGetSessionURLs(packet)
	case MethodUpdateSessionHost:
		go protocol.HandleUpdateSessionHost(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported MatchMaking method ID: %#v\n", request.MethodID())
	}
}

// NewMatchMakingProtocol returns a new MatchMakingProtocol
func NewMatchMakingProtocol(server *nex.Server) *MatchMakingProtocol {
	protocol := &MatchMakingProtocol{Server: server}

	protocol.Setup()

	return protocol
}
