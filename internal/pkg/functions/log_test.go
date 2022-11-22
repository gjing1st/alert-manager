// Path: internal/pkg/functions
// FileName: log_test.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/17$ 16:03$

package functions

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestAddLog(t *testing.T) {
	AddLog(log.Fields{"test": "1111111"})

}
