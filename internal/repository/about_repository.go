package repository

import "portfolio-arya-service/internal/domain"

type aboutRepository struct{}

var aboutData = &domain.About{
	Description: "Initial description",
	Summary:     "Initial summary",
	Photo:       "path/to/photo.jpg",
	Languages:   []string{"Go", "JavaScript"},
	Education:   []string{"Bachelor of Computer Science"},
	Projects:    []string{"Project A", "Project B"},
	Tools: []domain.Tool{
		{Id: "1", Name: "Git", Icon: "git-icon.png"},
	},
}

func (r *aboutRepository) UpdateAbout(about domain.About) (*domain.About, error) {
	aboutData = &about
	return aboutData, nil
}

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
