// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package redis

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/go-vela/types"
)

func TestRedis_Push(t *testing.T) {
	// setup types
	// use global variables in redis_test.go
	_item := &types.Item{
		Build:    _build,
		Pipeline: _steps,
		Repo:     _repo,
		User:     _user,
	}

	// setup queue item
	_bytes, err := json.Marshal(_item)
	if err != nil {
		t.Errorf("unable to marshal queue item: %v", err)
	}

	// setup redis mock
	_redis, err := NewTest("vela")
	if err != nil {
		t.Errorf("unable to create queue service: %v", err)
	}

	// setup redis mock
	badItem, err := NewTest("vela")
	if err != nil {
		t.Errorf("unable to create queue service: %v", err)
	}

	// setup tests
	tests := []struct {
		failure bool
		redis   *client
		bytes   []byte
	}{
		{
			failure: false,
			redis:   _redis,
			bytes:   _bytes,
		},
		{
			failure: true,
			redis:   badItem,
			bytes:   nil,
		},
	}

	// run tests
	for _, test := range tests {
		err := test.redis.Push(context.Background(), "vela", test.bytes)

		if test.failure {
			if err == nil {
				t.Errorf("Push should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Push returned err: %v", err)
		}
	}
}
