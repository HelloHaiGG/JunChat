package servers

import (
	common2 "JunChat/common"
	common "JunChat/common/discover"
	"JunChat/config"
	"JunChat/connect/models"
	core "JunChat/core/protocols"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

var HandleConn sync.Map
var NETServer string

type Connect struct {
	Uid         string
	Conn        *websocket.Conn
	ConnectTime int64
}

func NetConnect(port string) {
	http.HandleFunc("/jun_chat", HandleReq)
	err := http.ListenAndServe(net.JoinHostPort("", port), nil)
	if err != nil {
		log.Fatal("Http Err:", err)
	}
}

func HandleReq(w http.ResponseWriter, r *http.Request) {
	conn := &Connect{ConnectTime: time.Now().Unix()}
	userId, err := conn.upgrade(w, r)
	if err != nil || userId == "" {
		log.Println("Upgrade Err:", err)
		_, _ = io.WriteString(w, "链接失败!")
		return
	}
	// 汇报
	rpcConn := common.GetServerConn(config.APPConfig.Servers.Core)
	client := core.NewCenterServerClient(rpcConn)
	rsp, err := client.OnlineReport(context.Background(), &core.ReportOnlineParams{
		Id:       userId,
		ServerId: NETServer,
		Category: 0,
		Status:   common2.Online,
	})
	if err != nil {
		_ = conn.CloseConn()
		log.Println("Rpc Call Report Online Err:", err)
	}
	if rsp.Code != common2.Success {
		_ = conn.CloseConn()
		log.Println("Rpc Call Report Online Err:", rsp.Code)
		return
	}
	conn.Uid = userId
	//存储玩家链接
	HandleConn.Store(conn.Uid, conn)
	Send(userId)
	//监控,处理关闭请求
	go conn.HandlerConn()
	return
}

//http 升级为 websocket
func (p *Connect) upgrade(w http.ResponseWriter, r *http.Request) (string, error) {
	//解析TOKEN 验证用户信息
	protocols := websocket.Subprotocols(r)
	if len(protocols) < 1 {
		_, _ = io.WriteString(w, "Sub Protocol!")
		return "", nil
	}
	userId := protocols[0]
	//将http请求升级为websocket
	upgrade := &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}, Subprotocols: protocols}
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return "", err
	}
	p.Conn = conn
	return userId, nil
}

func (p *Connect) HandlerConn() {
	defer func() {
		_ = p.CloseConn()
	}()
	for {
		_, msg, err := p.Conn.ReadMessage()
		if err != nil {
			_ = p.CloseConn()
			return
		}
		body := &models.HeartBeat{}
		_ = json.Unmarshal(msg, body)
		fmt.Println("Heart-Beat:", body.Msg)
	}
}

//关闭链接,并将删除map内链接,并向Core层汇报
func (p *Connect) CloseConn() error {
	_ = p.Conn.Close()
	conn := common.GetServerConn(config.APPConfig.Servers.Core)
	client := core.NewCenterServerClient(conn)
	rsp, err := client.OnlineReport(context.Background(), &core.ReportOnlineParams{
		Id:       p.Uid,
		ServerId: NETServer,
		Category: 0,
		Status:   common2.Offline,
	})
	if err != nil {
		log.Println("Rpc Call Report Disconnect Err:", err)
	}
	if rsp.Code != common2.Success {
		return errors.New("Clear Connect Message Fail. ")
	}
	HandleConn.Delete(rsp.Id)
	return nil
}

func Send(uid string) {
	conn, ok := HandleConn.Load(uid)
	if !ok {
		log.Println("Send Message On Connect JunChat.")
		return
	}
	c, _ := conn.(*Connect)
	_ = c.Conn.WriteMessage(websocket.TextMessage, models.GetDailyMessage())
}
