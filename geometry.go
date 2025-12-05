package geometry
import(
"math"
)

type Point struct{
X,Y float64
}

func Media(p1,p2 Point)(P Point){
P.x=(p1.X+p2.X)/2
P.Y=(p1.Y+p2.Y)/2
return
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



