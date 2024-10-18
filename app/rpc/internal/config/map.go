package config

func MapCropIdToCrop(cropId string) string {
	m := map[string]string{
		"ww30bd30b1f2d387b0": CropYm,
		"wwb8145373758bf6e1": CropYx,
		"ww456c1aa6d205b167": CropZx,
	}
	return m[cropId]
}

func MapCropToCropId(crop string) string {
	m := map[string]string{
		CropYm: "ww30bd30b1f2d387b0",
		CropYx: "wwb8145373758bf6e1",
		CropZx: "ww456c1aa6d205b167",
	}
	return m[crop]
}

func MapHistoryPlatformToCrop(platform int) string {
	m := map[int]string{
		1: CropYm,
		2: CropYx,
		3: CropZx,
	}
	return m[platform]
}

func MapCropToHistoryPlatform(crop string) int {
	m := map[string]int{
		CropYm: 1,
		CropYx: 2,
		CropZx: 2,
	}
	return m[crop]
}
