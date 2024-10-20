package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

	parametros := mux.Vars(r)

	ID, erro := strconv.Atoi(parametros["id"])
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuário"))
		return
	}

	defer linha.Close()

	var usuario usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Idade, &usuario.DataInsercao); erro != nil {
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

// AtualizarUsuario atualiza um usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	ID, erro := strconv.Atoi(parametros["id"])
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

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

	statement, erro := db.Prepare("update usuarios set nome = ?, idade = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Idade, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario deleta um usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	ID, erro := strconv.Atoi(parametros["id"])
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
