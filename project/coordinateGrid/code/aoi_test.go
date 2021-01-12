package code

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	aoiMgr := NewAOIManager(100, 300, 4, 200, 450, 5)
	fmt.Println(aoiMgr)
}

func TestAOIManager_GetAroundGridsByGid(t *testing.T) {
	aoiMgr := NewAOIManager(100, 300, 4, 200, 450, 5)

	for k, _ := range aoiMgr.grids {
		gid := k
		grids := aoiMgr.GetAroundGridsByGidA(gid)

		// 得到九宫格所有的IDS
		fmt.Println("gid:", k, " grids len :", len(grids))
		gids := make([]int, 0)
		for _, grid := range grids {
			gids = append(gids, grid.Gid)
		}
		fmt.Printf("grid ID: %d, rounding grid IDS are %v \n", gid, gids)
	}
}
