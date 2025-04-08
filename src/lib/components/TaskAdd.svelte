<script lang="ts">
    import { tasks } from "$lib";


    export let addTask:boolean

    let form:HTMLFormElement

    async function submitTask(data:any):Promise<void>{
        data.preventDefault()
        const response = await fetch("/api/tasks", {
            method: "POST",
            body: new FormData(form)
        })
        const result = await response.json()
        if(!response.ok){
            console.error(result.error)
            return
        }
        tasks.set(result.tasks)
    }

</script>

<button type="button" onclick={() => addTask = !addTask}>Retour</button>

<form id="setTask" method="POST" onsubmit={submitTask} bind:this={form}>
    <input type="text" name="title" placeholder="Titre de la tâche" />
    <button type="submit">Valid</button>
</form>