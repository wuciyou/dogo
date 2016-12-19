package dogo

import (
	"net/http"
	"time"
)

type dogo struct {
	serveMux *http.ServeMux
}

var DoGo *dogo

func (d *dogo) handler(response http.ResponseWriter, request *http.Request) {
	// 解析请求参数
	request.ParseForm()

	context := &Context{}
	context.parse(response, request)

	checkpipelin := Commonpipeline.each(func(pipelin *pipelineNode) bool {
		DogoLog.Debugf("Begin call PipelineRun by name [%s]", pipelin.name)
		var beginTime, endTime int64
		// 请求运行时间
		if RunTimeConfig.IsDebug() {
			beginTime = time.Now().UnixNano()
		}
		pipelineRunData := pipelin.h.PipelineRun(context)

		// 请求运行时间
		if RunTimeConfig.IsDebug() {
			endTime = time.Now().UnixNano()
			DogoLog.Debugf("Run end PipelineRun by name [%s], begintime:%d, endTime:%d, runtime:[%c[0,0,%dm %d us] %c[0m ", pipelin.name, beginTime, endTime, 0x1B, 31, (endTime-beginTime)/1000, 0x1B)
		}

		return pipelineRunData
	})

	if !checkpipelin {
		return
	}
}

// start servers
func Start() {

	regisger_ipeline()

	if RunTimeConfig.UserSession {
		// UserSession
		session := &pipelineSession{}
		Commonpipeline.AddFirst(PIPELINE_SESSION, session)
	}
	DogoLog.Infof("Start Dogo in the port:%s", RunTimeConfig.Port)
	DoGo.start()
}

func regisger_ipeline() {
	// 添加日志记录
	request_log := &PipelineLog{}
	Commonpipeline.AddFirst(PIPELINE_LOG, request_log)

	// 添加路由解析
	prouter := &pipelineRouter{}
	Commonpipeline.AddLast(PIPELINE_ROUTER, prouter)

	// 将数据刷新到浏览器
	finishRequest := &pipelineFinishRequest{}
	Commonpipeline.AddLast(PIPELINE_FINISH_REQUEST, finishRequest)
}

func (t *dogo) start() {

	// 注册静态资源请求路径
	http.HandleFunc(RunTimeConfig.staticRequstPath, serverFileController)

	Router("/favicon.ico", faviconIcoController)

	http.HandleFunc("/", t.handler)

	http.ListenAndServe(":"+RunTimeConfig.Port, nil)
}

func init() {
	DoGo = &dogo{serveMux: http.NewServeMux()}
}
