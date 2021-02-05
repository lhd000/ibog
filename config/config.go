package config

/*

 框架的配置目录

*/
/*----------------------------------APPLICATION CONFIG------------------------------------------ */

//应用端口号
const GIN_PORT string = "8081"

//应用模板目录 ， 注意 */*
const GIN_TEMPLATE_DIR string = "templates/**/*"

//数据的存储
const BLOG_DATA string = "data/data.db"

/*---------------------------------- MYSQL CONFIG -----------------------------------------------*/

// mysql 主机
const MYSQL_HOST string = "127.0.0.1"

//mysql  数据库
const MYSQL_DATABASE string = "blog"

// mysql 端口号
const MYSQL_PORT string = "3306"

// mysql 用户名
const MYSQL_USER string = "root"

//mysql  密码
const MYSQL_PASS string = "root"

//链接池
const MYSQL_MAX_IDES int = 10
const MYSQL_MAX_CONNECTIONS int = 100

//基本目录
const BASE_DIR string = "./"

//cookie 加密字符串窜 ,加密 密钥
const COOKIE_SECRET_STR = "secret11111"

//session 名字
const SESSION_STORE_NAME = "mysession"
