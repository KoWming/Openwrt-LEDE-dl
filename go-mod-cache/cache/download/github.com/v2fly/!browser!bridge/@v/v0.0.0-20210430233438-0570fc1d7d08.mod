module github.com/v2fly/BrowserBridge

go 1.16

require (
	github.com/gopherjs/gopherjs v0.0.0-20210420193930-a4630ec28c79
	github.com/gopherjs/websocket v0.0.0-20191103002815-9a42957e2b3a
	github.com/gorilla/websocket v1.4.2
	github.com/lunixbochs/struc v0.0.0-20200707160740-784aaebc1d40
	github.com/xtaci/smux v1.5.15
)

replace github.com/gopherjs/websocket => ./vendor2/github.com/gopherjs/websocket/
