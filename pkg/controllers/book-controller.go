package controllers

import(
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
    "github.com/atifali-pm/go-bookstore/pkg/utils"
    "github.com/atifali-pm/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
    newBooks := models.GetAllBooks()
    res, _ :=json.Marshal(newBooks)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        fmt.PrintLn("error while parsing")
    }
    bookDetail, _ := models.GetBookById(ID)
    res, _ := json.Marshal(bookDetail)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
    CreateBook := &models.Book{}
    utils.ParseBody(r, CreateBook)
    b := CreateBook.CreateBook()
    res, _ := json.Marshal(b)
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    Id, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        fmt.PrintLn("error while parsing")
    }
    book := models.DeleteBook(Id)
    res, _ := json.Marshal(book)
    w.Header().Set("Content-Type", "pkglication/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
    var updateBook = &models.Book{}
    utils.ParseBody(r, updateBook)
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    Id, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        fmt.PrintLn("error while parsing")
    }
    bookDetail, db := models.GetBookById(Id)
    if updateBook.Name != ""{
        bookDetail.Name = updateBook.Name
    }
    if updateBook.Author != ""{
        bookDetail.Author = updateBook.Author
    }
    if updateBook.Publication != ""{
   		bookDetails.Publication = updateBook.Publication
   	}
   	db.Save(&bookDetails)
   	res, _ := json.Marshal(bookDetails)
   	w.Header().Set("Content-Type", "pkglication/json")
   	w.WriteHeader(http.StatusOK)
   	w.Write(res)
}