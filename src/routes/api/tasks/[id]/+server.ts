import pool from '$lib/server/db'
import { json } from '@sveltejs/kit'

export async function PATCH({params}):Promise<Response>{
    const [data]:any = await pool.query("SELECT * FROM tasks WHERE id = ?", [params.id])
    if(data.length === 0){
        return json({error: 'Task not found'}, {status: 404})
    }
    const [update]:any = await pool.query("UPDATE tasks SET is_done = NOT is_done WHERE id = ?", [params.id])
    if(update.affectedRows === 0){
        return json({error: "Task not updated"}, {status: 500})
    }
    const task = data[0]
    task.is_done = !task.is_done
    return json({
        message: 'Task updated successfully',
        task: task
    }, {
        status: 200
    })
}