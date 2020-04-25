package servers

import (
	"JunChat/config"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

var HandleConn sync.Map

type Connect struct {
	Uid         string
	Conn        *websocket.Conn
	ConnectTime int64
}

func NetConnect() {
	http.HandleFunc("/jun_chat", HandleReq)
	log.Println(net.JoinHostPort("", config.APPConfig.JC.Port))
	err := http.ListenAndServe(net.JoinHostPort("", config.APPConfig.JC.Port), nil)
	if err != nil {
		log.Fatal("Http Err:", err)
	}
}

func HandleReq(w http.ResponseWriter, r *http.Request) {
	conn := &Connect{ConnectTime: time.Now().Unix(), Uid: "XXX"}
	//TODO 解析TOKEN 验证玩家信息
	err := conn.upgrade(w, r)
	if err != nil {
		_, _ = w.Write([]byte("Fail"))
	}
	//存储玩家链接
	HandleConn.Store(conn.Uid, conn)
	//TODO 向Core汇报所在服务

	//1.玩家所在服务
	//2.聊天室所在服务
}



func HandleMsg()  {

}

//http 升级为 websocket
func (p *Connect) upgrade(w http.ResponseWriter, r *http.Request) error {
	//将http请求升级为websocket
	upgrade := &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	p.Conn = conn
	return nil
}
