package users

import (
		"community-go/utility"
	"community-go/model"
)

func GetUser(userID int) (model.Users, error) {

	users := model.Users{}
  	row := utility.DB.QueryRow("SELECT name,nail FROM user where userID = ?",userID)
  	//defer utility.DB.Close()
	row.Scan(&users.AuthorName, &users.AuthorIcon)
  	users.AuthorID = userID

  	return users, nil
	//if err != nil {
	//	fmt.Printf("rows === %p", rows)
	//	fmt.Println("查询错误")
	//	return Users{}, err
	//} else {
	//	users := Users{}
	//	for rows.Next() {
	//		rows.Scan(&users.AuthorID, &users.AuthorName, &users.AuthorIcon)
	//		err = nil
	//	}
	//	fmt.Println("users:", users)
	//	fmt.Println("查询正确")
	//	return users, err
	//}
}