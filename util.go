package faker

import (
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Arinzeokeke/tr"
)

type strslice []string

func replaceSymbols(s string) string {
	alpha := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	str := ""

	for _, v := range alpha {
		if strings.Compare(v, "#") == 0 {
			str += strconv.Itoa(random(10))
		} else if strings.Compare(v, "?") == 0 {
			str += alpha[random(len(alpha))]
		} else if strings.Compare(v, "*") == 0 {
			rn := random(2)
			if rn == 0 {
				str += strconv.Itoa(random(10))
			} else {
				str += alpha[random(len(alpha))]
			}
		} else {
			str += v
		}

	}
	return str
}

func random(i int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(i)
}

func randomIntRange(min, max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(max-min) + min
}
func randomFloatRange(min, max int) float64 {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return float64(r.Intn(max-min)) + float64(min) + r.Float64()
}

func directoryExists(dir string) bool {
	dir, _ = filepath.Abs(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func getList(n *tr.Engine, q, d string) ([]string, error) {
	w, err := n.Tr(q)
	if err != nil {
		w, err = n.Lang(d).Tr(q)
	}
	if err != nil {
		return nil, err
	}
	return strings.Split(w, "\n"), nil
}

// Find the first index at which an element exists in the array
func (slice strslice) indexOf(searchElement string) int {
	for index, val := range slice {
		if val == searchElement {
			return index
		}
	}
	return -1
}

// Return a random elemnet from an array
func selectElement(arr []string) string {
	return arr[random(len(arr))]
}

func slugify(word string) string {
	re := regexp.MustCompile(" ")
	rex := regexp.MustCompile(`[^\w\.\-]+`)

	replaceHyphen := re.ReplaceAllString(word, "-")
	replaceBlank := rex.ReplaceAllString(replaceHyphen, "")

	return replaceBlank
}

// Generate a 4-digit hash
func randHash() string {
	hash := ""
	for i := 0; i < 4; i++ {
		hash += selectElement([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"})
	}
	return hash
}

func mustache(str string, data map[string]string) string {
	for key, val := range data {
		re := regexp.MustCompile("{{" + key + "}}")
		str = re.ReplaceAllString(str, val)
	}
	return str
}
