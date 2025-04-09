package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/db"
	"go_web/jwt_plugin"
	"go_web/type/structType/user"
	"log"
	"strings"
)

const imageBaseURL = "http://localhost:8080/static/"

// GetUserInfoHandler 獲取用戶信息
func GetUserInfoHandler(c *gin.Context) {
	sqlxDB := db.ConnectDB()
	token := c.GetHeader("Authorization")
	// 檢驗token 獲取用戶登錄用戶id
	claims, err := jwt_plugin.ValidateToken(token)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "沒有查到此用戶",
		})
		return
	}
	var userInfo user.UserInfoStruct
	// 如果查到這個用戶 通過id返回用戶信息以及用戶權限返回菜單
	queryById := "SELECT email, name, phone, job_id, avatar, role FROM profile WHERE id = ?"
	err = sqlxDB.Get(&userInfo, queryById, claims.Id)
	if err != nil {
		// 处理查询结果为空的情况
		log.Println("User not found")
		return
	}
	// 把 Windows 路徑轉換為 URL 格式
	userInfo.Avatar = strings.Replace(userInfo.Avatar, `C:\project\Images\`, imageBaseURL, 1)
	userInfo.Avatar = strings.ReplaceAll(userInfo.Avatar, `\`, `/`) // 轉換反斜槓

	var menuData []byte
	// 返回菜單
	queryMenuById := "SELECT menu_json FROM menu WHERE role = ?"
	err = sqlxDB.Get(&menuData, queryMenuById, userInfo.Role)

	// 定義一個存儲 JSON 解析結果的變數
	type MenuItem struct {
		Icon     string     `json:"icon"`
		Path     string     `json:"path"`
		Label    string     `json:"label"`
		Key      string     `json:"key"`
		Children []MenuItem `json:"children,omitempty"` // 可選字段
	}
	var menu []MenuItem
	// 解析 JSON
	err = json.Unmarshal(menuData, &menu)
	if err != nil {
		fmt.Println("解析 JSON 失敗:", err)
		return
	}

	// 打印解析後的數據
	fmt.Println(menu) // 這裡會正常顯示 JSON 數據

	// 返回用戶的信息
	c.JSON(200, gin.H{
		"code":     200,
		"msg":      "success",
		"userInfo": userInfo,
		"menu":     menu,
	})
}

// GetUserInfoHandle 獲取部門信息
func GetUserInfoHandle(c *gin.Context) {
	sqlxDB := db.ConnectDB()
	var nodes []*user.Node
	err := sqlxDB.Select(&nodes, "SELECT  id, `key`, title, parent_id FROM org_list")
	if err != nil {
		log.Fatalln(err)
	}

	buildTree := func(nodes []*user.Node) []*user.Node {
		// 创建一个映射，存储节点 ID -> 节点 的映射
		nodeMap := make(map[int64]*user.Node)
		var rootNodes []*user.Node

		// 将所有节点按 ID 存入 nodeMap，并初始化每个节点的 children 为一个空数组
		for _, node := range nodes {
			nodeMap[int64(node.ID)] = node
			// 初始化 Children 字段，即使没有子节点
			if node.Children == nil {
				node.Children = []*user.Node{}
			}
		}

		// 组装树形结构
		for _, node := range nodes {
			if node.ParentID == nil {
				// 如果 ParentID 为 nil，说明是根节点
				rootNodes = append(rootNodes, node)
			} else {
				// 否则，将该节点加入到父节点的 children 列表中
				parentNode, exists := nodeMap[int64(*node.ParentID)]
				if exists {
					parentNode.Children = append(parentNode.Children, node)
				}
			}
		}
		return rootNodes
	}
	tree := buildTree(nodes)
	c.JSON(200, gin.H{
		"code":     200,
		"msg":      "success",
		"org_list": tree,
	})
}
