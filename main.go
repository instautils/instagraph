package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/schollz/progressbar"

	"github.com/ahmdrz/instagraph/src/graph"
	"github.com/ahmdrz/instagraph/src/instagram"
)

var configuration struct {
	NeoAddr  string        `required:"true" envconfig:"neo_addr"`
	Password string        `required:"true" envconfig:"password"`
	Username string        `required:"true" envconfig:"username"`
	Delay    time.Duration `default:"1s" envconfig:"delay"`
	Limit    int           `default:"300" envconfig:"limit"`
}

func main() {
	err := envconfig.Process("INSTAGRAPH", &configuration)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connecting to neo4j on", configuration.NeoAddr)
	var g *graph.Neo
	for {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		var err error
		g, err = graph.New(configuration.NeoAddr)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	if len(configuration.Username)*len(configuration.Password) == 0 {
		log.Fatal("username or password is empty")
		return
	}

	var instance *instagram.Instagram
	if fileExists(configuration.Username + ".json") {
		var err error
		log.Printf("Loading instagram as %s ...", configuration.Username)
		instance, err = instagram.Import(configuration.Username + ".json")
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		var err error
		log.Printf("Connecting to instagram as %s ...", configuration.Username)
		instance, err = instagram.New(configuration.Username, configuration.Password)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("Connected !")

		instance.Export(configuration.Username + ".json")
	}

	log.Println("Start fetching followers")
	currentUsers := instance.Followers()
	shuffle(currentUsers)

	limit := configuration.Limit
	if limit == -1 {
		limit = len(currentUsers)
	}

	err = g.AddNode(configuration.Username)
	if err != nil {
		log.Fatal(err)
	}

	bar := progressbar.NewOptions(limit, progressbar.OptionSetRenderBlankState(true))
	for i, user := range currentUsers {
		bar.Add(1)

		err = g.AddNode(user.Username)
		if err != nil {
			log.Fatal(err)
		}
		err = g.AddConnection(user.Username, configuration.Username)
		if err != nil {
			log.Fatal(err)
		}

		if i >= limit {
			break
		}

		users := user.Followers(instance)
		if len(users) > limit {
			users = users[:limit]
		}
		shuffle(users)

		for _, target := range users {
			if target.Username == configuration.Username {
				continue
			}
			g.AddNode(target.Username)
			g.AddConnection(target.Username, user.Username)
		}

		time.Sleep(configuration.Delay)
	}

	// newline after progressbar
	fmt.Println()
}

func shuffle(vals []instagram.User) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func getPWD() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
