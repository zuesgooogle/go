package main

import (
        "fmt"
        "time"
        "astar/core"
)

const (
        WIDTH  = 192
        HEIGHT = 192
)

func main() {
        count := 0
        for {
                count += 1
                if count > 1000 {
                        return
                }

                now := time.Now()

                points := CreatePoint()
                start := &core.ANode{P: &core.Point{1, 1}}
                end := &core.ANode{P: &core.Point{64, 64}}

                mapInfo := core.MapInfo{points, len(points[0]), len(points), start, end}

                core.FindPath(&mapInfo)

                fmt.Printf("use time: %d \n", time.Since(now).Nanoseconds() / 1e6)

                //PrintMap(&mapInfo)
        }
}

func PrintMap(mapInfo *core.MapInfo) {
        for i := 0; i < HEIGHT; i++ {
            for j := 0; j < WIDTH; j++ {
                    fmt.Printf("%d ", mapInfo.Points[i][j])
            }
            fmt.Println("")
        }
}

func CreatePoint() ([][]int) {
        points := make([][]int, HEIGHT)

        for i := range points {
                subArray := make([]int, WIDTH)
                points[i] = subArray
        }

        for i := 0; i < HEIGHT; i++ {
                for j := 0; j < WIDTH; j++ {
                        //init barrier
                        if i > 20 && j > 20 {
                                if i % 20 == 0 && j % 20 == 0 {
                                        points[i - 15][j] = 1
                                        points[i - 14][j] = 1
                                        points[i - 13][j] = 1
                                        points[i - 12][j] = 1
                                        points[i - 11][j] = 1
                                        points[i - 10][j] = 1
                                        points[i - 9][j] = 1
                                        points[i - 8][j] = 1
                                        points[i - 7][j] = 1
                                        points[i - 6][j] = 1
                                        points[i - 5][j] = 1
                                        points[i - 4][j] = 1
                                        points[i - 3][j] = 1
                                        points[i - 2][j] = 1
                                        points[i - 1][j] = 1
                                        points[i][j] = 1
                                }
                        }
                }
        }

        return points
}