package main

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/morum/gerador-parecer/parser"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx          context.Context
	relacaoFile  string
	modeloFile   string
	dirSaida     string
	nomePlanilha string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		nomePlanilha: "Relação de processos",
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) selectFile(fileType, extension string) string {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: fileType,
				Pattern:     extension,
			},
		},
	})
	if err != nil {
		return err.Error()
	}

	return file
}

// SelectRelacaoFile seleciona o arquivo com a relação
func (a *App) SelectRelacaoFile() string {
	a.relacaoFile = a.selectFile("Relação de processos (*.xlsx)", "*.xlsx")
	return a.relacaoFile
}

// SelectModelFile seleciona o arquivo com o modelo de parecer
func (a *App) SelectModelFile() string {
	a.modeloFile = a.selectFile("Modelo de Impugnação (*.docx)", "*.docx")
	return a.modeloFile
}

// SelectDirSaida seleciona o diretório de saída
func (a *App) SelectDirSaida() string {
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return err.Error()
	}

	a.dirSaida = dir
	return a.dirSaida
}

// ClearFiles limpa os arquivos selecionados
func (a *App) ClearFiles() {
	a.relacaoFile = ""
	a.modeloFile = ""
}

func partesData() (int, string, int) {
	now := time.Now()
	dia := now.Day()
	meses := []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}
	mes := meses[now.Month()-1]
	ano := now.Year()
	return dia, mes, ano
}

func getDefaultReplacements() []parser.Replacement {
	dia, nomeMes, ano := partesData()
	return []parser.Replacement{
		{
			Old: parser.ReplacementParam("dia"),
			New: fmt.Sprintf("%d", dia),
		},
		{
			Old: parser.ReplacementParam("mes"),
			New: nomeMes,
		},
		{
			Old: parser.ReplacementParam("ano"),
			New: fmt.Sprintf("%d", ano),
		},
	}
}

// GerarArquivos
func (a *App) GerarArquivos() string {
	slog.Info("Iniciando geração de arquivos")

	data, err := parser.GetDataFromXlsx(a.relacaoFile, a.nomePlanilha)
	if err != nil {
		return err.Error()
	}

	errs := make([]string, 0)
	replacements := make([]parser.Replacement, 0)
	replacements = append(replacements, getDefaultReplacements()...)
	for _, d := range data[:10] {
		for _, f := range d.Fields {
			replacements = append(replacements, parser.Replacement{
				Old: parser.ReplacementParam(f.VariableName),
				New: f.Value,
			})
		}

		err := parser.ReplaceAndSaveDoc(a.modeloFile, fmt.Sprintf("%s/%s.docx", a.dirSaida, d.Fields[0].Value), replacements...)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		return strings.Join(errs, "; ")
	}

	return "Arquivo gerado com sucesso!"
}
