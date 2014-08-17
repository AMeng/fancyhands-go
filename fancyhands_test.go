package fancyhands

import (
    "testing"
)

func assertEq(t *testing.T, lhs interface{}, rhs interface{}) {
    if lhs != rhs {
        t.Error(lhs, " != ", rhs)
    }
}

func TestNewClientCreation(t *testing.T) {
    client := NewClient("key", "secret")

    assertEq(t, client.test, false)
}

func TestNewTestClientCreation(t *testing.T) {
    client := NewTestClient("key", "secret")

    assertEq(t, client.test, true)
}
