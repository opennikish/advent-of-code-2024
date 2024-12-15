package main

import (
	"adventofcode2024/lib"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app := NewApp(101, 103)
	app.Start(ctx)
}

var rNums = regexp.MustCompile(`[\-0-9]+`)

type Command int

const (
	Pause Command = iota
	Next
	Prev
	Quit
)

var cmdNames = map[Command]string{
	Pause: "pause",
	Next:  "next",
	Prev:  "prev",
	Quit:  "quit",
}

func (c Command) String() string {
	return cmdNames[c]
}

type robot struct {
	x, y       int
	movX, movY int
	char       byte
}

type point struct {
	x, y int
}

type App struct {
	width     int
	height    int
	robots    []*robot
	roboCache map[point]*robot
	counter   int
	paused    bool
	dir       int
}

func NewApp(width, height int) *App {
	return &App{
		width:     width,
		height:    height,
		roboCache: map[point]*robot{},
		paused:    false,
		dir:       1,
	}
}

func (a *App) Start(ctx context.Context) {
	input, err := lib.GetInput(14)
	lib.Check(err)

	a.parseRobots(input)

	cmds, errc := a.runCMDReader(ctx)
	ticks := time.Tick(200 * time.Millisecond)

	gotArrowOnPause := false

	for {
		select {
		case cmd := <-cmds:
			log("cmd: %s", cmd)

			switch cmd {
			case Quit:
				fmt.Println("Bye")
				return
			case Pause:
				a.paused = a.paused != true
				a.dir = 1
				gotArrowOnPause = false
			case Prev:
				gotArrowOnPause = true
				a.dir = -1
			case Next:
				gotArrowOnPause = true
				a.dir = 1
			}

		case <-ticks:
			log("tick")

			if a.paused && !gotArrowOnPause {
				continue
			}

			clearScreen()

			a.counter += a.dir
			a.move()
			a.render()

			gotArrowOnPause = false

		case <-ctx.Done():
			fmt.Println("Bye")
			return
		case err := <-errc:
			fmt.Println(err)
			return
		}
	}
}

func (a *App) move() {
	transform := func(x, move, limit int) int {
		move *= a.dir
		steps := abs(move) % limit
		if move > 0 {
			return (x + steps) % limit
		}

		x -= steps
		if x < 0 {
			x += limit
		}
		return x
	}

	for _, r := range a.robots {
		r.x = transform(r.x, r.movX, a.width)
		r.y = transform(r.y, r.movY, a.height)
	}
}

func (a *App) render() {
	clear(a.roboCache)
	for _, r := range a.robots {
		a.roboCache[point{x: r.x, y: r.y}] = r
	}

	fmt.Printf("--- SEC: %d ---\n\n", a.counter)

	for i := range a.height {
		for j := range a.width {
			if r, ok := a.roboCache[point{y: i, x: j}]; ok {
				fmt.Print(string(r.char))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (a *App) parseRobots(input []byte) {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	robots := make([]*robot, len(lines))

	char := byte(33)
	nextChar := func() byte {
		char++
		if char == 46 { // skip dot
			char++
		}
		if char == 127 { // DEL, reset
			char = 33
		}
		return char
	}

	for i, l := range lines {
		nums := rNums.FindAllString(l, -1)
		if len(nums) != 4 {
			panic("corrupted robot: " + l)
		}

		robots[i] = &robot{
			x:    toInt(nums[0]),
			y:    toInt(nums[1]),
			movX: toInt(nums[2]),
			movY: toInt(nums[3]),
			char: nextChar(),
		}
	}

	a.robots = robots
}

func (a *App) runCMDReader(ctx context.Context) (<-chan Command, <-chan error) {
	// disable input buffering
	err := exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
	lib.Check(err)

	// do not display entered characters on the screen
	err = exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
	lib.Check(err)

	cmds, errc := make(chan Command), make(chan error, 1)
	buf := make([]byte, 3)

	cmdMap := map[string]Command{
		" ":                        Pause,
		string([]byte{27, 91, 67}): Next, // escape sequence for right arrow
		string([]byte{27, 91, 68}): Prev, // escape sequence for left arrow
		"q":                        Quit,
	}

	go func() {
		for {
			if ctx.Err() != nil {
				close(errc)
				close(cmds)
				return
			}

			n, err := os.Stdin.Read(buf)
			if err != nil {
				errc <- fmt.Errorf("read stdin: %w", err)
				close(errc)
				close(cmds)
			}

			if cmd, ok := cmdMap[string(buf[:n])]; ok {
				cmds <- cmd
			}
		}
	}()

	return cmds, errc
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	lib.Check(err)
	return n
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func log(format string, a ...any) {
	if len(a) == 0 {
		fmt.Fprint(os.Stderr, format+"\n")
	} else {
		fmt.Fprintf(os.Stderr, format+"\n", a)
	}
}
