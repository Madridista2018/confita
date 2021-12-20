package object

import (
	"github.com/casbin/confita/casdoor"
	"github.com/casdoor/casdoor-go-sdk/auth"
	"xorm.io/core"
)

type Reviewer struct {
	Owner           string    `xorm:"varchar(100) notnull pk" json:"owner"`
	Name            string    `xorm:"varchar(100) notnull pk" json:"name"`
	FirstName       string    `xorm:"varchar(100)" json:"firstName"`
	LastName        string    `xorm:"varchar(100)" json:"lastName"`
	SubmittionId    []*string `xorm:"varchar(1000)" json:"submittionId"` //稿件id
	TopicId         []*string `xorm:"varchar(1000)" json:"topicId"`      //会议子方向id
	Title           string    `xorm:"varchar(100)" json:"title"`         //审稿人职称
	VarifyCode      string    `xorm:"varchar(100)" json:"varifyCode"`
	CreatedTime     string    `xorm:"varchar(100)" json:"createdTime"`
	IsNotified      bool      `json:"isNotified"`
	IsCorresponding bool      `json:"isCorresponding"`
}

func GetReviewers(owner string) []*Reviewer {
	var reviewers []*Reviewer
	err := adapter.engine.Desc("created_time").Find(&reviewers, &Reviewer{Owner: owner})
	if err != nil {
		panic(err)
	}

	return reviewers
}
func GetReviewer(id string) *auth.User {
	user := casdoor.GetUserById(id)
	if user == nil || user.Tag != "reviewer" {
		return nil
	}
	return user
}

// func GetReviewers() []*auth.User {
// 	users := casdoor.GetUsers()
// 	return users
// }

func GetMemberEmailReminder(id string) (bool, string) {
	user := casdoor.GetUserById(id)
	if user == nil {
		return false, ""
	}

	return true, user.Email
}
func UpdateReviewer(owner string, reviewer *Reviewer) bool {
	// owner, name := util.GetOwnerAndNameFromId(reviewerId)
	// if getSubmission(owner, name) == nil {
	// 	return false
	// }

	// _, err := adapter.engine.ID(core.PK{owner, name}).AllCols().Update(submission)
	// if err != nil {
	// 	panic(err)
	// }

	// //return affected != 0
	return true
}

func AddReviewer(reviewer *Reviewer) bool {
	affected, err := adapter.engine.Insert(reviewer)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func DeleteReviewer(reviewer *Reviewer) bool {
	affected, err := adapter.engine.ID(core.PK{reviewer.Owner}).Delete(&Reviewer{})
	if err != nil {
		panic(err)
	}

	return affected != 0
}
