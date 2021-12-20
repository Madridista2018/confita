package object

type ReviewerAttrVal struct {
	Fid             int    `xorm:"int(10) notnull pk autoincr" json:"Fid"`
	FSubmissionName string `xorm:"varchar(100) notnull" json:"FSubmissionName"`
	FAttrId         int    `xorm:"int(100) notnull" json:"FAttrId"`
	FAttrVal        string `xorm:"varchar(1000)" json:"FAttrVal"`
	FReviewerName   string `xorm:"varchar(100)" json:"FReviewerName"`
}

func GetReviewerAttrVal(id int) []*ReviewerAttrVal {
	reviewerAttrVal := []*ReviewerAttrVal{}
	err := adapter.engine.Find(&reviewerAttrVal, &ReviewerAttrVal{Fid: id})
	if err != nil {
		panic(err)
	}

	return reviewerAttrVal
}

func GetReviewerAttrVals() []*ReviewerAttrVal {
	reviewerAttrVals := []*ReviewerAttrVal{}
	err := adapter.engine.Find(&reviewerAttrVals, &ReviewerAttrVal{})
	if err != nil {
		panic(err)
	}

	return reviewerAttrVals
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
func AddReviewerAttrVal(reviewerAttrVal *ReviewerAttrVal) bool {
	affected, err := adapter.engine.Insert(reviewerAttrVal)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

//func DeleteReviewer(reviewer *Reviewer) bool {
//	affected, err := adapter.engine.ID(core.PK{reviewer.Owner}).Delete(&Reviewer{})
//	if err != nil {
//		panic(err)
//	}
//
//	return affected != 0
//}
