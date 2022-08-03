package auth

type (
	LoginParams struct {
		Identity string `json:"identity" form:"identity"`
		Password string `json:"password" form:"password"`
	}
)
