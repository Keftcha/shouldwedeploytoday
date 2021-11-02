package main

import (
	"time"
)

type Day struct {
	Title            string `json:"title"`
	Image            string `json:"image"`
	Text             string `json:"text"`
	IsADeploymentDay bool   `json:"is_a_deployment_day"`
}

var days []Day = []Day{
	// Sunday, have index 0 (https://pkg.go.dev/time#Weekday)
	{
		Title:            "On Sunday ?",
		Image:            "/images/nap_time.gif",
		Text:             "NO.",
		IsADeploymentDay: false,
	},
	{
		Title:            "On Monday ?",
		Image:            "/images/sleepy_dog.gif",
		Text:             "Let's go ! (even if the return to work is hard)",
		IsADeploymentDay: true,
	},
	{
		Title:            "On Tuesday ?",
		Image:            "/images/sounds_good_to_me.gif",
		Text:             "Let's go !",
		IsADeploymentDay: true,
	},
	{
		Title:            "On Wednesday ?",
		Image:            "/images/ping_pong.gif",
		Text:             "Let's go !",
		IsADeploymentDay: true,
	},
	{
		Title:            "On Thursday ?",
		Image:            "/images/hot.gif",
		Text:             "In the morning ok. In the afternoon maybe not.",
		IsADeploymentDay: true,
	},
	{
		Title:            "On Friday ?",
		Image:            "/images/no.gif",
		Text:             "NO.",
		IsADeploymentDay: false,
	},
	{
		Title:            "On Saturday ?",
		Image:            "/images/chaos.gif",
		Text:             "NO.",
		IsADeploymentDay: false,
	},
}

// GetDay return the day given his id
// Sunday = 0, Monday = 1, ...
func GetDay(weekDay time.Weekday) Day {
	return days[weekDay]
}
