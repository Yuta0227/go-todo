package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"io"
	"os"
	"log"
)

func helloHander(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(c.Writer, "<h1>Helloaajjja Japaaaaaaaあaaan </h1>")
	fmt.Fprintf(c.Writer, "<h1>jjこんにちあは</h1>")
	fmt.Fprintf(c.Writer, "<div>avg%f</div>", avg([]int{12, 3, 4, 5}))
	min, max := minMax([]int{12, 3, 4, 5})
	fmt.Fprintf(c.Writer, "<div>min%dmax%d</div>", min, max)
	var yuta = Human{name: "yuta"}
	fmt.Fprintf(c.Writer, yuta.introduce())
	fmt.Fprintf(c.Writer, strings.Join(arrTest(), "")) // Convert []string to string
	go printNumbers()
	go printLetters()
	time.Sleep(time.Second)
	results := make(chan int)
	go worker(results) // start the worker in a goroutine

	for i := 0; i < 5; i++ {
		fmt.Println(<-results) // receive data from the channel
	}
}
func worker(results chan int) {
	for i := 0; i < 5; i++ {
		results <- i * 2 // send data on the channel
	}
}
func arrTest() []string {
	// var arr[2]string=[2]string{"hello","world"}
	// var slice3=arr[0:2]
	var slice4 = make([]string, 2, 3)
	//define slice4
	slice4[0] = "hello"
	slice4[1] = "world"
	slice4 = append(slice4, "yuta")
	slice4 = append(slice4, "yuta")
	return slice4
}

func dbConnect() {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected to the database")
}

func minMax(a []int) (int, int) {
	min := a[0]
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
		if a[i] > max {
			max = a[i]
		}
	}
	return min, max
}

func avg(a []int) float64 {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return float64(sum) / float64(len(a))
}

type Human struct {
	name string
}

func (h Human) introduce() string {
	return "I am " + h.name
}

func main() {
	engine := gin.Default()
	engine.ForwardedByClientIP = true
	engine.SetTrustedProxies([]string{"127.0.0.1"})
	dbConnect()
	engine.LoadHTMLGlob("views/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			// htmlに渡す変数を定義
			"message": "hello gin",
		})
	})
	engine.POST("/upload", func(c *gin.Context) {
        file,header, err :=  c.Request.FormFile("image")
        if err != nil {
            c.String(http.StatusBadRequest, "Bad request")
            return
        }
        fileName := header.Filename
        dir, _ := os.Getwd()
        out, err := os.Create(dir+"\\images\\"+fileName)
        if err != nil {
            log.Fatal(err)
        }
        defer out.Close()
        _, err = io.Copy(out, file)
        if err != nil {
            log.Fatal(err)
        }
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
        })
    })
	engine.Static("/static", "./static")

	engine.Run(":8080")
}
func printNumbers() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func printLetters() {
	for i := 'a'; i < 'a'+10; i++ {
		fmt.Println(string(i))
	}
}
