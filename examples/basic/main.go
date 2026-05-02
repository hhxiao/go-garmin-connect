// Command basic logs into Garmin Connect and prints the social profile +
// today's user-summary stats. Set GARMIN_LOGIN/GARMIN_PASSWORD before running.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	garmin "github.com/sealbro/go-garmin-connect"
	"github.com/sealbro/go-garmin-connect/auth"
)

func main() {
	login := os.Getenv("GARMIN_LOGIN")
	password := os.Getenv("GARMIN_PASSWORD")
	if login == "" || password == "" {
		log.Fatal("set GARMIN_LOGIN and GARMIN_PASSWORD")
	}

	params, err := auth.NewBasicAuth(login, password)
	if err != nil {
		log.Fatal(err)
	}

	home, _ := os.UserHomeDir()
	cache := auth.NewFileTokenCache(filepath.Join(home, ".garmin_token.json"))
	gctx := garmin.NewContext(nil, params, garmin.WithTokenCache(cache))
	client := garmin.NewClient(gctx)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	profile, err := client.GetSocialProfile(ctx)
	if err != nil {
		log.Fatalf("profile: %v", err)
	}
	fmt.Printf("hello %s (%s)\n", profile.DisplayName, profile.FullName)

	stats, err := client.GetUserSummary(ctx, time.Now().UTC())
	if err != nil {
		log.Fatalf("summary: %v", err)
	}
	fmt.Printf("steps=%d calories=%.0f resting_hr=%d\n",
		stats.TotalSteps, stats.TotalKilocalories, stats.RestingHeartRate)
}
