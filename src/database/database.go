package database

import (
	"ControlCenter-Server/src/config"
	"ControlCenter-Server/src/model"
	"fmt"

	//"gorm.io/driver/sqlite"
	//"gorm.io/gorm"
	//"gorm.io/gorm/logger"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// var mutex sync.Mutex

func initDatabase() *gorm.DB {
	conf := config.LoadConfig()
	var err error
	var db *gorm.DB
	db, err = gorm.Open("sqlite3", conf.Database)
	if conf.Debug {
		/*db, err = gorm.Open(sqlite.Open(conf.Database), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold: time.Second, // 慢 SQL 阈值
					LogLevel:      logger.Info, // Log level
					Colorful:      false,       // 禁用彩色打印
				},
			),
		})*/
		db.LogMode(true)
	} else {
		/*db, err = gorm.Open(sqlite.Open(conf.Database), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					LogLevel: logger.Silent,
				}),
		})*/
		db.LogMode(false)
	}
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return db
}

//Server
func AddServer(server model.Server) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return false
	}

	_ = tx.AutoMigrate(&model.Server{})
	if err := tx.Create(&server).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func UpdateServer(where model.Server, server model.Server) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		fmt.Println(tx.Error)
		return false
	}
	_ = tx.AutoMigrate(&model.Server{})
	//if err := tx.Model(&where).Updates(server).Error; err != nil {
	if err := tx.Model(&model.Server{}).Where(where).Update(server).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func GetServer(server model.Server) []model.Server {
	db := initDatabase()
	defer db.Close()
	_ = db.AutoMigrate(&model.Server{})
	servers := []model.Server{}
	db.Where(server).Find(&servers)
	return servers
}

func DelServer(id int64, uid int64) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.Server{})
	server := model.Server{ID: id, UID: uid}
	if err := tx.Where(server).Delete(model.Server{}).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//User
func AddUser(user model.User) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.User{})
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func UpdateUser(where model.User, user model.User) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.User{})
	if err := tx.Model(&model.User{}).Where(where).Update(user).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func GetUser(user model.User) []model.User {
	db := initDatabase()
	defer db.Close()
	_ = db.AutoMigrate(&model.User{})
	users := []model.User{}
	db.Where(user).Find(&users)
	return users
}

func DelUser(user model.User) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.User{})
	if err := tx.Where(user).Delete(model.User{}).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//Cer
func AddCer(certificate model.Certificate) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.Certificate{})
	if err := tx.Create(&certificate).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func DelCer(certificate model.Certificate) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.Certificate{})
	if err := tx.Where(certificate).Delete(model.Certificate{}).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func UpdateCer(where model.Certificate, certificate model.Certificate) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.Certificate{})
	if err := tx.Model(&model.Certificate{}).Where(where).Update(certificate).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func GetCer(certificate model.Certificate) []model.Certificate {
	db := initDatabase()
	defer db.Close()
	_ = db.AutoMigrate(&model.Certificate{})
	certificates := []model.Certificate{}
	db.Where(certificate).Find(&certificates)
	return certificates
}

// LinkCer link
func LinkServer(serverLink model.ServerLink) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.ServerLink{})
	if err := tx.Create(&serverLink).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func GetServerLinkedItem(serverLink model.ServerLink) []model.ServerLink {
	db := initDatabase()
	defer db.Close()
	_ = db.AutoMigrate(&model.ServerLink{})
	ServerLinks := []model.ServerLink{}
	db.Where(serverLink).Find(&ServerLinks)
	return ServerLinks
}

func UnLinkServer(serverLink model.ServerLink) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.ServerLink{})
	if err := tx.Where(serverLink).Delete(model.ServerLink{}).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//System Config
func SetConfig(config model.SysConfig) bool {
	if len(GetConfig(model.SysConfig{Name: config.Name})) == 0 {
		return addConfig(config)
	}
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.SysConfig{})
	if err := tx.Model(&model.SysConfig{}).Where(model.SysConfig{Name: config.Name}).Update(config).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func addConfig(config model.SysConfig) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.SysConfig{})
	if err := tx.Create(&config).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func GetConfig(config model.SysConfig) []model.SysConfig {
	db := initDatabase()
	defer db.Close()
	_ = db.AutoMigrate(&model.SysConfig{})
	configs := []model.SysConfig{}
	db.Where(config).Find(&configs)
	return configs
}

func AddLog(service string, event string, level int64) bool {
	newLog := model.LogInfo{Service: service, Info: event, Level: level, CreatedAt: time.Now()}
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.LogInfo{})
	if err := tx.Create(&newLog).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func AddEvent(eventType int64, target int64, code int64, info string) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.Event{})
	if len(GetEvent(eventType, target, code, info, 0)) != 0 {
		if err := tx.Model(&model.Event{}).Where(model.Event{Type: eventType, TargetID: target, Code: code, Info: info}).Update(model.Event{Active: 1}).Error; err != nil {
			tx.Rollback()
			return false
		}
	} else {
		if err := tx.Create(&model.Event{Type: eventType, TargetID: target, Code: code, Info: info, Active: 1}).Error; err != nil {
			tx.Rollback()
			panic(err)
			return false
		}
	}
	tx.Commit()
	return true
}

func GetEvent(eventType int64, target int64, code int64, info string, active int64) []model.Event {
	db := initDatabase()
	defer db.Close()
	_ = db.AutoMigrate(&model.Event{})
	events := []model.Event{}
	db.Where(model.Event{Type: eventType, TargetID: target, Code: code, Info: info, Active: active}).Find(&events)
	return events
}

func FinishEvent(id int64) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.Event{})
	event := model.Event{Active: 2}
	if err := tx.Model(&model.Event{}).Where(model.Event{ID: id, Active: 1}).Update(event).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

//Configuration
func AddConfiguration(conf model.Configuration) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.Configuration{})
	if err := tx.Create(&conf).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func GetConfiguration(conf model.Configuration) []model.Configuration {
	db := initDatabase()
	defer db.Close()
	_ = db.AutoMigrate(&model.Configuration{})
	confs := []model.Configuration{}
	db.Where(conf).Find(&confs)
	return confs
}

func DeleteConfiguration(conf model.Configuration) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.Configuration{})
	if err := tx.Where(conf).Delete(model.Configuration{}).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func UpdateConfiguration(conf model.Configuration, where model.Configuration) bool {
	db := initDatabase()
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false
	}
	_ = tx.AutoMigrate(&model.Configuration{})
	if err := tx.Model(&model.Configuration{}).Where(where).Update(conf).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}
