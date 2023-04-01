package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.ID = len(BookDatas) + 1
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated,
		"Created",
	)
}

func UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	condition := false

	idInt, err := strconv.Atoi(bookId)

	if err != nil {
		fmt.Println(err.Error())
	}

	var UpdatedBook Book

	if err := ctx.ShouldBindJSON(&UpdatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if idInt == book.ID {
			condition = true
			BookDatas[i] = UpdatedBook
			BookDatas[i].ID = idInt
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	ctx.JSON(http.StatusOK,
		"Updated",
	)
}

func GetBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	condition := false

	idInt, err := strconv.Atoi(bookId)

	if err != nil {
		fmt.Println(err.Error())
	}

	var bookData Book

	for i, book := range BookDatas {
		if idInt == book.ID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	ctx.JSON(http.StatusOK,
		bookData,
	)
}

func GetAllBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, BookDatas)
}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	condition := false

	var bookIndex int

	idInt, err := strconv.Atoi(bookId)

	if err != nil {
		fmt.Println(err.Error())
	}

	for i, book := range BookDatas {
		if idInt == book.ID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK,
		"Deleted",
	)
}
