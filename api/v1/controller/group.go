package controller

import (
	"errors"
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
	groupinfo.GroupId = uuid.New().String()[:8] // 提取前8个字符作为短UUID
	//创建群组
	err = database.CreateGroup(groupinfo)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var groupwithuser model.GroupWithUser
	groupwithuser.UserId = GroupOwnerId
	groupwithuser.GroupId = groupinfo.GroupId
	//加入到群组成员对应表
	err = database.AddUserToGroup(groupwithuser)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": groupinfo.GroupId, "name": groupinfo.GroupName, "description": groupinfo.Group_Description, "owner": groupinfo.GroupOwnerId})
}

// 获取已加入的群组列表
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

// 获取群组成员
func (gc *GroupController) GetGroupMembers(c *gin.Context) {
	GroupId := c.Param("id")
	members, count, err := database.GetGroupMembers(GroupId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count, "members": members})
}

// 获取群组邀请码
func (gc *GroupController) GetGroupCode(c *gin.Context) {
	GroupId := c.Param("id")
	Code, err := database.GetGroupCode(GroupId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	//群组本来没有邀请码，新建邀请码
	if Code == "" {
		Code = uuid.New().String()[:12]
		database.StoreGroupCode(GroupId, Code)
	}
	c.JSON(http.StatusOK, gin.H{"code": Code})
}

// 加入群组
func (gc *GroupController) JoinGroup(c *gin.Context) {
	code := c.Query("code")
	groupid, err := database.Search_Group_Invite(code)
	if err != nil {
		var ErrNotFound = errors.New("NOT FOUND")
		if errors.Is(err, ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
	} else {
		//添加用户到groupwithuser表
		var groupwithuser model.GroupWithUser
		groupwithuser.GroupId = groupid
		groupwithuser.UserId, _ = util.GetId_Session(c.Request)
		err = database.AddUserToGroup(groupwithuser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		} else {
			//获取群组信息并返回
			groupinfo, _ := database.GetGroupInfo(groupid)
			c.JSON(http.StatusOK, groupinfo)
		}
	}

}

// 退出群组
func (gc *GroupController) LeaveGroup(c *gin.Context) {
	var groupwithuser model.GroupWithUser
	groupwithuser.UserId, _ = util.GetId_Session(c.Request)
	groupwithuser.GroupId = c.Param("id")
	err := database.LeaveGroup(groupwithuser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "leave group successfully"})
	}
}

// 解散群组
func (gc *GroupController) DelGroup(c *gin.Context) {
	groupid := c.Param("id")
	err := database.DelGroup(groupid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Del group successfully"})
	}
}
