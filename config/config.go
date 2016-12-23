package config

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/wuciyou/dogo/dglog"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

type runConfig struct {
	c map[string]string
	m *sync.Mutex
}

var runTimeConfig = &runConfig{c: make(map[string]string), m: &sync.Mutex{}}

func parse(rd io.Reader) {
	fcBuf := bufio.NewReader(rd)

	var configPrefix string

	for {
		linData, _, err := fcBuf.ReadLine()
		if err != nil {

			if err == io.EOF {
				break
			}

			dglog.Errorf("Can't read the config  %v", err)
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

func Parse(confiPath string) {
	fc, err := os.Open(confiPath)
	if err != nil {
		dglog.Errorf("Can't open the file '%s' \n %s", confiPath, err)
		return
	}

	parse(fc)
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
	return "", errors.New(fmt.Sprintf("No found name:%s in the config container", name))
}

func GetInt(name string) (int, error) {
	values, err := GetString(name)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(values)
}
func GetBool(name string) (bool, error) {
	values, err := GetString(name)
	if err != nil {
		return false, err
	}
	switch values {
	case "true", "TRUE":
		return true, nil
	case "false", "FALSE":
		return false, nil
	default:
		return false, errors.New(fmt.Sprintf("Unknown values:%s, Only support(true,TRUE or false, FALSE)", values))
	}
}

func init() {
	parse(strings.NewReader(default_conf))

	f, err := os.Open("./config/default.ini")
	if os.IsNotExist(err) {
		dglog.Warning("No found default.ini in your web root")
	} else {
		parse(f)
	}
	// dglog.Debugf("pc:%v, file:%s, line:%d, ok:%v \n ", pc, file, line, ok)
}
