package handle

import (
	"fmt"
	"github.com/wuciyou/dogo/config"
	"github.com/wuciyou/dogo/dglog"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type fileStore struct {
	rootDir string
}

var FileStoreEntity = &fileStore{}

func (fs *fileStore) Open() error {

	rootDir, err := config.GetString("SESSION.FILE_STORE.ROOT_DIR")
	if err != nil {
		return err
	}

	rootDir = filepath.Clean(rootDir)
	rd, err := os.Open(rootDir)
	rd.Close()
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(rootDir, os.ModePerm); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	fs.rootDir = rootDir
	return nil
}

func (fs *fileStore) Close() {

}

func (fs *fileStore) getRootDir() string {
	clearRootDir := filepath.Clean(strings.TrimSpace(fs.rootDir))
	if clearRootDir == "" {
		fs.rootDir = "./session_store"
	}

	return fs.rootDir
}

func (fs *fileStore) Read(sid string) []byte {
	sessionFileName := fmt.Sprintf("%s/%s.session", fs.getRootDir(), sid)
	sessionData, err := ioutil.ReadFile(sessionFileName)
	if err != nil {

		if os.IsNotExist(err) {
			return nil
		} else {
			dglog.Errorf("Can't read session File[%s]: %v", sessionFileName, err)
		}

	}
	return sessionData
}

func (fs *fileStore) Write(sid string, data []byte) {
	sessionName := fmt.Sprintf("%s/%s.session", fs.getRootDir(), sid)

	dglog.Debugf("Write session[sid:%s, data:%s]", sid, string(data))
	ioutil.WriteFile(sessionName, data, os.ModePerm)

}

func (fs *fileStore) Delete(sid string) {

}

func (fs *fileStore) Gc() {

}
