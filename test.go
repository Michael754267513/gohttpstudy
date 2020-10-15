package main

import (
	"fmt"
	"strconv"
)
import apiv1 "k8s.io/api/core/v1"

type RedisMeta struct {
	Port int
	Name string
}

func AddEnvString(name string, value string) apiv1.EnvVar {
	var env apiv1.EnvVar
	env.Name = name
	env.Value = value
	return env
}

func main() {
	var env []apiv1.EnvVar
	nodeList := []RedisMeta{{Port: 23, Name: "aa"}, {Port: 33, Name: "cc"}, {Port: 222, Name: "waaa"}}
	for index, nodev := range nodeList {
		index += 1
		if nodev.Name == "waaa" {
			env = append(env, AddEnvString("Myid", strconv.Itoa(index)))
		}
		env = append(env, AddEnvString("node"+strconv.Itoa(index), nodev.Name))
		env = append(env, AddEnvString("port"+strconv.Itoa(index), strconv.Itoa(nodev.Port)))
	}
	fmt.Println(env)
}
