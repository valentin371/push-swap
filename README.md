# à faire 
 ctrl + shift + v (permet d'afficher le README.md en mieux)
# push-swap

Projet en Go avec deux programmes :

- **push-swap** : reçoit une liste d'entiers et affiche, une instruction par ligne, de quoi trier la pile `a`.
- **checker** : reçoit la même liste, lit des instructions sur l'entrée standard et affiche `OK` si la pile `a` est triée (et `b` vide), sinon `KO`.

On trie `a` par ordre croissant (le plus petit au sommet) avec une seconde pile `b` et ces instructions :
`sa` `sb` `ss` (échange), `pa` `pb` (déplacement entre piles), `ra` `rb` `rr` (rotation), `rra` `rrb` `rrr` (rotation inverse).

## Compiler

```
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker
```
## Tester dans le terminal

```
go run ./cmd/bench 100 20    # 20 listes de 100 nombres : réussite et nombre d'instructions 
On peut changer le 20 et 100 par exemple:      go run ./cmd/bench 200 10             
               
./push-swap "2 1 3 6 5 8"                                  # affiche les instructions
ARG="4 67 3 87 23"; ./push-swap "$ARG" | wc -l             # nombre d'instructions
ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"  # affiche OK
echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"       # affiche OK
./push-swap "0 one 2 3"                                    # affiche Error
```

Sans argument, les programmes n'affichent rien. `Error` est écrit sur la sortie d'erreur si un argument n'est pas un entier, s'il y a un doublon ou (checker) si une instruction n'existe pas.
