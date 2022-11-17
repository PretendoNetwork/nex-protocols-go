package nexproto

import (
	//"errors"
	"fmt"
	"encoding/hex"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// SubscriptionProtocolID is the protocol ID for the Subscription protocol
	SubscriptionProtocolID = 0x75

	// SubscriptionHelloID is the method ID for the method SubscriptionHello
	SubscriptionHelloID = 0x1

	// SubscriptionPostContentID is the method ID for the method SubscriptionPostContent
	SubscriptionPostContentID = 0x2

	// Unk1ID is the method ID for the method Unk1
	SubscriptionUnk1ID = 0x7

	// Unk2ID is the method ID for the method Unk2
	SubscriptionUnk2ID = 0x8

	// GetTimelineID is the method ID for the method GetTimeline
	SubscriptionGetTimelineID = 0x9

	// Unk3ID is the method ID for the method Unk3
	SubscriptionUnk3ID = 0xA

	// Unk4ID is the method ID for the method Unk4
	SubscriptionUnk4ID = 0xb

	// UpdateContentID is the method ID for the method UpdateContent
	SubscriptionUpdateContentID = 0xf
)

// SubscriptionProtocol handles the Subscription nex protocol
type SubscriptionProtocol struct {
	server                *nex.Server
	HelloHandler          func(err error, client *nex.Client, callID uint32, content []byte)
	PostContentHandler    func(err error, client *nex.Client, callID uint32, content []byte)
	Unk1Handler           func(err error, client *nex.Client, callID uint32)
	Unk2Handler           func(err error, client *nex.Client, callID uint32)
	GetTimelineHandler    func(err error, client *nex.Client, callID uint32)
	Unk3Handler           func(err error, client *nex.Client, callID uint32, pids []uint32)
	Unk4Handler           func(err error, client *nex.Client, callID uint32)
	UpdateContentHandler  func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (subscriptionProtocol *SubscriptionProtocol) Setup() {
	nexServer := subscriptionProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if SubscriptionProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case SubscriptionHelloID:
				go subscriptionProtocol.handleHello(packet)
			case SubscriptionPostContentID:
				go subscriptionProtocol.handlePostContent(packet)
			case SubscriptionUnk1ID:
				go subscriptionProtocol.handleUnk1(packet)
			case SubscriptionUnk2ID:
				go subscriptionProtocol.handleUnk2(packet)
			case SubscriptionGetTimelineID:
				go subscriptionProtocol.handleGetTimeline(packet)
			case SubscriptionUnk3ID:
				go subscriptionProtocol.handleUnk3(packet)
			case SubscriptionUnk4ID:
				go subscriptionProtocol.handleUnk4(packet)
			case SubscriptionUpdateContentID:
				go subscriptionProtocol.handleUpdateContent(packet)
			default:
				go respondNotImplemented(packet, SubscriptionProtocolID)
				fmt.Printf("Unsupported Subscription method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// Hello sets the Hello handler function
func (subscriptionProtocol *SubscriptionProtocol) Hello(handler func(err error, client *nex.Client, callID uint32, content []byte)) {
	subscriptionProtocol.HelloHandler = handler
}

// PostContent sets the PostContent handler function
func (subscriptionProtocol *SubscriptionProtocol) PostContent(handler func(err error, client *nex.Client, callID uint32, content []byte)) {
	subscriptionProtocol.PostContentHandler = handler
}

// Unk1 sets the Unk1 handler function
func (subscriptionProtocol *SubscriptionProtocol) Unk1(handler func(err error, client *nex.Client, callID uint32)) {
	subscriptionProtocol.Unk1Handler = handler
}

// Unk2 sets the Unk2 handler function
func (subscriptionProtocol *SubscriptionProtocol) Unk2(handler func(err error, client *nex.Client, callID uint32)) {
	subscriptionProtocol.Unk2Handler = handler
}

// GetTimeline sets the GetTimeline handler function
func (subscriptionProtocol *SubscriptionProtocol) GetTimeline(handler func(err error, client *nex.Client, callID uint32)) {
	subscriptionProtocol.GetTimelineHandler = handler
}

// Unk3 sets the Unk3 handler function
func (subscriptionProtocol *SubscriptionProtocol) Unk3(handler func(err error, client *nex.Client, callID uint32, pids []uint32)) {
	subscriptionProtocol.Unk3Handler = handler
}

// Unk4 sets the Unk4 handler function
func (subscriptionProtocol *SubscriptionProtocol) Unk4(handler func(err error, client *nex.Client, callID uint32)) {
	subscriptionProtocol.Unk4Handler = handler
}

// UpdateContent sets the UpdateContent handler function
func (subscriptionProtocol *SubscriptionProtocol) UpdateContent(handler func(err error, client *nex.Client, callID uint32)) {
	subscriptionProtocol.UpdateContentHandler = handler
}

func (subscriptionProtocol *SubscriptionProtocol) handleHello(packet nex.PacketInterface) {
	if subscriptionProtocol.HelloHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::Hello not implemented")
		go respondNotImplemented(packet, SubscriptionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	content := request.Parameters()[8:]

	//fmt.Println(hex.EncodeToString(parameters))
	_ = parameters

	go subscriptionProtocol.HelloHandler(nil, client, callID, content)
}

func (subscriptionProtocol *SubscriptionProtocol) handlePostContent(packet nex.PacketInterface) {
	if subscriptionProtocol.PostContentHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::PostContent not implemented")
		go respondNotImplemented(packet, SubscriptionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	fmt.Println(hex.EncodeToString(parameters))
	content := request.Parameters()[4:]

	go subscriptionProtocol.PostContentHandler(nil, client, callID, content)
}

func (subscriptionProtocol *SubscriptionProtocol) handleUnk1(packet nex.PacketInterface) {
	if subscriptionProtocol.Unk1Handler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::Unk1 not implemented")
		go respondNotImplemented(packet, SubscriptionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	//fmt.Println(hex.EncodeToString(parameters))
_ = parameters

	go subscriptionProtocol.Unk1Handler(nil, client, callID)
}

func (subscriptionProtocol *SubscriptionProtocol) handleUnk2(packet nex.PacketInterface) {
	if subscriptionProtocol.Unk2Handler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::Unk2 not implemented")
		go respondNotImplemented(packet, SubscriptionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	//fmt.Println(hex.EncodeToString(parameters))
_ = parameters

	go subscriptionProtocol.Unk2Handler(nil, client, callID)
}

func (subscriptionProtocol *SubscriptionProtocol) handleGetTimeline(packet nex.PacketInterface) {
	if subscriptionProtocol.GetTimelineHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetTimeline not implemented")
		go respondNotImplemented(packet, SubscriptionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	//fmt.Println(hex.EncodeToString(parameters))
_ = parameters

	go subscriptionProtocol.GetTimelineHandler(nil, client, callID)
}

func (subscriptionProtocol *SubscriptionProtocol) handleUnk3(packet nex.PacketInterface) {
	if subscriptionProtocol.Unk3Handler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::Unk3 not implemented")
		go respondNotImplemented(packet, SubscriptionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	fmt.Println(hex.EncodeToString(parameters))
_ = parameters

	parametersStream := nex.NewStreamIn(parameters, subscriptionProtocol.server)
	pids := parametersStream.ReadListUInt32LE()

	go subscriptionProtocol.Unk3Handler(nil, client, callID, pids)
}

func (subscriptionProtocol *SubscriptionProtocol) handleUnk4(packet nex.PacketInterface) {
	if subscriptionProtocol.Unk4Handler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::Unk4 not implemented")
		go respondNotImplemented(packet, SubscriptionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	//fmt.Println(hex.EncodeToString(parameters))
_ = parameters

	go subscriptionProtocol.Unk4Handler(nil, client, callID)
}

func (subscriptionProtocol *SubscriptionProtocol) handleUpdateContent(packet nex.PacketInterface) {
	if subscriptionProtocol.UpdateContentHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::UpdateContent not implemented")
		go respondNotImplemented(packet, SubscriptionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	//fmt.Println(hex.EncodeToString(parameters))
_ = parameters

	go subscriptionProtocol.UpdateContentHandler(nil, client, callID)
}

// NewAuthenticationProtocol returns a new AuthenticationProtocol
func NewSubscriptionProtocol(server *nex.Server) *SubscriptionProtocol {
	subscriptionProtocol := &SubscriptionProtocol{server: server}

	subscriptionProtocol.Setup()

	return subscriptionProtocol
}
