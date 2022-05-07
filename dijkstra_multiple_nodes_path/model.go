/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-05-07 12:07:13
 * @LastEditTime: 2022-05-07 12:21:36
 * @FilePath: /graphs-Rishabh-Mishra/dijkstra_multiple_nodes_path/model.go
 * @description: type some description
 */

package main

import (
	"math"

	geo "github.com/kellydunn/golang-geo"
)

type GeoNode interface {
	// Node identify
	ID() string
	Lng() float64
	Lat() float64
}

// DistanceBetween Use this repo -> https://github.com/kellydunn/golang-geo
func DistanceBetween(n1 GeoNode, n2 GeoNode) int {
	p1 := geo.NewPoint(n1.Lat(), n1.Lng())
	p2 := geo.NewPoint(n2.Lat(), n2.Lng())
	kM := p1.GreatCircleDistance(p2)
	meter := math.Trunc(kM*1000*1e0) * 1e-0
	return int(meter)
}
