# Noteify

---

Is a Personal Blogging Platform API

## Noteify API endpoints

|  Method  |    URL Pattern     |       Handler        | Action                          |
| :------: | :----------------: | :------------------: | :------------------------------ |
|  `GET`   | `/v1/healthcheck`  |  healthcheckHandler  | Display application information |
|  `GET`   |   `/v1/articles`   | listArticlesHandler  | Return a list of articles       |
|  `POST`  |   `/v1/articles`   | createArticleHandler | Create a new article            |
|  `GET`   | `/v1/articles/:id` |  showArticleHandler  | Return a single article         |
|  `PUT`   | `/v1/articles/:id` |  editArticleHandler  | Update a single article         |
| `DELETE` | `/v1/articles/:id` | deleteArticleHandler | Delete a single article         |
