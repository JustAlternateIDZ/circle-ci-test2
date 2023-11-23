package main

import (
    "testing"
)

func TestSayGoodBye(t *testing.T) {
    if SayGoodBye(&Person{"Name", 0}) != "Good Bye Name" {
        t.Errorf("Error saying Good Bye")
    }
}
