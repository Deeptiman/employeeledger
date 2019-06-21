package model

type UserData struct {

  ID	     string `json:"id"`
  Name	     string `json:"name"`
  Email	     string `json:"email"`
  Company    string `json:"company"`
  Occupation string `json:"occupation"`
  Salary     string `json:"salary"`
  UserType   string `json:"userType"`
}
