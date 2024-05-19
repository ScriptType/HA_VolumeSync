#ifndef __KEYLOGGER_H__
#define __KEYLOGGER_H__

#include <stdio.h>
#include <ApplicationServices/ApplicationServices.h>
#include <Carbon/Carbon.h>

// Function prototype for the CGEventCallback
static inline CGEventRef CGEventCallback(CGEventTapProxy, CGEventType, CGEventRef, void *);

#endif // __KEYLOGGER_H__
