package service

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/fuwensun/goms/eMysql/internal/dao"
	"github.com/fuwensun/goms/pkg/conf"
)

// Service.
type Service struct {
	cfg *config
	dao *dao.Dao
}

// config.
type config struct {
	Name    string `yaml:"name,omitempty"`
	Version string `yaml:"version,omitempty"`
}

// getConfig
func getConfig(cfgpath string) (*config, error) {
	cfg := &config{}
	filep := filepath.Join(cfgpath, "app.yml")
	if err := conf.GetConf(filep, cfg); err != nil {
		log.Printf("get config file: %v", err)
		err = fmt.Errorf("get config: %w", err)
		return cfg, err
	}
	log.Printf("config name: %v,version: %v", cfg.Name, cfg.Version)
	return cfg, nil
}

// New.
func New(cfgpath string, d *dao.Dao) *Service {
	cfg, err := getConfig(cfgpath)
	if err != nil {
		log.Panicf("failed to get config file: %v", err)
	}
	return &Service{
		cfg: cfg,
		dao: d,
	}
}

// Ping.
func (s *Service) Ping(c context.Context) (err error) {
	return s.dao.Ping(c)
}

// Close.
func (s *Service) Close() {
	s.dao.Close()
}
