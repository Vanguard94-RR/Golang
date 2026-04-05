# Terminal Formatting Library - Refactor Plan

## Goal
Convert `tittle-subTittle.go` into a reusable, configurable terminal formatting library.

## Phase 1: Restructure Files

```
Projects/
├── terminal-formatter/
│   ├── go.mod
│   ├── colors.go (ANSI constants)
│   ├── formatter.go (Core logic + types)
│   ├── styles.go (Predefined styles)
│   ├── examples/
│   │   └── example.go
│   └── README.md
```

## Phase 2: Define Core Types

### TextFormatter Struct
```go
type TextFormatter struct {
    Width      int
    PadChar    string // "-" by default
    TextColor  string
    Centered   bool
    Bold       bool
}
```

### Style Presets
```go
var (
    TitleStyle   = TextFormatter{Width: 80, TextColor: Green, Bold: true}
    SubTitleStyle = TextFormatter{Width: 80, TextColor: Cyan, Bold: false}
    ErrorStyle    = TextFormatter{Width: 80, TextColor: Red, Bold: true}
    SuccessStyle  = TextFormatter{Width: 80, TextColor: LightGreen, Bold: true}
)
```

## Phase 3: Core Functions

1. **Format(text string) string**
   - Returns formatted string (no printing)
   - Allows chaining

2. **Print(text string)**
   - Wrapper: fmt.Println(formatter.Format(text))

3. **Box(text string) string**
   - Creates bordered box around text

4. **Section(title, subtitle, content string) string**
   - Multi-line formatted section

5. **Colorize(text, color string) string**
   - Generic color wrapper

## Phase 4: API Design

```go
// Method chaining
formatter := New().
    WithWidth(100).
    WithColor(Blue).
    WithBold(true).
    Print("My Title")

// Quick helpers
Title("Welcome")
SubTitle("Introduction")
Box("Important Message")
Section("Section Title", "Subtitle", "Content...")

// Customizable
custom := NewFormatter(80, "-", Red, true)
result := custom.Format("Custom Title")
```

## Phase 5: Implementation Tasks

- [ ] Create directory structure
- [ ] Move ANSI constants to colors.go
- [ ] Fix `Italics` typo
- [ ] Implement TextFormatter struct
- [ ] Add builder methods (WithColor, WithWidth, etc)
- [ ] Return strings instead of printing
- [ ] Add style presets
- [ ] Create helper functions (Title, SubTitle, Box, Section)
- [ ] Add examples
- [ ] Add unit tests
- [ ] Write README with usage examples

## Phase 6: Testing

```go
func TestTitleFormat(t *testing.T) {
    result := Title("Hello")
    if !strings.Contains(result, "Hello") {
        t.Fail()
    }
}
```

## Phase 7: Documentation

- README.md with examples
- Function godoc comments
- Usage patterns

## Benefits After Refactor

✅ Configurable (width, colors, padding)
✅ Reusable across projects
✅ Chainable API
✅ No side effects (returns strings)
✅ Type-safe
✅ Testable
✅ Package-ready

## Priority Order

1. Core structure (Phase 1-2)
2. Basic functions (Phase 3)
3. Builder API (Phase 4)
4. Testing & docs (Phase 6-7)
