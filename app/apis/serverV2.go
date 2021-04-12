package apis

import (
	"ControlCenter-Server/app/database"
	"ControlCenter-Server/app/model"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var upgrader = websocket.Upgrader{}

func ServerV2(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	lock := sync.WaitGroup{}
	done := make(chan int64)
	defer func() {
		done <- 1
		lock.Wait()
		ws.Close()
	}()
	_, msg, err := ws.ReadMessage()
	if err != nil {
		c.Logger().Error(err)
	}
	token := string(msg)
	server := model.Server{Token: token}
	serverList := database.GetServer(server)
	if len(serverList) != 0 {
		server = serverList[0]
	} else {
		ws.Close()
		return nil
	}
	log.Println("[Websocket]", server.Ipv4, "connected")
	go func() {
		lock.Add(1)
		for {
			select {
			case <-done:
				break
			default:
				_, msg, err := ws.ReadMessage()
				if err != nil {
					log.Println("[Websocket]", server.Ipv4, "Disconnected")
					lock.Done()
					return
				}
				if strings.Contains(string(msg), "pushStatus#") {
					status := strings.Replace(string(msg), "pushStatus#", "", 1)
					database.UpdateServer(model.Server{Token: token}, model.Server{Status: status})
				}
				if strings.Contains(string(msg), "psStatus#") {
					status := strings.Replace(string(msg), "psStatus#", "", 1)
					database.UpdateServer(model.Server{Token: token}, model.Server{Ps: status})
				}
			}
		}
		lock.Done()
	}()
	lock.Wait()
	return nil
}

func OpenTerminal(c echo.Context) error {
	user := CheckAuth(c)
	if user.Level <= 1 {
		action, _ := strconv.ParseInt(c.Param("action"), 10, 64)
		serverID, _ := strconv.ParseInt(c.Param("serverid"), 10, 64)
		//fmt.Println(database.AddEvent(1, serverID, action, "OK"))
		if len(database.GetServer(model.Server{ID: serverID, UID: database.GetUser(model.User{Mail: user.Mail})[0].ID})) == 1 {
			d := uuid.New().String()
			if database.AddEvent(1, serverID, action, d) == false {
				log.Print("AddEvent Fail:" + c.Path())
				database.AddLog("Event", c.Path()+"|From:"+c.RealIP(), 2)
				return c.JSON(http.StatusOK, model.Callback{Code: 500, Info: "Internal Server Error"})
			}
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: d})
		}
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}