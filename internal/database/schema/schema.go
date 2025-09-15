package schema

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Actor -> tabela "actors"
type Actor struct {
	ID          int          `gorm:"primaryKey;autoIncrement"`
	Name        string       `gorm:"size:255;unique;not null"`
	MovieActors []MovieActor `gorm:"foreignKey:ActorID"`
	Movies      []Movie      `gorm:"many2many:movieactor;joinForeignKey:ActorID;joinReferences:MovieID"`
}

// Company -> tabela "companies"
type Company struct {
	ID             int            `gorm:"primaryKey;autoIncrement"`
	Name           string         `gorm:"size:255;unique;not null"`
	MovieCompanies []MovieCompany `gorm:"foreignKey:CompanyID"`
	Movies         []Movie        `gorm:"many2many:moviecompany;joinForeignKey:CompanyID;joinReferences:MovieID"`
}

func (Company) TableName() string { return "companies" }

// Director -> tabela "directors"
type Director struct {
	ID             int             `gorm:"primaryKey;autoIncrement"`
	Name           string          `gorm:"size:255;unique;not null"`
	MovieDirectors []MovieDirector `gorm:"foreignKey:DirectorID"`
	Movies         []Movie         `gorm:"many2many:moviedirector;joinForeignKey:DirectorID;joinReferences:MovieID"`
}

func (Director) TableName() string { return "directors" }

// Genre -> tabela "genres"
type Genre struct {
	ID          int          `gorm:"primaryKey;autoIncrement"`
	Name        string       `gorm:"size:255;unique;not null"`
	MovieGenres []MovieGenre `gorm:"foreignKey:GenreID"`
	Movies      []Movie      `gorm:"many2many:moviegenre;joinForeignKey:GenreID;joinReferences:MovieID"`
}

func (Genre) TableName() string { return "genres" }

// Movie -> tabela "movie"
type Movie struct {
	ID    int    `gorm:"primaryKey;autoIncrement"`
	Title string `gorm:"size:255;not null;uniqueIndex:ux_movie_title_release_date"`
	//Poster         string           `gorm:"type:longtext"`
	ReleaseDate    *time.Time       `gorm:"type:date;uniqueIndex:ux_movie_title_release_date"` // nullable
	Budget         *decimal.Decimal `gorm:"type:decimal(15,2)"`
	TicketOffice   *decimal.Decimal `gorm:"type:decimal(15,2)"`
	VoteAverage    *decimal.Decimal `gorm:"type:decimal(15,2)"`
	SearchMovies   []SearchMovie    `gorm:"foreignKey:MovieID"`
	MovieActors    []MovieActor     `gorm:"foreignKey:MovieID"`
	MovieCompanies []MovieCompany   `gorm:"foreignKey:MovieID"`
	MovieDirectors []MovieDirector  `gorm:"foreignKey:MovieID"`
	MovieGenres    []MovieGenre     `gorm:"foreignKey:MovieID"`
	ClassicGames   []ClassicGame    `gorm:"foreignKey:MovieID"`
}

func (Movie) TableName() string { return "movie" }

// SearchMovie -> tabela "searchmovie"
// movieId é PK (1:1 com movie.id na sua schema Prisma)
type SearchMovie struct {
	MovieID int    `gorm:"primaryKey"`
	Title   string `gorm:"size:255;not null"`
	Movie   Movie  `gorm:"foreignKey:MovieID;references:ID;constraint:OnUpdate:RESTRICT"`
}

func (SearchMovie) TableName() string { return "searchmovie" }

// MovieActor -> tabela de junção "movieactor" com PK composta (movieId, actorId)
type MovieActor struct {
	MovieID int   `gorm:"primaryKey;index:idx_movieactor_movieid"`
	ActorID int   `gorm:"primaryKey;index:idx_movieactor_actorid"`
	Movie   Movie `gorm:"foreignKey:MovieID;references:ID;constraint:OnUpdate:RESTRICT"`
	Actor   Actor `gorm:"foreignKey:ActorID;references:ID;constraint:OnUpdate:RESTRICT"`
}

func (MovieActor) TableName() string { return "movieactor" }

// MovieCompany -> join table (movieId, companyId)
type MovieCompany struct {
	MovieID   int     `gorm:"primaryKey;index:idx_moviecompany_movieid"`
	CompanyID int     `gorm:"primaryKey;index:idx_moviecompany_companyid"`
	Movie     Movie   `gorm:"foreignKey:MovieID;references:ID;constraint:OnUpdate:RESTRICT"`
	Company   Company `gorm:"foreignKey:CompanyID;references:ID;constraint:OnUpdate:RESTRICT"`
}

func (MovieCompany) TableName() string { return "moviecompany" }

// MovieDirector -> join table (movieId, directorId)
type MovieDirector struct {
	MovieID    int      `gorm:"primaryKey;index:idx_moviedirector_movieid"`
	DirectorID int      `gorm:"primaryKey;index:idx_moviedirector_directorid"`
	Movie      Movie    `gorm:"foreignKey:MovieID;references:ID;constraint:OnUpdate:RESTRICT"`
	Director   Director `gorm:"foreignKey:DirectorID;references:ID;constraint:OnUpdate:RESTRICT"`
}

func (MovieDirector) TableName() string { return "moviedirector" }

// MovieGenre -> join table (movieId, genreId)
type MovieGenre struct {
	MovieID int   `gorm:"primaryKey;index:idx_moviegenre_movieid"`
	GenreID int   `gorm:"primaryKey;index:idx_moviegenre_genreid"`
	Movie   Movie `gorm:"foreignKey:MovieID;references:ID;constraint:OnUpdate:RESTRICT"`
	Genre   Genre `gorm:"foreignKey:GenreID;references:ID;constraint:OnUpdate:RESTRICT"`
}

func (MovieGenre) TableName() string { return "moviegenre" }

// ClassicGame -> tabela "classicgame" com movieId como PK e relação com movie
type ClassicGame struct {
	MovieID int       `gorm:"primaryKey"`
	Title   string    `gorm:"size:255;not null"`
	Date    time.Time `gorm:"type:date;not null"`
	Movie   Movie     `gorm:"foreignKey:MovieID;references:ID;constraint:OnUpdate:RESTRICT"`
}

func (ClassicGame) TableName() string { return "classicgame" }

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(

		&Movie{},
	)
}
