package views

// ViewUser is used to hold some values of models.User. This struct is 
// used to pass user data to context and then to templates.
// It is used instead of models.User to pass only certain data.
type ViewUser struct {
	Name  string
	Email string
}

// ViewData is used to construct template data. It takes ViewUser data, 
// if there is user, error message if there are errors 
// and other data to pass to the template.
type ViewData struct {
	User   *ViewUser
	ErrMsg string
	Data   interface{}
}

// SetViewData initializes ViewData and then passes it to the template.
func SetViewData(u *ViewUser, m string, data interface{}) *ViewData {
	return &ViewData{
		User:   u,
		ErrMsg: m,
		Data:   data,
	}
}
