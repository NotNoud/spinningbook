package macro

type Platform string

const (
	PlatformAll     Platform = "all"
	PlatformWindows Platform = "windows"
	PlatformLinux   Platform = "linux"
)

type StepType string

const (
	StepCD  StepType = "cd"
	StepRun StepType = "run"
)

type Step struct {
	Type     StepType `toml:"type"`
	Value    string   `toml:"value"`
	Platform Platform `toml:"platform"`
}

type Macro struct {
	Name        string `toml:"name"`
	Description string `toml:"description"`
	Steps       []Step `toml:"steps"`
}
