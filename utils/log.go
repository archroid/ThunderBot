package utils

import (
	log "github.com/sirupsen/logrus"
)

type LogKV struct {
	key   string
	value string
}

func LogInfo(m string, kvs []LogKV) {
	fields := log.Fields{}
	for _, kv := range kvs {
		fields[kv.key] = kv.value
	}
	log.WithFields(fields).Info(m)
}

func LogError(m string, kvs []LogKV) {
	fields := log.Fields{}
	for _, kv := range kvs {
		fields[kv.key] = kv.value
	}
	log.WithFields(fields).Error(m)
}

func KvForHandler(guild string, handler string, kvs []LogKV) []LogKV {
	kvs = append(kvs, LogKV{"guildID", guild}, LogKV{"event", handler})

	return kvs
}

func KvForEvent(event string, kvs []LogKV) []LogKV {
	kvs = append(kvs, LogKV{"event", event})

	return kvs
}

func KVs(kvs ...string) []LogKV {
	var res []LogKV
	for i := 0; i < len(kvs); i += 2 {
		res = append(res, LogKV{kvs[i], kvs[i+1]})
	}
	return res
}
