### 开发目录结构
(按server目录下的一级目录进行逐个说明)
##### app目录
| 目录&文件 | 说明 | 详情 |
| ----- | ----- | ----- |
| app(父级目录)  | app应用服 | app应用相关服务端接口及逻辑都在此目录内 |
| config | 配置 | 配置文件及配置初始化（配置文件说明：_dev.yaml表示开发服、 _test.yaml表示测试服、 _prod.yaml表示线上服） |
| controller | 控制器 | 调度层 业务逻辑、增删改查 及 事务处理，根据不同的模块划分不同的子目录 例如：用户模块为cuser 点播模块为cvideo |
| docs | 接口文档配置 | swag生成的接口文档配置 |
| log | 程序运行生成的日志文件 | 当前进程运行生成的日志文件 |
| routers | 路由层 | 分模块、版本划分出接口路由 例如：用户模块为 routers/api/v1/user目录，v1表示版本|
| static | 静态文件目录 | 放置一些静态文件或html模版等等（暂未使用到）子目录 应划分为 css、images、js、plugins、templates|
| main | 程序的初始化和执行 | 包含对配置、日志、mysql、redis、性能监控、运行模式、路由等进行初始化 及 web服务启动 |


##### backend目录
| 目录&文件 | 说明 | 详情 |
| ----- | ----- | ----- |
| backend(父级目录)  | 管理后台后端服务 | 管理后台相关服务端接口及逻辑都在此目录内 |
| config | 配置 | 配置文件及配置初始化（配置文件说明：_dev.yaml表示开发服、 _test.yaml表示测试服、 _prod.yaml表示线上服） |
| controller | 控制器 | 调度层 业务逻辑、增删改查 及 事务处理，根据不同的模块划分不同的子目录 例如：用户模块为cuser 点播模块为cvideo |
| docs | 接口文档配置 | swag生成的接口文档配置 |
| log | 程序运行生成的日志文件 | 当前进程运行生成的日志文件 |
| routers | 路由层 | 分模块、版本划分出接口路由 例如：用户模块为 routers/api/v1/user目录，v1表示版本|
| static | 静态文件目录 | 放置一些静态文件或html模版等等（暂未使用到）子目录 应划分为 css、images、js、plugins、templates|
| main | 程序的初始化和执行 | 包含对配置、日志、mysql、redis、性能监控、运行模式、路由等进行初始化 及 web服务启动 |


##### barrage目录
| 目录&文件 | 说明 | 详情 |
| ----- | ----- | ----- |
| barrage(父级目录)  | 弹幕服务 | 采用websocket实现的实时弹幕 |
| config | 配置 | 配置文件及配置初始化（配置文件说明：_dev.yaml表示开发服、 _test.yaml表示测试服、 _prod.yaml表示线上服） |
| log | 程序运行生成的日志文件 | 当前进程运行生成的日志文件 |
| client | 模拟客户端 | 模拟客户端连接（发送、接收弹幕） |
| server | 模拟服务端 | 模拟服务端发送弹幕 |
| main | 程序的初始化和执行 | 包含对配置、日志、mysql、redis、性能监控、运行模式、路由等进行初始化 及 ws服务启动 |


##### dao目录
| 目录&文件 | 说明 | 详情 |
| ----- | ----- | ----- |
| dao(父级目录)  | 数据访问接口及对象 | 针对mysql、redis等数据库，全局的访问对象 |


##### global目录
| 目录&文件 | 说明 | 详情 |
| ----- | ----- | ----- |
| global(父级目录) | 全局公共目录 | 整个开发目录下的公用目录，包含每个服务的第三方日志接口、错误码定义 、接口响应封装、 公用的系统常量、业务模块常量 以及redis key的定义 |
| app | app应用服相关 | app应用服相关的错误码定义 及 接口响应封装 |
| errdef | 错误码 | 错误码定义 及 接口响应封装 |
| log | 第三方日志接口 | 第三方日志接口全局对象（可对不同的日志库实例化） |
| consts | 常量定义 | 包含系统常量及业务模块常量声明（文件名区分 ），例如：env开头的文件声明系统常量 、user开头的文件声明用户模块常量 |
| rdskey | redis key定义 | 所有业务及系统相关的redis key定义 |


##### job目录
| 目录&文件 | 说明 | 详情 |
| ----- | ----- | ----- |
| job(父级目录)  | 定时任务 | 异步执行定时任务 |


##### log目录
| 目录&文件 | 说明 | 详情 |
| ----- | ----- | ----- |
| log | 第三方日志 | 第三方日志封装 [zap]（interface）|


##### middleware目录
| 目录&文件 | 说明 | 详情 |
| ----- | ----- | ----- |
| middleware(父级目录) | 中间件 | 包含日志中间件、跨域处理、接口签名中间件、token校验中间件 以及部分中间件的初始化 |
| engineLog | 日志中间件 | 接口耗时、用户代理、访问的接口地址、请求参数、请求ip等输出到日志文件 |
| header | 跨域处理 | 声明请求头信息 跨域处理 |
| sign | 签名中间件 | 校验客户端请求接口时 所携带的签名是否合法（对称加密） |
| token | token中间件 | 校验用户token是否合法 |
| init | 中间件初始化 | 部分需初始化的中间件在init文件中进行处理 |


##### models目录
| 目录&文件 | 说明 | 详情 |
| ----- | ----- | ----- |
| models(父级目录) | 数据模型及数据库curd | 包含数据模型、数据库增删改查操作  数据模型使用xorm生成，数据库操作按业务模块划分 例如：用户模块 划分文件夹为muser |
| modelTpl | 指令生成数据模型的配置及脚本 | 指令生成数据模型的配置及脚本 |
| pprof | 性能监控 | 性能监控、堆栈信息及golang相关信息 |
| sql | 建表sql | 业务相关建表sql |


##### tools目录
| 目录&文件 | 说明 | 详情 |
| ----- | ----- | ----- |
| tools(父级目录) | 工具目录 | 包含第三方登录封装(微信、微博、QQ)、以及MobTech 一键登陆服务端接入 |
| mobTech | mob开发者服务平台服务端接入 | 接入了mob平台free login |
| thirdLogin | 第三方登陆 | 第三方登录封装(微信、微博、QQ) |
| filter | 脏词过滤 | 封装脏词过滤 |
| nsq | 消息队列 | 高效、轻量的消息队列 |
| tencentCloud | 腾讯云 | 调用腾讯云sdk 包含但不限于 点播、文本检测等 |


##### util目录
| 目录&文件 | 说明 | 详情 |
| ----- | ----- | ----- |
| util(父级目录) | 公共方法目录 | 公共方法封装 按功能划分文件 便于使用 |


##### web目录
| 目录&文件 | 说明 | 详情 |
| ----- | ----- | ----- |
| web(父级目录) | 管理后台前端 | vue + element ui |
