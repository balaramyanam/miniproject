package utils

import(

	"encoding/json"

	"io/ioutil"

	"net/http"
)


func ParseBody(Res *http.Request,Req interface{}){


	if body,err := ioutil.ReadAll(Res.Body);err== nil{

		if err := json.Unmarshal([]byte(body),Req);err != nil{

			return
		}
	}
}