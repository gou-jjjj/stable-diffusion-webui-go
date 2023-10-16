package main

type Img2ImgRequest struct {
	BatchSize             int           `json:"batch_size"`
	CfgScale              int           `json:"cfg_scale"`
	DenoisingStrength     float64       `json:"denoising_strength"`
	Eta                   int           `json:"eta"`
	Height                int           `json:"height"`
	IncludeInitImages     interface{}   `json:"include_init_images"`
	InitImages            []interface{} `json:"init_images"`
	InpaintFullRes        interface{}   `json:"inpaint_full_res"`
	InpaintFullResPadding int           `json:"inpaint_full_res_padding"`
	InpaintingFill        int           `json:"inpainting_fill"`
	InpaintingMaskInvert  interface{}   `json:"inpainting_mask_invert"`
	Mask                  interface{}   `json:"mask"`
	MaskBlur              int           `json:"mask_blur"`
	NIter                 int           `json:"n_iter"`
	NegativePrompt        string        `json:"negative_prompt"`
	OverrideSettings      struct{}      `json:"override_settings"`
	Prompt                string        `json:"prompt"`
	ResizeMode            int           `json:"resize_mode"`
	RestoreFaces          interface{}   `json:"restore_faces"`
	SChurn                int           `json:"s_churn"`
	SNoise                int           `json:"s_noise"`
	STmax                 int           `json:"s_tmax"`
	STmin                 int           `json:"s_tmin"`
	SamplerIndex          string        `json:"sampler_index"`
	Seed                  int           `json:"seed"`
	SeedResizeFromH       int           `json:"seed_resize_from_h"`
	SeedResizeFromW       int           `json:"seed_resize_from_w"`
	Steps                 int           `json:"steps"`
	Styles                []interface{} `json:"styles"`
	Subseed               int           `json:"subseed"`
	SubseedStrength       int           `json:"subseed_strength"`
	Tiling                interface{}   `json:"tiling"`
	Width                 int           `json:"width"`
}

func (r *Img2ImgRequest) SetImage(I *Image) {
	r.InitImages[0] = I.Images[0]
}

type Img2ImgResponse struct {
	Images    []string     `json:"images"`
	Parameter ImgParameter `json:"parameter"`
	Info      string       `json:"info"`
}

type ImgParameter struct {
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
	OverrideSettings                  struct{}      `json:"override_settings"`
	OverrideSettingsRestoreAfterwards bool          `json:"override_settings_restore_afterwards"`
	RefinerCheckpoint                 interface{}   `json:"refiner_checkpoint"`
	RefinerSwitchAt                   interface{}   `json:"refiner_switch_at"`
	DisableExtraNetworks              bool          `json:"disable_extra_networks"`
	Comments                          interface{}   `json:"comments"`
	InitImages                        interface{}   `json:"init_images"`
	ResizeMode                        int           `json:"resize_mode"`
	ImageCfgScale                     interface{}   `json:"image_cfg_scale"`
	Mask                              interface{}   `json:"mask"`
	MaskBlurX                         int           `json:"mask_blur_x"`
	MaskBlurY                         int           `json:"mask_blur_y"`
	MaskBlur                          int           `json:"mask_blur"`
	InpaintingFill                    int           `json:"inpainting_fill"`
	InpaintFullRes                    bool          `json:"inpaint_full_res"`
	InpaintFullResPadding             int           `json:"inpaint_full_res_padding"`
	InpaintingMaskInvert              int           `json:"inpainting_mask_invert"`
	InitialNoiseMultiplier            interface{}   `json:"initial_noise_multiplier"`
	LatentMask                        interface{}   `json:"latent_mask"`
	SamplerIndex                      string        `json:"sampler_index"`
	IncludeInitImages                 bool          `json:"include_init_images"`
	ScriptName                        interface{}   `json:"script_name"`
	ScriptArgs                        []interface{} `json:"script_args"`
	SendImages                        bool          `json:"send_images"`
	SaveImages                        bool          `json:"save_images"`
	AlwaysonScripts                   struct{}      `json:"alwayson_scripts"`
}
