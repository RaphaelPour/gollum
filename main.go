// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
)

var (
	//input = `How long does it take for the failover mechanism to fail-over? How long does it take for my server to be up again? I recently had the challenge of measuring downtime durations. ping on the one hand “makes me see” when paket loss occurs and vanishes but lacks in aggregation. mtr on the other side aggregates the replies, but not for the metric I was looking for. So I made up my own tool namely downtime. Aside the plain measurement logic, I tried a few new things like external lib interfaces and mocks which will shortly be addressed in the end. Ping the host via icmp using a customized timeout and a small interval to pinpoint the start and end timestamp of the downtime. This can then be used to calculate the duration. The timeout should be as low as possible. The RTT of my service was not higher than 25ms, so 50ms timeout should be good to go. It should also tolerate some spikes that might appear randomly without falsifying our measurements. The interval should be pretty low and defines the error of the downtime. The icmp logic is based on my colleagues mtr go-rewrite providing me the most high-level icmp interface I could find. This enables me to use raw icmp pakets without caring about the low level stuff like encoding the PID into the request body to fish out the corresponding replies from the ocean of icmp traffic.`
	rawInput = `Hallo Raphael!`
	model    = map[string][]string{}
	prompt   = "Hallo"
)

type LanguageModel struct {
	model       map[string][]string
	tokenLength int
}

func NewLanguageModel(tokenLength int) LanguageModel {
	return LanguageModel{
		model:       make(map[string][]string),
		tokenLength: tokenLength,
	}
}

func (lm *LanguageModel) Train(input string) {
	for i, raw := range input {
		c := string(raw)
		follower := ""
		for j := 1; j <= lm.tokenLength; j++ {
			follower += string(input[(i+j)%len(input)])
		}
		if _, ok := lm.model[c]; !ok {
			lm.model[c] = make([]string, 0)
		} else if stringInSlice(follower, lm.model[c]) {
			continue
		}

		lm.model[c] = append(lm.model[c], follower)
	}
}

func (lm LanguageModel) Prompt(input string) {
	fmt.Printf("%s> ", input)
	for _, raw := range input {
		p := string(raw)
		fmt.Print(sample(lm.model[p]))
	}
}

func (lm LanguageModel) String() string {
	return fmt.Sprintf("%+v", lm.model)
}

func stringInSlice(needle string, haystack []string) bool {
	for _, c := range haystack {
		if string(c) == needle {
			return true
		}
	}
	return false
}

func sample(arr []string) string {
	return string(arr[rand.Intn(len(arr))])
}

func main() {
	lm := NewLanguageModel(3)
	lm.Train(rawInput)
	fmt.Println(lm)
	lm.Prompt(prompt)
}
