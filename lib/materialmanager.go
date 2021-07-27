package lib

type Contract int

const (
	SignTypeNew    Contract = iota + 1 //新签
	SignTypeReNew                      //续签
	SignTypeChange                     //变更
)

func (s Contract) SignTypeName() (name string) {
	switch s {
	case SignTypeNew:
		name = "新签"
	case SignTypeReNew:
		name = "续签"
	case SignTypeChange:
		name = "变更"
	default:
		name = "未知"
	}
	return
}

const (
	YtTypeCityService Contract = iota + 1 //城市服务
	YtTypeNotHouse                        //存量非住
	YtTypeHouse                           //存量住宅
	YtTypeNewNotHouse                     //新建非住
	YtTypeNewHouse                        //新建住宅
)

func (s Contract) YtName() string {
	var (
		name string
	)
	switch s {
	case YtTypeCityService:
		name = "城市服务"
	case YtTypeNotHouse:
		name = "存量非住"
	case YtTypeHouse:
		name = "存量住宅"
	case YtTypeNewNotHouse:
		name = "城市服务"
	case YtTypeNewHouse:
		name = "新建住宅"
	default:
		name = "未知"
	}
	return name
}

const (
	YtDetailTypeMuliHouse Contract = iota + 1 //多层住宅
	YtDetailTypeHighHouse                     //高层住宅
	YtDetailTypeVilla                         //别墅
	YtDetalTypeBGProperty                     //办公物业
	YtDetalTypeSYProperty                     //商业物业
)

func (s Contract) YtDetailName() string {
	var (
		name string
	)
	switch s {
	case YtDetailTypeMuliHouse:
		name = "多层住宅"
	case YtDetailTypeHighHouse:
		name = "高层住宅"
	case YtDetailTypeVilla:
		name = "别墅"
	case YtDetalTypeBGProperty:
		name = "办公物业"
	case YtDetalTypeSYProperty:
		name = "商业物业"
	default:
		name = "未知"
	}
	return name
}

const (
	ContractTypePWYService Contract = iota + 1 //前期物业合同
	ContractTypeWYService                      //物业服务合同
	ContractTypeACService                      //案场服务合同
	ContractTypeGWService                      //顾问服务合同
	ContractTypeZXService                      //专项服务合同
)

func (s Contract) ContractTypeName() string {
	var (
		name string
	)
	switch s {
	case ContractTypePWYService:
		name = "前期物业合同"
	case ContractTypeWYService:
		name = "物业服务合同"
	case ContractTypeACService:
		name = "案场服务合同"
	case ContractTypeGWService:
		name = "顾问服务合同"
	case ContractTypeZXService:
		name = "专项服务合同"
	default:
		name = "未知"
	}
	return name
}

const (
	ContractServiceTypeQW   Contract = iota + 1 //全委
	ContractServiceTypeCity                     //城市
	ContractServiceTypeZX                       //专项
)

// ServiceTypeName 合同服务
func (s Contract) ServiceTypeName() string {
	var (
		name string
	)
	switch s {
	case ContractServiceTypeQW:
		name = "全委"
	case ContractServiceTypeCity:
		name = "城市"
	case ContractServiceTypeZX:
		name = "专项"
	default:
		name = "未知"
	}
	return name
}

const (
	ContractStatusTypeNLvYue  Contract = iota + 1 //待履约
	ContractStatusTypeNXueYue                     //待续约
	ContractStatusTypeLvYue                       //履约中
	ContractStatusTypeXuYue                       //已续约
	ContractStatusTypeStop                        //已终止
)

// ConstractStatusName 合同状态
func (s Contract) ConstractStatusName() string {
	var (
		name string
	)
	switch s {
	case ContractStatusTypeNLvYue:
		name = "待履约"
	case ContractStatusTypeNXueYue:
		name = "待续约"
	case ContractStatusTypeLvYue:
		name = "履约中"
	case ContractStatusTypeXuYue:
		name = "已续约"
	case ContractStatusTypeStop:
		name = "已终止"
	default:
		name = "未知"
	}
	return name
}

const (
	CustomerTypeGQ   Contract = iota + 1 //国企
	CustomerTypeMQ                       //民企
	CustomerTypeSYDW                     //事业单位
	CustomerTypeZFJG                     //政府机关
)

// CustomerTypeName 客户类型
func (s Contract) CustomerTypeName() string {
	var (
		name string
	)
	switch s {
	case CustomerTypeGQ:
		name = "国企"
	case CustomerTypeMQ:
		name = "民企"
	case CustomerTypeSYDW:
		name = "事业单位"
	case CustomerTypeZFJG:
		name = "政府机关"
	default:
		name = "未知"
	}
	return name
}

const (
	ChargeTypeBGZ Contract = iota + 1
)

// ChargeTypeName 合同收费类型
func (s Contract) ChargeTypeName() string {
	var (
		name string
	)
	switch s {
	case ChargeTypeBGZ:
		name = "包干制"
	default:
		name = "未知"
	}
	return name
}

const (
	SignCompanyTypeXG  Contract = iota + 1 //收购
	SignCompanyTypeHZ                      //合资
	SignCompanyTypeZC                      //自行注册
	SignCompanyType55                      //555旗下
	SignCompanyTypeSub                     //子公司
)

// SignCompanyTypeName 签约公司类型
func (s Contract) SignCompanyTypeName() string {
	var (
		name string
	)
	switch s {
	case SignCompanyTypeXG:
		name = "收购"
	case SignCompanyTypeHZ:
		name = "合资"
	case SignCompanyTypeZC:
		name = "自行注册"
	case SignCompanyType55:
		name = "555旗下"
	case SignCompanyTypeSub:
		name = "子公司"
	default:
		name = "未知"
	}
	return name
}

const (
	ContractAbnormalTypeAZ   = iota + 1 //甲方主观原因
	ContractAbnormalTypeAK              //甲方客观原因
	ContractAbnormalTypeYYTC            //运营退出
	ContractAbnormalTypNJC              //未运营
)

// ContractAbnormalTypeName 合同异常类型
func (s Contract) ContractAbnormalTypeName() string {
	var (
		name string
	)
	switch s {
	case ContractAbnormalTypeAZ:
		name = "甲方主观原因"
	case ContractAbnormalTypeAK:
		name = "甲方客观原因"
	case ContractAbnormalTypeYYTC:
		name = "运营退出"
	case ContractAbnormalTypNJC:
		name = "未运营"
	default:
		name = "未知"
	}
	return name
}
