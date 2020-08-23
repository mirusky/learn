# URI

uri folder refers to questions obtained from [URI Online Judge](https://www.urionlinejudge.com.br/judge/en)

## Dependencies

All dependencies are installed using [go modules](https://blog.golang.org/using-go-modules) so you just need to use a golang compatible version. To download the dependencies you can run the following:

```bash
cd golang
go mod download
```

## Usage

All questions are under the question difficulty level folder and then the question number.To run any program just type the following:
```bash
cd golang/cmd/uri/<difficulty level>/<question number>
go run main.go
```

To find the wording of the question:
```
https://www.urionlinejudge.com.br/judge/en/problems/view/<question number>
```

## License
[MIT](../../../LICENSE)