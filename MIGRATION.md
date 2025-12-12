# Migration Guide

This guide covers breaking changes in the html2text package due to upgrading tablewriter from v0.0.5 to v1.1.2.

## Module Path Change

The module path has been updated to v2 to reflect the breaking changes:

**Before:**
```go
import "github.com/DefendaSolutions/fileconversion-go/html2text"
```

**After:**
```go
import "github.com/DefendaSolutions/fileconversion-go/v2/html2text"
```

Update your `go.mod`:
```bash
go get github.com/DefendaSolutions/fileconversion-go/v2
```

## Who needs to migrate?

Only if you create `PrettyTablesOptions` with custom values. If you use `NewPrettyTablesOptions()` or `Options{PrettyTables: true}`, no changes are needed.

## Breaking Changes

### 1. Alignment types changed from `int` to `tw.Align`

**Before:**
```go
import "github.com/olekukonko/tablewriter"

opts := &html2text.PrettyTablesOptions{
    HeaderAlignment: tablewriter.ALIGN_CENTER,
    FooterAlignment: tablewriter.ALIGN_RIGHT,
    Alignment:       tablewriter.ALIGN_LEFT,
}
```

**After:**
```go
import "github.com/olekukonko/tablewriter/tw"

opts := &html2text.PrettyTablesOptions{
    HeaderAlignment: tw.AlignCenter,
    FooterAlignment: tw.AlignRight,
    Alignment:       tw.AlignLeft,
}
```

**Mapping:**

| Old (int)                   | New (tw.Align)    |
|-----------------------------|-------------------|
| `tablewriter.ALIGN_DEFAULT` | `tw.AlignDefault` |
| `tablewriter.ALIGN_CENTER`  | `tw.AlignCenter`  |
| `tablewriter.ALIGN_RIGHT`   | `tw.AlignRight`   |
| `tablewriter.ALIGN_LEFT`    | `tw.AlignLeft`    |

### 2. ColumnAlignment type changed from `[]int` to `[]tw.Align`

**Before:**
```go
opts := &html2text.PrettyTablesOptions{
    ColumnAlignment: []int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_CENTER},
}
```

**After:**
```go
opts := &html2text.PrettyTablesOptions{
    ColumnAlignment: []tw.Align{tw.AlignLeft, tw.AlignCenter},
}
```

### 3. Borders type changed from `tablewriter.Border` to `tw.Border`

**Before:**
```go
opts := &html2text.PrettyTablesOptions{
    Borders: tablewriter.Border{Left: true, Right: true, Top: true, Bottom: true},
}
```

**After:**
```go
opts := &html2text.PrettyTablesOptions{
    Borders: tw.Border{Left: tw.On, Right: tw.On, Top: tw.On, Bottom: tw.On},
}
```

**Mapping:**

| Old (bool) | New (tw.State) |
|------------|----------------|
| `true`     | `tw.On`        |
| `false`    | `tw.Off`       |

### 4. AutoFormatHeader and AutoFormatFooter changed from `bool` to `tw.State`

**Before:**
```go
opts := &html2text.PrettyTablesOptions{
    AutoFormatHeader: true,
    AutoFormatFooter: false,
}
```

**After:**
```go
opts := &html2text.PrettyTablesOptions{
    AutoFormatHeader: tw.On,
    AutoFormatFooter: tw.Off,
}
```

### 5. AutoWrapText changed from `bool` to `int`

The `AutoWrapText` field now uses tablewriter's wrap constants for finer control.

**Before:**
```go
opts := &html2text.PrettyTablesOptions{
    AutoWrapText: true,  // or false
}
```

**After:**
```go
opts := &html2text.PrettyTablesOptions{
    AutoWrapText: tw.WrapNormal,  // or tw.WrapNone, tw.WrapTruncate, tw.WrapBreak
}
```

**Mapping:**

| Old (bool) | New (int)          |
|------------|--------------------|
| `true`     | `tw.WrapNormal`    |
| `false`    | `tw.WrapNone`      |

**Available wrap modes:**

| Constant          | Description                      |
|-------------------|----------------------------------|
| `tw.WrapNone`     | No wrapping                      |
| `tw.WrapNormal`   | Standard word wrapping           |
| `tw.WrapTruncate` | Truncate text with ellipsis      |
| `tw.WrapBreak`    | Break words to fit               |

