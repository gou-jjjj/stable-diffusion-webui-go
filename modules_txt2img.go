package main

type Txt2ImgRequest struct {
	BatchSize         int           `json:"batch_size"`
	CfgScale          int           `json:"cfg_scale"`
	DenoisingStrength int           `json:"denoising_strength"`
	EnableHr          bool          `json:"enable_hr"`
	Eta               int           `json:"eta"`
	FirstphaseHeight  int           `json:"firstphase_height"`
	FirstphaseWidth   int           `json:"firstphase_width"`
	Height            int           `json:"height"`
	NIter             int           `json:"n_iter"`
	NegativePrompt    string        `json:"negative_prompt"`
	Prompt            string        `json:"prompt"`
	RestoreFaces      bool          `json:"restore_faces"`
	SChurn            int           `json:"s_churn"`
	SNoise            int           `json:"s_noise"`
	STmax             int           `json:"s_tmax"`
	STmin             int           `json:"s_tmin"`
	SamplerIndex      string        `json:"sampler_index"`
	Seed              int           `json:"seed"`
	SeedResizeFromH   int           `json:"seed_resize_from_h"`
	SeedResizeFromW   int           `json:"seed_resize_from_w"`
	Steps             int           `json:"steps"`
	Styles            []interface{} `json:"styles"`
	Subseed           int           `json:"subseed"`
	SubseedStrength   int           `json:"subseed_strength"`
	Tiling            bool          `json:"tiling"`
	Width             int           `json:"width"`
}

type Txt2imgResponse struct {
	Images    []string     `json:"images"`
	Parameter TxtParameter `json:"parameters"`
	Info      string       `json:"info"`
}

type TxtParameter struct {
	Prompt                            string        `json:"prompt"`
	NegativePrompt                    string        `json:"negative_prompt"`
	Styles                            []interface{} `json:"styles"`
	Seed                              int           `json:"seed"`
	Subseed                           int           `json:"subseed"`
	SubseedStrength                   float64       `json:"subseed_strength"`
	SeedResizeFromH                   int           `json:"seed_resize_from_h"`
	SeedResizeFromW                   int           `json:"seed_resize_from_w"`
	SamplerName                       interface{}   `json:"sampler_name"`
	BatchSize                         int           `json:"batch_size"`
	NIter                             int           `json:"n_iter"`
	Steps                             int           `json:"steps"`
	CfgScale                          float64       `json:"cfg_scale"`
	Width                             int           `json:"width"`
	Height                            int           `json:"height"`
	RestoreFaces                      bool          `json:"restore_faces"`
	Tiling                            bool          `json:"tiling"`
	DoNotSaveSamples                  bool          `json:"do_not_save_samples"`
	DoNotSaveGrid                     bool          `json:"do_not_save_grid"`
	Eta                               float64       `json:"eta"`
	DenoisingStrength                 float64       `json:"denoising_strength"`
	SMinUncond                        interface{}   `json:"s_min_uncond"`
	SChurn                            float64       `json:"s_churn"`
	STmax                             float64       `json:"s_tmax"`
	STmin                             float64       `json:"s_tmin"`
	SNoise                            float64       `json:"s_noise"`
	OverrideSettings                  interface{}   `json:"override_settings"`
	OverrideSettingsRestoreAfterwards bool          `json:"override_settings_restore_afterwards"`
	RefinerCheckpoint                 interface{}   `json:"refiner_checkpoint"`
	RefinerSwitchAt                   interface{}   `json:"refiner_switch_at"`
	DisableExtraNetworks              bool          `json:"disable_extra_networks"`
	Comments                          interface{}   `json:"comments"`
	EnableHr                          bool          `json:"enable_hr"`
	FirstphaseWidth                   int           `json:"firstphase_width"`
	FirstphaseHeight                  int           `json:"firstphase_height"`
	HrScale                           float64       `json:"hr_scale"`
	HrUpscaler                        interface{}   `json:"hr_upscaler"`
	HrSecondPassSteps                 int           `json:"hr_second_pass_steps"`
	HrResizeX                         int           `json:"hr_resize_x"`
	HrResizeY                         int           `json:"hr_resize_y"`
	HrCheckpointName                  interface{}   `json:"hr_checkpoint_name"`
	HrSamplerName                     interface{}   `json:"hr_sampler_name"`
	HrPrompt                          string        `json:"hr_prompt"`
	HrNegativePrompt                  string        `json:"hr_negative_prompt"`
	SamplerIndex                      string        `json:"sampler_index"`
	ScriptName                        interface{}   `json:"script_name"`
	ScriptArgs                        []interface{} `json:"script_args"`
	SendImages                        bool          `json:"send_images"`
	SaveImages                        bool          `json:"save_images"`
	AlwaysonScripts                   struct{}      `json:"alwayson_scripts"`
}
