package surface

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func PrintSurface() {

	var (
		fileName = "tutorial/3.2_surface/surface.svg"
		content  = ""
		err      error
		//file *os.File
	)
	s1 := "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' " + "width='"
	s2 := "' hieght='"
	s3 := "'>"
	content = fmt.Sprintf("%s%d%s%d%s", s1, width, s2, height, s3)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			new_string := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			content = fmt.Sprintf("%s%s", content, new_string)
		}
	}
	s4 := "</svg>"
	content = fmt.Sprintf("%s%s", content, s4)

	// fmt.Printf("%s", content)
	if err = os.WriteFile(fileName, []byte(content), 0666); err != nil {
		fmt.Println("Write File Error: ", err)
	}

}

/* code annotation of corener() */
func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy /* 每个网格顶点的坐标参数 */
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
