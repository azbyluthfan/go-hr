# go-hr
Simple employee notice API

### Running Test
`go test ./... -cover | grep -v mocks`

### How to run
`docker-compose up`

If you need to change default ports, update ports and environment variables specified in `docker/docker-compose.yaml`

### Data
DB client is available at `localhost:8080` with Adminer. 
Default user and password is `root:p4ssw0rd` as configured in docker-compose.yaml.

### Password Implementation
Password is stored with following algorithm:

`hex64(sha256({password}.{salt})=)`

meaning, if a user password is `123456` with salt of `ABCD`, it will first encrypt with SHA256("123456.ABCD") and returns `db5cef685915495b5ec84606b859228d1a71d5e91e76da366e94a8899ede3071`
Hashed password of `db5cef685915495b5ec84606b859228d1a71d5e91e76da366e94a8899ede3071` will then be added with `=` at the end, and will be base64 encoded as `ZGI1Y2VmNjg1OTE1NDk1YjVlYzg0NjA2Yjg1OTIyOGQxYTcxZDVlOTFlNzZkYTM2NmU5NGE4ODk5ZWRlMzA3MT0=`

### Endpoints

#### Login
Will return access token, expires in one hour

example as admin:

`curl -X POST http://localhost:9000/auth/login \
    -H 'Content-Type: application/json' \
    -d '{"companyId": "a0b8dc73-67e9-11e9-91f4-0242ac120002", "employeeNo": "10000", "password":"admin"}'`

example as normal employee:

`curl -X POST http://localhost:9000/auth/login \
    -H 'Content-Type: application/json' \
    -d '{"companyId": "a0b8dc73-67e9-11e9-91f4-0242ac120002", "employeeNo": "10001", "password":"123456"}'` 
      
#### Hello
Will print "Hello" to logged in user

`curl http://localhost:9000/employee/hello \
    -H 'Content-Type: application/json'
    -H 'Authorization: Bearer {token}'`

#### Get Notice
Will return all notice from logged in user

`curl localhost:9000/employee/notice \
    -H 'Content-Type: application/json'
    -H 'Authorization: Bearer {token}'`

#### Get Company Notice
Will return all public notices from employees in a company.
If logged in user role is Admin, it will return both public and private notices.

`curl localhost:9000/employee/noticecompanyId=a0b8dc73-67e9-11e9-91f4-0242ac120002 \
    -H 'Content-Type: application/json'
    -H 'Authorization: Bearer {token}'`

#### Create Notice
Will create a notice for an employee.
Admin can create a notice for other employee.
Can not create notice with overlapping date.

`curl -X POST http://localhost:9000/employee/notice \
    -H 'Content-Type: application/json' \
    -H 'Authorization: Bearer {token}' \
    -d '{"companyId": "a0b8dc73-67e9-11e9-91f4-0242ac120002", "employeeNo": "10000", "type": "sick", "visibility": "private", "periodStart": "2019-04-20", "periodEnd": "2019-04-20"}' ` 
    
