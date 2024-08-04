package po

type Model struct {
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64 `gorm:"column:updated_at;autoUpdateTime:milli"`
	DeletedAt int64 `gorm:"column:deleted_at"`
}
