package model

type MusicSearchReq struct {
	Keyword  string `json:"keyword"` // 模糊搜索字段（从歌曲名、歌手名称中模糊搜索）
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
}

type MusicSearchResp struct {
	Infos []*AuthedMetaInfo `json:"infos"`
	Total int64             `json:"total"`
}

// AuthedMetaInfo 对应 message MusicSearchResp 中的 message AuthedMetaInfo
type AuthedMetaInfo struct {
	MetadataId string `json:"metadata_id"` // metadata_id
	Title      string `json:"title"`       // 歌曲名
	ArtistName string `json:"artist_name"` // 歌手名称
}

type MusicMetaListReq struct {
	MetadataIds []string `json:"metadata_ids"` // metadata_id 列表
}

// MusicMetaInfo 音乐相关信息
type MusicMetaInfo struct {
	MetadataId string `json:"metadata_id"`
	Title      string `json:"title"`       // 歌曲标题
	ArtistName string `json:"artist_name"` // 歌手名称
	Album      string `json:"album"`       // 专辑名称
	CoverURL   string `json:"cover_url"`   // 封面地址
}

// MusicMetaListResp 对应 message MusicMetaListResp
type MusicMetaListResp struct {
	Infos map[string]*MusicMetaInfo `json:"infos"`
}

type MusicListReq struct {
	MetadataIds []string `json:"metadata_ids"`
}

type MusicListResp struct {
	Infos map[string]*TranscodeResourceList `json:"info"`
}

type TranscodeResourceList struct {
	Infos []*TranscodeResourceInfo `json:"transcode_resources_info"`
}

type TranscodeResourceInfo struct {
	ResourceId   string `json:"resource_id"`   // 资源ID
	TranscodeURL string `json:"transcode_url"` // 转码链接
	Duration     int64  `json:"duration"`      // 时长（微妙）
	Size         int64  `json:"size"`          // 资源大小
	Format       string `json:"format"`        // 码率格式
}
