package main

/* 
#cgo LDFLAGS: -framework ApplicationServices -framework Foundation

#include <ApplicationServices/ApplicationServices.h>
#include <unistd.h>
#include <stdio.h>

CGPoint get_curr_pos();

void move_mouse(float x, float y) {
    CGPoint current = get_curr_pos();
    CGPoint move_to = CGPointMake(x, y);
    CGEventRef move = CGEventCreateMouseEvent(NULL, kCGEventMouseMoved, move_to,
                                              kCGMouseButtonLeft);
    CGEventPost(kCGHIDEventTap, move);
    CFRelease(move);
}

void mouse_click_down(uint32_t button) {
    CGEventRef click = CGEventCreateMouseEvent(NULL, kCGEventLeftMouseDown,
                                               get_curr_pos(), button);
    CGEventPost(kCGHIDEventTap, click);
    CFRelease(click);
}

void mouse_click_up(uint32_t button) {
    CGEventRef click = CGEventCreateMouseEvent(NULL, kCGEventLeftMouseUp,
                                               get_curr_pos(), button);
    CGEventPost(kCGHIDEventTap, click);
    CFRelease(click);
}

CGPoint get_curr_pos() {
    CGEventRef ourEvent = CGEventCreate(NULL);
    CGPoint point = CGEventGetLocation(ourEvent);
    CFRelease(ourEvent);
    return point;
}
*/
import "C"

func GetDisplayWidth() int {
	return int(C.CGDisplayPixelsWide(C.CGMainDisplayID()));
}

func GetDisplayHeight() int {
	return int(C.CGDisplayPixelsHigh(C.CGMainDisplayID()));
}

func MoveMouse(movement Movement) {
	C.move_mouse(C.float(movement.x), C.float(movement.y))
}
