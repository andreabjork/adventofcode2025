package day7

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
	"sort"
)

func Day7(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Largest area: %d\n", findLargest(inputFile, false))
	} else {
		fmt.Printf("Largest area: %d\n", findLargest(inputFile, true))
	}
}

type Point struct {
	x, y int
}

type Segment struct {
	p, q *Point
}

type Box struct {
	p, q *Point
	area int
}

type Polygon struct {
	vertSegments []*Segment
	horizSegments []*Segment
	memo map[int]map[int]bool 
}

var MAX int

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
		
	return boxes[0].area
}

func uncontained(pts []*Point) (*Polygon, []*Box) {
	pgn := makePolygon(pts)
	boxes := []*Box{}
	max := 0
	for i := 0; i < len(pts); i++ {
		for j := 0; j < len(pts); j++ {
			if i == j {
				continue
			}
			A := area(pts[i], pts[j])
			boxes = append(boxes, &Box{pts[i],pts[j], A})
			if A > max {
				max = A
			}
		}
	}

	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i].area > boxes[j].area
	})

	return pgn, boxes
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
			 return b.area // Break on the first box we find because it's the largest
		 }
	}
	return -1
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

func (pgn *Polygon) containsBox(p *Point, q *Point) bool {
  // First check the other two corners for fast exit
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

	//t = &Point{q.x, p.y}
	//if !pgn.contains(t) {
	//	return false
	//}

	//startX,startY := p.x, p.y
	//endX,endY := q.x, q.y
	//if p.x > q.x {
	//	startX  = q.x
	//	endX = p.x
	//} 
	//if p.y > q.y {
	//	startY = q.y
	//	endY = p.y
	//}

	//for x := startX ; x <= endX; x++ {
	//	for y := startY ; y <= endY; y++ {
	//		if !pgn.contains(&Point{x,y}) {
	//			return false
	//		}
	//	}
	//}  

	return true
}

// Checks if a vertical or horizontal segment is contained
func (pgn *Polygon) containsSegment(s *Segment) bool {
	if s.p.x == s.q.x  { // Process vertical segments
		startY := s.p.y
		endY := s.q.y
		if s.p.y > s.q.y {
			startY = s.q.y
			endY = s.p.y
		}
		for y := startY; y <= endY; y++ {
			if !pgn.contains(&Point{s.p.x,y}) {
				return false
			}
		}
	} else { // horizontal segments
		startX := s.p.x 
		endX := s.q.x
		if s.p.x > s.q.x {
			startX = s.q.x
			endX = s.p.x
		}
		for x := startX; x <= endX; x++ {
			if !pgn.contains(&Point{x,s.p.y}) {
				return false
			}
		}
	}
	return true
}

// true if p is inside the polygon bounded by pts
func (pgn *Polygon) contains(p *Point) bool {
	if _, ok := pgn.memo[p.x]; ok {
		if c, ok := pgn.memo[p.x][p.y]; ok {
			return c
		} 
	}

	ints := 0
	// We walk to the right from p always and check how many intersects it makes
	// to segment
	for k := 0; p.x+k <= MAX; k++ {
		for _, s := range pgn.vertSegments {
			if s.onSegment(&Point{p.x + k, p.y}) {
				if k == 0 {
					pgn.memorize(p.x+k,p.y, true)
					return true // p is on one of the segments
				}
				ints++ // a line from p to the right intersects with one of the segments
			}
		}

		// Check also if any is on horizontal segments
		if k == 0 {
			for _, s := range pgn.horizSegments {
				if s.onSegment(&Point{p.x, p.y}) {
					pgn.memorize(p.x,p.y, true)
					return true
				}
			}
		}
	}

	if ints%2 == 0 { // even number of intersects -> point is outside
		pgn.memorize(p.x,p.y, false)
		return false
	} else {
		pgn.memorize(p.x,p.y, true)
		return true
	}
}

func (pgn *Polygon) memorize(x,y int, contained bool) {
	if _, ok := pgn.memo[x]; !ok {
		pgn.memo[x] = make(map[int]bool)
	}
	pgn.memo[x][y] = contained
}

func (s *Segment) onSegment(p *Point) bool {
	if s.p.x == s.q.x && p.x == s.p.x { // Vertical segment, matching x
		from := s.p.y
		to := s.q.y
		if from > to {
			from = s.q.y
			to = s.p.y
		}

		for y := from; y <= to; y++ {
			if p.y == y {
				return true
			}
		}
	}
	
	if s.p.y == s.q.y && p.y == s.p.y { // Horizontal segment, matching y
		from := s.p.x
		to := s.q.x
		if from > to {
			from = s.q.x
			to = s.p.x
		}

		for x := from; x <= to; x++ {
			if p.x == x {
				return true
			}
		}
	}
	return false
}

func area(p *Point, q *Point) int {
	return (util.Abs(p.x-q.x) + 1) * (util.Abs(p.y-q.y) + 1)
}
