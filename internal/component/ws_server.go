package component

import (
	"bytes"
	"compress/gzip"
	"github.com/gorilla/websocket"
	"io"
	"landlord/common/config"
	"landlord/common/constant"
	"landlord/common/token"
	"landlord/common/utils"
	"landlord/db"
	"log"
	"net/http"
	"sync"
	"time"
)

type UserConn struct {
	*websocket.Conn
	mu         sync.Mutex
	UserId     string
	IsCompress bool
	IsOnline   bool
}

var WS wsServer

type wsServer struct {
	Port       string //8080
	MaxConnNum int
	UpGrader   *websocket.Upgrader
	//UserConnMap map[string]*UserConn
	UserConnMap sync.Map
}

func Start() {
	WS.onInit()
	WS.run()
}

func (ws *wsServer) onInit() {
	conf := config.Config.Websocket
	ws.Port = conf.Port[0]
	ws.MaxConnNum = conf.MaxConnNum
	ws.UpGrader = &websocket.Upgrader{
		HandshakeTimeout: time.Duration(conf.HandshakeTimeOut) * time.Second,
		ReadBufferSize:   conf.MaxMsgLen,
		CheckOrigin:      func(r *http.Request) bool { return true },
	}
}

func (ws *wsServer) run() {
	http.HandleFunc("/ws", ws.Handler)
	err := http.ListenAndServe(":"+ws.Port, nil)
	if err != nil {
		panic("ws listening err:" + err.Error())
	}
}

func (ws *wsServer) Handler(w http.ResponseWriter, r *http.Request) {
	//query := r.URL.Query()

	if isPass, userId := ws.headerCheck(w, r); isPass {
		conn, err := ws.UpGrader.Upgrade(w, r, nil)
		if err == nil {
			uc := &UserConn{UserId: userId, Conn: conn, IsOnline: true}
			ws.UserConnMap.Store(userId, uc)
			go ws.readMsg(uc)

			//ticker := time.NewTicker(time.Duration(config.Config.Websocket.OnlineTimeOut) * time.Second)
			//go uc.Ping(ticker)
		}
	}
}

func (uc *UserConn) Ping(ticker *time.Ticker) {
	for {
		<-ticker.C
		uc.SetWriteDeadline(time.Now().Add(time.Duration(config.Config.Websocket.OnlineTimeOut) * time.Second))

		uc.mu.Lock()
		if err := uc.WriteMessage(websocket.PingMessage, nil); err != nil {
			log.Printf("ping error: %s\n", err.Error())
			uc.IsOnline = false
			//todo
			Push("ExitRoom", &db.User{Id: uc.UserId})
			//biz.ExitRoom(&db.User{Id: uc.UserId})
			uc.mu.Unlock()
			break
		} else {
			uc.mu.Unlock()
		}

	}
}

func (ws *wsServer) headerCheck(w http.ResponseWriter, r *http.Request) (isPass bool, userId string) {
	status := http.StatusUnauthorized
	query := r.URL.Query()
	if len(query["token"]) != 0 {
		if ok, userId := token.GetUserIdFromToken(query["token"][0]); !ok {
			w.Header().Set("Sec-Websocket-Version", "13")
			w.Header().Set("ws_err_msg", "decode token failure")
			http.Error(w, "error token", status)
			return false, ""
		} else {
			return true, userId
		}
	} else {
		status = int(constant.ErrArgs.ErrCode)
		w.Header().Set("Sec-Websocket-Version", "13")
		errMsg := "args err, need token, userId"
		w.Header().Set("ws_err_msg", errMsg)
		http.Error(w, errMsg, status)
		return false, ""
	}
}

func (ws *wsServer) readMsg(conn *UserConn) {
	for {
		messageType, msg, err := conn.ReadMessage()

		if err != nil {
			//ws.UserConnMap.Delete(conn.UserId)
			return
		}

		switch messageType {
		case 0:

		}

		if conn.IsCompress {
			buff := bytes.NewBuffer(msg)
			reader, err := gzip.NewReader(buff)
			if err != nil {

				continue
			}
			msg, err = io.ReadAll(reader)
			if err != nil {

				continue
			}
			err = reader.Close()
			if err != nil {

			}
		}

		println(msg)
		//ws.msgParse(conn, msg)
	}
}

func (ws *wsServer) writeMsg(conn *UserConn, msgType int, msg []byte) error {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	if conn.IsCompress {
		var buffer bytes.Buffer
		gz := gzip.NewWriter(&buffer)
		if _, err := gz.Write(msg); err != nil {
			return utils.Wrap(err, "")
		}
		if err := gz.Close(); err != nil {
			return utils.Wrap(err, "")
		}
		msg = buffer.Bytes()
	}
	conn.SetWriteDeadline(time.Now().Add(time.Duration(60) * time.Second))
	return conn.WriteMessage(msgType, msg)
}

func (ws *wsServer) IsOnline(userId string) bool {
	if value, ok := ws.UserConnMap.Load(userId); !ok {
		return false
	} else {
		conn := value.(*UserConn)
		conn.mu.Lock()
		defer conn.mu.Unlock()
		return true
		//if conn.IsOnline {
		//	return true
		//} else {
		//	conn = nil
		//	ws.UserConnMap.Delete(userId)
		//	return false
		//}
	}
	return true
}

func (ws *wsServer) Send2Users(userIds []string, content string) bool {
	res := true

	for _, id := range userIds {
		res = res && ws.Send2User(id, content)
	}

	return res
}

func (ws *wsServer) Send2User(userId string, content string) bool {
	if value, ok := ws.UserConnMap.Load(userId); ok {
		conn := value.(*UserConn)
		if !conn.IsOnline {
			log.Printf("玩家不在线userId:%s\n", userId)
		} else {
			log.Printf("websocket:%s\n", content)
			err := conn.WriteMessage(websocket.TextMessage, []byte(content))
			if err != nil {
				log.Printf("消息推送异常userId:%s,err:%s\n", userId, err.Error())
			} else {
				return true
			}
		}
	}

	return false

}

func (ws *wsServer) Send2AllUser(content string) {
	ws.UserConnMap.Range(func(key, value any) bool {
		conn := value.(*UserConn)
		if !conn.IsOnline {
			log.Println("玩家不在线")
		} else {
			err := conn.WriteMessage(websocket.TextMessage, []byte(content))
			if err != nil {
				log.Printf("消息推送异常userId:%s,err:%s\n", key, err.Error())
			}
		}
		return true
	})
}
