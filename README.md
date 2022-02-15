# go-mini-starter

[ 🛠 ***this is under construction*** ]

## Golang minimal web starter boilerplate with MongoDB.

- [x] Routing with ```go-chi/chi```

- [x] Request logging with ```go-chi/chi/middleware/Logger```

- [x] Simple templating structure

- [x] ```Tailwind CSS``` for styling

- [x] Vanilla JavaScript client side form validations

- [x] Simple error handling

- [x] Server side validations with ```go-ozzo/ozzo-validation```

- [x] Password hash with ```x/crypto/bcrypt```

- [x] Sessions and cookies with ```x/crypto/rand``` and ```x/crypto/hmac```

- [x] CSRF/XSRF with ```gorilla/csrf```

- [x] CSS/XSS

## App structure

```shell
|---contexts
|   |---usercontext.go
|---handlers
|   |---signinwithcookie.go
|   |---static.go
|   |---user.go
|---helpers
|   |---errors.go
|   |---hashstring.go
|   |---normalize.go
|   |---tokens.go
|   |---validate.go
|---middlewares
|   |---checkuser.go
|   |---loggeduser.go
|   |---requireuser.go
|---models
|   |---dbconnect.go
|   |---user.go
|---static
|   |---css
|       |---style.css
|   |---img
|   |   |---forest.jpg
|   |---js
|   |   |---main.js
|   |---favicon.ico
|---views
|   |---templates
|   |   |---layouts
|   |   |   |---base.html
|   |   |   |---footer.html
|   |   |   |---navbar.html
|   |   |---user
|   |   |   |---dashboard.html
|   |   |   |---login.html
|   |   |   |---signup.html
|   |   |---contacts.html
|   |   |---home.html
|   |---view.go
|   |---viewdata.go
|---.env
|---.gitignore
|---create-user.png
|---main.go
|---Makefile
|---README.md
|---routes.go
|---static.png
```

## Request-Response cycle of the static page

![static page request-response](/static.png "Static page Request-Response")

## Request-Response cycle of the signup page

![signup page request-response](/create-user.png "Signup page Request-Response")

## License

MIT license

