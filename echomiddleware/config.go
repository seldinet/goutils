package echomiddleware

import "github.com/devarchi33/goutils/kafka"

type KafkaConfig = kafka.Config

type ZipkinConfig struct {
	Collector struct {
		Url string
	}
	Addr, Service string
	Kafka         KafkaConfig
}
