package config

import "time"

type Conf struct {
	Server    Server    `yaml:"server"`
	AliYunOSS AliYunOSS `yaml:"aliYunOSS"`
}

type Server struct {
	Mode         string        `yaml:"mode"`
	Addr         int           `yaml:"addr"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

type AliYunOSS struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	BucketName      string `yaml:"bucketName"`
	BucketUrl       string `yaml:"bucketUrl"`
}
