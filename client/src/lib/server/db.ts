import mysql from 'mysql2/promise'
import { env } from '$env/dynamic/private'


const pool:mysql.Pool = mysql.createPool({
    host: env.DB_HOST,
    user: env.DB_USER,
    password: env.DB_PASSWORD,
    database: env.DB_NAME,
})

export default pool