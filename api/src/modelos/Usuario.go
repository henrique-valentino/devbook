package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (u *Usuario) Preparar(etapa string) error {
	if erro := u.validar(etapa); erro != nil {
		return erro
	}
	if erro := u.formartar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (u *Usuario) validar(etapa string) error {

	if u.Nome == "" {
		return errors.New("o campo nome é obrigatório")
	}

	if u.Nick == "" {
		return errors.New("o campo nick é obrigatório")
	}

	if u.Email == "" {
		return errors.New("o campo email é obrigatório")
	}

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("o email informado é inválido")
	}

	if etapa == "cadastro" && u.Senha == "" {
		return errors.New("o campo senha é obrigatório")
	}

	return nil
}

func (u *Usuario) formartar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(u.Senha)
		if erro != nil {
			return erro
		}

		u.Senha = string(senhaComHash)
	}
	return nil
}
