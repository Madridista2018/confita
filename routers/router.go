package routers

import (
	"github.com/astaxie/beego"

	"github.com/casbin/confita/controllers"
)

func init() {
	initAPI()
}

func initAPI() {
	ns :=
		beego.NewNamespace("/api",
			beego.NSInclude(
				&controllers.ApiController{},
			),
		)
	beego.AddNamespace(ns)

	beego.Router("/api/signin", &controllers.ApiController{}, "POST:Signin")
	beego.Router("/api/signout", &controllers.ApiController{}, "POST:Signout")
	beego.Router("/api/get-account", &controllers.ApiController{}, "GET:GetAccount")

	beego.Router("/api/get-global-conferences", &controllers.ApiController{}, "GET:GetGlobalConferences")
	beego.Router("/api/get-conferences", &controllers.ApiController{}, "GET:GetConferences")
	beego.Router("/api/get-conference", &controllers.ApiController{}, "GET:GetConference")
	beego.Router("/api/update-conference", &controllers.ApiController{}, "POST:UpdateConference")
	beego.Router("/api/add-conference", &controllers.ApiController{}, "POST:AddConference")
	beego.Router("/api/delete-conference", &controllers.ApiController{}, "POST:DeleteConference")

	beego.Router("/api/get-submissions", &controllers.ApiController{}, "GET:GetSubmissions")
	beego.Router("/api/get-submission", &controllers.ApiController{}, "GET:GetSubmission")
	beego.Router("/api/update-submission", &controllers.ApiController{}, "POST:UpdateSubmission")
	beego.Router("/api/add-submission", &controllers.ApiController{}, "POST:AddSubmission")
	beego.Router("/api/delete-submission", &controllers.ApiController{}, "POST:DeleteSubmission")
	beego.Router("/api/upload-submission-file", &controllers.ApiController{}, "POST:UploadSubmissionFile")

	beego.Router("/api/get-reviewer", &controllers.ApiController{}, "GET:GetReviewer")
	beego.Router("/api/send-signup-mail", &controllers.ApiController{}, "GET:SendSignupMail")
	beego.Router("/api/review-paper", &controllers.ApiController{}, "GET:ReviewPaper")

}
