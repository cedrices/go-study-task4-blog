# go-study-task4-blog
个人博客系统后段


1.使用idea的版本
IntelliJ IDEA 2023.3.8 (Ultimate Edition)

2.启动文件
go-study-task4-blog/startUp.go



url 

//注册 response header中 返回 Authorization
curl --request POST \
--url http://localhost:8080/v1/user/register \
--header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Connection: keep-alive' \
--header 'Content-Type: application/json' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
--data '{
"username":"haos",
"password":"haos123",
"email":"123@163.com"
}'


//登录
curl --request POST \
--url http://localhost:8080/v1/user/login \
--header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Connection: keep-alive' \
--header 'Content-Type: application/json' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
--data '{
"username":"haos",
"password":"haos123"
}'

//创建文章
curl --request POST \
--url http://127.0.0.1:8080/v1/post/create \
--header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiaGFvcyIsIlBhc3N3b3JkIjoiJDJhJDE0JGY4S3VSQzdBc0lvRlBWOWhMeWJLdHV6aEdldTE2WXdtSEJRQXh4VW9XWVlVL0xzenFpLzZHIiwiaXNzIjoiaHNzIiwic3ViIjoiZ28tYmxvZyIsImF1ZCI6WyJoc3MiXSwiZXhwIjoxNzU4MzYwNDUwLCJuYmYiOjE3NTgyNzQwNTAsImlhdCI6MTc1ODI3NDA1MCwianRpIjoiZ28tYXBpLTEifQ.vwS-qbWhBOmzx9uLsMwhZwYJkQKN5DWbtg03KYJwQZU' \
--header 'Connection: keep-alive' \
--header 'Content-Type: application/json' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
--data '{
"title":"文章标题",
"content":"文章内容11233"
}'

//获取文章列表
curl --request GET \
--url http://127.0.0.1:8080/v1/post/list \
--header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiaGFvcyIsIlBhc3N3b3JkIjoiJDJhJDE0JGY4S3VSQzdBc0lvRlBWOWhMeWJLdHV6aEdldTE2WXdtSEJRQXh4VW9XWVlVL0xzenFpLzZHIiwiaXNzIjoiaHNzIiwic3ViIjoiZ28tYmxvZyIsImF1ZCI6WyJoc3MiXSwiZXhwIjoxNzU4MzYwNDUwLCJuYmYiOjE3NTgyNzQwNTAsImlhdCI6MTc1ODI3NDA1MCwianRpIjoiZ28tYXBpLTEifQ.vwS-qbWhBOmzx9uLsMwhZwYJkQKN5DWbtg03KYJwQZU' \
--header 'Connection: keep-alive' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0'

//修改文章
curl --request PUT \
--url http://127.0.0.1:8080/v1/post/update \
--header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiaGFvcyIsIlBhc3N3b3JkIjoiJDJhJDE0JGY4S3VSQzdBc0lvRlBWOWhMeWJLdHV6aEdldTE2WXdtSEJRQXh4VW9XWVlVL0xzenFpLzZHIiwiaXNzIjoiaHNzIiwic3ViIjoiZ28tYmxvZyIsImF1ZCI6WyJoc3MiXSwiZXhwIjoxNzU4MzYwNDUwLCJuYmYiOjE3NTgyNzQwNTAsImlhdCI6MTc1ODI3NDA1MCwianRpIjoiZ28tYXBpLTEifQ.vwS-qbWhBOmzx9uLsMwhZwYJkQKN5DWbtg03KYJwQZU' \
--header 'Connection: keep-alive' \
--header 'Content-Type: application/json' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
--data '{
"id":1,
"title":"测试修改文章标题"
}'

//逻辑删除文章
curl --request DELETE \
--url http://127.0.0.1:8080/v1/post/delete/1 \
--header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiaGFvcyIsIlBhc3N3b3JkIjoiJDJhJDE0JGY4S3VSQzdBc0lvRlBWOWhMeWJLdHV6aEdldTE2WXdtSEJRQXh4VW9XWVlVL0xzenFpLzZHIiwiaXNzIjoiaHNzIiwic3ViIjoiZ28tYmxvZyIsImF1ZCI6WyJoc3MiXSwiZXhwIjoxNzU4MzYwNDUwLCJuYmYiOjE3NTgyNzQwNTAsImlhdCI6MTc1ODI3NDA1MCwianRpIjoiZ28tYXBpLTEifQ.vwS-qbWhBOmzx9uLsMwhZwYJkQKN5DWbtg03KYJwQZU' \
--header 'Connection: keep-alive' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0'

//发表评论
curl --request POST \
--url http://127.0.0.1:8080/v1/comment/publish \
--header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Connection: keep-alive' \
--header 'Content-Type: application/json' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
--data '{
"content":"文章棒棒的！！！",
"postId":2
}'

//获取评论
curl --request GET \
--url 'http://127.0.0.1:8080/v1/comment/list?postId=2' \
--header 'Accept: */*' \
--header 'Accept-Encoding: gzip, deflate, br' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsIlVzZXJuYW1lIjoiaGFvcyIsIlBhc3N3b3JkIjoiJDJhJDE0JGY4S3VSQzdBc0lvRlBWOWhMeWJLdHV6aEdldTE2WXdtSEJRQXh4VW9XWVlVL0xzenFpLzZHIiwiaXNzIjoiaHNzIiwic3ViIjoiZ28tYmxvZyIsImF1ZCI6WyJoc3MiXSwiZXhwIjoxNzU4MzYwNDUwLCJuYmYiOjE3NTgyNzQwNTAsImlhdCI6MTc1ODI3NDA1MCwianRpIjoiZ28tYXBpLTEifQ.vwS-qbWhBOmzx9uLsMwhZwYJkQKN5DWbtg03KYJwQZU' \
--header 'Connection: keep-alive' \
--header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0'