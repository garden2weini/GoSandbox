package main

/**
# 开发单文件执行
>nohup go run big-screen.go &
# 开发多文件执行
>go run *.go
# 打包执行（打包名称任意）
>go build -o "bigscreen"
>./bigscreen
*/
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DataV中的列表显示，可以显示多列
type Data4List struct {
	Key        string `json:"key"`
	Attribute1 string `json:"attribute1"`
	Attribute2 string `json:"attribute2"`
}

// DataV排名展示的数据结构，Value用于数值排序，Content展示对应内容
type Data4Sort struct {
	Value   string `json:"value"`
	Content string `json:"content"`
}

type Data4Axis struct {
	X string `json:"x"`
	Y string `json:"y"`
	Z string `json:"z"`
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// curl -l http://localhost:4200/datav/top5device
	// curl -l http://localhost:4200/datav/hotsku
	// curl -l http://localhost:4200/datav/grossProfitX
	s := r.PathPrefix("/datav/").Subrouter()
	s.HandleFunc("/top5device", getTop5device).Methods("GET")
	s.HandleFunc("/hotsku", getHotSKU).Methods("GET")
	s.HandleFunc("/grossProfitX", getGrossProfitX).Methods("GET")
	s.HandleFunc("/deviceDistribute", getDeviceDistribute).Methods("GET")

	s.HandleFunc("/todo", getTODO).Methods("GET")

	r.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":4200", r))

}

// 设备量分布
func getDeviceDistribute(w http.ResponseWriter, r *http.Request) {
	var areas = [...]string{"朝阳", "海淀", "东城", "西城", "丰台", "崇文", "宣武", "门头沟"}

	fmt.Println("calling DeviceDistribute handler")
	w.Header().Set("Content-Type", "application/json")
	var datas = []Data4Axis{}

	for i := 0; i < len(areas); i++ {
		var x = areas[i]
		var y = rand.Intn(100)
		data1 := Data4Axis{X: x, Y: strconv.Itoa(y)}
		datas = append(datas, data1)
	}

	json.NewEncoder(w).Encode(datas)
}

// TOP毛利*周转SKU(排序展示)
func getGrossProfitX(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling GrossProfitX handler")
	w.Header().Set("Content-Type", "application/json")
	var datas = []Data4Sort{}

	for i := 0; i < 10; i++ {
		var value = rand.Intn(1000)
		var content = strconv.Itoa(value) + ": 红牛饮料 500ml"
		data1 := Data4Sort{Value: strconv.Itoa(value), Content: content}
		datas = append(datas, data1)
	}

	json.NewEncoder(w).Encode(datas)
}

// 近30天 热销SKU(排序展示)
func getHotSKU(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling HotSKU handler")
	w.Header().Set("Content-Type", "application/json")
	var datas = []Data4Sort{}

	for i := 0; i < 10; i++ {
		var value = rand.Intn(1000)
		var content = strconv.Itoa(value) + ": ¥3213/1200笔"
		data1 := Data4Sort{Value: strconv.Itoa(value), Content: content}
		datas = append(datas, data1)
	}

	json.NewEncoder(w).Encode(datas)
}

// 近30天 TOPn设备
func getTop5device(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling Top5device handler")
	w.Header().Set("Content-Type", "application/json")
	var datas = []Data4List{}

	datas = append(datas, Data4List{Key: "100", Attribute1: "¥1213/1200笔"})
	for i := 0; i < 10; i++ {
		tmp := "点位地址" + strconv.Itoa(i)
		data1 := Data4List{Key: tmp, Attribute1: "¥3213/1200笔"}
		datas = append(datas, data1)
	}

	json.NewEncoder(w).Encode(datas)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
