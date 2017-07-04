# Notification
Notification implementation (like in macOS/iOS)

# Example: How to use
```Go
package main

import (
	"Notification"
	"fmt"
)

var notificationCenter = (&Notification.Center{}).Init()
var object1 *Type1
var object2 *Type2

//------- Observer for Type1 -------

type Type1 struct {
	continent string
	value int
}

func (t1 *Type1) Init() *Type1 {
	notificationCenter.Register(t1, "NotificationDidFire", Type1_Observer)
	return t1
}

func Type1_Observer(object interface{}, name string, data Notification.Data) {
	object1 := object.(*Type1)
	fmt.Printf(">>> continent: %s, value: %d, notification: %s\n", object1.continent, object1.value, name)
}

//------- Observer for Type2 -------

type Type2 struct {
	name string
	number int
}

func (t2 *Type2) Init() *Type2 {
	notificationCenter.Register(t2, "NotificationDidFire", Type2_Observer)
	return t2
}

func Type2_Observer(object interface{}, name string, data Notification.Data) {
	object2 := object.(*Type2)
	fmt.Printf("### name: %s, number: %d, notification: %s\n", object2.name, object2.number, name)	
}

//------- Standalone observer -------

func standaloneObserver(object interface{}, name string, data Notification.Data) {
	fmt.Printf("*** standalone observer notification: %s\n", name)	
}


func main() {
	// register observers
	object1 = (&Type1{continent: "Europa", value: 1}).Init()
	object2 = (&Type2{name: "Asia", number: 2}).Init()
	notificationCenter.Register(nil, "NotificationDidFire", standaloneObserver)


	data := Notification.Data{}

	// notify registered observers
	notificationCenter.Post("NotificationDidFire", data)

	fmt.Println("OK")
}
```
