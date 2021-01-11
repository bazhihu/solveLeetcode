package code

import (
	"fmt"
	"sync"
)

type hGrid interface {
}

/*
	地图中的格子类
*/
type Grid struct {
	Gid       int          // 格子ID
	MinX      int          // 格子左边界坐标
	MaxX      int          // 格子右边界坐标
	MinY      int          // 格子上边界坐标
	MaxY      int          // 格子下边界坐标
	playerIds map[int]bool // 当前格子内的玩家或物体成员ID
	pIdLock   sync.RWMutex //playerIds的保护map 的锁
}

// 初始化一个格子
func NewGrid(gid, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		Gid:       gid,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		playerIds: make(map[int]bool),
	}
}

// 向当前格子中添加一个玩家
func (g *Grid) Add(playerId int) {
	g.pIdLock.Lock()
	defer g.pIdLock.Unlock()

	g.playerIds[playerId] = true
}

// 从格子中删除一个玩家
func (g *Grid) Remove(playerId int) {
	g.pIdLock.Lock()
	defer g.pIdLock.Unlock()

	delete(g.playerIds, playerId)
}

// 得到当前格子中所有的玩家
func (g *Grid) GetPlayerIds() (playerIds []int) {
	g.pIdLock.RLock()
	defer g.pIdLock.RUnlock()

	for k, _ := range g.playerIds {
		playerIds = append(playerIds, k)
	}
	return
}

// 打印信息方法
func (g *Grid) String() string {
	return fmt.Sprintf("Gird id: %d, minX: %d, maxX: %d, minY: %d, maxY: %d, playerIds: %v",
		g.Gid, g.MinX, g.MaxX, g.MinY, g.MaxY, g.playerIds)
}
