package controller

import "golearn/bluebell/models"

// bluebell/controller/docs_models.go

//_ResponsePostList 帖子列表接口响应数据
type _ResponsePostList struct {
	Code    ResCode                 `json:"code"`    // 业务响应状态码
	Message string                  `json:"message"` // 提示信息
	Data    []*models.ApiPostDetail `json:"data"`    // 数据
}

//_ResponseCommunityList 帖子列表接口响应数据
type _ResponseCommunityList struct {
	Code    ResCode             `json:"code"`    // 业务响应状态码
	Message string              `json:"message"` // 提示信息
	Data    []*models.Community `json:"data"`    // 数据
}

//_ResponseCommunityDetail 帖子社区详细接口响应数据
type _ResponseCommunityDetail struct {
	Code    ResCode                   `json:"code"`    // 业务响应状态码
	Message string                    `json:"message"` // 提示信息
	Data    []*models.CommunityDetail `json:"data"`    // 数据
}

//_ResponsePost 创建帖子数据接口
type _ResponsePost struct {
	Code    ResCode        `json:"code"`    // 业务响应状态码
	Message string         `json:"message"` // 提示信息
	Data    []*models.Post `json:"data"`    // 数据
}

//_ResponseGetPostDetail 查询帖子详细信息
type _ResponseGetPostDetail struct {
	Code    ResCode                 `json:"code"`    // 业务响应状态码
	Message string                  `json:"message"` // 提示信息
	Data    []*models.ApiPostDetail `json:"data"`    // 数据
}
type _ResponseCommunity struct {
	Code    ResCode           `json:"code"`    // 业务响应状态码
	Message string            `json:"message"` // 提示信息
	Data    *models.Community `json:"data"`    // 数据
}
