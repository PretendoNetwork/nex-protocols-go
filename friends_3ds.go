package nexproto

import (
	"fmt"
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// Friends3DSProtocolID is the protocol ID for the Friends (3DS) protocol
	Friends3DSProtocolID = 0x65
	
	// Friends3DSMethodGetFriendMii is the method ID for method GetFriendMii
	Friends3DSMethodGetFriendMii = 0x6
	
	// Friends3DSMethodAddFriendByPrincipalID is the method ID for method AddFriendByPrincipalID
	Friends3DSMethodAddFriendByPrincipalID = 0xb
	
	// Friends3DSMethodSyncFriend is the method ID for method SyncFriend
	Friends3DSMethodSyncFriend = 0x11
	
	// Friends3DSMethodUpdatePresence is the method ID for method UpdatePresence
	Friends3DSMethodUpdatePresence = 0x12
	
	// Friends3DSMethodGetFriendPresence is the method ID for method GetFriendPresence
	Friends3DSMethodGetFriendPresence = 0x16
	
	// Friends3DSMethodGetFriendPersistentInfo is the method ID for method GetFriendPersistentInfo
	Friends3DSMethodGetFriendPersistentInfo = 0x19
)

// Friends3DSProtocol handles the Friends (3DS) nex protocol
type Friends3DSProtocol struct {
	server                              *nex.Server
	UpdatePresenceHandler                func(err error, client *nex.Client, callID uint32, presence *NintendoPresence)
	SyncFriendHandler                    func(err error, client *nex.Client, callID uint32, unknown1 uint64, unknown2 []uint32, unknown3 []uint64)
	AddFriendByPrincipalIDHandler        func(err error, client *nex.Client, callID uint32, unknown1 uint64, principalID uint32)
	GetFriendPersistentInfoHandler       func(err error, client *nex.Client, callID uint32, pidList []uint32)
	GetFriendMiiHandler                  func(err error, client *nex.Client, callID uint32, pidList []uint32)
	GetFriendPresenceHandler             func(err error, client *nex.Client, callID uint32, pidList []uint32)
}

