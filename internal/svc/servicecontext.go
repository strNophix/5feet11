package svc

import (
	"5feet11/internal/config"
	"5feet11/internal/db"

	"github.com/bwmarrin/snowflake"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config    config.Config
	DB        gocqlx.Session
	Snowflake *snowflake.Node
}

func NewServiceContext(c config.Config) *ServiceContext {
	cluster := gocql.NewCluster(c.ScyllaDB.Hosts...)

	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		logx.Error(err)
	}

	if err := db.Seed(session); err != nil {
		logx.Error(err)
	}
	logx.Info("Schema's are up to date!")

	node, err := snowflake.NewNode(1)
	if err != nil {
		logx.Error(err)
	}

	return &ServiceContext{
		Config:    c,
		DB:        session,
		Snowflake: node,
	}
}

func (svcCtx *ServiceContext) Close() {
	svcCtx.DB.Close()
}
