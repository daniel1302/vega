package api

import (
	context "context"
	"net"
	"net/http"
	"time"

	"code.vegaprotocol.io/vega/internal/logging"
	uuid "github.com/satori/go.uuid"
)

func unixTimestamp(datetime time.Time) uint64 {
	return uint64(datetime.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)))
}

func newGuid() string {
	return uuid.NewV4().String()
}

func RemoteAddrMiddleware(log *logging.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		found := false
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Warn("Remote address is not splittable in middleware",
				logging.String("remote-addr", r.RemoteAddr))
		} else {
			userIP := net.ParseIP(ip)
			if userIP == nil {
				log.Warn("Remote address is not IP:port format in middleware",
					logging.String("remote-addr", r.RemoteAddr))
			} else {
				found = true

				// Only defined when site is accessed via non-anonymous proxy
				// and takes precedence over RemoteAddr
				forward := r.Header.Get("X-Forwarded-For")
				if forward != "" {
					ip = forward
				}
			}
		}

		if found {
			ctx := context.WithValue(r.Context(), "remote-ip-addr", ip)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