### 6. RowLine changed from `bool` to `tw.State`

**Before:**
```go
opts := &html2text.PrettyTablesOptions{
    RowLine: false,
}
```

**After:**
```go
opts := &html2text.PrettyTablesOptions{
    RowLine: tw.Off,
}
```

### 7. Separator fields removed in favor of CustomSymbols

The `ColumnSeparator`, `RowSeparator`, and `CenterSeparator` fields have been removed. Use the `CustomSymbols` field instead for full control over table border characters.

**Before:**
```go
opts := &html2text.PrettyTablesOptions{
    ColumnSeparator: "|",
    RowSeparator:    "-",
    CenterSeparator: "+",
}
```

**After:**
```go
opts := html2text.NewPrettyTablesOptions()
opts.CustomSymbols = tw.NewSymbolCustom("custom").
    WithColumn("|").      // Vertical lines
    WithRow("-").         // Horizontal lines
    WithCenter("+").      // Junctions
    WithTopLeft("+").
    WithTopMid("+").
    WithTopRight("+").
    WithMidLeft("+").
    WithMidRight("+").
    WithBottomLeft("+").
    WithBottomMid("+").
    WithBottomRight("+").
    WithHeaderLeft("+").
    WithHeaderMid("+").
    WithHeaderRight("+")

output, _ := html2text.FromString(htmlInput, html2text.Options{
    PrettyTables:        true,
    PrettyTablesOptions: opts,
})
```

**Why this change?**
- `CustomSymbols` provides complete control over all 14 border characters (corners, junctions, lines)
- The old fields only controlled 3 characters and were never wired up in the implementation
- More flexible for creating beautiful tables with Unicode box drawing characters

**Example with Unicode box drawing:**
```go
customSymbols := tw.NewSymbolCustom("unicode").
    WithColumn("│").
    WithRow("─").
    WithCenter("┼").
    WithTopLeft("┌").
    WithTopMid("┬").
    WithTopRight("┐").
    WithMidLeft("├").
    WithMidRight("┤").
    WithBottomLeft("└").
    WithBottomMid("┴").
    WithBottomRight("┘").
    WithHeaderLeft("├").
    WithHeaderMid("┼").
    WithHeaderRight("┤")

opts := html2text.NewPrettyTablesOptions()
opts.CustomSymbols = customSymbols

output, _ := html2text.FromString(htmlInput, html2text.Options{
    PrettyTables:        true,
    PrettyTablesOptions: opts,
})

// Output:
// ┌───────┬─────┐
// │ NAME  │ AGE │
// ├───────┼─────┤
// │ Alice │ 30  │
// │ Bob   │ 25  │
// └───────┴─────┘
```

### 8. Additional fields removed in favor of new custom fields

The following fields have been removed and replaced with more flexible alternatives:

| Removed Field          | Replaced By            | Notes                                                    |
|------------------------|------------------------|----------------------------------------------------------|
| `ReflowDuringAutoWrap` | N/A                    | No direct equivalent in tablewriter v1                   |
| `HeaderLine`           | `CustomSeparators`     | Use `CustomSeparators.ShowHeader` for full control       |
| `AutoMergeCells`       | `CustomCellMerging`    | Use `CustomCellMerging` for advanced cell merging config |

**Example with CustomSeparators:**
```go
opts := html2text.NewPrettyTablesOptions()
opts.CustomSeparators = tw.Separators{
    ShowHeader:     tw.On,  // Show separator after header
    ShowFooter:     tw.On,  // Show separator before footer
    BetweenRows:    tw.On,  // Show separator between rows
    BetweenColumns: tw.On,  // Show column separators
}

output, _ := html2text.FromString(htmlInput, html2text.Options{
    PrettyTables:        true,
    PrettyTablesOptions: opts,
})
```

**Example with CustomCellMerging:**
```go
opts := html2text.NewPrettyTablesOptions()
opts.CustomCellMerging = tw.CellMerging{
    Mode: tw.MergeVertical | tw.MergeHorizontal, // Merge both ways
}

output, _ := html2text.FromString(htmlInput, html2text.Options{
    PrettyTables:        true,
    PrettyTablesOptions: opts,
})
```
