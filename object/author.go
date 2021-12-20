package object

import (
	"xorm.io/core"
)

type Author struct {
	Owner           string    `xorm:"varchar(100) notnull pk" json:"owner"`
	Name            string    `xorm:"varchar(100) notnull pk" json:"name"`
	FirstName       string    `xorm:"varchar(100)" json:"firstName"`
	LastName        string    `xorm:"varchar(100)" json:"lastName"`
	Affiliation     string    `xorm:"varchar(100)" json:"affiliation"`
	Country         string    `xorm:"varchar(100)" json:"country"`
	SubmittionId    []*string `xorm:"varchar(10000)" json:"submittionId"`
	Telphone        string    `xorm:"varchar(100)" json:"telphone"`
	Email           string    `xorm:"varchar(100)" json:"email"`
	CreatedTime     string    `xorm:"varchar(100)" json:"createdTime"`
	IsNotified      bool      `json:"isNotified"`
	IsCorresponding bool      `json:"isCorresponding"`
}

func GetAuthors(owner string) []*Author {
	authors := []*Author{}
	err := adapter.engine.Desc("created_time").Find(&authors, &Author{Owner: owner})
	if err != nil {
		panic(err)
	}

	return authors
}

// func getAuthor(owner string, name string) *Submission {
// 	submission := Submission{Owner: owner, Name: name}
// 	existed, err := adapter.engine.Get(&submission)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if existed {
// 		return &submission
// 	} else {
// 		return nil
// 	}
// }

// func GetAuthor(id string) *Submission {
// 	owner, name := util.GetOwnerAndNameFromId(id)
// 	return getSubmission(owner, name)
// }

func UpdateAuthor(owner string, author *Author) bool {
	// owner, name := util.GetOwnerAndNameFromId(id)
	// if getSubmission(owner, name) == nil {
	// 	return false
	// }

	// _, err := adapter.engine.ID(core.PK{owner, name}).AllCols().Update(author)
	// if err != nil {
	// 	panic(err)
	// }

	//return affected != 0
	return true
}

func AddAuthor(author *Author) bool {
	affected, err := adapter.engine.Insert(author)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func DeleteAuthor(author *Author) bool {
	affected, err := adapter.engine.ID(core.PK{author.Owner}).Delete(&Author{})
	if err != nil {
		panic(err)
	}

	return affected != 0
}
