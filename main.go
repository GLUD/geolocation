package main

import (
    "log"
    "fmt"
	"github.com/kataras/iris"
    "googlemaps.github.io/maps"
    "golang.org/x/net/context"
    )

func Rutas(ctx *iris.Context){
    des := ctx.URLParam("des")
    lat := ctx.URLParam("lat")
    lng := ctx.URLParam("lng")
    c, err := maps.NewClient(maps.WithAPIKey("AIzaSyAL-vQX-Gk-7bb0IPv5Xgx8MnZCZGUZ8IE"))
    if err != nil {
       log.Fatalf("fatal error: %s", err)
    }
     r := &maps.DirectionsRequest{
       Origin: fmt.Sprintf("%s,%s",lat,lng),
       Destination: des,
    }
    res, _, err := c.Directions(context.Background(), r)
    if err != nil {
       log.Fatalf("fatal error: %s", err)
    }
    ctx.JSON(200,res)
}
func main() {
    iris.Get("/rutas",Rutas)
	iris.Listen(":8080")
}
