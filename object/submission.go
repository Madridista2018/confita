package object

import (
	"github.com/casbin/confita/util"
	"xorm.io/core"
)

type AuthorItem struct {
	Name            string `xorm:"varchar(100)" json:"name"`
	Affiliation     string `xorm:"varchar(100)" json:"affiliation"`
	Email           string `xorm:"varchar(100)" json:"email"`
	IsNotified      bool   `json:"isNotified"`
	IsCorresponding bool   `json:"isCorresponding"`
}
type ReviewerItem struct {
	Name            string `xorm:"varchar(100)" json:"name"`
	Affiliation     string `xorm:"varchar(100)" json:"affiliation"`
	Email           string `xorm:"varchar(100)" json:"email"`
	IsNotified      bool   `json:"isNotified"`
	IsCorresponding bool   `json:"isCorresponding"`
}

type Submission struct {
	Owner       string          `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string          `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string          `xorm:"varchar(100)" json:"createdTime"`
	Conference  string          `xorm:"varchar(100)" json:"conference"`
	Title       string          `xorm:"varchar(100)" json:"title"`
	Authors     []*AuthorItem   `xorm:"varchar(10000)" json:"authors"`
	Reviewers   []*ReviewerItem `xorm:"varchar(10000)" json:"reviewers"`
	// Comments    []*string       `xorm:"varchar(1000)" json:"comments"`
	Type        string `xorm:"varchar(100)" json:"type"`
	SubType     string `xorm:"varchar(100)" json:"subType"`
	WordFileUrl string `xorm:"varchar(100)" json:"wordFileUrl"`
	PdfFileUrl  string `xorm:"varchar(100)" json:"pdfFileUrl"`
	Status      string `xorm:"varchar(100)" json:"status"`
}

func GetSubmissions(owner string) []*Submission {
	submissions := []*Submission{}
	err := adapter.engine.Desc("created_time").Find(&submissions, &Submission{Owner: owner})
	if err != nil {
		panic(err)
	}

	return submissions
}

func GetSubmission(owner string, name string) *Submission {
	submission := Submission{Owner: owner, Name: name}
	existed, err := adapter.engine.Get(&submission)
	if err != nil {
		panic(err)
	}

	if existed {
		return &submission
	} else {
		return nil
	}
}

func GetSubmissionByName(name string) *Submission {
	submission := Submission{Name: name}
	existed, err := adapter.engine.Get(&submission)
	if err != nil {
		panic(err)
	}

	if existed {
		return &submission
	} else {
		return nil
	}
}

func UpdateSubmission(id string, submission *Submission) bool {
	owner, name := util.GetOwnerAndNameFromId(id)
	if GetSubmission(owner, name) == nil {
		return false
	}

	_, err := adapter.engine.ID(core.PK{owner, name}).AllCols().Update(submission)
	if err != nil {
		panic(err)
	}

	//return affected != 0
	return true
}

func AddSubmission(submission *Submission) bool {
	affected, err := adapter.engine.Insert(submission)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func DeleteSubmission(submission *Submission) bool {
	affected, err := adapter.engine.ID(core.PK{submission.Owner, submission.Name}).Delete(&Submission{})
	if err != nil {
		panic(err)
	}

	return affected != 0
}
