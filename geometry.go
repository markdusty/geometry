package geometry
import(
"fmt"
"math"
)

type Point struct{
X,Y float64
}

func CreateAPoint(x,y float64)  (P Point){
P.X = x
P.Y = y
return
}

func Distance(p1,p2 Point) (D float64){
D = math.Sqrt((p1.X-p2.X)*(p1.X-p2.X)+(p1.Y-p2.Y)*(p1.Y-p2.Y))
return
}

func main(){
var x,y float64
var p1,p2 Point
fmt.Scan(&x,&y)
p1 =CreateAPoint(x,y)
fmt.Scan(&x,&y)
p2 =CreateAPoint(x,y)
fmt.Println(Distance(p1,p2))
}

