package responseutil

import "net/http"

// 通用异常
var (
	BadRequest              = 400000000
	Unauthorized            = 401000000
	Forbidden               = 403000000
	NotFound                = 404000000
	Conflict                = 409000000
	InternalServerError     = 500000000
	StatusExpectationFailed = 417000000
)

var (
	// Conncfg sap链接配置模块
	Conncfg                        = 100000
	ConncfgBadRequest              = BadRequest + Conncfg
	ConncfgUnauthorized            = Unauthorized + Conncfg
	ConncfgForbidden               = Forbidden + Conncfg
	ConncfgNotFound                = NotFound + Conncfg
	ConncfgConflict                = Conflict + Conncfg
	ConncfgInternalServer          = InternalServerError + Conncfg
	ConncfgStatusExpectationFailed = StatusExpectationFailed + Conncfg

	// scene 场景模块
	Scene               = 200000
	SceneBadRequest     = BadRequest + Scene
	SceneUnauthorized   = Unauthorized + Scene
	SceneForbidden      = Forbidden + Scene
	SceneNotFound       = NotFound + Scene
	SceneConflict       = Conflict + Scene
	SceneInternalServer = InternalServerError + Scene

	// meta 元数据
	Meta               = 300000
	MetaBadRequest     = BadRequest + Meta
	MetaUnauthorized   = Unauthorized + Meta
	MetaForbidden      = Forbidden + Meta
	MetaNotFound       = NotFound + Meta
	MetaConflict       = Conflict + Meta
	MetaInternalServer = InternalServerError + Meta
)

var (
	// commmon
	CommInternalServer = NewCode(http.StatusInternalServerError, InternalServerError, "Internal Server Error")
	CommNotFound       = NewCode(http.StatusNotFound, NotFound, "Record does not exist")
	CommUnauthorized   = NewCode(http.StatusUnauthorized, Unauthorized, "Unauthorized")
	CommBadRequest     = NewCode(http.StatusBadRequest, Unauthorized, "Invalid parameter")
	ComConflictRequest = NewCode(http.StatusConflict, Conflict, "The request is conflicts")

	// cfg
	CodeCfgBadRequest          = NewCode(http.StatusBadRequest, ConncfgBadRequest, "Invalid parameter")
	CodeCfgUnauthorized        = NewCode(http.StatusUnauthorized, ConncfgUnauthorized, "Unauthorized")
	CodeCfgForbidden           = NewCode(http.StatusForbidden, ConncfgForbidden, "Forbidden")
	CodeCfgNotFound            = NewCode(http.StatusNotFound, ConncfgNotFound, "Record does not exist")
	CodeCfgConflict            = NewCode(http.StatusConflict, ConncfgConflict, "The request is conflicts")
	CodeCfgInternalServer      = NewCode(http.StatusInternalServerError, ConncfgInternalServer, "Internal Server Error")
	CodeCfgCheckInternalServer = NewCode(http.StatusExpectationFailed, ConncfgStatusExpectationFailed, "Check sap Invalid parameter")

	// scene/sap
	CodeSceneBadRequest     = NewCode(http.StatusBadRequest, SceneBadRequest, "Invalid parameter")
	CodeSceneUnauthorized   = NewCode(http.StatusUnauthorized, SceneUnauthorized, "Unauthorized")
	CodeSceneForbidden      = NewCode(http.StatusForbidden, SceneForbidden, "Forbidden")
	CodeSceneNotFound       = NewCode(http.StatusNotFound, SceneNotFound, "Record does not exist")
	CodeSceneConflict       = NewCode(http.StatusConflict, SceneConflict, "The request is conflicts")
	CodeSceneInternalServer = NewCode(http.StatusInternalServerError, SceneInternalServer, "Internal Server Error")

	// meta
	CodeMetaBadRequest     = NewCode(http.StatusBadRequest, MetaBadRequest, "Invalid parameter")
	CodeMetanauthorized    = NewCode(http.StatusUnauthorized, MetaUnauthorized, "Unauthorized")
	CodeMetaForbidden      = NewCode(http.StatusForbidden, MetaForbidden, "Forbidden")
	CodeMetaNotFound       = NewCode(http.StatusNotFound, MetaNotFound, "Record does not exist")
	CodeMetaConflict       = NewCode(http.StatusConflict, MetaConflict, "The request is conflicts")
	CodeMetaInternalServer = NewCode(http.StatusInternalServerError, MetaInternalServer, "Internal Server Error")
)
