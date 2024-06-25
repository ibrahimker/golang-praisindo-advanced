package entity

import "time"

type ProfileRiskCategory string

const (
	ProfileRiskCategoryConservative ProfileRiskCategory = "Conservative"
	ProfileRiskCategoryModerate     ProfileRiskCategory = "Moderate"
	ProfileRiskCategoryBalanced     ProfileRiskCategory = "Balanced"
	ProfileRiskCategoryGrowth       ProfileRiskCategory = "Growth"
	ProfileRiskCategoryAggresive    ProfileRiskCategory = "Aggresive"
)

type Submission struct {
	ID             int                 `gorm:"primaryKey" json:"id"`
	UserID         int                 `json:"user_id"`
	User           User                `gorm:"foreignKey:UserID" json:"user"`
	RiskScore      int                 `json:"risk_score"`
	RiskCategory   ProfileRiskCategory `json:"risk_category"`
	RiskDefinition string              `json:"risk_definition"`
	Answers        []Answer            `gorm:"type:jsonb" json:"answers"`
	CreatedAt      time.Time           `json:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"`
}

type Answer struct {
	QuestionID int    `json:"question_id"` // ID pertanyaan atau nomor urut pertanyaan
	UserID     int    `json:"user_id"`     // ID pengguna yang menjawab pertanyaan ini
	Answer     string `json:"answer"`      // Jawaban dari pengguna, misalnya "a", "b", "c", dst.
}

type ProfileRisk struct {
	MinScore   int
	MaxScore   int
	Category   ProfileRiskCategory
	Definition string
}

var RiskMapping = []ProfileRisk{
	{
		MinScore: 0,
		MaxScore: 11,
		Category: ProfileRiskCategoryConservative,
		Definition: "Tujuan utama Anda adalah untuk melindungi modal/dana yang ditempatkan dan Anda tidak memiliki toleransi " +
			"sama sekali terhadap perubahan harga/nilai dari dana investasinya tersebut. " +
			"Anda memiliki pengalaman yang sangat terbatas atau tidak memiliki pengalaman sama sekali mengenai produk investasi.",
	},
	{
		MinScore:   12,
		MaxScore:   19,
		Category:   ProfileRiskCategoryModerate,
		Definition: "Anda memiliki toleransi yang rendah dengan perubahan harga/nilai dari dana investasi dan risiko investasi.",
	},
	{
		MinScore: 20,
		MaxScore: 28,
		Category: ProfileRiskCategoryBalanced,
		Definition: "Anda memiliki toleransi yang cukup terhadap produk investasi dan dapat menerima perubahan yang besar dari " +
			"harga/nilai dari harga yang diinvestasikan.",
	},
	{
		MinScore: 29,
		MaxScore: 35,
		Category: ProfileRiskCategoryGrowth,
		Definition: "Anda memiliki toleransi yang cukup tinggi dan dapat menerima perubahan yang besar dari harga/nilai portfolio" +
			"pada produk investasi yang diinvestasikan." +
			"Pada umumnya Anda sudah pernah atau berpengalaman dalam berinvestasi di produk investasi.",
	},
	{
		MinScore: 36,
		MaxScore: 40,
		Category: ProfileRiskCategoryAggresive,
		Definition: "Anda sangat berpengalaman terhadap produk investasi dan memiliki toleransi yang sangat tinggi atas" +
			"produk-produk investasi. Anda bahkan dapat menerima perubahan signifikan pada modal/nilai investasi." +
			"Pada umumnya portfolio Anda sebagian besar dialokasikan pada produk investasi.",
	},
}
