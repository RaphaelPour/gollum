// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	//rawInput = `How long does it take for the failover mechanism to fail-over? How long does it take for my server to be up again? I recently had the challenge of measuring downtime durations. ping on the one hand “makes me see” when paket loss occurs and vanishes but lacks in aggregation. mtr on the other side aggregates the replies, but not for the metric I was looking for. So I made up my own tool namely downtime. Aside the plain measurement logic, I tried a few new things like external lib interfaces and mocks which will shortly be addressed in the end. Ping the host via icmp using a customized timeout and a small interval to pinpoint the start and end timestamp of the downtime. This can then be used to calculate the duration. The timeout should be as low as possible. The RTT of my service was not higher than 25ms, so 50ms timeout should be good to go. It should also tolerate some spikes that might appear randomly without falsifying our measurements. The interval should be pretty low and defines the error of the downtime. The icmp logic is based on my colleagues mtr go-rewrite providing me the most high-level icmp interface I could find. This enables me to use raw icmp pakets without caring about the low level stuff like encoding the PID into the request body to fish out the corresponding replies from the ocean of icmp traffic.`
	//rawInput = `Hallo Raphael!`
	//rawInput = `This project is the site you are currently reading this article from. Not just the back- or the frontend, but both combined. This site is built with Nuxt.js, which is a Vue framework, that enables server-side rendering and gives much other neat utility. My previous site was also built with Vue, but with client-side rendering and worse styling. The main problem that bothered me was that the site needed way too long to load, and if you clicked on a post, it would take an entire second to load the actual article. While waiting for the content for loading, you would sit in front of a blank article screen, with nothing indicating that anything is happening. Another problem was that there was no real data management. By this, I mean that on the front page every single article was already loaded with every single bit of its information. But when you entered a certain article, the entire block of data was fetched again. When I discovered Nuxt.js I thought that this would be exactly what I needed. And that is still my opinion. Compared to my old site, the new site is loading blazingly fast and the user has to fetch way less data, because now that the backend and the frontend can communicate better with each other, the data to be fetched can be handled smartly. When it comes to styling, I started using Tailwind CSS and tried to think more about the design, which seems to have worked out, because I see fewer flaws with the design on this site. Previously I have always named my main-site "JP". But when I realised that I rewrite the entire page every now and then, I thought I might name it individually so I don't have to overwrite the repository as I did with JP-old. I named this project "JP-nuxt" because I think that this will be the only version of my site that is built with Nuxt.js. Not because I don't like Nuxt.js, but because with my website, I always test out new technology. So my next project will maybe be made using Next.js, Svelte Kit or I might go back to the roots and generate simple static sites like evilcookie.de. When setting up my project, I decided to use Nuxt 3, which may not have been the wisest decision, because it is so new and unstable that nearly none of the plugins for Nuxt.js works. So I had to make most of the basic stuff on my own. What annoyed me, was that I wasn't able to use express.js in the backend, and I had to handle requests via the default http module in node.js, which isn't as nice to use as express.js. So to abstract things a bit, I started building a little router, that at least makes me feel like I was using express.js. The router looked basically like this: Now I could start talking about the styling, the backend or the upload page, but I think I made my point clear. I saw flaws all over the site.`
	rawInput = "chemnitzer linuxtage"
	model    = map[string][]string{}
)

type LanguageModel struct {
	model       map[string][]string
}

func NewLanguageModel() LanguageModel {
	return LanguageModel{
		model:       make(map[string][]string),
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
		}

		lm.model[c] = append(lm.model[c], follower)
	}
}

func (lm LanguageModel) Prompt(input string, length int) {
	fmt.Printf("%s> ", input)

	start = sample[int]find(rawInput, input)

	for _, raw := range input {
		p := string(raw)
		fmt.Println(sample(lm.model[p]))
	}
}

func (lm LanguageModel) String() string {
	return fmt.Sprintf("%+v", lm.model)
}

func find(needle string, haystack string) []int {
	result := []int{}
	for i := 0; i < len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			result = append(result, i)
			i += len(needle) - 1
		}
	}
	return result
}

func stringInSlice(needle string, haystack []string) bool {
	for _, c := range haystack {
		if string(c) == needle {
			return true
		}
	}
	return false
}

func sample[K any](arr []K) K {
	return K(arr[rand.Intn(len(arr))])
}

func main() {
	lm := NewLanguageModel()

	prompt := ""
	fmt.Println(strings.Join(lm.model[prompt], ""))
	lm.Prompt(prompt, 1)
}

