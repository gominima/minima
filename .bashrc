alias test='go test'
alias run='go run _examples/main.go'
alias lint='go fmt ./'

if [ -f $HOME/.bash_aliases ]
then
  . $HOME/.bash_aliases
f