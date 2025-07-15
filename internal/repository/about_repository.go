package repository

import "portfolio-arya-service/internal/domain"

type aboutRepository struct{}

func NewAboutRepository() domain.AboutRepository {
	return &aboutRepository{}
}

func (r *aboutRepository) GetAbout() (*domain.About, error) {
	about := &domain.About{
		Summary:   "I am a Frontend and Mobile Developer.",
		Languages: []string{"Go", "JavaScript", "Dart"},
		Education: []string{"Bachelor of Computer Science"},
		Projects:  []string{"Project A", "Project B"},
		Tools: []domain.Tool{
			{Name: "Android Studio", Icon: "android-studio-icon.png"},
			{Name: "VS Code", Icon: "vscode-icon.png"},
		},
	}

	return about, nil
}
