# User_Management

## Problématique
Développer une API REST permettant de gérer des utilisateurs. On doit pouvoir créer/mettre à jour/supprimer des utilisateurs, mais aussi de récupérer un ou la liste des utilisateurs. Ces utilisateurs seront stockés sur MongoDB. En plus de ça, il faudrait avoir une API permettant de se connecter sur son profil afin de pouvoir accéder à toutes les données qui lui sont attribuées.

## Arborescence projet
Ce projet est composé de :
- Un dossier **Auth** contenant le code source pour la partie authentification et gestion de token.
- Un dossier **DataBase** contenant le code source pour la connexion et les APIs de communication avec la DataBase.
- Un dossier **Routes** contenant le code source pour la déclaration des routes.
- Un dossier **Users** contenant les fichiers : 
    - **create.go** pour la création et l'ajout des utilisateurs.
    - **delete.go** pour la suppression d'utilisateurs.
    - **read.go** pour la récupération d'un ou tous les utilisateurs.
    - **update.go** pour la mise à jour d'utilisateur.
    - **login.go** pour se connecter sur son profile.
    - **create_test.go** test le create de plusieurs users
    - **login_test.go** test le login
    - **read_test.go** test la récupération d'un ou tous les utilisateurs, avec les uses case où le token est faux ou pas fourni.
    - **update_test.go** test d'un users, avec les uses case où le token est faux ou pas fourni. Et si on essaye d'ajouter un champ qui n'existe pas
    - **delete_test.go** test le delete d'un utilisateur, avec les uses case où le token est faux ou pas fourni.
- Un fichier **main.go** contenant la fonction principal 
- Un fichier **Makefile** conenant comme cible : 
    - **run_server** pour lancer le server
    - **docker_compose_up** pour lancer le container mongo
    - **docker_compose_down** pour arrêter le container mongo
    - **test** pour lancer les tests

## Mode d'emploie
Pour Tester le projet il faut :
1. Lancer d'abord `` make docker_compose_up ``  dans un terminal
2. Lancer ensuite `` make run_server `` dans un autre terminal
3. Lancer enfin `` make test `` pour les tests (les outputs s'afficheront dans stdout)
    - Pour le test de l'update, il y a un GetUser qui est fait ensuite pour voir le changement de l'age du user qui devient **50**, son nom qui devient **Linus Torvalds** et la Data qui devient **Test_file_change**, et qu'on peut remarquer aussi le changement dans le fichier text correspondant.

## Implementation
L'implémentation a été fait comme suite:
- Pour la création, lorsqu'on reçoit une requête, on crée un tableau de User, on supprime les doublons de la liste et on crée une goroutine qui s'occupera de vérifier si le user existe dans la DB, et si ce n'est pas le cas, de l'insérer. Il faut savoir ici que :
    - Le UpdateOne a été choisi comme API mongo pour l'insértion afin d'éviter d'insérer des doublons.
    - Dans la fonction **hashAndInsertUser**, l'appel à la fonction **addUserInDatabase** est protégé par un mutex pour éviter de se retrouver dans le cas où on a deux clients qui font une requête, et que les listes contiennent des users en commun. Si un changement de context se fait entre le **FindOne** et le **UpdateOne**, on peut se retrouver avec deux threads qui vont insérer le même user.
- Pour le login, on utilise la lib **jwt-go** permettant de générer des token à partir de l'id et le mot de passe de l'utilisateur. On regarde d'abord si l'utilisateur existe et que le mot de passe est bon pour retourner un token, qui sera utilisé ensuite pour pouvoir accéder aux autres routines (**GET /users/list**, **GET /user/:id**, **PATCH /user/:id** et **DELETE /delete/user/:id**)
- Pour le Read, on utilise l'API **Find** avec un objet bson vide pour récupérer tous les utilisateurs, et un **FindOne** avec un filtrage sur l'Id pour récupérer le user correspondant. On retourne ensuite un JSON contenant le résultat.
- Pour l'update, on utilise l'API **UpdateOne** avec un Upsert == false pour éviter d'insérer des documents qui ne trouve pas. En input, on aura un JSON contenant les champs modifiés de l'utilisateur. On créera une Map à partir de ce JSON et on mettra à jour la BD. Si le champ Data a été modifié, on modifiera le fichier text concerné. Si on essaye d'envoyer un champ qui n'existe pas, ce champ sera ignoré et non ajouté au user.
- Pour le delete, on utilise l'API **DeleteOne** pour supprimer le user de la DB.
- Tous les cas d'erreurs ont été gérés.
- **DataSet.json** contient un tableau d'utilisateurs avec des doublons.