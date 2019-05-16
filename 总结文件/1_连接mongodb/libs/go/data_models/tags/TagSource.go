package tags

type TagSource string

const TagSourceFromIntelligence TagSource = "intelligence"

func TagSourceFromPTD(ptd_tag string) TagSource {
	return TagSource("/source/ptd/tags/" + ptd_tag)
}
