package presenter

import (
	"fmt"
	"net/http"

	"github.com/katsu105/clean-todo/entity"
	"github.com/katsu105/clean-todo/usecase/port"
)

type User struct {
	w http.ResponseWriter
}

func NewUserOutputPort(w http.ResponseWriter) port.UserOutputPort {
	return &User{
		w: w,
	}
}

func (u *User) Render(user *entity.User) {
	u.w.WriteHeader(http.StatusOK)

	fmt.Fprint(u.w, user.Name)
}

func (u *User) RenderError(err error) {
	u.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(u.w, err)
}
