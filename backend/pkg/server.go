package pkg

import (
	"jangle/backend/auth"
)

type Server struct {
	auth.UnimplementedAuthenticationServer
}
