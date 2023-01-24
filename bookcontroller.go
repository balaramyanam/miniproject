package controllers

import (
	"encoding/json"

	"fmt"

	"github.com/gorilla/mux"

	"net/http"
	"strconv"

	"github.com/balaram/bookstore/pkg/models"
	"github.com/balaram/bookstore/pkg/utils"
)

var NewBook models.Book

func GetBook(Res http.ResponseWriter, Req *http.Request) {

	newBooks := models.GetAllBooks()                   //this is a input and output of the function 

	res, _ := json.Marshal(newBooks) //response               //get operation

	Res.Header().Set("Content-Type", "pkglication/json")

	Res.WriteHeader(http.StatusOK)
	Res.Write(res)

}

func GetBookById(Res http.ResponseWriter, Req *http.Request) {

	Vars := mux.Vars(Req)

	bookId := Vars["bookId"]                                    //this is a input and output of the function 

	ID, err := strconv.ParseInt(bookId, 0, 0)                      //get operation


	if err != nil {

		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails) //response

	Res.Header().Set("Content-Type", "pkglication/json")

	Res.WriteHeader(http.StatusOK)
	Res.Write(res)

}

func CreateBook(Res http.ResponseWriter, Req *http.Request) {

	CreateBook := &models.Book{}

	utils.ParseBody(Req, CreateBook)              //post operation
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b) //response
	Res.WriteHeader(http.StatusOK)
	Res.Write(res)

}

func DeleteBook(Res http.ResponseWriter, Req *http.Request) {

	Vars := mux.Vars(Req)

	bookId := Vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)      //delete operation

	if err != nil {

		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(ID)

	res, _ := json.Marshal(book) //response

	Res.Header().Set("Content-Type", "pkglication/json")

	Res.WriteHeader(http.StatusOK)
	Res.Write(res)

}

func UpdateBook(Res http.ResponseWriter, Req *http.Request) {

	var UpdateBook = &models.Book{}

	utils.ParseBody(Req, UpdateBook)
	Vars := mux.Vars(Req)

	bookId := Vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {

		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if UpdateBook.Name != "" {
		bookDetails.Name = UpdateBook.Name

	}                                                  //put operation
          //the below three lines are validation                                            
	if UpdateBook.Author != "" {

		bookDetails.Author = UpdateBook.Author
	}

	if UpdateBook.Publication != "" {

		bookDetails.Publication = UpdateBook.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)

	Res.Header().Set("Content-Type", "pkglication/json")

	Res.WriteHeader(http.StatusOK)
	Res.Write(res)

}
