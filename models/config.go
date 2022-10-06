package models

type Config struct {
	Device struct {
		DeviceName string `json:"device_name"`
		DeviceType int    `json:"device_type"`
	} `json:"device"`
	Jetson struct {
		ModelName string `json:"model_name"`
	} `json:"jetson"`
	Torch struct {
		ModelName         string `json:"model_name"`
		ModelNameSpecific string `json:"model_name_specific"`
	} `json:"torch"`
	Tensorflow struct {
		ModelName   string `json:"model_name"`
		CacheFolder string `json:"cache_folder"`
	} `json:"tensorflow"`
	SourceReader struct {
		ResizeImg  bool `json:"resize_img"`
		BufferSize int  `json:"buffer_size"`
		MaxRetry   int  `json:"max_retry"`
		MaxRetryIn int  `json:"max_retry_in"`
	} `json:"source_reader"`
	General struct {
		RootFolderPath    string `json:"root_folder_path"`
		HeartbeatInterval int    `json:"heartbeat_interval"`
	} `json:"general"`
	Db struct {
		Type             int    `json:"type"` // 0 is SQLite, 1 is MongoDB
		ConnectionString string `json:"connection_string"`
	} `json:"db"`
	FFmpeg struct {
		UseDoubleQuotesForPath         bool    `json:"use_double_quotes_for_path"`
		MaxOperationRetryCount         int     `json:"max_operation_retry_count"`
		RtmpServerInitInterval         float32 `json:"rtmp_server_init_interval"`
		WatchDogInterval               int     `json:"watch_dog_interval"`
		WatchDogFailedWaitInterval     float32 `json:"watch_dog_failed_wait_interval"`
		StartTaskWaitForInterval       float32 `json:"start_task_wait_for_interval"`
		RecordConcatLimit              int     `json:"record_concat_limit"`
		RecordVideoFileIndexerInterval int     `json:"record_video_file_indexer_interval"`
	} `json:"ffmpeg"`
	Ai struct {
		Overlay                 bool    `json:"overlay"`
		VideoClipDuration       int     `json:"video_clip_duration"`
		FaceRecogMtcnnThreshold float32 `json:"face_recog_mtcnn_threshold"`
		FaceRecogProbThreshold  float32 `json:"face_recog_prob_threshold"`
		PlateRecogInstanceCount int     `json:"plate_recog_instance_count"`
	} `json:"ai"`
	Ui struct {
		GsWidth                int     `json:"gs_width"`
		GsHeight               int     `json:"gs_height"`
		BoosterInterval        float32 `json:"booster_interval"`
		SeekToLiveEdgeInternal int     `json:"seek_to_live_edge_internal"`
	} `json:"ui"`
	Jobs struct {
		MacIpMatchingEnabled  bool `json:"mac_ip_matching_enabled"`
		MacIpMatchingInterval int  `json:"mac_ip_matching_interval"`
	} `json:"jobs"`
	DeepStack struct {
		ServerUrl       string  `json:"server_url"`
		ServerPort      int     `json:"server_port"`
		PerformanceMode int     `json:"performance_mode"`
		ApiKey          string  `json:"api_key"`
		OdEnabled       bool    `json:"od_enabled"`
		OdThreshold     float32 `json:"od_threshold"`
		FrEnabled       bool    `json:"fr_enabled"`
		FrThreshold     float32 `json:"fr_threshold"`
		DockerType      int     `json:"docker_type"`
	} `json:"deep_stack"`
	Archive struct {
		LimitPercent int    `json:"limit_percent"`
		ActionType   int    `json:"action_type"`
		MoveLocation string `json:"move_location"`
	} `json:"archive"`
	Snapshot struct {
		ProcessCount int `json:"process_count"`
	} `json:"snapshot"`
}
