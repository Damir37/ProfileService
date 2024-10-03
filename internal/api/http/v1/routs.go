package v1

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, profileHandler *ProfileHandler) {
	router.Use(metricsMiddleware())

	profileGroupV1 := router.Group("/v1")
	{
		profileGroupV1.GET("/profile/:id", profileHandler.GetProfile)
		profileGroupV1.GET("/stream-key/:id", profileHandler.GetStreamKey)
		profileGroupV1.GET("/profile-ava/:id", profileHandler.GetAvatarProfile)
		profileGroupV1.GET("/profile-stat/:id", profileHandler.GetProfileStats)

		profileGroupV1.PATCH("/profile-setadmin/:id", profileHandler.SetAdmin)
		profileGroupV1.PATCH("/profile-unsetadmin/:id", profileHandler.UnsetAdmin)
		profileGroupV1.PATCH("/stream-rndkey/:id", profileHandler.RandomStreamKey)
		profileGroupV1.PATCH("/profile-avarnd", profileHandler.EditAvatar)
		profileGroupV1.PATCH("/profile-edit", profileHandler.EditProfile)
		profileGroupV1.PATCH("/profile-modifyelo", profileHandler.ModifyElo)
	}
}
