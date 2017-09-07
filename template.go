package dogo

import (
	htmlTemplate "html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var beginSymbol = "{{"
var endSymbol = "}}"

type templateCache struct {
	cacheFileInfo map[string]map[string]os.FileInfo
	cacheData     map[string]*htmlTemplate.Template
}

var tCache = &templateCache{
	cacheFileInfo: make(map[string]map[string]os.FileInfo),
	cacheData:     make(map[string]*htmlTemplate.Template),
}

func (c *templateCache) set(name string, infos map[string]os.FileInfo, data *htmlTemplate.Template) {
	c.cacheFileInfo[name] = infos
	c.cacheData[name] = data

}

func (c *templateCache) check(name string) bool {
	//return false
	Dglog.Debugf("check cache file name :%s", name)
	if infos, ok := c.cacheFileInfo[name]; ok {
		for fileName, fileInfo := range infos {
			info, err := os.Stat(fileName)
			if err != nil {
				Dglog.Warningf("缓存模板文件信息获取失败:%v", err)
				c.del(name)
				return false
			}
			Dglog.Debugf("cache file name :%s, oldTime:%d, newTime:%d, eq:%b", name, fileInfo.ModTime().Unix(), info.ModTime().Unix(), info.ModTime().Equal(fileInfo.ModTime()))
			if !info.ModTime().Equal(fileInfo.ModTime()) {
				c.del(name)
				return false
			}
		}
		Dglog.Debugf("user cache file name :%s", name)
		return true
	}
	return false
}
func (c *templateCache) del(name string) {
	delete(c.cacheFileInfo, name)
	delete(c.cacheData, name)
}
func (c *templateCache) get(name string) *htmlTemplate.Template {
	return c.cacheData[name]
}

type template struct {
	data           []byte
	dataLength     int
	path           string
	startPoint     int
	len            int
	childs         []*template
	childFileNames map[string]os.FileInfo
	ht             *htmlTemplate.Template
	assignData     map[string]interface{}
}

func NewTemplate(fileName string) *template {

	startTime := time.Now()
	Dglog.Debugf("new template file[%s]", fileName)
	t := &template{}
	t.childFileNames = make(map[string]os.FileInfo)
	t.assignData = make(map[string]interface{})
	t.path = filepath.Dir(fileName)

	if tCache.check(fileName) {
		Dglog.Debugf("Have cache file:%s", fileName)
		t.ht = tCache.get(fileName)
	} else {
		t.Open(fileName)
		t.ht = htmlTemplate.New(fileName)
		t.ht.Parse(string(t.data))
		tCache.set(fileName, t.childFileNames, t.ht)
	}

	endTime := time.Now()
	Dglog.Debugf("total runtime %d", endTime.Nanosecond()-startTime.Nanosecond())
	return t
}

func (t *template) Display(w io.Writer) {
	Dglog.Debugf("childFileNamds%+v", t.childFileNames)
	t.ht.Execute(w, t.assignData)
	t.assignData["name"] = "今天的天气不错"
	t.ht.Execute(w, t.assignData)
}

func (t *template) Assign(name string, data interface{}) {
	t.assignData[name] = data
}

func (t *template) Open(fileName string) {

	f, err := os.Open(fileName)
	if err != nil {
		Dglog.Errorf("Can't open template file by name %s \n %+v", fileName, err)
		return
	}
	fData, err := ioutil.ReadAll(f)
	if err != nil {
		Dglog.Errorf("Can't read template file by name %s \n %+v", fileName, err)
		return
	}
	fileInfo, _ := os.Stat(fileName)
	t.childFileNames[fileName] = fileInfo
	t.data = fData
	t.parser()
}

func (t *template) Include(data map[string]string) []byte {

	if data == nil || len(data) <= 0 {
		return nil
	}
	if fileName, ok := data["file"]; ok {
		if !filepath.IsAbs(fileName) {
			fileName = t.path + "/" + fileName
		}
		Dglog.Debugf("include name:%s", fileName)
		childTemplate := NewTemplate(fileName)

		Dglog.Debugf("1 new template file[%s], childTemplate:%v", fileName, childTemplate.data)
		for tFileName, tFileInfo := range childTemplate.childFileNames {
			t.childFileNames[tFileName] = tFileInfo
		}
		return childTemplate.data
	}
	return []byte{}
}

func (t *template) parser() {
	dataLen := len(t.data)
	if dataLen <= 0 {
		return
	}
	t.dataLength = dataLen
	var beginPoint = 0
	var commandData []byte
	for i := 0; i < dataLen; {
		switch t.data[i] {
		case beginSymbol[0]:
			if i+len(beginSymbol) > dataLen {
				break
			}
			if strings.Compare(string(t.data[i:i+len(beginSymbol)]), beginSymbol) == 0 {
				i = i + len(beginSymbol)
				beginPoint = i
			}
			break
		case endSymbol[0]:
			if beginPoint <= 0 {
				break
			}
			if len(endSymbol) > dataLen-i {
				break
			}
			if !(strings.Compare(string(t.data[i:i+len(endSymbol)]), endSymbol) == 0) {
				break
			}
			commandData = t.data[beginPoint:i]
			i = i + len(endSymbol)
			commandName, attrMap := t.parserCommand(commandData)
			Dglog.Debugf("commandName:%s", commandName)
			switch strings.ToLower(commandName) {
			case "include":
				commandContext := t.Include(attrMap)
				t.data = append(t.data[:beginPoint-len(beginSymbol)], append(commandContext[:], t.data[i:]...)...)
				dataLen = len(t.data)
				Dglog.Debugf("1 i:%d", i)
				i = i + (len(commandContext) - (i - (beginPoint - len(beginSymbol)))) - 1
				Dglog.Debugf("2 i:%d", i)
			}
			beginPoint = 0
			break
		default:
		}
		i++
	}
}

func (t *template) parserCommand(commandData []byte) (commandName string, attrMap map[string]string) {
	commandDataLen := len(commandData)
	if commandDataLen <= 0 {
		return
	}
	attrMap = make(map[string]string)
	var attrName string
	var attrValue string
	var startPoint = 0
	var endPoint = 0
	for i := 0; i < commandDataLen; {
		switch {
		case commandData[i] == '=':
			startPoint++
			endPoint++
			if endPoint >= startPoint && commandName != "" {
				attrName = string(commandData[startPoint:endPoint])
			}
		case commandData[i] == '"' || commandData[i] == '\'':
			i++
			startPoint = i
			for i < commandDataLen && attrName != "" {

				if (commandData[i] == '\'' || commandData[i] == '"') && commandData[i-1] != '\\' {
					endPoint = i
					attrValue = string(commandData[startPoint:endPoint])
					if attrName != "" && attrValue != "" {
						startPoint = i
						attrMap[attrName] = attrValue
						attrName = ""
						attrValue = ""
						break
					}
				} else {
					i++
				}
			}

		case commandData[i] == ' ' || commandData[i] == '	':
			if endPoint >= startPoint && commandName == "" {
				endPoint = i
				commandName = string(commandData[startPoint:endPoint])
			}
			startPoint = i
		//case (commandData[i] >= 'A' && commandData[i] < 'Z') || (commandData[i] >= 'a' && commandData[i] <= 'z'):
		default:
			endPoint = i
		}
		i++
	}
	return
}
