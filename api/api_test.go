package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGenreComics(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Test response body"))
		if err != nil {
			t.Errorf("w.Write エラー: %v", err)
		}
	}))
	defer server.Close()

	api := &comicAPI{baseURL: server.URL + "?"}

	t.Run("正常系", func(t *testing.T) {
		resp, err := api.GetGenreComics("shonen", 1)
		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("HTTPステータスコードが異常です: %d", resp.StatusCode)
		}
	})

	t.Run("ジャンルコードが不正", func(t *testing.T) {
		_, err := api.GetGenreComics("不正ジャンル", 1)
		if err == nil {
			t.Error("エラーが発生するはずでしたが、エラーが発生しませんでした")
		}
	})
}

func TestGetSearchComics(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Test response body"))
		if err != nil {
			t.Errorf("w.Write エラー: %v", err)
		}
	}))
	defer server.Close()

	api := &comicAPI{baseURL: server.URL + "?"}

	t.Run("正常系", func(t *testing.T) {
		resp, err := api.GetSearchComics("test", 1)
		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("HTTPステータスコードが異常です: %d", resp.StatusCode)
		}
	})
}
