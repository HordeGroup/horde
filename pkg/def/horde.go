package def

type HordeCreateRequest struct {
	Name      string `json:"name" form:"name" binding:"required"`
	Desc      string `json:"desc" form:"name" binding:"required"`
	CreatorId uint32 `json:"creator_id" form:"creator_id" binding:"required"`
}

type HordeCreateData struct {
	Id uint32 `json:"id"`
}
