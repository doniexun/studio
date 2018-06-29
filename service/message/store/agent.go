// Copyright 2018. All rights reserved.
// This file is part of duguying project
// Created by duguying on 2018/6/29.

package store

import (
	"encoding/json"
	"log"
	"time"
)

type AgentInfo struct {
	Online      bool      `json:"online"`
	ClientID    string    `json:"client_id"`
	OnlineTime  time.Time `json:"online_time"`
	OfflineTime time.Time `json:"offline_time"`
}

func (ai *AgentInfo) String() string {
	c, _ := json.Marshal(ai)
	return string(c)
}

func PutAgent(clientId string, info *AgentInfo) error {
	value := info.String()
	return put("agent", clientId, []byte(value))
}

func ListAllAgent() (list []*AgentInfo, err error) {
	tx, err := boltDB.Begin(true)
	if err != nil {
		return nil, err
	}

	bkt := tx.Bucket([]byte("agent"))

	c := bkt.Cursor()
	list = []*AgentInfo{}

	for k, v := c.First(); k != nil; k, v = c.Next() {
		info := &AgentInfo{}
		err := json.Unmarshal(v, info)
		if err != nil {
			log.Println("marshal agent info failed, err:", err.Error())
		} else {
			list = append(list, info)
		}
	}

	return list, tx.Commit()
}
