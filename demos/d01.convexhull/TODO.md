# A regarder

Pour aller plus loin, on peut tenter d'implémenter un algorithme de calcul
d'enveloppe concave (concave hull algorithme). La méthode la plus documentée est
la méthode utilisant les alpha-shape, qui conciste à faire une triangulation de
delaynay de l'ensemble des points et de ne garder que les les triangles qui
respecte un critère sur la longueur des arrêtes (à l'intérieur du cercle
circonscrit de rayon alpha)

On peut trouver néanmoins des méthodes alternative, comme par exemple celle-ci:

* https://medium.com/data-science/the-concave-hull-c649795c0f0f

Fondé sur la méthode décrite dans l'article: "Concave hull: A k-nearest
neighbours approach for the computation of the region occupied by a set of
points", de A. Moreira, M. Y. Santos:

* https://www.semanticscholar.org/paper/Concave-hull%3A-A-k-nearest-neighbours-approach-for-a-Moreira-Santos/319a3450f9909043d46eb7ceb4299efceb984d4f?p2df