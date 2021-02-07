package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("20060102", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	if err != nil {
		return
	}
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}

//func main(){
//	if err:=Init("20210124",1);err!=nil{
//		fmt.Printf("init failed,err:=%v\n",err)
//		return
//	}
//	id:=GenID()
//	var t time.Time
//
//	fmt.Println(id,t.UnixNano())
//}
