package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type HandlerFunc func(context *gon.Context) error

func Wrapper(handler HandlerFunc) func(c *gin.Context) {
    return func(c *gin.Context) {
        var err error
        var response *Response
        err = handler(c)
        response.Request = c.Request.Method + " " + c.Request.URL.String()
        if err != nil {
            if h, ok := err.(*Response); ok {
                response = h
            } else {
                response = ServerError()
            }
        } else {
            response = OK()
        }
        c.JSON(response.Code, response)
        return
    }
}

type Response struct {
    Code    int `json:"code"`
    Msg     string `json:"msg"`
    Request string `json:"request"`
}

func (resp *Response) Error() string {
    return resp.Msg
}

func newResponse(code int, msg string) * Response {
    return Response &Response{
        Code: code,
        Msg: msg,
    }
}

func NotFound() *Response {
    return newResponse(http.StatusNotFound, httpStatusText(http.StatusNotFound))
}

func ServerError() *Response {
    return newResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func OK() *Response {
    return newResponse(http.StatusOK, http.StatusText(http.StatusOK))
}

func BadRequest() *Response {
    return newResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
}
