package ahmember

import (
	"github.com/melvinvoetberg/ah-scanner-api/pkg/ah/client"
	"encoding/json"
)

type Member struct {
	Id int `json:"memberId"`
	Email string `json:"email"`
	Name Name `json:"name"`
}

type Name struct {
	Title string `json:"title"`
	FirstName string `json:"firstName"`
	Initials string `json:"initials"`
	MiddleName string `json:"middleName"`
	SurName string `json:"surName"`
}

func GetMember() Member {
	ac := ahclient.Client()
	resBody := ac.Get("member/v3/member")

	var m Member
  json.Unmarshal([]byte(resBody), &m)

	return m
}
