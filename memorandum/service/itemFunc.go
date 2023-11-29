package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"memorandum/MiddleWare"
	"memorandum/model"
	"memorandum/util"
	"time"
)

type returnTodos struct {
	Id         int       `json:"id"`
	Content    string    `json:"content"`
	Status     string    `json:"status"`
	Created_at time.Time `json:"created_At"`
	Updated_at time.Time `json:"updated_At"`
}

// Test 测试
//
//		@Tags	item_CRUD
//		@Summary		测试
//		@Description	Description 创建待办事项
//		@Accept			application/json
//		@Produce		application/json
//	 	@Param Authorization header string false "Bearer JWT-TOKEN"
//		@Success 200 {string} json{"data","msg"}
//		@Router			/user/test [get]
func Test(ctx context.Context, c *app.RequestContext) {
	c.JSON(200, utils.H{
		"message": "测试成功",
	})
}

// CreateItem 创建待办事项
//
//		@Tags	item_CRUD
//		@Summary		增
//		@Description	Description 创建待办事项
//		@Accept			application/json
//		@Produce		application/json
//	 	@Param Authorization header string false "Bearer JWT-TOKEN"
//		@Param item query string false "item"
//		@Success 200 {string} json{"data","msg"}
//		@Router			/user/createitem [get]
func CreateItem(ctx context.Context, c *app.RequestContext) {
	item := c.Query("item")
	todo := &model.Todo{}
	todo.Status = "未完成"
	username := MiddleWare.UserMiddleware.IdentityHandler(ctx, c)
	isIdExist, id := model.FindIdByName(username)
	if isIdExist {
		todo.UserId = id
		todo.Item = item
		_ = util.DB.Create(todo)
		c.JSON(200, utils.H{
			"message": "创建成功",
		})
		return
	}
	c.JSON(200, utils.H{
		"message": "创建失败",
	})
}

// FindItem 查询待办事项
//
//		@Tags	item_CRUD
//		@Summary		查
//		@Description	Description 查询待办事项
//		@Accept			application/json
//		@Produce		application/json
//	 	@Param Authorization header string false "Bearer JWT-TOKEN"
//		@Success 200 {string} json{"data","msg"}
//		@Router			/user/finditem [get]
func FindItem(ctx context.Context, c *app.RequestContext) {
	username := MiddleWare.UserMiddleware.IdentityHandler(ctx, c)
	var returnDatas []returnTodos
	count := 0
	fmt.Println("用户名: ", username)
	isIdExist, todos := model.FindItemsByUsername(username)
	if isIdExist {
		for i, todo := range todos {
			count++
			var returnData returnTodos
			returnData.Id = i + 1
			returnData.Status = todo.Status
			returnData.Content = todo.Item
			returnData.Created_at = todo.CreatedAt
			returnData.Updated_at = todo.UpdatedAt
			returnDatas = append(returnDatas, returnData)
		}
		c.JSON(200, utils.H{
			"status": 200,
			"date": utils.H{
				"items": returnDatas,
				"total": count,
			},
			"message": "查询成功!",
			"error":   "",
		})
		return
	}
	c.JSON(200, utils.H{
		"message": "暂无待办事项",
	})
}

// EditItemStatus 修改待办事项状态
//
//		@Tags	item_CRUD
//		@Summary		改
//		@Description	Description 修改待办事项状态
//	 	@Param Authorization header string false "Bearer JWT-TOKEN"
//		@Param item formData string false "item"
//		@Param status formData string false "status"
//		@Success 200 {string} json{"data","msg"}
//		@Router			/user/edititemstatus [post]
func EditItemStatus(ctx context.Context, c *app.RequestContext) {
	item := c.PostForm("item")
	ok := c.PostForm("status")
	username := MiddleWare.UserMiddleware.IdentityHandler(ctx, c)
	isIdExist, id := model.FindIdByName(username)
	fmt.Println("item is:", item, "ok is:", ok)
	if isIdExist {
		isExist, todo := model.FindTodoUserByItemAndUserId(item, id)
		if isExist {
			result := util.DB.Model(todo).Updates(model.Todo{
				Status: ok,
			})
			if result.Error == nil {
				c.JSON(200, utils.H{
					"message": "修改成功",
				})
				return
			} else {
				fmt.Println(result.Error)
				c.JSON(200, utils.H{
					"message": "修改失败",
				})
				return
			}
		} else {
			fmt.Println("待办事项不存在")
			return
		}
	} else {
		fmt.Println("用户不存在")
		return
	}
}

// DeleteItem 删除待办事项
//
//		@Tags	item_CRUD
//		@Summary		删
//		@Description	Description 删除待办事项
//	 	@Param Authorization header string false "Bearer JWT-TOKEN"
//		@Param item formData string false "item"
//		@Success 200 {string} json{"data","msg"}
//		@Router			/user/deleteitem [post]
func DeleteItem(ctx context.Context, c *app.RequestContext) {
	item := c.PostForm("item")
	username := MiddleWare.UserMiddleware.IdentityHandler(ctx, c)
	isIdExist, id := model.FindIdByName(username)
	if isIdExist {
		isExist, todo := model.FindTodoUserByItemAndUserId(item, id)
		if isExist {
			result := util.DB.Table("todos").Delete(&todo)
			if result.Error == nil {
				c.JSON(200, utils.H{
					"message": "删除成功",
				})
				return
			} else {
				fmt.Println(result.Error)
				c.JSON(200, utils.H{
					"message": "删除失败",
				})
				return
			}
		} else {
			fmt.Println("待办事项不存在")
			c.JSON(200, utils.H{
				"message": "待办事项不存在",
			})
			return
		}
	} else {
		fmt.Println("用户不存在")
		c.JSON(200, utils.H{
			"message": "用户不存在",
		})
		return
	}
}
