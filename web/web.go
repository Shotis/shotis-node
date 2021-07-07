// The web server used to host the forward facing paths
package web

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/shotis/shotis-node/config"
	"github.com/shotis/shotis-node/network"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	cryptoutil "github.com/shotis/shotis-node/crypto"
)

type WebServer struct {
	ctx context.Context

	conf *config.NodeConfig

	e *echo.Echo

	rpcConnection *grpc.ClientConn
	rpcClient     network.ShotisServiceClient
}

func addRoutes(router *echo.Router) {
	router.Add("GET", "/api/status/network", func(c echo.Context) error {
		return nil
	})

	router.Add("POST", "/api/upload", func(c echo.Context) error {
		// file := c.FormFile("content")

		return nil
	})
}

func (ws *WebServer) Start(ctx context.Context) error {

	return ws.e.StartTLS(ws.conf.Server.Host, ws.conf.Server.TLS.CertPath, ws.conf.Server.TLS.KeyPath)
}

func Init(c *config.NodeConfig) (*WebServer, error) {
	e := echo.New()
	router := e.Router()

	addRoutes(router)

	certPool, err := cryptoutil.SingleCertificatePool(c.Server.TLS.CertPath)

	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(c.Server.RPC.Host, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(certPool, "")))

	if err != nil {
		return nil, err
	}

	client := network.NewShotisServiceClient(conn)

	return &WebServer{
		conf:          c,
		e:             e,
		rpcConnection: conn,
		rpcClient:     client,
	}, nil
}
