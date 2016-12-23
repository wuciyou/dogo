package dogo

import (
	"fmt"
	"github.com/wuciyou/dogo/common"
	"github.com/wuciyou/dogo/config"
	"github.com/wuciyou/dogo/context"
	"github.com/wuciyou/dogo/dglog"
	"github.com/wuciyou/dogo/hooks"
	"github.com/wuciyou/dogo/pipeline"
	pipelineHandle "github.com/wuciyou/dogo/pipeline/handle"
	"github.com/wuciyou/dogo/router"
	"net/http"
	"time"
)

type dogo struct {
	serveMux *http.ServeMux
	isDebug  bool
}

var DoGo *dogo

func (d *dogo) handler(response http.ResponseWriter, request *http.Request) {
	// 解析请求参数
	request.ParseForm()

	context := &context.Context{}
	context.Parse(response, request)

	isDebug, err := config.EqualFold("LOG.LEVEL", "debug")
	if err != nil {
		dglog.Error(err)
	} else {
		d.isDebug = isDebug
	}

	checkpipelin := pipeline.Each(func(name common.PipelineKey, handle pipeline.PipelineHandle) bool {
		dglog.Debugf("Begin call PipelineRun by name [%s]", name)
		hooks.Listen(common.COMMONPIPELINE_BEGIN, name)
		var beginTime, endTime int64
		// 请求运行时间
		if isDebug {
			beginTime = time.Now().UnixNano()
		}
		pipelineRunData := handle.PipelineRun(context)

		// 请求运行时间
		if isDebug {
			endTime = time.Now().UnixNano()
			dglog.Debugf("Run end PipelineRun by name [%s], begintime:%d, endTime:%d, runtime:[%c[0,0,%dm %d us] %c[0m ", name, beginTime, endTime, 0x1B, 31, (endTime-beginTime)/1000, 0x1B)
		}
		hooks.Listen(common.COMMONPIPELINE_END, name)
		return pipelineRunData
	})

	if !checkpipelin {
		return
	}
}

// start servers
func Start() {
	regisger_pipeline()

	DoGo.start()
}

func regisger_pipeline() {
	isAutoStartSession, err := config.GetBool("SESSION.IS_AUTO_START")
	if err != nil {
		dglog.Error(err)
	}
	if isAutoStartSession {
		// UserSession
		session := &pipelineHandle.Session{}
		pipeline.AddFirst(common.PIPELINE_SESSION, session)
	}

	// 添加日志记录
	request_log := &pipelineHandle.Log{}
	pipeline.AddFirst(common.PIPELINE_LOG, request_log)

	// 添加路由解析
	prouter := &pipelineHandle.Router{}
	pipeline.AddLast(common.PIPELINE_ROUTER, prouter)

	// 将数据刷新到浏览器
	finishRequest := &pipelineHandle.FinishRequest{}
	pipeline.AddLast(common.PIPELINE_FINISH_REQUEST, finishRequest)

	hooks.Add(common.COMMONPIPELINE_END, func(param ...interface{}) {
		dglog.Debugf("Run hooks COMMONPIPELINE_END :%v \n ", param)
	})
}

func (t *dogo) start() {

	// 注册静态资源请求路径
	static_path, err := config.GetString("STATIC_REQUST_PATH")
	if err != nil {
		dglog.Error(err)
	}
	http.HandleFunc(static_path, serverFileController)

	router.Router("/favicon.ico", faviconIcoController)

	http.HandleFunc("/", t.handler)

	listenPort, err := config.GetInt("LISTEN_PORT")
	if err != nil {
		dglog.Error(err)
	}
	listenIp, _ := config.GetString("LISTEN_IP")

	dglog.Infof("Dogo Listen in the port[%s:%d]", listenIp, listenPort)

	http.ListenAndServe(fmt.Sprintf("%s:%d", listenIp, listenPort), nil)
}

func init() {
	DoGo = &dogo{serveMux: http.NewServeMux()}
}
