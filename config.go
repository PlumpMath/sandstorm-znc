package main

import (
	"net"
	"os"
)

const (
	lo = "127.0.0.1"
)

var (
	zncPort       = os.Getenv("SANDSTORM_ZNC_PORT")
	ipNetworkPort = os.Getenv("SANDSTORM_IP_NETWORK_PORT")
	appDir        = os.Getenv("SANDSTORM_APP_DIR")

	zncAddr       = net.JoinHostPort(lo, zncPort)
	ipNetworkAddr = net.JoinHostPort(lo, ipNetworkPort)
)
