package api

import (
	"blockchain/blockchain"
	"blockchain/blockchain/blocks"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//P2pServer for p2p websocket connections
type P2pServer struct {
	service *blockchain.Blockchain
	sockets []*websocket.Conn
}

// NewP2pServer returns P2pServer Struct and also tries to connect to all the peer networks that are present
func NewP2pServer(bl *blockchain.Blockchain) *P2pServer {
	fmt.Println("Listening already started, connect to Peers")
	toReturn := &P2pServer{bl, []*websocket.Conn{}}
	go toReturn.connectToPeer()
	return toReturn
}

func (p2p *P2pServer) connectToPeer() {

	listing := strings.Split(os.Args[2], ",")
	for i := range listing {
		go p2p.peer(listing[i])
	}
}

// Here we are connecting to all the peers in the network
func (p2p *P2pServer) peer(ad string) {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: ad, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return
	}
	p2p.sockets = append(p2p.sockets, c)
	defer c.Close()

	done := make(chan struct{})
	c.WriteJSON(p2p.service.Chain)
	go func() {

		for {

			_, message, err := c.ReadMessage()

			if err != nil {

				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
					// os.Exit(0)
				}
				log.Println("read:", err)

				return
			}

			log.Printf("recv: %s", string(message))
			fmt.Println("Hi we have received a new chain")

			p2p.messageHandler(c)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			// os.Exit(0)
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}

// Listen function : the peers are trying to connect to this server
func (p2p *P2pServer) Listen(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	p2p.connectSocket(ws)

	p2p.messageHandler(ws)
}

func (p2p *P2pServer) connectSocket(ws *websocket.Conn) {
	p2p.sockets = append(p2p.sockets, ws)
	writeToSocket(ws, p2p.service.Chain)
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
// messageHandler : the servers that are trying to connect to us , once they send message it will be captured in this function's scope
func (p2p *P2pServer) messageHandler(conn *websocket.Conn) {
	for {
		// read in a message
		_, p, err := conn.ReadMessage()
		if err != nil {

			log.Println(err)
			return
		}
		// print out that message for clarity
		var d []*blocks.Block
		if err := json.Unmarshal(p, &d); err != nil {
			panic(err)
		}

		x, v := p2p.service.ReplaceChain(d)
		fmt.Println(x, v)

	}
}

//Syncing all the chains
func (p2p *P2pServer) syncChain() {
	for i := range p2p.sockets {
		p2p.sockets[i].WriteJSON(p2p.service.Chain)
	}
}

func writeToSocket(socket *websocket.Conn, chain []*blocks.Block) {
	socket.WriteJSON(chain)
}
