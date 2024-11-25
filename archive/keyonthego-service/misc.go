package main

import "fmt"

var Port int = 1233
var Banner = fmt.Sprintf(`
                  _   _           
 _ __   ___  _ __| |_(_) ___ _ __ 
| '_ \ / _ \| '__| __| |/ _ \ '__|
| |_) | (_) | |  | |_| |  __/ |   
| .__/ \___/|_|   \__|_|\___|_|   
|_|                               Vision %s - KeyOnTheGo Service
`, Version)
var NetworkInterfaces []NetworkInterface
