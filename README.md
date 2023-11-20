# fun-blog
fun的个人博客


```shell
docker run -itd --name mysql \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=test \
    --restart unless-stopped \
    mysql:latest
```
