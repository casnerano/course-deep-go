package interfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDIContainer(t *testing.T) {
	type UserService struct {
		NotEmptyStruct bool
	}
	type MessageService struct {
		NotEmptyStruct bool
	}
	type ExternalConnection struct {
		NotEmptyStruct bool
	}

	container := NewContainer()
	container.RegisterType("UserService", func() interface{} {
		return &UserService{}
	})
	container.RegisterType("MessageService", func() interface{} {
		return &MessageService{}
	})

	userService1, err := container.Resolve("UserService")
	assert.NoError(t, err)
	userService2, err := container.Resolve("UserService")
	assert.NoError(t, err)

	u1 := userService1.(*UserService)
	u2 := userService2.(*UserService)
	assert.False(t, u1 == u2)

	messageService, err := container.Resolve("MessageService")
	assert.NoError(t, err)
	assert.NotNil(t, messageService)

	paymentService, err := container.Resolve("PaymentService")
	assert.Error(t, err)
	assert.Nil(t, paymentService)

	container.RegisterSingletonType("ExternalConnection", func() interface{} {
		return &ExternalConnection{}
	})

	conn1, err := container.Resolve("ExternalConnection")
	assert.NoError(t, err)
	assert.NotNil(t, conn1)

	conn2, err := container.Resolve("ExternalConnection")
	assert.NoError(t, err)
	assert.NotNil(t, conn2)

	assert.Equal(t, conn1, conn2)
}
