package core

import (
        "strconv"
)

type Point struct {
        X int
        Y int
}

func GetKey(p *Point) string {
        return strconv.Itoa(p.X) + "-" + strconv.Itoa(p.Y)
}
