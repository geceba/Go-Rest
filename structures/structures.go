package structures

type User struct {
  Id int `json:"id"`
  Username string `json:"username"`
  First_Name string `json:"first_name"`
  Last_Name string `json:"last_name"`
}

type Response struct {
  Status string `json:"Status"`
  Data User `json:"data"`
  Message string `json:"message"`
}
