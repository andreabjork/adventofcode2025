package day9

import (
	"fmt"
	"strconv"
	"adventofcode/m/v2/util"
	"strings"
	"sort"
)

func Day9(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Largest area: %d\n", findLargest(inputFile, false))
	} else {
		fmt.Printf("Largest area: %d\n", findLargest(inputFile, true))
	}
}

var MAX int // Tracks the largest X-index
func findLargest(inputFile string, areaContained bool) int {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	pts := []*Point{}
	for ok {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		if x > MAX {
			MAX = x
		}
		pts = append(pts, &Point{x, y})
		line, ok = util.Read(ls)
	}
	
	pgn, boxes := uncontained(pts)
	if areaContained {
		return pgn.maxContained(boxes)
	} 
		
	return boxes[0].area()
}

// Returns the polygon and all boxes in a sorted list, by descending area
func uncontained(pts []*Point) (*Polygon, []*Box) {
	pgn := makePolygon(pts)
	boxes := []*Box{}
	max := 0
	for i := 0; i < len(pts); i++ {
		for j := 0; j < len(pts); j++ {
			if i == j {
				continue
			}
			b := &Box{pts[i],pts[j]}
			boxes = append(boxes, b)
			if b.area() > max {
				max = b.area() 
			}
		}
	}

	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i].area() > boxes[j].area()
	})

	return pgn, boxes
}

// =======
// POLYGON 
// =======
type Polygon struct {
	vertSegments []*Segment
	horizSegments []*Segment
	memo map[int]map[int]bool 
}

func makePolygon(pts []*Point) *Polygon {
	vertSegments := []*Segment{}
	horizSegments := []*Segment{}
	for i := 0; i < len(pts)-1; i++ {
		if pts[i].x == pts[i+1].x { // Vertical segment, matching x
			vertSegments = append(vertSegments, &Segment{pts[i], pts[i+1]})
		} else {
			horizSegments = append(horizSegments, &Segment{pts[i], pts[i+1]})
		}
	}

	if pts[len(pts)-1].x == pts[0].x { // Vertical segment, matching x
		vertSegments = append(vertSegments, &Segment{pts[len(pts)-1], pts[0]})
	} else {
		horizSegments = append(horizSegments, &Segment{pts[len(pts)-1], pts[0]})
	}
	
	return &Polygon{
		vertSegments: vertSegments, 
		horizSegments: horizSegments, 
		memo: make(map[int]map[int]bool),
	}
}

// ..............
// .......#XXX#..
// .......X...X..
// ..#XXXX#...X..
// ..X........X..
// ..#XXXXXX#.X..
// .........X.X..
// .........#X#..
// ..............
func (pgn *Polygon) maxContained(boxes []*Box) int {
	for i, b := range boxes { // This is sorted in descending order
		 fmt.Printf("%d/%d\n", i, len(boxes))
		 if pgn.containsBox(b.p, b.q) {
			 return b.area() // Break on the first box that is contained because the list is sorted
		 }
	}
	return -1
}

func (pgn *Polygon) containsBox(p *Point, q *Point) bool {
	 fmt.Printf("--------------- Check box %+v to %+v\n", p, q)
  // We simply check if any point on the box border lies 
	// outside the polygon
	a := &Point{p.x, p.y} 
	b := &Point{p.x, q.y}
	c := &Point{q.x, p.y}
	d := &Point{q.x, q.y}

	if !pgn.containsSegment(&Segment{a,b}) || 
	!pgn.containsSegment(&Segment{b,c}) || 
	!pgn.containsSegment(&Segment{c,d}) || 
	!pgn.containsSegment(&Segment{d,a}) {
		return false
	} 

	return true
}

func (pgn *Polygon) onBorder(p *Point) bool {
	for _, s := range pgn.horizSegments { // y is fixed
		if p.y == s.p.y && ((p.x >= s.p.x && p.x <= s.q.x) || (p.x <= s.p.x && p.x >= s.q.x)) {
			return true
		}
	} 

	for _, s := range pgn.vertSegments {  // x is fixed
		if p.x == s.p.x && ((p.y >= s.p.y && p.y <= s.q.y) || (p.y <= s.p.y && p.y >= s.q.y)) {
			return true
		}
	}

	return false
}

// Checks if a vertical or horizontal segment is contained
func (pgn *Polygon) containsSegment(s *Segment) bool {
	fmt.Printf("Contains segment %d,%d to %d,%d?\n", s.p.x, s.p.y, s.q.x, s.q.y)
	if s.p.x == s.q.x  { // Process vertical segments
		for _, r := range pgn.horizSegments {
			return s.verticalIntersectsHorizontal(r)
		}
	} else if s.p.y == s.q.y { // horizontal segments
		for _, r := range pgn.vertSegments { 
			return r.verticalIntersectsHorizontal(s)
		}
	}
	return true
}

// ===
// BOX
// ===
type Box struct {
	p, q *Point
}

func (b *Box) area() int {
	return (util.Abs(b.p.x-b.q.x) + 1) * (util.Abs(b.p.y-b.q.y) + 1)
}

// =======
// SEGMENT
// =======
type Segment struct {
	p, q *Point
}


func (s *Segment) verticalIntersectsHorizontal(r *Segment) bool {
	if (s.p.x > r.p.x && s.p.x < r.q.x) || 
		(s.p.x < r.p.x && s.p.x > r.q.x) { // x value is within the segment r
		if	(s.p.y > r.p.y && s.q.y < r.p.y) || 
		(s.p.y < r.p.y && s.q.y > r.p.y) { // y values are above and below the segment r
			fmt.Printf("NO: cuts into horizontal segment %d,%d to %d,%d\n", r.p.x, r.p.y, r.q.x, r.q.y)
			return false
		}
	}
	return true
}

// =====
// POINT
// =====
type Point struct {
	x, y int
}