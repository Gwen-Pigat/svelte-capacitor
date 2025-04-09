<script lang="ts">
    import { tasks } from "$lib";

    let form:HTMLFormElement
    let title:string = ""

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
        title = ""
        tasks.set(result.tasks)
    }

</script>

<form id="setTask" method="POST" onsubmit={submitTask} bind:this={form}>
    <input type="text" name="title" placeholder="Titre de la tâche" bind:value={title} />
    <button type="submit" disabled={title === ""}>Valid</button>
</form>