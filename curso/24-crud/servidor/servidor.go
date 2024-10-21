package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type usuario struct {
	ID           int       `json:"id"`
	Nome         string    `json:"nome"`
	Idade        uint32    `json:"idade"`
	DataInsercao time.Time `json:"data_insercao"`
}

// CriarUsuario cria um novo usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpo, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpo, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuário"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	// prepare statement
	statement, erro := db.Prepare("insert into usuarios (nome, idade) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Idade)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()

	if erro != nil {
		w.Write([]byte("Erro ao obter o ID inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso. ID: %d", idInserido)))
}

// BuscarUsuarios busca todos os usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte(erro.Error()))
		return
	}

	defer db.Close()

	linhas, erro := db.Query("select id AS ID, nome AS Nome, idade AS Idade, data_insercao DataInsercao from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuários"))
		return
	}

	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Idade, &usuario.DataInsercao); erro != nil {
			w.Write([]byte("Erro ao escanear usuário"))
			w.Write([]byte(erro.Error()))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	if erro = json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuários"))
		return
	}
}

// BuscarUsuario busca um usuário
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	linhas, erro := db.Query("select * from usuarios where id = ?", 1)
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuário"))
		return
	}

	defer linhas.Close()

	var usuario usuario
	if linhas.Next() {
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Idade); erro != nil {
			w.Write([]byte("Erro ao escanear usuário"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)

	if erro = json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuário"))
		return
	}
}
