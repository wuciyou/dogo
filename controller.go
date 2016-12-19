package dogo

import (
	"encoding/base64"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
	"strings"
)

func serverFileController(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, RunTimeConfig.staticRootPath+r.URL.Path)
}

func faviconIcoController(xtx *Context) {
	var icoReader io.Reader
	var err error
	icoReader, err = os.Open("./favicon.ico")
	if err != nil {
		// if err ==
		if os.IsNotExist(err) {
			icoReader = base64.NewDecoder(base64.StdEncoding, strings.NewReader(DEFAULT_FAVICON_ICO))
		} else {
			DogoLog.Errorf("Can't open ./favicon.ico file, msg:%+v", err)
		}
	}

	if icoReader != nil {
		i, _, _ := image.Decode(icoReader)
		png.Encode(xtx.response.writeBuf, i)
	}

}
