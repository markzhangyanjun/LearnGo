package tags

//标签

type TagDetails interface{}

type Tags struct {
	TagsName string      `json:"tag_name" bson:"tag_name"` 
	Source   []TagSource `json:"source" bson:"source"` 
	Detail   TagDetails  `json:"detail,omitempty" bson:"detail,omitempty"` 
}
