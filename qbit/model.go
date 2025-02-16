package qbit

type Torrent struct {
	AddedOn           int     `json:"added_on"`
	AmountLeft        int     `json:"amount_left"`
	AutoTMM           bool    `json:"auto_tmm"`
	Availability      float64 `json:"availability"`
	Category          string  `json:"category"`
	Completed         int     `json:"completed"`
	CompletionOn      int     `json:"completion_on"`
	ContentPath       string  `json:"content_path"`
	DlLimit           int     `json:"dl_limit"`
	Dlspeed           int     `json:"dlspeed"`
	Downloaded        int     `json:"downloaded"`
	DownloadedSession int     `json:"downloaded_session"`
	Eta               int     `json:"eta"`
	FLPiecePrio       bool    `json:"f_l_piece_prio"`
	ForceStart        bool    `json:"force_start"`
	Hash              string  `json:"hash"`
	IsPrivate         bool    `json:"isPrivate"`
	LastActivity      int     `json:"last_activity"`
	MagnetURI         string  `json:"magnet_uri"`
	MaxRatio          float64 `json:"max_ratio"`
	MaxSeedingTime    int     `json:"max_seeding_time"`
	Name              string  `json:"name"`
	NumComplete       int     `json:"num_complete"`
	NumIncomplete     int     `json:"num_incomplete"`
	NumLeechs         int     `json:"num_leechs"`
	NumSeeds          int     `json:"num_seeds"`
	Priority          int     `json:"priority"`
	Progress          float64 `json:"progress"`
	Ratio             float64 `json:"ratio"`
	RatioLimit        float64 `json:"ratio_limit"`
	SavePath          string  `json:"save_path"`
	SeedingTime       int     `json:"seeding_time"`
	SeedingTimeLimit  int     `json:"seeding_time_limit"`
	SeenComplete      int     `json:"seen_complete"`
	SeqDL             bool    `json:"seq_dl"`
	Size              int     `json:"size"`
	State             string  `json:"state"`
	SuperSeeding      bool    `json:"super_seeding"`
	Tags              string  `json:"tags"`
	TimeActive        int     `json:"time_active"`
	TotalSize         int     `json:"total_size"`
	Tracker           string  `json:"tracker"`
	UpLimit           int     `json:"up_limit"`
	Uploaded          int     `json:"uploaded"`
	UploadedSession   int     `json:"uploaded_session"`
	Upspeed           int     `json:"upspeed"`
}

type TorrentGeneric struct {
	SavePath               string  `json:"save_path"`
	CreationDate           int     `json:"creation_date"`
	PieceSize              int     `json:"piece_size"`
	Comment                string  `json:"comment"`
	TotalWasted            int     `json:"total_wasted"`
	TotalUploaded          int     `json:"total_uploaded"`
	TotalUploadedSession   int     `json:"total_uploaded_session"`
	TotalDownloaded        int     `json:"total_downloaded"`
	TotalDownloadedSession int     `json:"total_downloaded_session"`
	UpLimit                int     `json:"up_limit"`
	DlLimit                int     `json:"dl_limit"`
	TimeElapsed            int     `json:"time_elapsed"`
	SeedingTime            int     `json:"seeding_time"`
	NbConnections          int     `json:"nb_connections"`
	NbConnectionsLimit     int     `json:"nb_connections_limit"`
	ShareRatio             float64 `json:"share_ratio"`
	AdditionDate           int     `json:"addition_date"`
	CompletionDate         int     `json:"completion_date"`
	CreatedBy              string  `json:"created_by"`
	DlSpeedAvg             int     `json:"dl_speed_avg"`
	DlSpeed                int     `json:"dl_speed"`
	Eta                    int     `json:"eta"`
	LastSeen               int     `json:"last_seen"`
	Peers                  int     `json:"peers"`
	PeersTotal             int     `json:"peers_total"`
	PiecesHave             int     `json:"pieces_have"`
	PiecesNum              int     `json:"pieces_num"`
	Reannounce             int     `json:"reannounce"`
	Seeds                  int     `json:"seeds"`
	SeedsTotal             int     `json:"seeds_total"`
	TotalSize              int     `json:"total_size"`
	UpSpeedAvg             int     `json:"up_speed_avg"`
	UpSpeed                int     `json:"up_speed"`
	IsPrivate              bool    `json:"isPrivate"`
}

type Tracker struct {
	URL           string `json:"url"`
	Status        int    `json:"status"`
	Tier          int    `json:"tier"`
	NumPeers      int    `json:"num_peers"`
	NumSeeds      int    `json:"num_seeds"`
	NumLeeches    int    `json:"num_leeches"`
	NumDownloaded int    `json:"num_downloaded"`
	Msg           string `json:"msg"`
}

type WebSeed struct {
	URL string `json:"url"`
}

