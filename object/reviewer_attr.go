package object

type ReviewerAttr struct {
	Fid         int    `xorm:"int(10) notnull pk autoincr" json:"id"`
	FTitle      string `xorm:"varchar(100) notnull" json:"fTitle"`
	FType       string `xorm:"varchar(100) notnull" json:"fType"`
	FDefaultVal string `xorm:"varchar(100)" json:"fDefaultVal"`
	FVal        string `xorm:"varchar(1000)" json:"fVal"`
	FIsNeed     int    `xorm:"tinyint(1)" json:"fIsNeed"`
	FTips       string `xorm:"varchar(100)" json:"fTips"`
	FOrder      int    `xorm:"int(5)" json:"fOrder"`
	FName       string `xorm:"varchar(100)" json:"fName"`
}

func GetReviewerAttr(id int) []*ReviewerAttr {
	reviewerAttr := []*ReviewerAttr{}
	err := adapter.engine.Find(&reviewerAttr, &ReviewerAttr{Fid: id})
	if err != nil {
		panic(err)
	}

	return reviewerAttr
}

func GetReviewerAttrs() []*ReviewerAttr {
	reviewerAttrs := []*ReviewerAttr{}
	err := adapter.engine.Desc("f_order").Find(&reviewerAttrs, &ReviewerAttr{})
	if err != nil {
		panic(err)
	}

	return reviewerAttrs
}

//// func GetReviewers() []*auth.User {
//// 	users := casdoor.GetUsers()
//// 	return users
//// }
//
//func GetMemberEmailReminder(id string) (bool, string) {
//	user := casdoor.GetUserById(id)
//	if user == nil {
//		return false, ""
//	}
//
//	return true, user.Email
//}
//func UpdateReviewer(owner string, reviewer *Reviewer) bool {
//	// owner, name := util.GetOwnerAndNameFromId(reviewerId)
//	// if getSubmission(owner, name) == nil {
//	// 	return false
//	// }
//
//	// _, err := adapter.engine.ID(core.PK{owner, name}).AllCols().Update(submission)
//	// if err != nil {
//	// 	panic(err)
//	// }
//
//	// //return affected != 0
//	return true
//}
//
//func AddReviewer(reviewer *Reviewer) bool {
//	affected, err := adapter.engine.Insert(reviewer)
//	if err != nil {
//		panic(err)
//	}
//
//	return affected != 0
//}
//
//func DeleteReviewer(reviewer *Reviewer) bool {
//	affected, err := adapter.engine.ID(core.PK{reviewer.Owner}).Delete(&Reviewer{})
//	if err != nil {
//		panic(err)
//	}
//
//	return affected != 0
//}
