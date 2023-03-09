package exception

import "github.com/gofiber/fiber/v2"

type ErrorPayload struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// ThrowBadRequestException 400
func ThrowBadRequestException(c *fiber.Ctx, payload ...ErrorPayload) error {
	mappedData := getMappedErrorPayload("Bad request!", payload)
	return c.Status(fiber.ErrBadRequest.Code).JSON(mappedData)
}

// ThrowUnauthorizedException 401
func ThrowUnauthorizedException(c *fiber.Ctx, payload ...ErrorPayload) error {
	mappedData := getMappedErrorPayload("Unauthorized!", payload)
	return c.Status(fiber.ErrUnauthorized.Code).JSON(mappedData)
}

// ThrowForbiddenException 403
func ThrowForbiddenException(c *fiber.Ctx, payload ...ErrorPayload) error {
	mappedData := getMappedErrorPayload("Forbidden!", payload)
	return c.Status(fiber.ErrForbidden.Code).JSON(mappedData)
}

// ThrowNotFoundException 404
func ThrowNotFoundException(c *fiber.Ctx, payload ...ErrorPayload) error {
	mappedData := getMappedErrorPayload("Not found!", payload)
	return c.Status(fiber.ErrNotFound.Code).JSON(mappedData)
}

// ThrowUnprocessableEntityException 422
func ThrowUnprocessableEntityException(c *fiber.Ctx, payload ...ErrorPayload) error {
	mappedData := getMappedErrorPayload("Unprocessable entity!", payload)
	return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(mappedData)
}

func ThrowCustomException(c *fiber.Ctx, statusCode int, payload any) error {
	return c.Status(statusCode).JSON(payload)
}

func getMappedErrorPayload(message string, payload []ErrorPayload) ErrorPayload {
	mappedErrorPayload := ErrorPayload{
		Message: message,
	}

	if len(payload) > 0 {
		mappedErrorPayload.Data = payload[0].Data
	}

	return mappedErrorPayload
}