// Setup initializes the protocol
func (friends3DSProtocol *Friends3DSProtocol) Setup() {
	nexServer := friends3DSProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.GetRMCRequest()

		if Friends3DSProtocolID == request.GetProtocolID() {
			switch request.GetMethodID() {
			case Friends3DSMethodSyncFriend:
				go friends3DSProtocol.handleSyncFriend(packet)
			case Friends3DSMethodUpdatePresence:
				go friends3DSProtocol.handleUpdatePresence(packet)
			case Friends3DSMethodAddFriendByPrincipalID:
				go friends3DSProtocol.handleAddFriendByPrincipalID(packet)
			case Friends3DSMethodGetFriendPersistentInfo:
				go friends3DSProtocol.handleGetFriendPersistentInfo(packet)
			case Friends3DSMethodGetFriendMii:
				go friends3DSProtocol.handleGetFriendMii(packet)
			case Friends3DSMethodGetFriendPresence:
				go friends3DSProtocol.handleGetFriendPresence(packet)
			default:
				fmt.Printf("Unsupported Friends (3DS) method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

func (friends3DSProtocol *Friends3DSProtocol) respondNotImplemented(packet nex.PacketInterface) {
	client := packet.GetSender()
	request := packet.GetRMCRequest()

	rmcResponse := nex.NewRMCResponse(Friends3DSProtocolID, request.GetCallID())
	rmcResponse.SetError(0x80010002)

	rmcResponseBytes := rmcResponse.Bytes()

	var responsePacket nex.PacketInterface
	if packet.GetVersion() == 1 {
		responsePacket, _ = nex.NewPacketV1(client, nil)
	} else {
		responsePacket, _ = nex.NewPacketV0(client, nil)
	}

	responsePacket.SetVersion(packet.GetVersion())
	responsePacket.SetSource(packet.GetDestination())
	responsePacket.SetDestination(packet.GetSource())
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	friends3DSProtocol.server.Send(responsePacket)
}

func (friends3DSProtocol *Friends3DSProtocol) handleSyncFriend(packet nex.PacketInterface) {
	if friends3DSProtocol.SyncFriendHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::SyncFriend not implemented")
		go friends3DSProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	Unknown1 := parametersStream.ReadUInt64LE()
	Unknown2 := parametersStream.ReadListUInt32LE()
	Unknown3 := parametersStream.ReadListUInt64LE()

	go friends3DSProtocol.SyncFriendHandler(nil, client, callID, Unknown1, Unknown2, Unknown3)
}

func (friends3DSProtocol *Friends3DSProtocol) handleGetFriendPersistentInfo(packet nex.PacketInterface) {
	if friends3DSProtocol.GetFriendPersistentInfoHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::GetFriendPersistentInfo not implemented")
		go friends3DSProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)
	
	PidList := parametersStream.ReadListUInt32LE()

	go friends3DSProtocol.GetFriendPersistentInfoHandler(nil, client, callID, PidList)
}

func (friends3DSProtocol *Friends3DSProtocol) handleUpdatePresence(packet nex.PacketInterface) {
	if friends3DSProtocol.UpdatePresenceHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::UpdatePresence not implemented")
		go friends3DSProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	nintendoPresenceStructureInterface, err := parametersStream.ReadStructure(NewNintendoPresence())
	if err != nil {
		go friends3DSProtocol.UpdatePresenceHandler(err, client, callID, nil)
		return
	}

	nintendoPresence := nintendoPresenceStructureInterface.(*NintendoPresence)

	go friends3DSProtocol.UpdatePresenceHandler(nil, client, callID, nintendoPresence)
}

func (friends3DSProtocol *Friends3DSProtocol) handleAddFriendByPrincipalID(packet nex.PacketInterface) {
	if friends3DSProtocol.AddFriendByPrincipalIDHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::AddFriendByPrincipalID not implemented")
		go friends3DSProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	Unknown1 := parametersStream.ReadUInt64LE()
	PrincipalID := parametersStream.ReadUInt32LE()

	go friends3DSProtocol.AddFriendByPrincipalIDHandler(nil, client, callID, Unknown1, PrincipalID)
}

func (friends3DSProtocol *Friends3DSProtocol) handleGetFriendMii(packet nex.PacketInterface) {
	if friends3DSProtocol.GetFriendMiiHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::GetFriendMiiHandler not implemented")
		go friends3DSProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)
	
	PidList := parametersStream.ReadListUInt32LE()

	go friends3DSProtocol.GetFriendMiiHandler(nil, client, callID, PidList)
}

func (friends3DSProtocol *Friends3DSProtocol) handleGetFriendPresence(packet nex.PacketInterface) {
	if friends3DSProtocol.GetFriendPresenceHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::GetFriendPresenceHandler not implemented")
		go friends3DSProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)
	
	PidList := parametersStream.ReadListUInt32LE()

	go friends3DSProtocol.GetFriendPresenceHandler(nil, client, callID, PidList)
}

// AddFriendByPrincipalID sets the AddFriendByPrincipalID handler function
func (friends3DSProtocol *Friends3DSProtocol) AddFriendByPrincipalID(handler func(err error, client *nex.Client, callID uint32, unknown1 uint64, principalID uint32)) {
	friends3DSProtocol.AddFriendByPrincipalIDHandler = handler
}

// SyncFriend sets the SyncFriend handler function
func (friends3DSProtocol *Friends3DSProtocol) SyncFriend(handler func(err error, client *nex.Client, callID uint32, unknown1 uint64, unknown2 []uint32, unknown3 []uint64)) {
	friends3DSProtocol.SyncFriendHandler = handler
}

// UpdatePresence sets the UpdatePresence handler function
func (friends3DSProtocol *Friends3DSProtocol) UpdatePresence(handler func(err error, client *nex.Client, callID uint32, presence *NintendoPresence)) {
	friends3DSProtocol.UpdatePresenceHandler = handler
}

// GetFriendPersistentInfo sets the GetFriendPersistentInfo handler function
func (friends3DSProtocol *Friends3DSProtocol) GetFriendPersistentInfo(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	friends3DSProtocol.GetFriendPersistentInfoHandler = handler
}

// GetFriendMii sets the GetFriendMii handler function
func (friends3DSProtocol *Friends3DSProtocol) GetFriendMii(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	friends3DSProtocol.GetFriendMiiHandler = handler
}

// GetFriendPresence sets the GetFriendPresence handler function
func (friends3DSProtocol *Friends3DSProtocol) GetFriendPresence(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	friends3DSProtocol.GetFriendPresenceHandler = handler
}

// NewFriends3DSProtocol returns a new Friends3DSProtocol
func NewFriends3DSProtocol(server *nex.Server) *Friends3DSProtocol {
	Friends3DSProtocol := &Friends3DSProtocol{server: server}

	Friends3DSProtocol.Setup()

	return Friends3DSProtocol
}