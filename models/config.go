package models

type Config struct {
	Device struct {
		DeviceName     string `json:"device_name"`
		DeviceType     int    `json:"device_type"`
		DeviceServices []int  `json:"device_services"`
	} `json:"device"`
	Redis struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"redis"`
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
	OnceDetector struct {
		ImagehashThreshold int     `json:"imagehash_threshold"`
		PsnrThreshold      float64 `json:"psnr_threshold"`
		SsimThreshold      float64 `json:"ssim_threshold"`
	} `json:"once_detector"`
	SourceReader struct {
		Fps        int `json:"fps"`
		BufferSize int `json:"buffer_size"`
		MaxRetry   int `json:"max_retry"`
		MaxRetryIn int `json:"max_retry_in"`
	} `json:"source_reader"`
	General struct {
		RootFolderPath    string `json:"root_folder_path"`
		HeartbeatInterval int    `json:"heartbeat_interval"`
	} `json:"general"`
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
}
