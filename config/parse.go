package config

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/wuciyou/dogo/dglog"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type runConfig struct {
	c map[string]string
	m *sync.Mutex
}

var runTimeConfig = &runConfig{c: make(map[string]string), m: &sync.Mutex{}}

func Parse(confiPath string) {
	fc, err := os.Open(confiPath)
	if err != nil {
		dglog.Errorf("Can't open the file '%s' \n %s", confiPath, err)
		return
	}

	fcBuf := bufio.NewReader(fc)

	var configPrefix string

	for {
		linData, _, err := fcBuf.ReadLine()
		if err != nil {

			if err == io.EOF {
				break
			}

			dglog.Errorf("Can't read the file '%s' \n %v", confiPath, err)
			break
		}

		linstr := strings.TrimSpace(string(linData))

		if linstr == "" || strings.HasPrefix(linstr, "#") {
			continue
		}

		bracketsBeginIndex := strings.IndexAny(linstr, "[")
		bracketsEndIndex := strings.IndexAny(linstr, "]")
		if bracketsBeginIndex >= 0 && bracketsEndIndex > 1 {
			configPrefix = strings.TrimSpace(linstr[bracketsBeginIndex+1:bracketsEndIndex]) + "."
			continue
		}
		equalIndex := strings.IndexAny(linstr, "=")
		if equalIndex > 0 && len(linstr) > equalIndex {
			configKey := strings.TrimSpace(linstr[:equalIndex])
			configValue := strings.TrimSpace(linstr[equalIndex+1:])
			Add(configPrefix+configKey, configValue)
			continue
		}
	}
}

func Add(name string, value string) error {
	runTimeConfig.m.Lock()
	defer runTimeConfig.m.Unlock()
	name = strings.TrimSpace(name)
	if len(name) <= 0 || len(value) <= 0 {
		return errors.New(fmt.Sprintf("Can't add '%s' to config map", name))
	}
	runTimeConfig.c[name] = value
	return nil
}

func GetAll() map[string]string {
	return runTimeConfig.c
}

func EqualFold(name string, chars string) (bool, error) {
	values, err := GetString(name)
	if err != nil {
		return false, err
	}
	return strings.EqualFold(strings.ToUpper(values), strings.ToUpper(chars)), nil
}

func GetString(name string) (string, error) {
	runTimeConfig.m.Lock()
	defer runTimeConfig.m.Unlock()
	if value, ok := runTimeConfig.c[name]; ok {
		return value, nil
	}
	return "", errors.New(fmt.Sprintf("No found name:%s", name))
}

func GetInt(name string) (int, error) {
	values, err := GetString(name)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(values)
}

func init() {
	_, file, _, ok := runtime.Caller(1)
	if ok {
		configDir := filepath.Dir(file)
		Parse(configDir + "/default.ini")
	}
	// dglog.Debugf("pc:%v, file:%s, line:%d, ok:%v \n ", pc, file, line, ok)
}
