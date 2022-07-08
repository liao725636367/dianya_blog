package model

type CommonPageReq struct {
	Page int `json:"page" d:"1" v:"min:1"`
	Size int `json:"size" d:"10" v:"in:10,20"`
}
