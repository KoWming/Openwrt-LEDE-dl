package bridge

import (
	"fmt"
	"github.com/gopherjs/websocket"
	"github.com/v2fly/BrowserBridge/proto"
	"github.com/xtaci/smux"
	"io"
	"net"
	"time"
)

type Settings struct {
	DialAddr string
}

func Bridge(s *Settings) {
	for {
		delay := time.NewTimer(time.Second)
		DoConnect := func() {
			conn, err := websocket.Dial(s.DialAddr)
			if err != nil {
				fmt.Println(err, s.DialAddr)
				return
			}
			smuxc, err := smux.Client(conn, nil)
			if err != nil {
				return
			}
			for {
				stream, err := smuxc.Accept()
				if err != nil {
					fmt.Println(err)
					return
				}
				go func() {
					err, req := proto.ReadRequest(stream)
					if err != nil {
						fmt.Println(err)
						return
					}
					dialfunc := func() (net.Conn, error) {
						return websocket.Dial(req.Destination)
					}
					if req.ProtocolStringSize != 0 {
						dialfunc = func() (net.Conn, error) {
							return websocket.Dial2(req.Destination, req.ProtocolString)
						}
					}
					conn2, err := dialfunc()
					if err != nil {
						fmt.Println(err)
						stream.Close()
						return
					}

					go io.Copy(stream, conn2)
					io.Copy(conn2, stream)
					stream.Close()
				}()

			}
		}
		DoConnect()
		<-delay.C
	}

}
