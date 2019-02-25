package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
)

const (
	port            = "8080"
	cookieMaxAge    = 60 * 60 * 48
)

type conferenceWebsiteServer struct {
	confDetailsSvcAddr string
	confDetailsSvcConn *grpc.ClientConn

	scheduleSvcAddr string
	scheduleSvcConn *grpc.ClientConn
}

func main() {
	srvPort := port
	if os.Getenv("PORT") != "" {
		srvPort = os.Getenv("PORT")
	}
	addr := os.Getenv("LISTEN_ADDR")
	svc := new(conferenceWebsiteServer)
	
        //mustMapEnv(&svc.confDetailsSvcAddr, "CONFERENCE_DETAILS_SERVICE_ADDR")
	//mustMapEnv(&svc.scheduleSvcAddr, "SCHEDULE_SERVICE_ADDR")

	//mustConnGRPC(ctx, &svc.confDetailsSvcConn, svc.confDetailsSvcAddr)
	//mustConnGRPC(ctx, &svc.scheduleSvcConn, svc.scheduleSvcAddr)

	r := mux.NewRouter()
	r.HandleFunc("/", svc.homeHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })

	var handler http.Handler = r
        
        http.ListenAndServe(addr+":"+srvPort, handler)
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second*3),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}
