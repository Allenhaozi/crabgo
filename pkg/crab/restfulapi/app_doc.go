package restfulapi

import (
	"net/http"
	"reflect"
	"sync"

	crabcorev1 "github.com/allenhaozi/crabgo/pkg/apis/core/v1"

	"github.com/allenhaozi/crabgo/pkg/crab/manager"

	"github.com/allenhaozi/crabgo/pkg/apis/common"
	"github.com/allenhaozi/crabgo/pkg/register"
)

type appDocApi struct {
	ApiController
	*manager.AppDocManager
}

func SetupAppDoc(cfg *register.Config, carrier *sync.Map) (string, []common.RestfulApiMeta) {
	api := NewAppDocApi(cfg)
	k := reflect.TypeOf(api).Name()
	carrier.Store(k, api)
	return k, []common.RestfulApiMeta{
		common.RestfulApiMeta{}.Gen(k, http.MethodGet, "GetAppDoc", "/app/doc/:appId"),
		common.RestfulApiMeta{}.Gen(k, http.MethodPost, "CreateAppDoc", "/app/doc"),
	}
}

func NewAppDocApi(cfg *register.Config) *appDocApi {
	as := &appDocApi{}
	as.initial(cfg)
	as.AppDocManager = manager.NewAppDocManager(cfg)
	return as
}

func (as *appDocApi) GetAppDocAction(ctx register.Context) register.CrabResponseIf {
	appId := ctx.Param("appId")
	resp := as.GetAppDoc(ctx, appId)
	return resp
}

func (as *appDocApi) CreateAppDocAction(ctx register.Context) register.CrabResponseIf {
	req := &crabcorev1.AppDoc{}
	ctx.Bind(req)
	resp := as.CreateAppDoc(ctx, req)
	return resp
}
