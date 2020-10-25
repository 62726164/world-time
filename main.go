package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func main() {
	// wt
	if len(os.Args) == 1 {
		for _, name := range []string{
			"UTC",
			"Local",
			"Atlantic/Reykjavik",
			"Atlantic/Cape_Verde",
			"America/Noronha",
			"America/Buenos_Aires",
			"America/New_York",
			"America/Chicago",
			"America/Denver",
			"America/Los_Angeles",
			"America/Anchorage",
			"America/Adak",
			"Pacific/Honolulu",
			"Pacific/Midway",
			"Pacific/Wake",
			"Pacific/Guadalcanal",
			"Pacific/Guam",
			"Asia/Tokyo",
			"Asia/Shanghai",
			"Asia/Saigon",
			"Antarctica/Vostok",
			"Asia/Ashkhabad",
			"Asia/Dubai",
			"Europe/Moscow",
			"Africa/Johannesburg",
			"Europe/London",
		} {
			t, err := TimeIn(time.Now(), name)
			if err == nil {
				fmt.Println(t.Format(time.RFC3339), t.Location())
			} else {
				fmt.Println(name, "<time unknown>")
			}
		}
	}

	// wt ls
	if len(os.Args) == 2 {
		cmd := exec.Command(os.Args[1])
		output, err := cmd.Output()
		if err != nil {
			return
		} else {
			fmt.Printf("%s", output)
		}
	}

	// wt cat file
	if len(os.Args) == 3 {
		cmd := exec.Command(os.Args[1], os.Args[2])
		output, err := cmd.Output()
		if err != nil {
			return
		} else {
			fmt.Printf("%s", output)
		}
	}

	// wt chown root:root file
	// wt chmod 777 file
	// wt cp src dst
	// wt ls -alh file
	if len(os.Args) == 4 {
		cmd := exec.Command(os.Args[1], os.Args[2], os.Args[3])
		output, err := cmd.Output()
		if err != nil {
			return
		} else {
			fmt.Printf("%s", output)
		}
	}

	// wt ls -alh file file2
	if len(os.Args) == 5 {
		cmd := exec.Command(os.Args[1], os.Args[2], os.Args[3], os.Args[4])
		output, err := cmd.Output()
		if err != nil {
			return
		} else {
			fmt.Printf("%s", output)
		}
	}

	// wt ls -alh file file2 file3
	if len(os.Args) == 6 {
		cmd := exec.Command(os.Args[1], os.Args[2], os.Args[3], os.Args[4], os.Args[5])
		output, err := cmd.Output()
		if err != nil {
			return
		} else {
			fmt.Printf("%s", output)
		}
	}
}
