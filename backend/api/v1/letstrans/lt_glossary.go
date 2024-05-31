package letstrans

import (
	"github.com/firwoodlin/letstrans/model/common/response"
	"github.com/firwoodlin/letstrans/model/letstrans"
	GResponse "github.com/firwoodlin/letstrans/model/letstrans/response"
	"github.com/firwoodlin/letstrans/utils"
	"github.com/gin-gonic/gin"
)

type GlossaryApi struct{}

// CreateGlossary 创建术语库
func (a *GlossaryApi) CreateGlossary(c *gin.Context) {
	var glossary letstrans.Glossary
	if err := c.ShouldBindJSON(&glossary); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	glossary.AuthorID = uid
	if err := glossaryService.CreateGlossary(glossary); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

// GetGlossaryList 获取术语库列表
func (a *GlossaryApi) GetGlossaryList(c *gin.Context) {
	jwtId := utils.GetUserID(c)

	glossaries, err := glossaryService.GetGlossaryByUserID(jwtId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(gin.H{
		"glossaries": glossaries,
	}, c)

}

// GetGlossary 获取单个术语库
func (a *GlossaryApi) GetGlossary(c *gin.Context) {
	glossaryID := utils.Param2Uint(c, "glossary_id")
	if glossaryID == 0 {
		response.FailWithMessage("glossary_id is invalid", c)
		return
	}

	glossary, err := glossaryService.GetGlossaryByID(glossaryID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	terms, err := glossaryService.GetTermsByGlossaryID(glossaryID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(GResponse.GlossaryResponse{
		Glossary: glossary,
		Terms:    terms,
	}, c)
}

// UpdateGlossary 更新术语库
func (a *GlossaryApi) UpdateGlossary(c *gin.Context) {
	var glossary letstrans.Glossary
	if err := c.ShouldBindJSON(&glossary); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	glossaryID := utils.Param2Uint(c, "glossary_id")

	if err := glossaryService.UpdateGlossary(glossaryID, glossary); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

// DeleteGlossary 删除术语库
func (a *GlossaryApi) DeleteGlossary(c *gin.Context) {
	glossaryID := utils.Param2Uint(c, "glossary_id")
	if glossaryID == 0 {
		response.FailWithMessage("glossary_id is invalid", c)
		return
	}

	if err := glossaryService.DeleteGlossary(glossaryID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

// CreateTerm 创建术语
func (a *GlossaryApi) CreateTerm(c *gin.Context) {

	var term letstrans.Term
	if err := c.ShouldBindJSON(&term); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	gId := utils.Param2Uint(c, "glossary_id")
	term.GlossaryID = gId
	if term.GlossaryID == 0 {
		response.FailWithMessage("glossary_id is invalid", c)
		return

	}

	if err := glossaryService.CreateTerm(term); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

// GetTermsByGlossary 获取术语库下的所有术语
func (a *GlossaryApi) GetTermsByGlossary(c *gin.Context) {
	glossaryID := utils.Param2Uint(c, "glossary_id")
	if glossaryID == 0 {
		response.FailWithMessage("glossary_id is invalid", c)
		return
	}

	terms, err := glossaryService.GetTermsByGlossaryID(glossaryID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(terms, c)
}

// UpdateTerm 更新术语
func (a *GlossaryApi) UpdateTerm(c *gin.Context) {
	var term letstrans.Term
	if err := c.ShouldBindJSON(&term); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	termID := utils.Param2Uint(c, "term_id")
	if err := glossaryService.UpdateTerm(termID, term); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

// DeleteTerm 删除术语
func (a *GlossaryApi) DeleteTerm(c *gin.Context) {
	termID := utils.Param2Uint(c, "term_id")
	if termID == 0 {
		response.FailWithMessage("term_id is invalid", c)
		return
	}

	if err := glossaryService.DeleteTerm(termID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}