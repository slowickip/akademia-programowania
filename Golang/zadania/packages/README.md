# Packages, exported names, go mod

Celem tego zadania jest napisanie prostej aplikacji, która wyśle zgłoszenie
z (opcjonalną) pracą domową.

Aplikacja musi zawierać wyeksportowany struct Student o następujących
polach:
* wyeksportowane pole firstName o typie string
* wyeksportowane pole lastName o typie string
* niewyksportowane pole applicationID o typie uuid.UUID

Struct ten musi mieć również dwie wyeksportowane metody:
* applicationID zwracającą wartość pola applicationID jako string
* fullName zwracającą firstName z dodanym po spacji lastName

Zaimplementowanie tych dwóch metod sprawi że struct Student będzie implementował
interfejs `appdispatcher.Application` i będzie go można przekazać funkcji `Submit`.

Do wygenerowania UUID służy paczka `github.com/google/uuid`.

W celu zaliczenia zadania należy:
1. Zaimplementować wspomniany struct i metody
2. Zaimportować `github.com/grupawp/appdispatcher`
3. Zainicjować struct z waszym imieniem i nazwiskiem
4. Przekazać go funkcji `Submit` z paczki `appdispatcher`

Metoda `Submit` zwraca dwie wartości - HTTP status code oraz błąd.
Błąd zawiera szczegółowe informacje o przyczynie niepowiedzienia i należy
go zalogować.

O powodzeniu wysyłki świadczy status code `200`. Możecie również otrzymać
`208`, jeśli wasza aplikacja została już wcześniej poprawnie wysłana.

Aplikację należy uruchomić przy użyciu `go run main.go`.