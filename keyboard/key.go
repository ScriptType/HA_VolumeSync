// Copy of github.com/azul3d/engine/keyboard

package keyboard

// Key represents an single keyboard button.
//
// It should be noted that it does not represent an character that pressing an
// keyboard button would otherwise generate (hence you will find no capital
// keys defined).
type Key int

// Keyboard key constants. These are just for button state detection -- not for
// representing a character / text input being generated by pressing a key (for
// that, use TypedEvent).
//
// The buttons are mapped onto a traditional U.S. keyboard layout, which you
// can find a diagram of here:
//
// http://en.wikipedia.org/wiki/File:KB_United_States-NoAltGr.svg
//
// The Invalid key is defined strictly to allow users to detect uninitialized
// variables.
const (
	Invalid      Key = iota
	Tilde            // "~"
	Dash             // "-"
	Equals           // "="
	Semicolon        // ";"
	Apostrophe       // "'"
	Comma            // ","
	Period           // "."
	ForwardSlash     // "/"
	BackSlash        // "\"
	Backspace
	Tab // "\t"
	CapsLock
	Space // " "
	Enter // "\r", "\n", "\r\n"
	Escape
	Insert
	PrintScreen
	Delete
	PageUp
	PageDown
	Home
	End
	Pause
	Sleep
	Clear
	Select
	Print
	Execute
	Help
	Applications
	ScrollLock
	Play
	Zoom

	// Arrow keys
	ArrowLeft
	ArrowRight
	ArrowDown
	ArrowUp

	// Lefties
	LeftBracket // [
	LeftShift
	LeftCtrl
	LeftSuper
	LeftAlt

	// Righties
	RightBracket // ]
	RightShift
	RightCtrl
	RightSuper
	RightAlt

	// Numbers
	Zero  // "0"
	One   // "1"
	Two   // "2"
	Three // "3"
	Four  // "4"
	Five  // "5"
	Six   // "6"
	Seven // "7"
	Eight // "8"
	Nine  // "9"

	// Functions
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	F16
	F17
	F18
	F19
	F20
	F21
	F22
	F23
	F24
	F25

	// English characters
	A // "a"
	B // "b"
	C // "c"
	D // "d"
	E // "e"
	F // "f"
	G // "g"
	H // "h"
	I // "i"
	J // "j"
	K // "k"
	L // "l"
	M // "m"
	N // "n"
	O // "o"
	P // "p"
	Q // "q"
	R // "r"
	S // "s"
	T // "t"
	U // "u"
	V // "v"
	W // "w"
	X // "x"
	Y // "y"
	Z // "z"

	// Number pads
	NumLock
	NumMultiply // "*"
	NumDivide   // "/"
	NumAdd      // "+"
	NumSubtract // "-"
	NumZero     // "0"
	NumOne      // "1"
	NumTwo      // "2"
	NumThree    // "3"
	NumFour     // "4"
	NumFive     // "5"
	NumSix      // "6"
	NumSeven    // "7"
	NumEight    // "8"
	NumNine     // "9"
	NumDecimal  // "."
	NumComma    // ","
	NumEnter

	// Browser key buttons.
	BrowserBack
	BrowserForward
	BrowserRefresh
	BrowserStop
	BrowserSearch
	BrowserFavorites
	BrowserHome

	// Media key buttons.
	MediaNext
	MediaPrevious
	MediaStop
	MediaPlayPause
	MediaMute
	MediaVolumeDown
	MediaVolumeUp

	// Launcher key buttons.
	LaunchMail
	LaunchMedia
	LaunchAppOne
	LaunchAppTwo

	// Expanded int. key buttons.
	Kana
	Kanji
	Junja
	Attn
	CrSel
	ExSel
	EraseEOF
)

var keyCodeKeyMapping = map[int]Key{
	0x00: A,
	0x01: S,
	0x02: D,
	0x03: F,
	0x04: H,
	0x05: G,
	0x06: Z,
	0x07: X,
	0x08: C,
	0x09: V,
	0x0B: B,
	0x0C: Q,
	0x0D: W,
	0x0E: E,
	0x0F: R,
	0x10: Y,
	0x11: T,
	0x12: One,
	0x13: Two,
	0x14: Three,
	0x15: Four,
	0x16: Six,
	0x17: Five,
	0x18: Equals,
	0x19: Nine,
	0x1A: Seven,
	0x1B: Dash,
	0x1C: Eight,
	0x1D: Zero,
	0x1E: RightBracket,
	0x1F: O,
	0x20: U,
	0x21: LeftBracket,
	0x22: I,
	0x23: P,
	0x25: L,
	0x26: J,
	0x27: Apostrophe,
	0x28: K,
	0x29: Semicolon,
	0x2A: BackSlash,
	0x2B: Comma,
	0x2C: ForwardSlash,
	0x2D: N,
	0x2E: M,
	0x2F: Period,
	0x32: Tilde,
	0x24: Enter,
	0x30: Tab,
	0x31: Space,
	0x33: Delete,
	0x35: Escape,
	0x36: RightSuper,
	0x37: LeftSuper,
	0x38: LeftShift,
	0x39: CapsLock,
	0x3A: LeftAlt,
	0x3B: LeftCtrl,
	0x3C: RightShift,
	0x3D: RightAlt,
	0x3E: RightCtrl,
	0x3F: Invalid, // Function
	0x40: F17,
	0x4F: F18,
	0x50: F19,
	0x5A: F20,
	0x60: F5,
	0x61: F6,
	0x62: F7,
	0x63: F3,
	0x64: F8,
	0x65: F9,
	0x67: F11,
	0x69: F13,
	0x6A: F16,
	0x6B: F14,
	0x6D: F10,
	0x6F: F12,
	0x71: F15,
	0x72: Help,
	0x73: Home,
	0x74: PageUp,
	0x75: Delete,
	0x76: F4,
	0x77: End,
	0x78: F2,
	0x79: PageDown,
	0x7A: F1,
	0x7B: ArrowLeft,
	0x7C: ArrowRight,
	0x7D: ArrowDown,
	0x7E: ArrowUp,
	0x7F: MediaPrevious,
	0x80: MediaPlayPause,
	0x81: MediaNext,
	0x82: MediaMute,
	0x83: MediaVolumeDown,
	0x84: MediaVolumeUp,
}

func ConvertKeyCode(keyCode int) Key {
	k, ok := keyCodeKeyMapping[keyCode]
	if !ok {
		return Invalid
	}

	return k
}
