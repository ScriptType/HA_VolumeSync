package keylogger

/*
#cgo LDFLAGS: -framework Foundation -framework Foundation -framework Carbon
#include "keylogger.h"

typedef enum State { Up, Down, Invalid } State;

extern void handleButtonEvent(int k, State s);

static inline void listen() {
    CGEventMask eventMask = CGEventMaskBit(kCGEventKeyDown) |
                            CGEventMaskBit(kCGEventKeyUp) |
                            CGEventMaskBit(NX_SYSDEFINED);  // Listening to system-defined events

    CFMachPortRef eventTap = CGEventTapCreate(
        kCGSessionEventTap, kCGHeadInsertEventTap, 0, eventMask, CGEventCallback, NULL
    );

    if (!eventTap) {
        fprintf(stderr, "ERROR: Unable to create event tap.\n");
        exit(1);
    }

    CFRunLoopSourceRef runLoopSource = CFMachPortCreateRunLoopSource(kCFAllocatorDefault, eventTap, 0);
    CFRunLoopAddSource(CFRunLoopGetCurrent(), runLoopSource, kCFRunLoopCommonModes);
    CGEventTapEnable(eventTap, true);

    CFRunLoopRun();
}

// The following callback method is invoked on every keypress.
static inline CGEventRef CGEventCallback(CGEventTapProxy proxy, CGEventType type, CGEventRef event, void *refcon) {
    if (type == kCGEventKeyDown || type == kCGEventKeyUp) {
        CGKeyCode keyCode = (CGKeyCode)CGEventGetIntegerValueField(event, kCGKeyboardEventKeycode);
        State s = (type == kCGEventKeyDown) ? Down : Up;
        handleButtonEvent((int)keyCode, s);
    } else if (type == NX_SYSDEFINED) {
        // Experiment with different integer fields to determine which might hold relevant data
        uint64_t fieldData = CGEventGetIntegerValueField(event, 149); // Example: testing field 1

	if (fieldData == 2560 || fieldData == 2561) {
        handleButtonEvent(132, 1); // Button 132 pressed
    }
    if (fieldData == 2816) {
        handleButtonEvent(132, 2); // Button 132 released
    }
    if (fieldData == 68096 || fieldData == 68097) {
        handleButtonEvent(131, 1); // Button 131 pressed
    }
    if (fieldData == 68352) {
        handleButtonEvent(131, 2); // Button 131 released
    }

    if (fieldData == 461312 || fieldData == 461313) {
        handleButtonEvent(130, 1); // Button 130 pressed
    }
    if (fieldData == 461568) {
        handleButtonEvent(130, 2); // Button 130 released
    }
	if (fieldData == 1247744 || fieldData == 1247745) {
        handleButtonEvent(129, 1); // Button 129 pressed
    }
    if (fieldData == 1248000) {
        handleButtonEvent(129, 2); // Button 129 released
    }

	if (fieldData == 1051136 || fieldData == 1051137) {
        handleButtonEvent(128, 1); // Button 128 pressed
    }
    if (fieldData == 1051392) {
        handleButtonEvent(128, 2); // Button 128 released
    }

	if (fieldData == 1313280 || fieldData == 1313281) {
        handleButtonEvent(127, 1); // Button 127 pressed
    }
    if (fieldData == 1313536) {
        handleButtonEvent(127, 2); // Button 127 released
    }
    }
    return event;
}
*/
import "C"

import (
	"errors"
	"os/user"

	"github.com/ScriptType/HA_VolumeSync/keyboard"
)

type (
	ListenFunc func(key keyboard.Key, state keyboard.State)
	listenFunc func(keyCode C.int, stateCode C.State)
)

type KeyLogger struct{}

var _f listenFunc

//export handleButtonEvent
func handleButtonEvent(keyCode C.int, stateCode C.State) {
	_f(keyCode, stateCode)
}

func New() (*KeyLogger, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	if u.Uid != "0" {
		return nil, errors.New("capturing key logs requires to run the script with root user")
	}

	return &KeyLogger{}, nil
}

func (k *KeyLogger) Listen(f ListenFunc) {
	k.listen(func(keyCode C.int, stateCode C.State) {
		key := keyboard.ConvertKeyCode(int(keyCode))
		state := keyboard.State(stateCode)
		f(key, state)
	})
}

func (k *KeyLogger) listen(f listenFunc) {
	_f = f
	C.listen()
}
