package factory

import (
	"fmt"
	"time"
)

type Content interface {
	Play()
}

type CloudContent struct{}

func (c *CloudContent) Play() {
	fmt.Println("this is cloud content")
}

type ProgrammingContent struct{}

func (c *ProgrammingContent) Play() {
	fmt.Println("this is programming content")
}

type MarketingContent struct{}

func (c *MarketingContent) Play() {
	fmt.Println("this is marketing content")
}

type ContentType string

const (
	Bisnis  ContentType = "bisnis"
	Edukasi ContentType = "edukasi"
)

type ContentCreator interface {
	Produce(time time.Time) Content
	Type() ContentType
}

type Imre struct{}

func (i *Imre) Produce(time time.Time) Content {

	if time.Weekday().String() == "Tuesday" {
		return &CloudContent{}
	} else if time.Weekday().String() == "Saturday" {
		return &ProgrammingContent{}
	}

	return nil
}

func (i Imre) Type() ContentType {
	return Edukasi
}

type DewaPrayoga struct{}

func (d *DewaPrayoga) Produce(time time.Time) Content {
	return &MarketingContent{}
}

func (d DewaPrayoga) Type() ContentType {
	return Bisnis
}
