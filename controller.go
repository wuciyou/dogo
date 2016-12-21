package dogo

import (
	"encoding/base64"
	"github.com/wuciyou/dogo/common"
	"github.com/wuciyou/dogo/config"
	"github.com/wuciyou/dogo/context"
	"github.com/wuciyou/dogo/dglog"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
	"strings"
)

func serverFileController(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, config.RunTimeConfig.StaticRootPath()+r.URL.Path)
}

func faviconIcoController(ctx *context.Context) {
	var icoReader io.Reader
	var err error
	icoReader, err = os.Open("./favicon.ico")
	if err != nil {
		// if err ==
		if os.IsNotExist(err) {
			icoReader = base64.NewDecoder(base64.StdEncoding, strings.NewReader(common.DEFAULT_FAVICON_ICO))
		} else {
			dglog.Errorf("Can't open ./favicon.ico file, msg:%+v", err)
		}
	}

	if icoReader != nil {
		img, _, _ := image.Decode(icoReader)
		png.Encode(ctx.GetWrite(), img)
	}

}
