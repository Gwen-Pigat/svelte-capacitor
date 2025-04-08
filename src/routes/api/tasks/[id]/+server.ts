import pool from '$lib/server/db'
import { json } from '@sveltejs/kit'


export async function PATCH({params}):Promise<Response>{
    const [data]:any = await pool.query("SELECT * FROM tasks WHERE id = ?", [params.id])
    if(data.length === 0){
        return json({error: 'Task not found'}, {status: 404})
    }
    const task = data[0]
    let dateNow:Date|null = new Date()
    if(task.is_done){
        dateNow = null
    }
    const [update]:any = await pool.query(
        "UPDATE tasks SET is_done = NOT is_done, date_to = ? WHERE id = ?", 
        [dateNow, task.id]
    )
    if(update.affectedRows === 0){
        return json({error: "Task not updated"}, {status: 500})
    }
    task.is_done = !task.is_done
    task.date_to = dateNow
    return json({
        message: 'Task updated successfully',
        task: task
    }, {
        status: 200
    })
}

export async function DELETE({params}):Promise<Response>{
    const [data]:any = await pool.query("SELECT * FROM tasks WHERE id = ?", [params.id])
    if(data.length === 0){
        return json({error: 'Task not found'}, {status: 404})
    }
    const [update]:any = await pool.query("DELETE FROM tasks WHERE id = ?", [params.id])
    if(update.affectedRows === 0){
        return json({error: "Task not deleted"}, {status: 500})
    }
    return json({
        message: 'Task deleted successfully',
    }, {
        status: 200
    })
}