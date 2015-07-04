package main

/* 
#cgo LDFLAGS: -framework ApplicationServices -framework Foundation

#include <ApplicationServices/ApplicationServices.h>
#include <unistd.h>
#include <stdio.h>

CGPoint get_curr_pos();

void mouse_move(float dx, float dy) {
    CGPoint current = get_curr_pos();
    CGPoint move_to = CGPointMake(current.x + dx, current.y + dy);
    CGEventRef move = CGEventCreateMouseEvent(NULL, kCGEventMouseMoved, move_to,
                                              kCGMouseButtonLeft);
    CGEventPost(kCGHIDEventTap, move);
    CFRelease(move);
}

void mouse_click_down(CGPoint position, uint32_t button) {
    CGEventRef click = CGEventCreateMouseEvent(NULL, kCGEventLeftMouseDown, position,
                                               button);
    CGEventPost(kCGHIDEventTap, click);
    CFRelease(click);
}

void mouse_click_up(CGPoint position, uint32_t button) {
    CGEventRef click = CGEventCreateMouseEvent(NULL, kCGEventLeftMouseUp, position,
                                               button);
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
import (
	"fmt"
	"time"
)

func isNegative(num float64) bool {
	if num < 0 {
		return true
	}
	return false
}

func processComponent(component float64, last *float64) float64 {
	if isNegative(component) != isNegative(*last) {
		component := 0.0
		if *last < component {
			component -= *last
			*last = component
		}
		return component
	}
	*last = component
	return component
}

func Loop(signal chan Acceleration) {
	var value, acceleration Acceleration
	last := Acceleration{0, 0}
	for {
		value = <-signal
		acceleration = Acceleration{processComponent(value.x, &last.x),
      			processComponent(value.y, &last.y)}
		MoveMouse(acceleration)
		last = acceleration
		fmt.Println(acceleration)
		time.Sleep(1/time.Second * time.Millisecond)
	}
}

func MoveMouse(value Acceleration) {
	C.mouse_move(C.float(value.x), C.float(value.y))
}
