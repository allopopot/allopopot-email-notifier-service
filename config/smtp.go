package config

import "strconv"

var SMTP_HOST = ParseEnv("SMTP_HOST", "", false)
var SMTP_PORT, _ = strconv.Atoi(ParseEnv("SMTP_PORT", "", false))
var SMTP_USERNAME = ParseEnv("SMTP_USERNAME", "", false)
var SMTP_PASSWORD = ParseEnv("SMTP_PASSWORD", "", false)
var SMTP_SENDER = ParseEnv("SMTP_SENDER", "", false)
