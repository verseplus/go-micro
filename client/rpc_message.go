package client

type message struct {
	topic       string
	headerTopic string
	contentType string
	payload     interface{}
}

func newMessage(topic string, headerTopic string, payload interface{}, contentType string, opts ...MessageOption) Message {
	var options MessageOptions
	for _, o := range opts {
		o(&options)
	}

	if len(options.ContentType) > 0 {
		contentType = options.ContentType
	}

	return &message{
		payload:     payload,
		topic:       topic,
		headerTopic: headerTopic,
		contentType: contentType,
	}
}

func (m *message) ContentType() string {
	return m.contentType
}

func (m *message) Topic() string {
	return m.topic
}
func (m *message) HeaderTopic() string {
	return m.headerTopic
}
func (m *message) Payload() interface{} {
	return m.payload
}
