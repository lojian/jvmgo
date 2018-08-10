package main

import (
	"flag"
	"fmt"
	"testing"
)

var (
	hanahost = flag.Bool("hanahost-aws", false, "request a HANA host on AWS")
	kvstore  = flag.Bool("kvstore", false, "request setup KV store")
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain")
	flag.Parse()

	if *hanahost {
		fmt.Println("HANA Host is requested, setup...")
	}

	if *kvstore {
		fmt.Println("KV Store is requested, setup...")
	}
}

func TestMethod1(t *testing.T) {
	fmt.Println("Tested")
}
