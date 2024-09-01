package ln

import (
	"math/rand"
)

var lns map[string]shortlink

var chars = []string {
    "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", 
    "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
    "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
    "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
    "1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
}

func makePath(length int) string {
    var path string

    if length == 0 {
        length = 7
    }

    for i := 0; i < length; i++ {
        path += chars[rand.Intn(len(chars))]
    }

    return path
}

func pathExists(path string) bool {
    return false
}

func getNextPath() string {
    return ""
}

func registerShortlink(s shortlink) {
    lns[s.shortURL.Path] = s
}
