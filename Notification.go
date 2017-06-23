package Notification

// Data - alias for data passed via notification
type Data map[string]interface{}

// Callback - type of observer function/method
type Callback func(name string, data Data)

// Center - notifications manager
type Center struct {
	data map[string][]Callback
}

// Init - data allocation
func (nc *Center) Init() *Center {
	nc.data = make(map[string][]Callback)
	return nc
}

// Register - registers callback for name of notification
func (nc *Center) Register(name string, callback Callback) {
	array, ok := nc.data[name]
	if !ok {
		array = []Callback{}
	}
	array = append(array, callback)
	nc.data[name] = array
}

// Post - posts notification with data
func (nc *Center) Post(name string, data Data) {
	array := nc.data[name]
	for _, callback := range array {
		callback(name, data)
	}
}
