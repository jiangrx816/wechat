package response

// 转换结果响应参数
type NetCDFConvertResult struct {
	BatchId string `json:"batch_id"`
}

// 方案的数据内容规则返回结构体
type PlanRegulationResult struct {
	Message string             `json:"message"`
	Code    int                `json:"code"`
	Data    PlanRegulationData `json:"data"` // 方案的数据内容规则
}

// 方案的数据内容规则
type PlanRegulationData struct {
	// 节点信息-里面包含站点信息
	Nodes []NodeItem `json:"nodes"`
	// 面信息-里面包含相关的多边形区域信息
	Subcs []SubcItem `json:"subcs"`
	// 线关系
	Links []LinkItem `json:"links"`
}

// 方案的数据内容规则——节点信息
type NodeItem struct {
	ID           string                 `json:"id"`
	PlanID       string                 `json:"planId"`
	NodeID       string                 `json:"nodeId"`
	NodeCurve    map[string]interface{} `json:"nodeCurve"`
	IsOutfall    string                 `json:"isOutfall"`
	Bscd         string                 `json:"bscd"`
	Lgtd         float64                `json:"lgtd"`
	Lttd         float64                `json:"lttd"`
	Code         string                 `json:"code"`
	Elev         float64                `json:"elev"`
	Type         string                 `json:"type"`
	MlID         string                 `json:"mlId"`
	MlType       string                 `json:"mlType"`
	MlName       string                 `json:"mlName"`
	Stcd         string                 `json:"stcd"`
	Sttp         string                 `json:"sttp"`
	Stnm         string                 `json:"stnm"`
	IsRegu       string                 `json:"isRegu"`
	IsDivider    string                 `json:"isDivider"`
	IsReplace    string                 `json:"isReplace"`
	SttpType     string                 `json:"sttpType"`
	NodeRelation string                 `json:"nodeRelation"`
	ReguMethod   interface{}            `json:"reguMethod"`
	Drna         float64                `json:"drna"`
}

// NodeRelation 节点关系
type NodeRelation struct {
	Q      string `json:"q"`
	Sttp   string `json:"sttp"`
	ToNext string `json:"to_next"`
}

// 方案的数据内容规则——面信息
type SubcItem struct {
	ID         string  `json:"id"`
	SubcatID   string  `json:"subcatId"`
	SubcatName string  `json:"subcatName"`
	FromNodeID string  `json:"fromNodeId"`
	ToNodeID   string  `json:"toNodeId"`
	Area       float64 `json:"area"`
	Elev       float64 `json:"elev"`
	Slope      float64 `json:"slope"`
	Lgtd1      float64 `json:"lgtd1,string"`
	Lttd1      float64 `json:"lttd1,string"`
	RelID1     string  `json:"relId1"`
	MlID1      string  `json:"mlId1"`
	MlType1    string  `json:"mlType1"`
	MlName1    string  `json:"mlName1"`
	Type1      string  `json:"type1"`
	Drna1      float64 `json:"drna1"`
	Lgtd2      float64 `json:"lgtd2,string"`
	Lttd2      float64 `json:"lttd2,string"`
	RelID2     string  `json:"relId2"`
	MlID2      string  `json:"mlId2"`
	MlType2    string  `json:"mlType2"`
	MlName2    string  `json:"mlName2"`
	Type2      string  `json:"type2"`
	Drna2      float64 `json:"drna2"`
	IsMerge    string  `json:"isMerge"`
	SubcatJson string  `json:"subcatJson"`
}

// 线关系
type LinkItem struct {
	ID         string  `json:"id"`
	RiverID    string  `json:"riverId"`
	Code       string  `json:"code"`
	FromNodeID string  `json:"fromNodeId"`
	ToNodeID   string  `json:"toNodeId"`
	Elev       float64 `json:"elev"`
	Slope      float64 `json:"slope"`
	Lgtd1      string  `json:"lgtd1"`
	Lttd1      string  `json:"lttd1"`
	Elev1      float64 `json:"elev1"`
	RelId1     string  `json:"relId1"`
	MlId1      string  `json:"mlId1"`
	MlType1    string  `json:"mlType1"`
	MlName1    string  `json:"mlName1"`
	Type1      string  `json:"type1"`
	Lgtd2      string  `json:"lgtd2"`
	Lttd2      string  `json:"lttd2"`
	Elev2      float64 `json:"elev2"`
}

// 河道水情表-等时段标准化数据返回数据结构体
// 最外层响应结构
type StandardRiverResponse struct {
	Msg     string                   `json:"msg"`
	Code    interface{}              `json:"code"`
	Data    map[string][]StationData `json:"data"` // 动态站点ID映射
	Success bool                     `json:"success"`
}

// 站点监测数据结构
type StationData struct {
	Q    *float64 `json:"q"`    // 可能为null的字段使用指针类型
	Stcd string   `json:"stcd"` // 站点编码
	Tm   string   `json:"tm"`   // 时间字符串（如需时间类型可改为 time.Time）
	Z    float64  `json:"z"`    // 水位值
}

