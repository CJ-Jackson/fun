package prGenModel

import "fun/commonUtil"

type PrGenModel struct {
	Jira         string
	LinkedPr     string
	RelatedPr    string
	Cms          bool
	Db           bool
	LiveFollowUp bool
	Parameters   bool
	Cloudinary   bool
}

func (m PrGenModel) JiraClean() []string      { return commonUtil.CleanAndSplitString(m.Jira) }
func (m PrGenModel) LinkedPrClean() []string  { return commonUtil.CleanAndSplitString(m.LinkedPr) }
func (m PrGenModel) RelatedPrClean() []string { return commonUtil.CleanAndSplitString(m.RelatedPr) }
