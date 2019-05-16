package evaluate

//威胁量化评分情况

type ThreatEvaluations struct {
	Score        int         `json:"score" bson:"score"` 
	ByDimensions ByDimension `json:"by_dimension" bson:"by_dimension"` 
	ByAbilitys   ByAbility   `json:"by_ability" bson:"by_ability"` 
}

//威胁分维度评分
type ByDimension struct {
	Ability                   int `json:"ability" bson:"ability"` 
	Assets_impacts_directly   int `json:"assets_impacts_directly" bson:"assets_impacts_directly"` 
	Assets_impacts_potentialy int `json:"assets_impacts_potentialy" bson:"assets_impacts_potentialy"` 
	Governance                int `json:"governance" bson:"governance"` 
}

//威胁分能力评分

type ByAbility struct {

	//传播移动
	//对抗逃逸
	//埋伏预制
	//权限获取
	//潜伏隐蔽
	//监控窃取
	//杀伤破坏
	//侦查探测

}
