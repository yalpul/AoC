package main

import (
        "fmt"
        "os"
        "strings"
        "strconv"
        "bufio"
)

const (
        value = "value"
        bot = "bot"
        totalBot = 250
        goal1 = 61
        goal2 = 17
)

type botNode struct {
        low int
        lowToBot bool
        high int
        highToBot bool
        has [2]int
        hasLen int
}

var bots [totalBot]botNode

func main() {
        f, _ := os.Open(os.Args[1])
        reader := bufio.NewScanner(f)

        for reader.Scan() {
                line := reader.Text()
                words := strings.Fields(line)
                if words[0] == value {
                        val, _ := strconv.Atoi(words[1])
                        botNum, _ := strconv.Atoi(words[5])
                        give(val, botNum)
                } else if words[0] == bot {
                        botNum, _ := strconv.Atoi(words[1])
                        pref1, _ := strconv.Atoi(words[6])
                        pref2, _ := strconv.Atoi(words[11])
                        setPrefs(botNum, words[5] == bot, pref1, words[10] == bot, pref2)
                } else {
                        panic("invalid key")
                }
        }

        for i, v := range bots {
                if v.hasLen == 2 {
                        num := react(i)
                        if num != -1 {
                                fmt.Println(num)
                                return
                        }
                }
        }
}

func react(num int) int {
        bot := bots[num]
        if bot.has[0] == goal1 && bot.has[1] == goal2 || bot.has[0] == goal2 && bot.has[1] == goal1 {
                return num
        }
        var l, h int
        if bot.has[0] < bot.has[1] {
                l = bot.has[0]
                h = bot.has[1]
        } else {
                h = bot.has[0]
                l = bot.has[1]
        }
        if bot.lowToBot {
                give(l, bot.low)
                if bots[bot.low].hasLen == 2 {
                        result := react(bot.low)
                        if result != -1 {
                                return result
                        }
                }
        }
        if bot.highToBot {
                give(h, bot.high)
                if bots[bot.high].hasLen == 2 {
                        result := react(bot.high)
                        if result != -1 {
                                return result
                        }
                }
        }
        return -1
}

func give(val, botNum int) {
        bot := bots[botNum]
        if bot.hasLen == 0 {
                bot.has[0] = val
                bot.hasLen++
        } else if bot.hasLen == 1 {
                bot.has[1] = val
                bot.hasLen++
        } else {
                panic("bot " + strconv.Itoa(botNum) + " has more than 2")
        }
        bots[botNum] = bot
}

func setPrefs(botNum int, toBot1 bool, pref1 int, toBot2 bool, pref2 int) {
        bot := bots[botNum]
        bot.low = pref1
        bot.lowToBot = toBot1
        bot.high = pref2
        bot.highToBot = toBot2
        bots[botNum] = bot
}
