package main

import "testing"

func TestLoadConfig(t *testing.T) {
	t.Log(loadConfig())
}

func TestGetData(t *testing.T) {
	t.Log(getData())
}
