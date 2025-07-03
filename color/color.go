package color

import (
	"fmt"
	"strconv"
)

// Basic ANSI foreground colors and resets
const (
	Reset = "\033[0m"

	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"
)

// Semantic/HTTP method colors
const (
	GetColor     = BrightGreen   // GET = bright green (safe/read)
	PostColor    = BrightBlue    // POST = bright blue (create)
	PutColor     = BrightYellow  // PUT = bright yellow (update/replace)
	PatchColor   = BrightMagenta // PATCH = bright magenta (partial update)
	DeleteColor  = BrightRed     // DELETE = bright red (danger)
	HeadColor    = Cyan          // HEAD = cyan
	OptionsColor = White         // OPTIONS = white

	SuccessColor = BrightGreen  // generic success
	WarnColor    = BrightYellow // generic warning
	ErrorColor   = BrightRed    // generic error
	InfoColor    = BrightBlue   // generic info
	DebugColor   = BrightBlack  // debug
)

// Example map for programmatic access if needed
var HTTPMethodToColor = map[string]string{
	"GET":     GetColor,
	"POST":    PostColor,
	"PUT":     PutColor,
	"PATCH":   PatchColor,
	"DELETE":  DeleteColor,
	"HEAD":    HeadColor,
	"OPTIONS": OptionsColor,
}

// Extended named colors and ANSI helpers remain the same as before...
var NameToCode = map[string]string{
	"reset":          Reset,
	"black":          Black,
	"red":            Red,
	"green":          Green,
	"yellow":         Yellow,
	"blue":           Blue,
	"magenta":        Magenta,
	"cyan":           Cyan,
	"white":          White,
	"bright_black":   BrightBlack,
	"bright_red":     BrightRed,
	"bright_green":   BrightGreen,
	"bright_yellow":  BrightYellow,
	"bright_blue":    BrightBlue,
	"bright_magenta": BrightMagenta,
	"bright_cyan":    BrightCyan,
	"bright_white":   BrightWhite,
	// Semantic aliases for easy reference
	"get":     GetColor,
	"post":    PostColor,
	"put":     PutColor,
	"patch":   PatchColor,
	"delete":  DeleteColor,
	"head":    HeadColor,
	"options": OptionsColor,
	"success": SuccessColor,
	"warn":    WarnColor,
	"error":   ErrorColor,
	"info":    InfoColor,
	"debug":   DebugColor,
}

func Paint(s, color string) string {
	return color + s + Reset
}

func ANSI256(n int) string {
	return "\033[38;5;" + strconv.Itoa(n) + "m"
}

func RGB(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}
