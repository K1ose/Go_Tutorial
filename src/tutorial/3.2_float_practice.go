package tutorial

/* TODO */

/*  如果f函数返回的是无限制的float64值，那么SVG文件可能输出无效的多边形元素（虽然许多SVG渲染器会妥善处理这类问题）。跳过无效的多边形。 */
func SkipInvalid() {

}

/* 输出Egg Box图案 */
func PrintEggBox() {

}

/* 输出Moguls图案 */
func PrintMoguls() {

}

/* 输出Saddle图案 */
func PrintSaddle() {

}

/* 根据高度给每个多边形上色，那样峰值部将是红色（#ff0000），谷部将是蓝色（#0000ff） */
func PaintSurface() {

}

/* 构造一个web服务器，用于计算函数曲面然后返回SVG数据给客户端。服务器必须设置Content-Type头部：
w.Header().Set("Content-Type", "image/svg+xml")
允许客户端通过HTTP请求参数设置高度、宽度和颜色等参数。 */
func PrintByHttp() {

}
