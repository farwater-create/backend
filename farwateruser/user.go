package farwateruser

import "github.com/farwater-create/backend/models"

type UserEvent string

const (
	EventRegister UserEvent = "register"
)

type UserEventListener = func(user *models.User)

var userEventListeners = map[UserEvent][]UserEventListener{}

func AddEventListener(event UserEvent, listener UserEventListener) {
	userEventListeners[event] = append(userEventListeners[event], listener)
}

func TriggerEvent(event UserEvent, user *models.User) {
	for _, listener := range userEventListeners[event] {
		listener(user)
	}
}
