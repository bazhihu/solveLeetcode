package code

import "fmt"

/*
	AOI管理模块
*/
type AOIManager struct {
	MinX  int // 区域左边界坐标
	MaxX  int // 区域右边界坐标
	CntsX int // X方向格子的数量

	MinY  int // 区域上边界坐标
	MaxY  int // 区域下边界坐标
	CntsY int //Y方向的格子数量

	grids map[int]*Grid // 当前区域中都有哪些格子，key:格子ID val: 格子对象
}

/*
	初始化一个AOI区域
*/
func NewAOIManager(minX, maxX, cntsX, minY, maxY, cntsY int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntsX: cntsX,
		MinY:  minY,
		MaxY:  maxY,
		CntsY: cntsY,
		grids: make(map[int]*Grid),
	}

	// 给AOI初始化区域中所有的格子
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			// 计算格子ID
			gid := y*cntsX + x

			// 初始化一个格子放在AOI中的MAP里
			aoiMgr.grids[gid] = NewGrid(gid,
				minX+x*aoiMgr.gridWidth(),
				minX+(x+1)*aoiMgr.gridWidth(),
				minY+y*aoiMgr.gridLength(),
				minY+(y+1)*aoiMgr.gridLength())
		}
	}

	return aoiMgr
}

// 得到每个格子在X轴方向的宽度
func (a *AOIManager) gridWidth() int {
	return (a.MaxX - a.MinX) / a.CntsX
}

// 得到每个格子在X轴方向的长度
func (a *AOIManager) gridLength() int {
	return (a.MaxY - a.MinY) / a.CntsY
}

// 打印信息方法
func (a *AOIManager) String() string {
	s := fmt.Sprintf("AOIManager:\n minX: %d, maxX: %d, cntsX:%d, minY: %d, maxY: %d, cntsY: %d \n Grids in AOI Manager: \n",
		a.MinX, a.MaxX, a.CntsX, a.MinY, a.MaxY, a.CntsY)

	for k, _ := range a.grids {
		s += fmt.Sprintln(a.grids[k])
	}

	return s
}

// 根据格子的gid 得到当前周边的九宫格信息
func (a *AOIManager) GetAroundGridsByGid(gid int) (grids []*Grid) {
	// 判断gid 是否存在
	if _, ok := a.grids[gid]; !ok {
		return
	}

	// 将gid添加进九宫格
	grids = append(grids, a.grids[gid])

	// 根据gid, 得到格子所在的坐标
	x, y := gid%a.CntsX, gid/a.CntsY
	surRoundGid := make([]int, 0)

	// 8个方向向量
	// 左上：(-1, -1) 左中：(-1, 0) 左下：(-1, 1)
	// 中上：(0, -1) 中下：(0, 1)
	// 右上：(1, -1) 右中：(1, 0) 右下：(1, 1)
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < 8; i++ {
		newX := x + dx[i]
		newY := y + dy[i]

		if newX >= 0 && newX < a.CntsX && newY >= 0 && newY < a.CntsY {
			surRoundGid = append(surRoundGid, newY*a.CntsX+newX)
		}
	}

	// 根据没有越界的gid,得到格子的信息
	for _, gid := range surRoundGid {
		grids = append(grids, a.grids[gid])
	}

	return
}

// 根据格子的gid 得到当前周边的九宫格信息 || 两次遍历版
func (a *AOIManager) GetAroundGridsByGidA(gid int) (grids []*Grid) {
	// 判断gid 是否存在
	if _, ok := a.grids[gid]; !ok {
		return
	}

	// 将gid添加进九宫格
	grids = append(grids, a.grids[gid])

	// 根据gid, 得到格子所在的坐标
	y := gid / a.CntsY
	surRoundGid := make([]int, 0)

	// 竖向找
	casualGid := []int{gid}
	if y-1 > 0 {
		casualGid = append(casualGid, gid-a.CntsX)
		surRoundGid = append(surRoundGid, gid-a.CntsX)
	}

	if y+1 < a.MaxY {
		casualGid = append(casualGid, gid+a.CntsX)
		surRoundGid = append(surRoundGid, gid+a.CntsX)
	}

	// 横向找
	for _, val := range casualGid {
		newX := val % a.CntsX
		if newX-1 > 0 {
			surRoundGid = append(surRoundGid, val-1)
		}

		if newX+1 < a.MaxX {
			surRoundGid = append(surRoundGid, val+1)
		}
	}

	// 根据没有越界的gid,得到格子的信息
	for _, gid := range surRoundGid {
		grids = append(grids, a.grids[gid])
	}

	return
}

// 通过横纵坐标获取对应的格子ID

// 通过横纵坐标得到周边九宫格内的全部PlayerIDs

// 通过GID 获取当前格子的全部playerID

// 移除一个格子的playerId

// 添加一个player到一个格子中

// 通过横纵坐标添加一个player到一个格子中

// 通过横纵坐标把一个player对应的格子中删除
