package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/alex-laycalvert/chattn/internal/client"
	"github.com/alex-laycalvert/chattn/internal/server"
)

const (
	MODE_SERVER  = iota
	MODE_CLIENT  = iota
	DEFAULT_PORT = 8889
	SERVER_FLAG  = "--server"
)

type Config struct {
	// Either MODE_SERVER or MODE_CLIENT
	mode uint8
	// Host to connect to if in MODE_CLIENT
	host string
	// Port to use for either MODE_SERVER or MODE_CLIENT
	port int
}

func main() {
	args := os.Args
	config, err := parseArgs(args)
	if err != nil {
		usage()
		fmt.Printf("\nerror: %s\n", err.Error())
		os.Exit(1)
	}
	if config.mode == MODE_SERVER {
		server := server.New(config.port)
		if err := server.Start(); err != nil {
			fmt.Printf("error: %s", err.Error())
			os.Exit(1)
		}
	} else if config.mode == MODE_CLIENT {
		client := client.New(config.host, config.port)
		if err := client.Connect(); err != nil {
			fmt.Printf("error: %s", err.Error())
			os.Exit(1)
		}
	} else {
		fmt.Printf("error: invalid mode, must be either MODE_SERVER or MODE_CLIENT")
		os.Exit(1)
	}
}

func parseArgs(args []string) (Config, error) {
	config := Config{mode: MODE_CLIENT, port: 8889}
	length := len(args)
	if length < 2 {
		return config, errors.New("not enough arguments")
	}
	if length > 3 {
		return config, errors.New("too many arguments")
	}
	if args[1] == SERVER_FLAG {
		config.mode = MODE_SERVER
	} else {
		config.host = args[1]
	}
	if length > 2 {
		port, err := strconv.Atoi(args[2])
		if err != nil {
			return config, errors.New("failed to parse port - " + err.Error())
		}
		if port < 0 {
			return config, errors.New(fmt.Sprintf("invalid port %d, cannot be less than or equal to 0", port))
		}
		if port > 65535 {
			return config, errors.New(fmt.Sprintf("invalid port %d, cannot be greater than 65535", port))
		}
		config.port = port
	}
	return config, nil
}

func usage() {
	fmt.Println("usage:")
	fmt.Println("\tchattn <host>[:<port>]")
	fmt.Println("\tchattn --server [<port>]")
	fmt.Println("")
	fmt.Println("<host>\t\tthe chattn server to connect to (only in client mode)")
	fmt.Println("<port>\t\tthe port to connect to chattn on (default 8889)")
	fmt.Println("--server, -s\tif provided, will start a chattn server on the given")
	fmt.Println("\t\tport (default 8889)")
}