type File struct {
	Index        int     `json:"index"`
	Name         string  `json:"name"`
	Size         int64   `json:"size"`
	Progress     float64 `json:"progress"`
	Priority     int     `json:"priority"`
	IsSeed       bool    `json:"is_seed"`
	PieceRange   []int   `json:"piece_range"`
	Availability float64 `json:"availability"`
}

type Preferences struct {
	AddTrackers                      string         `json:"add_trackers"`
	AddTrackersEnabled               bool           `json:"add_trackers_enabled"`
	AltDlLimit                       int            `json:"alt_dl_limit"`
	AltUpLimit                       int            `json:"alt_up_limit"`
	AlternativeWebuiEnabled          bool           `json:"alternative_webui_enabled"`
	AlternativeWebuiPath             string         `json:"alternative_webui_path"`
	AnnounceIP                       string         `json:"announce_ip"`
	AnnounceToAllTiers               bool           `json:"announce_to_all_tiers"`
	AnnounceToAllTrackers            bool           `json:"announce_to_all_trackers"`
	AnonymousMode                    bool           `json:"anonymous_mode"`
	AsyncIOThreads                   int            `json:"async_io_threads"`
	AutoDeleteMode                   int            `json:"auto_delete_mode"`
	AutoTMMEnabled                   bool           `json:"auto_tmm_enabled"`
	AutorunEnabled                   bool           `json:"autorun_enabled"`
	AutorunProgram                   string         `json:"autorun_program"`
	BannedIPs                        string         `json:"banned_IPs"`
	BittorrentProtocol               int            `json:"bittorrent_protocol"`
	ByPassAuthSubnetWhitelist        string         `json:"bypass_auth_subnet_whitelist"`
	ByPassAuthSubnetWhitelistEnabled bool           `json:"bypass_auth_subnet_whitelist_enabled"`
	ByPassLocalAuth                  bool           `json:"bypass_local_auth"`
	CategoryChangedTMMEnabled        bool           `json:"category_changed_tmm_enabled"`
	CheckingMemoryUse                int            `json:"checking_memory_use"`
	CreateSubfolderEnabled           bool           `json:"create_subfolder_enabled"`
	CurrentInterfaceAddress          string         `json:"current_interface_address"`
	CurrentNetworkInterface          string         `json:"current_network_interface"`
	DHT                              bool           `json:"dht"`
	DiskCache                        int            `json:"disk_cache"`
	DiskCacheTTL                     int            `json:"disk_cache_ttl"`
	DLLimit                          int            `json:"dl_limit"`
	DontCountSlowTorrents            bool           `json:"dont_count_slow_torrents"`
	DyndnsDomain                     string         `json:"dyndns_domain"`
	DyndnsEnabled                    bool           `json:"dyndns_enabled"`
	DyndnsPassword                   string         `json:"dyndns_password"`
	DyndnsService                    int            `json:"dyndns_service"`
	DyndnsUsername                   string         `json:"dyndns_username"`
	EmbeddedTrackerPort              int            `json:"embedded_tracker_port"`
	EnableCoalesceReadWrite          bool           `json:"enable_coalesce_read_write"`
	EnableEmbeddedTracker            bool           `json:"enable_embedded_tracker"`
	EnableMultiConnectionsFromSameIP bool           `json:"enable_multi_connections_from_same_ip"`
	EnableOSCache                    bool           `json:"enable_os_cache"`
	EnablePieceExtentAffinity        bool           `json:"enable_piece_extent_affinity"`
	EnableUploadSuggestions          bool           `json:"enable_upload_suggestions"`
	Encryption                       int            `json:"encryption"`
	ExportDir                        string         `json:"export_dir"`
	ExportDirFin                     string         `json:"export_dir_fin"`
	FilePoolSize                     int            `json:"file_pool_size"`
	IncompleteFilesExt               bool           `json:"incomplete_files_ext"`
	IPFilterEnabled                  bool           `json:"ip_filter_enabled"`
	IPFilterPath                     string         `json:"ip_filter_path"`
	IPFilterTrackers                 bool           `json:"ip_filter_trackers"`
	LimitLANPeers                    bool           `json:"limit_lan_peers"`
	LimitTCPOverhead                 bool           `json:"limit_tcp_overhead"`
	LimitUTPRate                     bool           `json:"limit_utp_rate"`
	ListenPort                       int            `json:"listen_port"`
	Locale                           string         `json:"locale"`
	LSD                              bool           `json:"lsd"`
	MailNotificationAuthEnabled      bool           `json:"mail_notification_auth_enabled"`
	MailNotificationEmail            string         `json:"mail_notification_email"`
	MailNotificationEnabled          bool           `json:"mail_notification_enabled"`
	MailNotificationPassword         string         `json:"mail_notification_password"`
	MailNotificationSender           string         `json:"mail_notification_sender"`
	MailNotificationSMTP             string         `json:"mail_notification_smtp"`
	MailNotificationSSLEnabled       bool           `json:"mail_notification_ssl_enabled"`
	MailNotificationUsername         string         `json:"mail_notification_username"`
	MaxActiveDownloads               int            `json:"max_active_downloads"`
	MaxActiveTorrents                int            `json:"max_active_torrents"`
	MaxActiveUploads                 int            `json:"max_active_uploads"`
	MaxConnec                        int            `json:"max_connec"`
	MaxConnecPerTorrent              int            `json:"max_connec_per_torrent"`
	MaxRatio                         float64        `json:"max_ratio"`
	MaxRatioAct                      int            `json:"max_ratio_act"`
	MaxRatioEnabled                  bool           `json:"max_ratio_enabled"`
	MaxSeedingTime                   int            `json:"max_seeding_time"`
	MaxSeedingTimeEnabled            bool           `json:"max_seeding_time_enabled"`
	MaxUploads                       int            `json:"max_uploads"`
	MaxUploadsPerTorrent             int            `json:"max_uploads_per_torrent"`
	OutgoingPortsMax                 int            `json:"outgoing_ports_max"`
	OutgoingPortsMin                 int            `json:"outgoing_ports_min"`
	PEX                              bool           `json:"pex"`
	PreallocateAll                   bool           `json:"preallocate_all"`
	ProxyAuthEnabled                 bool           `json:"proxy_auth_enabled"`
	ProxyIP                          string         `json:"proxy_ip"`
	ProxyPassword                    string         `json:"proxy_password"`
	ProxyPeerConnections             bool           `json:"proxy_peer_connections"`
	ProxyPort                        int            `json:"proxy_port"`
	ProxyTorrentsOnly                bool           `json:"proxy_torrents_only"`
	ProxyType                        int            `json:"proxy_type"`
	ProxyUsername                    string         `json:"proxy_username"`
	QueueingEnabled                  bool           `json:"queueing_enabled"`
	RandomPort                       bool           `json:"random_port"`
	RecheckCompletedTorrents         bool           `json:"recheck_completed_torrents"`
	ResolvePeerCountries             bool           `json:"resolve_peer_countries"`
	RSSAutoDownloadingEnabled        bool           `json:"rss_auto_downloading_enabled"`
	RSSDownloadRepackProperEpisodes  bool           `json:"rss_download_repack_proper_episodes"`
	RSSMaxArticlesPerFeed            int            `json:"rss_max_articles_per_feed"`
	RSSProcessingEnabled             bool           `json:"rss_processing_enabled"`
	RSSRefreshInterval               int            `json:"rss_refresh_interval"`
	RSSSmartEpisodeFilters           string         `json:"rss_smart_episode_filters"`
	SavePath                         string         `json:"save_path"`
	SavePathChangedTMMEnabled        bool           `json:"save_path_changed_tmm_enabled"`
	SaveResumeDataInterval           int            `json:"save_resume_data_interval"`
	ScanDirs                         map[string]int `json:"scan_dirs"`
	SchedulerEnabled                 bool           `json:"scheduler_enabled"`
	SchedulerDays                    int            `json:"scheduler_days"`
	ScheduleFromHour                 int            `json:"schedule_from_hour"`
	ScheduleFromMin                  int            `json:"schedule_from_min"`
	ScheduleToHour                   int            `json:"schedule_to_hour"`
	ScheduleToMin                    int            `json:"schedule_to_min"`
	SendBufferLowWatermark           int            `json:"send_buffer_low_watermark"`
	SendBufferWatermark              int            `json:"send_buffer_watermark"`
	SendBufferWatermarkFactor        int            `json:"send_buffer_watermark_factor"`
	SlowTorrentDLRateThreshold       int            `json:"slow_torrent_dl_rate_threshold"`
	SlowTorrentInactiveTimer         int            `json:"slow_torrent_inactive_timer"`
	SlowTorrentULRateThreshold       int            `json:"slow_torrent_ul_rate_threshold"`
	SocketBacklogSize                int            `json:"socket_backlog_size"`
	StartPausedEnabled               bool           `json:"start_paused_enabled"`
	StopTrackerTimeout               int            `json:"stop_tracker_timeout"`
	TempPath                         string         `json:"temp_path"`
	TempPathEnabled                  bool           `json:"temp_path_enabled"`
	TorrentChangedTMMEnabled         bool           `json:"torrent_changed_tmm_enabled"`
	UpLimit                          int            `json:"up_limit"`
	UploadChokingAlgorithm           int            `json:"upload_choking_algorithm"`
}

type TransferInfo struct {
	DlInfoSpeed      int    `json:"dl_info_speed"`
	DlInfoData       int    `json:"dl_info_data"`
	UpInfoSpeed      int    `json:"up_info_speed"`
	UpInfoData       int    `json:"up_info_data"`
	DlRateLimit      int    `json:"dl_rate_limit"`
	UpRateLimit      int    `json:"up_rate_limit"`
	DhtNodes         int    `json:"dht_nodes"`
	ConnectionStatus string `json:"connection_status"`
}

type SpeedLimitsMode struct {
	Enabled int `json:"enabled"`
}

type Limit struct {
	Limit int `json:"limit"`
}
