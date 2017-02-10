package config

var default_conf = `
###########
# 默认配置 #
###########

# 端口
LISTEN_PORT=8080

# IP 一般为空
LISTEN_IP=

# ajax请求返回的数据格式
# 支持json,xml,auto
# 如果值为auto,系统将根据url后缀智能选择返回类型
AJAX_RETURN_FORMAT=json

# 服务名称
# 系统在返回header时候，会把这个值添加到Server上
SERVER_NAME=DoGoServerv1

# 静态资源路径
STATIC_REQUST_PATH=/static/

# 表态资源保存目录
# 注意：
# 实际在项目中的存放目录： STATIC_ROOT_PATH + STATIC_REQUST_PATH
STATIC_ROOT_PATH=./


##########################
[SESSION]

NAME=DOGO_SESSION_ID
#session名称
NAME=DogoSessionID

[SESSION.FILE_STORE]
ROOT_DIR=./session_store

[LOG]
DATA_CHAN_SIZE = 0
# 日志记录级别
# 目前仅支持 info 和 debug 两种级别
LEVEL = debug

[TEMPLATE]
VIEW_DIR=./view
`
