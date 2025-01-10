package config

var AMQP_HOST = ParseEnv("AMQP_HOST", "", false)
var AMQP_PORT = ParseEnv("AMQP_PORT", "5672", true)
var AMQP_USERNAME = ParseEnv("AMQP_USERNAME", "", false)
var AMQP_PASSWORD = ParseEnv("AMQP_PASSWORD", "", false)

var AMQP_EXCHANGE_NAME = ParseEnv("AMQP_EXCHANGE_NAME", "EMAIL-SERVICE-EXCHANGE", true)
