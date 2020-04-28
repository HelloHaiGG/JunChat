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
		_, _ = io.WriteString(w, "链接失败!")
		return
	}
	conn.Uid = userId
	//存储玩家链接
	HandleConn.Store(conn.Uid, conn)
	return
}

//http 升级为 websocket
func (p *Connect) upgrade(w http.ResponseWriter, r *http.Request) (string, error) {
	//解析TOKEN 验证用户信息
	protocols := websocket.Subprotocols(r)
	if len(protocols) < 2 {
		_, _ = io.WriteString(w, "Sub Protocol!")
		return "",nil
	}
	serverId := protocols[0]
	userId := protocols[1]
	//判断用户时候连到对的节点
	if NETServer != serverId {
		_, _ = io.WriteString(w, "节点错误!")
		return "", nil
	}

	//将http请求升级为websocket
	upgrade := &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	},Subprotocols:protocols}
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return "", err
	}
	p.Conn = conn
	return userId, nil
}

//监听socket链接心跳
func (p *Connect) HandlerConn() {
	defer func() {
		_ = p.CloseConn()
	}()
	t := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-t.C:
			//掉线处理
			_ = p.CloseConn()
			return
		default:
			_, msg, err := p.Conn.ReadMessage()
			if err != nil {
				_ = p.CloseConn()
			}
			body := &models.HeartBeat{}
			_ = json.Unmarshal(msg, body)
			fmt.Println("Heart-Beat:", body.Msg)
		}
	}
}

//关闭链接,并将删除map内链接,并向Core层汇报
func (p *Connect) CloseConn() error {
	_ = p.Conn.Close()

	conn := common.GetServerConn(config.APPConfig.Servers.Core)
	client := core.NewCenterServerClient(conn)
	rsp, err := client.OnDisconnectReport(context.Background(), &core.ReportDisconnectParams{
		Id:       p.Uid,
		Category: 0,
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
