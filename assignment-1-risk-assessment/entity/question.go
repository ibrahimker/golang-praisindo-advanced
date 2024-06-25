package entity

// Question adalah struktur untuk menyimpan pertanyaan beserta opsi jawaban dan bobotnya
type Question struct {
	ID       int
	Question string
	Options  []Option
}

// Option adalah struktur untuk menyimpan opsi jawaban beserta bobotnya
type Option struct {
	Answer string
	Weight int
}

// Questions adalah slice dari Question yang berisi daftar pertanyaan beserta opsi jawaban dan bobotnya
var Questions = []Question{
	{
		ID:       1,
		Question: "Apakah tujuan investasi Anda?",
		Options: []Option{
			{Answer: "Pertumbuhan kekayaan untuk jangka panjang", Weight: 5},
			{Answer: "Pendapatan dan pertumbuhan dalam jangka panjang", Weight: 4},
			{Answer: "Pendapatan berkala", Weight: 3},
			{Answer: "Pendapatan dan keamanan dana investasi", Weight: 2},
			{Answer: "Keamanan dana investasi", Weight: 1},
		},
	},
	{
		ID:       2,
		Question: "Berdasarkan tujuan investasi Anda, dana Anda akan diinvestasikan untuk jangka waktu?",
		Options: []Option{
			{Answer: "≥ 10 tahun", Weight: 5},
			{Answer: "7 - 10 tahun", Weight: 4},
			{Answer: "4 - ≥ 6 tahun", Weight: 3},
			{Answer: "1 - ≥ 3 tahun", Weight: 2},
			{Answer: "< 1 tahun", Weight: 1},
		},
	},
	{
		ID:       3,
		Question: "Berapa lama pengalaman Anda berinvestasi dalam produk yang nilainya berfluktuasi?",
		Options: []Option{
			{Answer: "> 10 tahun", Weight: 5},
			{Answer: "8 - 10 tahun", Weight: 4},
			{Answer: "4 - 7 tahun", Weight: 3},
			{Answer: "< 4 tahun", Weight: 2},
			{Answer: "0 tahun (tidak memiliki pengalaman)", Weight: 1},
		},
	},
	{
		ID:       4,
		Question: "Jenis investasi apa yang pernah Anda miliki?",
		Options: []Option{
			{Answer: "Saham, Reksa Dana terbuka, equity linked structure product", Weight: 5},
			{Answer: "Mata uang asing, currency linked structured product", Weight: 4},
			{Answer: "Uang tunai, deposito, produk dengan proteksi modal", Weight: 3},
		},
	},
	{
		ID:       5,
		Question: "Berapa persen dari aset Anda yang disimpan dalam produk investasi berfluktuasi?",
		Options: []Option{
			{Answer: "> 50%", Weight: 5},
			{Answer: "> 25% - ≥ 50%", Weight: 4},
			{Answer: "> 10% - ≥ 25%", Weight: 3},
			{Answer: "> 0% - ≥ 10%", Weight: 2},
			{Answer: "0%", Weight: 1},
		},
	},
	{
		ID:       6,
		Question: "Tingkat kenaikan dan penurunan nilai investasi yang dapat Anda terima?",
		Options: []Option{
			{Answer: "< -20% - > +20%", Weight: 5},
			{Answer: "-20% - +20%", Weight: 4},
			{Answer: "-15% - +15%", Weight: 3},
			{Answer: "-10% - +10%", Weight: 2},
			{Answer: "-5% - +5%", Weight: 1},
		},
	},
	{
		ID:       7,
		Question: "Ketergantungan Anda pada hasil investasi untuk biaya hidup sehari-hari?",
		Options: []Option{
			{Answer: "Tidak bergantung pada hasil investasi", Weight: 5},
			{Answer: "Tidak bergantung pada hasil investasi, minimal 5 tahun ke depan", Weight: 4},
			{Answer: "Sedikit bergantung pada hasil investasi", Weight: 3},
			{Answer: "Bergantung pada hasil investasi", Weight: 2},
			{Answer: "Sangat bergantung pada hasil investasi", Weight: 1},
		},
	},
	{
		ID:       8,
		Question: "Persentase pendapatan bulanan yang dapat Anda sisihkan untuk investasi/tabungan?",
		Options: []Option{
			{Answer: "> 50%", Weight: 5},
			{Answer: "> 25% - 50%", Weight: 4},
			{Answer: "> 10% - 25%", Weight: 3},
			{Answer: "> 0% - 10%", Weight: 2},
			{Answer: "0%", Weight: 1},
		},
	},
}
