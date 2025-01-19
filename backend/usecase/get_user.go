package usecase

import (
	"github.com/labstack/echo/v4"

	"github.com/Yamako76/go-react-sample/domain"
)

type GetUserRequest struct {
	ID int64 `param:"id"`
}

type GetUserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser(c echo.Context, userRepository domain.UserRepository) error {
	ctx := c.Request().Context()

	var request GetUserRequest
	if err := c.Bind(&request); err != nil {
		return badRequest(c, "リクエストが不正です", err)
	}

	user, err := userRepository.Get(ctx, request.ID)
	if err != nil {
		if err == domain.ErrUserNotFound {
			return newErrorResponse(c, 400, "UserNotFound", "ユーザーが見つかりませんでした", err)
		}
		return internalServerError(c, "ユーザーの取得に失敗しました", err)
	}

	response := GetUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(200, response)
}
