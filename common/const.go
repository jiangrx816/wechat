package common

const DEFAULT_PAGE_SIZE = 12

const (
	// 定义最终生成的nc文件的目录
	NC_DIR = "file"

	// 定义生成站点ID的nc文件的临时目录
	STATION_TEMP_DIR = "station_temp"

	// 定义预览时的临时目录
	TMP_DIR = "tmp"
)

const (
	// 多个站点合并成一个nc文件命名规则
	STATION_IDS_NC_NAME = "station_ids.nc"

	// 区域降水nc文件命名规则
	TP_NC_NAME = "tp.nc"

	// 多个站点的json文件命名规则
	STATION_IDS_JSON_NAME = "station_ids.json"

	// 多个降水对应的json文件命名规则
	TP_IDS_JSON_NAME = "tp_ids.json"

	// 多个降水合并成一个nc文件命名规则
	TP_IDS_NC_NAME = "tp_ids.nc"

	// 拓扑类型的上下游json文件命名规则
	TOPOLOGY_JSON_NAME = "topology.json"
)

const (
	// 定义 .nc 后缀
	Suffix_Extension_Nc = ".nc"

	// 定义.shp 后缀
	Suffix_Extension_Shp = ".shp"
)

// 定义 河道水情表-等时段标准化数据 字段
const (
	Src  = "1"
	Step = 3600
)

// 定义Step 字段类型
const (
	Step_ZZ = "ZZ" //河道
	Step_ZQ = "ZQ" //河道
	Step_DD = "DD" //堰闸
	Step_RR = "RR" //水库
)

const (
	// 没有上传的文件
	MessageNotFileWasUploaded = "No file was uploaded"

	//文件缺失
	MessageFileMissing = "File missing"

	//批次号为空
	MessageBatchNoEmpty = "The batch number is empty"
)
