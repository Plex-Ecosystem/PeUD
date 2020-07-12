# PeUD - Plex ecosystem User Database
PeUD is in alpha development

### [API Documentation](https://api.peud.io)

## Goals
* Single api to keep track of all user information on all platforms
* Extensibility to the apps it aggregates from
* "Source Of Truth" for all user data
* Extended fields for custom attributes
* Modularity - Api doesnt require UI nor Syncing to work

## Aggregated Ecosystems
* Plex - [Homepage](https://www.plex.tv/) - [Unofficial API Documentation](https://github.com/Arcanemagus/plex-api/wiki)
* Tautulli - [Homepage](https://tautulli.com/) - [API Documentation](https://github.com/Tautulli/Tautulli/blob/master/API.md)
* Ombi - [Homepage](https://ombi.io/) - [API Documentation](https://demo.ombi.io/swagger)
* Organizr - [Homepage](https://organizr.app/) - [API Documentation](https://dev.organizr.app/api/docs/)

## Dev Environment
Want to start helping develop? Most commands are in a Makefile. You need golang 1.14 for the api
and openapi for the docs. then just run
```
make run
```
