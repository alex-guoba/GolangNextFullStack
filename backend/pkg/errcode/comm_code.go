package errcode

import "net/http"

var (
	Success                   = NewError(0, "Success", http.StatusOK)
	ServerError               = NewError(100000, "Server error", http.StatusInternalServerError)
	InvalidParams             = NewError(100001, "Invalid params", http.StatusBadRequest)
	NotFound                  = NewError(100002, "Not found", http.StatusNotFound)
	UnauthorizedAuthNotExist  = NewError(100003, "Auth failed, not exist", http.StatusUnauthorized)
	UnauthorizedTokenError    = NewError(100004, "Auth failed, token error", http.StatusUnauthorized)
	UnauthorizedTokenTimeout  = NewError(100005, "Auth failed, token timeout", http.StatusUnauthorized)
	UnauthorizedTokenGenerate = NewError(100006, "Auth failed, token generate", http.StatusUnauthorized)
	TooManyRequests           = NewError(100007, "Too many requests", http.StatusTooManyRequests)
	NotImplemented            = NewError(100010, "Not Implemented Yet.", http.StatusNotImplemented)
)

var (
	// 10001000 - 10002000: blog error code.
	ErrTagIDForbidden   = NewError(10001000, "`Tag` in used by other article", http.StatusInternalServerError)
	ErrorGetTagListFail = NewError(20010001, "Get tag list fail", http.StatusInternalServerError)
	ErrorCreateTagFail  = NewError(20010002, "Create tag fail", http.StatusInternalServerError)
	ErrorUpdateTagFail  = NewError(20010003, "Update tag fail", http.StatusInternalServerError)
	ErrorDeleteTagFail  = NewError(20010004, "Delete tag fail", http.StatusInternalServerError)
	ErrorCountTagFail   = NewError(20010005, "Count tag fail", http.StatusInternalServerError)

	ErrorGetArticleFail    = NewError(20020001, "Get article fail", http.StatusInternalServerError)
	ErrorGetArticlesFail   = NewError(20020002, "Get articles fail", http.StatusInternalServerError)
	ErrorCreateArticleFail = NewError(20020003, "Create article fail", http.StatusInternalServerError)
	ErrorUpdateArticleFail = NewError(20020004, "Update article fail", http.StatusInternalServerError)
	ErrorDeleteArticleFail = NewError(20020005, "Delete article fail", http.StatusInternalServerError)

	ErrorUploadFileFail = NewError(20030001, "File upload fail", http.StatusInternalServerError)
)
