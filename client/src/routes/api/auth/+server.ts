import pool from '$lib/server/db.js'
import { error, json } from '@sveltejs/kit'

const USERNAME = "Gwen"

export async function GET({cookies}):Promise<any>{
    if(cookies.get("user") !== USERNAME){
        return json({
            "message":"You have to connect"
        }, {status: 200})
    }
    return json({
        "user":cookies.get("user")
    }, {status: 200})
}

export async function POST({request,cookies}):Promise<any>{
    const post = await request.formData()
    if(post.get("user") === ""){
        error(400, "You have to specify a username")
    }
    const [data]:any = await pool.query(
        "SELECT * FROM user WHERE username = ?",
        [post.get("user")]
    )
    if(data.length === 0){
        try{
            await pool.execute(
                'INSERT INTO user(username,date_add,is_active) VALUES(?,?,?)', 
                [post.get("user"), new Date(), 1]
            )
        }catch(err:any){
            error(400, err.message)
        }
    }
    let username:any = post.get("user")
    cookies.set("user", username, {path: "/"})
    return json({
        "user":cookies.get("user")
    }, {status: 200})
}