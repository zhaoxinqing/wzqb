#业务模块保存日志功能使用方法：
(可以参照log/下使用示例)

##1.添加中间件
修改api文件，在需要生成业务日志的接口上方配置中间件，如下：
    @server (
        middleware: LogAdd
    ）
使用goctl api生成代码
将lib/middleware/logaddmiddleware.go中的代码复制到当前middleware/下生成的logaddmiddleware.go中，
在其中27行，配置该模块功能名称（27： Name: "日志列表",）

##2.配置中间件
在当前模块svc/servicecontext.go中添加中间件，如下：

    导入lib.model
    import libModel "gemdale-server/lib/model"

    type ServiceContext struct {
        Config   config.Config
        DBEngine *gorm.DB
        Middle   *middleware.LogAddMiddleware // 日志
        LogAdd   rest.Middleware // 日志
        UserRpc 	userclient.User // 用户
    }
    libModel.NewDbEnginLog(db)
    func NewServiceContext(c config.Config) *ServiceContext {
        ......
    
        middle := middleware.NewLogAddMiddleware() // 日志
        model.NewDbEnginLog(db)
        return &ServiceContext{
            svcCtx = &ServiceContext{
            Config:   c,
            DBEngine: db,
            Middle:   middle, // 日志
            LogAdd:   middle.Handle, // 日志
            UserRpc: userclient.NewUser(zrpc.MustNewClient(c.User)), // 用户
        }
    }
由于日志需要记录当前用户，调用了UserRpc，也需加入UserRpc,已加请忽略
materialmanager-api.yaml

    User:
        Etcd:
            Hosts:
                - 127.0.0.1:2379
            Key: user.rpc
internal/config/config.go

    User    zrpc.RpcClientConf
##3.配置handler（添加//旁代码）

    func LogListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
        var req types.LogListReq
        var log = ctx.Middle.BussLog // 获得日志model
        // 获取用户信息
        token := lib.GetHeaderToken(r)
        userinfo, _ := ctx.UserRpc.UserInfo(r.Context(), &userclient.UserInfoReq{
        Token: token,
        })
        userinfo, _ := ctx.UserRpc.UserInfo(r.Context(), &userclient.UserInfoReq{
        Token: token,
        })
        if err := httpx.Parse(r, &req); err != nil {
        xhttp.ParamErrorResult(r, w, err)
         // 传入日志参数并保存
        log.BuildParamAndSava(req, 400, userinfo.UserName, userinfo.UserGuid)
        return
        }
        l := logic.NewLogListLogic(r.Context(), ctx)
        infos, err := l.LogList(req)
        code := xhttp.HttpResult(r, w, infos, err)
        // 传入日志参数并保存
        log.BuildParamAndSava(req, code, userinfo.UserName, userinfo.UserGuid)
        }
    }

Done