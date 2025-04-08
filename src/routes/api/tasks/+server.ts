import pool from '$lib/server/db'
import { error, json } from '@sveltejs/kit'

const TASK_GET_SQL = "SELECT * FROM tasks ORDER BY id DESC"

export async function GET():Promise<Response>{
    const [data]:any = await pool.query(TASK_GET_SQL) 
    return json(data, {
        status: 200
    })
}

export async function POST(event):Promise<Response>{
    const post = await event.request.formData()
    if(post.get("title") === ""){
        error(400, "Wrong request")
    }
    try{
        await pool.execute('INSERT INTO tasks(title) VALUES(?)', [post.get("title")])
    }catch(err:any){
        error(400, err.message)
    }
    const [data]:any = await pool.query(TASK_GET_SQL) 
    return json({
        "message":"New task has been aded",
        "tasks": data
    }, {status: 200})
}