package models

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id, omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formartar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatorio")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatorio")
	}

	if usuario.Email == "" {
		return errors.New("O email é obrigatorio")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O email é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatoria")
	}

	return nil
}

func (usuario *Usuario) formartar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}
	return nil
}
