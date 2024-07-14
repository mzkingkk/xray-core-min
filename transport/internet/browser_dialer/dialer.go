package browser_dialer

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/common/platform"
	"github.com/xtls/xray-core/common/uuid"
)

//go:embed dialer.html
var webpage []byte

func init() {
	addr := platform.NewEnvFlag(platform.BrowserDialerAddress).GetValue(func() string { return "" })
	if addr != "" {
		token := uuid.New()
		csrfToken := token.String()
		webpage = bytes.ReplaceAll(webpage, []byte("csrfToken"), []byte(csrfToken))
		go http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(webpage)
		}))
	}
}

func HasBrowserDialer() bool {
	return conns != nil
}