// 水库水情表-等时段标准化数据返回数据结构体
// 最外层响应结构
type StandardRiverWaterResponse struct {
	Msg     string                        `json:"msg"`
	Code    interface{}                   `json:"code"`
	Data    map[string][]StationWaterData `json:"data"` // 动态站点ID映射
	Success bool                          `json:"success"`
}

// 站点监测数据结构
type StationWaterData struct {
	Rwchrcd *string  `json:"rwchrcd"` // 库水特征码
	Otqty   int      `json:"otqty"`   // 出库流量插值方式
	Blrzty  int      `json:"blrzty"`  // 库下水位插值方式
	W       float64  `json:"w"`       // 蓄水量106m3
	Rz      float64  `json:"rz"`      // 库上水位m --- 使用的值
	Rzty    int      `json:"rzty"`    // 库上水位插值方式
	Tm      string   `json:"tm"`      // 时间字符串（如需时间类型可改为 time.Time）
	Inqty   int      `json:"inqty"`   // 入库流量插值方式
	Blrz    *float64 `json:"blrz"`    // 库下水位m
	Inq     *float64 `json:"inq"`     // 入库流量m3/s
	Otq     *float64 `json:"otq"`     // 出库流量m3/s --- 使用的值
	Wty     int      `json:"wty"`     // 蓄水量插值方式
}

// 堰闸标准化数据返回数据结构体
// 最外层响应结构
type StandardWeirGateResponse struct {
	Msg     string                    `json:"msg"`
	Code    interface{}               `json:"code"`
	Data    map[string][]WeirGateData `json:"data"` // 动态站点ID映射
	Success bool                      `json:"success"`
}

// 堰闸数据结构
type WeirGateData struct {
	Stcd string   `json:"stcd"` // 站ID
	Dwz  *float64 `json:"dwz"`  // 闸下水位
	Tgtq *float64 `json:"tgtq"` // 总过闸流量
	Tm   string   `json:"tm"`   // 时间字符串（如需时间类型可改为 time.Time）
	Upz  float64  `json:"upz"`  // 闸上水位
}

// 查询预报面雨量数据返回数据结构体
// 最外层响应结构
type SurfaceRainfallResponse struct {
	Msg     string         `json:"msg"`
	Code    interface{}    `json:"code"`
	Data    []RainfallItem `json:"data"` // 动态站点ID映射
	Success bool           `json:"success"`
}

// 面雨量数据单个元素结构
type RainfallItem struct {
	Tm string `json:"tm"` // 时间戳
	P  string `json:"p"`  // 参数值
}

// 面雨量处理元素结构
type DealRainfallData struct {
	Tm   string `json:"tm"`   // 时间戳
	Stcd string `json:"stcd"` //站点id
	Tp   string `json:"tp"`   // 参数值
}

// 转换记录列表返回
type NCLogListResponse struct {
	Page  int         `json:"page"`
	Total int64       `json:"total"`
	List  []NCLogData `json:"list"`
}

type NCLogData struct {
	Id             int64  `json:"id"`
	BatchId        string `json:"batch_id"`
	Type           int    `json:"type"`
	DataType       int    `json:"data_type"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	Status         string `json:"status"`
	NcPath         string `json:"nc_path"`
	TpPath         string `json:"tp_path"`
	StationIdsPath string `json:"station_ids_path"`
	Remark         string `json:"remark"`
	FailedReason   string `json:"failed_reason"`
	CreateTime     string `json:"create_time"`
	FinishTime     string `json:"finish_time"`
}

// NC文件预览前上传返回参数
type NetCDFViewResult struct {
	FilePath string `json:"file_path"`
}

// NC文件预览
type NCFilePreview struct {
	Dimensions  map[string]int `json:"dimensions"`
	Coordinates map[string]struct {
		Dtype  string        `json:"dtype"`
		Values []interface{} `json:"values"`
	} `json:"coordinates"`
	Variables map[string]struct {
		Dtype      string            `json:"dtype"`
		Dimensions []string          `json:"dimensions"`
		Attributes map[string]string `json:"attributes"`
		SampleData []interface{}     `json:"sample_data"`
	} `json:"variables"`
	GlobalAttributes map[string]string `json:"global_attributes"`
}
type NCFilePreviewResult struct {
	NetCDFVariableField string        `json:"netcdf_variable_field"`
	NetCDFFileContent   NCFilePreview `json:"netcdf_file_content"`
}

// NC文件预览文件文本内容
type NCDumpResult struct {
	Dimensions  map[string]int `json:"dimensions"`
	Variables   []Variable     `json:"variables"`
	GlobalAttrs []Attribute    `json:"global_attrs"`
	Unlimited   []string       `json:"unlimited"`
}
type Variable struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Dims  []string    `json:"dims"`
	Attrs []Attribute `json:"attrs"`
}
type Attribute struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// stcd 对应 fromNodeId、toNodeId
type StcdToNodeId struct {
	Stcd       string `json:"stcd"`
	FromNodeId string `json:"fromNodeId"`
	ToNodeId   string `json:"toNodeId"`
}
