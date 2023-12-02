package database

import (
	"github.com/andrepriyanto10/server_favaa/configs/env"
	log "github.com/andrepriyanto10/server_favaa/configs/logger"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestConfigConn_Open(t *testing.T) {
	type fields struct {
		env *viper.Viper
		log *log.Log
	}
	tests := []struct {
		name   string
		fields fields
		want   *gorm.DB
	}{
		{
			name: "Test Open Connection",
			fields: fields{
				env: env.LoadEnv("config", "../../"),
				log: log.NewLogger(),
			},
			want: &gorm.DB{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConnection(tt.fields.env, tt.fields.log)
			got := c.Open()

			assert.Equal(t, reflect.TypeOf(tt.want), reflect.TypeOf(got))
		})
	}
}
