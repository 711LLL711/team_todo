package controller

import (
	"net/http"
	"team_todo/database"
	"team_todo/model"
	"team_todo/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GroupController struct {
}

// 创建群组
func (gc *GroupController) CreateGroup(c *gin.Context) {
	//需要从session中获取用户id
	GroupOwnerId, err := util.GetId_Session(c.Request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var groupinfo model.Group
	groupinfo.GroupName = c.PostForm("name")
	groupinfo.GroupOwnerId = GroupOwnerId
	groupinfo.Group_Description = c.PostForm("description")
	groupinfo.GroupId = uuid.New().String()
	err = database.CreateGroup(groupinfo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": groupinfo.GroupId, "name": groupinfo.GroupName, "description": groupinfo.Group_Description, "owner": groupinfo.GroupOwnerId})
	return
}

// 获取群组列表
func (gc *GroupController) GetGroupList(c *gin.Context) {
	Id, err := util.GetId_Session(c.Request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	groups, count, err := database.GetUserGroups(Id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count, "groups": groups})
}

// 获取群组信息
func (gc *GroupController) GetGroupInfo(c *gin.Context) {
	GetId := c.Param("id") //群组id
	groupinfo, err := database.GetGroupInfo(GetId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": groupinfo.GroupId, "name": groupinfo.GroupName, "description": groupinfo.Group_Description, "owner": groupinfo.GroupOwnerId})
}

func (gc *GroupController) GetGroupMembers(c *gin.Context) {
	GroupId := c.Param("id")
	members, count, err := database.GetGroupMembers(GroupId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count, "members": members})
}
