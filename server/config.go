package server

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port     int
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func NewConfig(path string) (Config, error) {
	// Инициализируем viper для работы с конфигурационным файлом
	viper.SetConfigName("config") // Имя файла без расширения
	viper.SetConfigType("yaml")   // Тип файла (yaml, json, toml и т. д.)
	viper.AddConfigPath(path)     // Путь к файлу конфигурации (текущая директория)

	// Прочитаем конфигурационный файл
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения конфигурационного файла: %s", err)
	}

	// Создадим экземпляр структуры Config для хранения настроек
	var cfg Config

	// Прочитаем настройки из конфигурационного файла
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Ошибка разбора конфигурации: %s", err)
	}

	return cfg, cfg.valid()
}

func (c Config) valid() error {
	if c.Port == 0 {
		return errors.New("empty PORT")
	}
	if emptyString(c.Database.Host) {
		return errors.New("empty DB.Host")
	}
	if emptyString(c.Database.Name) {
		return errors.New("empty DB.Name")
	}
	if emptyString(c.Database.Password) {
		return errors.New("empty DB.Password")
	}
	if emptyString(c.Database.Username) {
		return errors.New("empty DB.Username")
	}
	if c.Database.Port == 0 {
		return errors.New("empty DB.Port")
	}
	return nil
}

func emptyString(s string) bool {
	if len(strings.TrimSpace(s)) == 0 {
		return true
	}
	return false
}

func (d DatabaseConfig) ConnStr() string {
	return fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable pool_max_conns=10",
		d.Port,
		d.Host,
		d.Username,
		d.Password,
		d.Name)
}

func (c Config) ToString() string {
	return fmt.Sprintf(":%d", c.Port)
}
