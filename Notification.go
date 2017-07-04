package Notification

import "fmt"

type info struct {
	object interface{}
	callback Callback
}

// Data - alias for data passed via notification
type Data map[string]interface{}

// Callback - type of observer function/method
type Callback func(object interface{}, name string, data Data)

// Center - notifications manager
type Center struct {
	data map[string][]info
}

// Init - data allocation
func (nc *Center) Init() *Center {
	nc.data = make(map[string][]info)
	fmt.Println("Notification center initialized")
	return nc
}

// Register - registers callback for name of notification
func (nc *Center) Register(object interface{}, name string, callback Callback) {
	array, ok := nc.data[name]
	if !ok {
		array = []info{}
	}
	var infoItem = info{object, callback}
	array = append(array, infoItem)
	nc.data[name] = array
}

// Post - posts notification with data
func (nc *Center) Post(name string, data Data) {
	array := nc.data[name]
	for _, infoItem := range array {
		infoItem.callback(infoItem.object, name, data)
	}
}
