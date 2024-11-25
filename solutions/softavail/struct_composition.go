package softavail

import (
	"errors"
	"slices"
	"strconv"
	"strings"
	"time"
)

type OS struct {
	name   string
	isFree bool
}

type Windows struct {
	OS
	expiryDate time.Time
}

func (w Windows) details() []string {
	return []string{
		w.name,
		strconv.FormatBool(w.isFree),
		w.expiryDate.Format("2006-01-02 15:04:05 -0700 MST"),
	}
}

type Linux struct {
	OS
	aptBased, yumBased bool
}

func (l Linux) details() []string {
	return []string{
		l.name,
		strconv.FormatBool(l.isFree),
		strconv.FormatBool(l.aptBased),
		strconv.FormatBool(l.yumBased),
	}
}

// Here the idea is to propose composition of struct so that they contain both
// common information and specific to particular os.
// Following rules apply:
// - windows is not free, linux is
// - windows endOfSupport table (this is not full, I can't access this recruitment exercise now):
// - linux apt/yum table (this is not full, I can't access this recruitment exercise now):
//
// Example:
// Input: "Windows 8"
// Output: ["Windows", "false", "2008-01-01 17:00:00 +0000 UTC"] (os type, isFree, expiryDate)
//
// Input: Ubuntu
// Output: ["Linux", "true", "true", "false"] (os type, aptBased, yumBased)
func getOsDetails(name string) ([]string, error) {
	if strings.Contains(name, "Windows") {
		expiry, err := getExpiryDate(name)
		if err != nil {
			return nil, err
		}
		windows := Windows{
			OS: OS{
				name:   "Windows",
				isFree: false,
			},
			expiryDate: expiry,
		}
		return windows.details(), nil
	}
	linux := Linux{
		OS: OS{
			name:   "Linux",
			isFree: true,
		},
		aptBased: isAptBased(name),
		yumBased: isYumBased(name),
	}
	return linux.details(), nil
}

func getExpiryDate(name string) (time.Time, error) {
	nameVersion := strings.Split(name, " ")
	if len(nameVersion) != 2 {
		return time.Time{}, errors.New("invalid input received")
	}
	switch nameVersion[1] {
	case "8":
		return time.Date(2009, 2, 12, 15, 0, 0, 0, time.UTC), nil
	case "Vista":
		return time.Date(2012, 4, 11, 15, 0, 0, 0, time.UTC), nil
	}
	return time.Time{}, errors.New("unsupported version")
}

var yumBased = []string{"redhat"}
var aptBased = []string{"ubuntu"}

func isYumBased(name string) bool {
	distro := strings.Split(name, " ")[0]
	return slices.Contains(yumBased, strings.ToLower(distro))
}

func isAptBased(name string) bool {
	distro := strings.Split(name, " ")[0]
	return slices.Contains(aptBased, strings.ToLower(distro))
}
