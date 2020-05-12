package main

import (
	"log"
	"testing"
)

func TestPush(t *testing.T) {
	sq := &StringQueue{}
	expected := "foobar"
	sq.Push(expected)
	got, err := sq.Front()
	if err != nil {
		log.Fatal(err)
	}
	if got != expected {
		log.Fatalf("Front() returned %s, expected %s", got, expected)
	}
}

func TestFrontEmpty(t *testing.T) {
	sq := &StringQueue{}
	_, err := sq.Front()
	if err == nil {
		log.Fatal("Front() on an empty queue does not error")
	}
}

func TestFrontNonEmpty(t *testing.T) {
	sq := &StringQueue{}
	actual := "foobar"
	sq.Push(actual)
	got, err := sq.Front()
	if err != nil {
		log.Fatal(err)
	}
	if got != actual {
		log.Fatalf("Front() returned %s, expected %s", got, actual)
	}
}

func TestPopEmpty(t *testing.T) {
	sq := &StringQueue{}
	_, err := sq.Pop()
	if err == nil {
		log.Fatal("Pop() on an empty queue does not error")
	}
}

func TestPop(t *testing.T) {
	sq := &StringQueue{}
	actual := "foobar"
	sq.Push(actual)
	got, err := sq.Pop()
	if err != nil {
		log.Fatal(err)
	}
	if got != actual {
		log.Fatalf("Pop() returned %s, expected %s", got, actual)
	}
}

func TestSizeEmpty(t *testing.T) {
	sq := &StringQueue{}
	got := sq.Size()
	expected := 0
	if got != expected {
		log.Fatalf("Size() returned %d, expected %d", got, expected)
	}
}

func TestSizeNonEmpty(t *testing.T) {
	sq := &StringQueue{}
	sq.Push("foo")
	got := sq.Size()
	expected := 1
	if got != expected {
		log.Fatalf("Size() returned %d, expected %d", got, expected)
	}
}

func TestEmptyEmpty(t *testing.T) {
	sq := &StringQueue{}
	expected := true
	got := sq.Empty()
	if got != expected {
		log.Fatalf("Empty() returned %t, expected %t", got, expected)
	}
}

func TestEmptyNonEmpty(t *testing.T) {
	sq := &StringQueue{}
	sq.Push("foo")
	expected := false
	got := sq.Empty()
	if got != expected {
		log.Fatalf("Empty() returned %t, expected %t", got, expected)
	}
}
