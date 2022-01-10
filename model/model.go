package model
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)
 func checkError( err error ){
	 if err != nil {
		 fmt.Println(err)
	 }
 }
 type History struct{
	Result string   `json:"result"`
	Count string     `json: "count"`
	Event Event
}
type Event struct{
	Event string    `json: "event"`
	Date string      `json: "date"`
	Description string `json: "description"`
	Lang string       `json: "lang"`
	Category1 string  `json: "category1"`
	Granularity string `json: "granularity"`

}
func ShowAllUsers()