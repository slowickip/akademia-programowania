# Fetcher

Celem zadania jest napisanie aplikacji, która pobierze listę wpisów z portalu reddit, i umożliwi zapisanie do pliku, bufora lub innych structów spełniających interfejs `io.Writer`

* Adres URL do odpytywania API:
https://www.reddit.com/r/golang.json
* Struct reprezentujący format danych oraz interfejs znajdują się w pliku `redditFetcher.go`

### Wymagania:
* Aplikacja musi poprawnie pobierać, (de)serializować i zapisywać dane z redditowego API:
  * wyeksportowane pole Title o typie string
  * wyeksportowane pole URL o typie string
* Struct fetcher'a musi implementować interfejs `RedditFetcher`
* Obsłużyć błędy np. poprzez logowanie

#### W celu zaliczenia zadania należy:
1. Zaimplementować własny struct i metody tak by spełniać interfejs RedditFetcher
2. Zapisać pobrany listing do wybranego przez siebie io.Writer'a (plik lub bufor - przykład w pliku `example_output`)


#### Dodatkowo można poeksperymentować:
1. Klient używa Context z timeoutem do połączeń HTTP z wykorzystaniem http.NewRequestWithContext()
2. Współbieżne pobieranie innych wybranych subredditów SFW
3. Można rozszerzyć interfejs o obsługę Contextu

