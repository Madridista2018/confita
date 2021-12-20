package controllers

import (
	"github.com/astaxie/beego"
	"github.com/casbin/confita/object"
	"github.com/casbin/confita/service"
)

func (c *ApiController) GetReviewer() {
	id := c.Input().Get("id")
	reviewer := object.GetReviewer(id)
	if reviewer == nil {
		c.ResponseError("该用户不存在")
	}
	c.ResponseOk(reviewer)
}

func (c *ApiController) SendSignupMail() {
	id := c.Input().Get("id")
	email := object.GetReviewer(id).Email
	if email == "" {
		c.ResponseError("邮箱地址错误")
	}
	title := ""
	content := ""
	domain := beego.AppConfig.String("domain")
	err := service.SendSignupMail(title, content, email, domain)
	if err != nil {
		panic(err)
	}
	c.ResponseOk()
}

func CheckSubmissionReviewer(submissionName string, reviewerName string) bool {

	reviewers := object.GetSubmissionByName(submissionName).Reviewers
	for i := range reviewers {
		if reviewerName == reviewers[i].Name {
			return true
		}
	}
	return false
}

func (c *ApiController) ReviewPaper() {
	if beego.AppConfig.String("AT_review_open") == "0" {
		c.ResponseError("评审功能未开启")
	}
	submissionName := c.Input().Get("submissionName")
	reviewerName := c.GetSessionClaims().Name

	if CheckSubmissionReviewer(submissionName, reviewerName) {
		flag := true
		num := 0
		for _, attr := range object.GetReviewerAttrs() {
			reviewinfo := object.ReviewerAttrVal{}
			if flag == false {
				break
			}
			field := attr.FName
			reviewinfo.FSubmissionName = submissionName
			reviewinfo.FAttrId = attr.Fid
			reviewinfo.FReviewerName = reviewerName
			val := c.Input().Get(field)
			if attr.FIsNeed == 1 {
				if val != "" {
					reviewinfo.FAttrVal = val
				} else {
					flag = false
				}
			} else {
				reviewinfo.FAttrVal = val
			}
			if val != "" {
				object.AddReviewerAttrVal(&reviewinfo)
				num++
			}
		}
		if num > 0 {
			c.ResponseOk(num)
		}
		if !flag {
			c.ResponseError("必选项为填写")
		}

	}
	c.ResponseError("没有权限审阅该文章")
}
