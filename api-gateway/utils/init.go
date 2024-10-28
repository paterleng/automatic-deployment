package utils

import (
	"api-gateway/model"
	"database/sql"
	"fmt"
	"github.com/fsnotify/fsnotify"
	_ "github.com/go-sql-driver/mysql"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Utils struct {
	LG *zap.Logger
	DB *gorm.DB
	PB *Pb
}

type Config struct {
	*ProjectConfig `mapstructure:"project"`
	*LogConfig     `mapstructure:"log"`
	*MySQLConfig   `mapstructure:"mysql"`
	*Etcd          `mapstructure:"etcd"`
}

type Etcd struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type ProjectConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Address   string `mapstructure:"address"`
	Port      string `mapstructure:"port"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

var Tools Utils
var Conf = new(Config)

func init() {
	NewTools()
	if err := ViperInit(); err != nil {
		fmt.Errorf("初始化viper失败：", err)
		return
	}

	if err := LoggerInit(); err != nil {
		fmt.Errorf("初始化日志对象失败：", err)
		return
	}
	Tools.LG.Info("初始化logger成功")

	if err := MysqlInit(); err != nil {
		Tools.LG.Error("初始化MySQL失败：", zap.Error(err))
		panic(err)
	}
	if err := MysqlInit(); err != nil {
		Tools.LG.Error("初始化MySQL失败：", zap.Error(err))
		return
	}
	Tools.LG.Info("初始化mysql成功")
	//创建表
	TableInit()
	DiscoveryService()
}
func NewTools() {
	Tools.PB = &Pb{}
}

func TableInit() (err error) {
	err = Tools.DB.AutoMigrate(&model.Secret{})
	err = Tools.DB.AutoMigrate(&model.Cluster{})
	err = Tools.DB.AutoMigrate(&model.Node{})
	return
}

func MysqlInit() (err error) {
	//判断库是否存在不存在就创建
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", Conf.MySQLConfig.User, Conf.MySQLConfig.Password, Conf.MySQLConfig.Host, Conf.MySQLConfig.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '%s'", Conf.MySQLConfig.DB)
	var name string
	err = db.QueryRow(query).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			//数据库不存在创建数据库
			err = nil
			_, err = db.Exec("CREATE DATABASE " + Conf.MySQLConfig.DB)
			if err != nil {
				return
			}
		}
		return
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", Conf.MySQLConfig.User, Conf.MySQLConfig.Password, Conf.MySQLConfig.Host, Conf.MySQLConfig.Port, Conf.MySQLConfig.DB)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	Tools.DB, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return
	} else {
		sqlDB, _ := Tools.DB.DB()
		sqlDB.SetMaxOpenConns(Conf.MySQLConfig.MaxOpenConns)
		sqlDB.SetMaxIdleConns(Conf.MySQLConfig.MaxIdleConns)
	}
	return
}

func ViperInit() error {
	viper.SetConfigFile("./config.yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("修改了配置文件...")
		viper.Unmarshal(&Conf)
	})
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("readconfig failed,err: %v", err)
		return err
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		fmt.Errorf("unmarshal to Conf failed, err: %v", err)
		return err
	}
	return err
}

func LoggerInit() (err error) {
	writeSyncer := getLogWriter(Conf.LogConfig.Filename, Conf.LogConfig.MaxSize, Conf.LogConfig.MaxBackups, Conf.LogConfig.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(Conf.LogConfig.Level))
	if err != nil {
		return
	}
	var core zapcore.Core
	if Conf.ProjectConfig.Mode == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	Tools.LG = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(Tools.LG)
	zap.L().Info("init logger success")
	return

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
